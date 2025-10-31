# configuraciones

# --- Configuración de MinIO ---
MINIO_ENDPOINT = 'ws-minio:9000'  #  'your_minio_server:9000'
MINIO_ACCESS_KEY = 'minio-root-user'   # Reemplace con su nombre de usuario MinIO
MINIO_SECRET_KEY = 'minio-root-password'   # Reemplace con su contraseña MinIO
MINIO_SECURE = False #  True si usa HTTPS
bucket_name = 'fan-engagement-bucket'    # El bucket donde se almacenarán los datos

# --- Configuración de Kafka ---
KAFKA_BROKERS = 'ws-kafka:19092'  # direccion del Kafka broker, 'server:port'
KAFKA_TOPIC = 'fan_engagement_notifications'     #  Kafka topic

# --- Configuración de Airflow ---
AIRFLOW_HOME = '/opt/airflow'
DAGS_FOLDER = '/opt/airflow/dags'