from datetime import datetime, timedelta
from airflow import DAG
from airflow.operators.python import PythonOperator
from airflow.operators.bash import BashOperator
import json
import sys
import os

# Agregar el directorio src al path para importar settings
sys.path.append('/opt/airflow/src')
from settings import KAFKA_BROKERS, KAFKA_TOPIC

default_args = {
    'owner': 'airflow',
    'depends_on_past': False,
    'start_date': datetime(2025, 1, 1),
    'email_on_failure': False,
    'email_on_retry': False,
    'retries': 1,
    'retry_delay': timedelta(minutes=5),
}

def send_kafka_notification(**context):
    """
    Envía notificación a Kafka cuando el pipeline se completa exitosamente
    """
    try:
        from kafka import KafkaProducer
        import json
        from datetime import datetime
        
        # Obtener información del contexto de Airflow
        task_instance = context['task_instance']
        dag_id = context['dag'].dag_id
        
        # Configurar el productor de Kafka
        producer = KafkaProducer(
            bootstrap_servers=KAFKA_BROKERS,
            value_serializer=lambda v: json.dumps(v).encode('utf-8'),
            retries=3,
            acks='all'
        )
        
        # Crear el mensaje de notificación
        notification_message = {
            "event_type": "data_processing_completed",
            "data_entity": "FanEngagement",
            "status": "success",
            "location": "/opt/airflow/output/fan_engagement_*.avro",
            "processed_at": datetime.now().strftime('%Y-%m-%d %H:%M:%S'),
            "source_system": dag_id,
            "dag_run_id": context['dag_run'].run_id
        }
        
        # Enviar mensaje a Kafka
        producer.send(KAFKA_TOPIC, notification_message)
        producer.flush()
        producer.close()
        
        print(f"Notificación enviada a Kafka: {notification_message}")
        return "Notificación enviada exitosamente"
        
    except Exception as e:
        print(f"Error enviando notificación a Kafka: {str(e)}")
        raise e

def verify_output_files(**context):
    """
    Verifica que los archivos de salida se hayan generado correctamente
    """
    import glob
    import os
    
    output_pattern = "/opt/airflow/output/fan_engagement*.avro"
    files = glob.glob(output_pattern)
    
    if not files:
        raise Exception(f"No se encontraron archivos de salida con el patrón: {output_pattern}")
    
    print(f"Archivos de salida encontrados: {files}")
    return f"Se encontraron {len(files)} archivos de salida"

# Definir el DAG
dag = DAG(
    'fan_engagement_etl_dag',
    default_args=default_args,
    description='ETL Pipeline para convertir datos de fan engagement de JSON a Avro',
    schedule_interval='0 2 * * *',  # Ejecutar diariamente a las 2:00 AM
    catchup=False,
    tags=['etl', 'fan_engagement', 'json_to_avro']
)

# Tarea 1: Ejecutar el pipeline de Beam
run_beam_pipeline = BashOperator(
    task_id='run_beam_pipeline',
    bash_command="""
    cd /opt/airflow && \
    python src/beam/fan_engagement_etl.py \
        --input "/opt/airflow/data/fan_engagement.jsonl" \
        --output "/opt/airflow/output/fan_engagement" \
        --schema "/opt/airflow/schemas/fan_engagement_schema.json"
    """,
    dag=dag
)

# Tarea 2: Verificar archivos de salida
verify_output = PythonOperator(
    task_id='verify_output_files',
    python_callable=verify_output_files,
    dag=dag
)

# Tarea 3: Enviar notificación a Kafka
send_notification = PythonOperator(
    task_id='send_kafka_notification',
    python_callable=send_kafka_notification,
    dag=dag
)

# Definir el flujo de tareas
run_beam_pipeline >> verify_output >> send_notification 