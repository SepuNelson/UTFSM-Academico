# Tarea 3 - ETL Fan Engagement HRL

## Descripción

Este workspace implementa una ETL (Extract, Transform, Load) para convertir datos de participación de fans de la Liga de Carreras de Helicópteros (HRL) de formato JSON a Avro, utilizando Apache Beam y Apache Airflow.

## Arquitectura

- **Apache Beam**: Pipeline de ETL que lee archivos JSON, transforma los datos agregando el campo `Timestamp_unix` y escribe en formato Avro
- **Apache Airflow**: Orquestador que ejecuta el pipeline diariamente y envía notificaciones a Kafka
- **Kafka**: Sistema de mensajería para notificaciones de finalización de procesamiento
- **MinIO**: Almacenamiento de objetos compatible con S3 (opcional)

## Estructura del Proyecto

```
workspace_hrl/
├── .devcontainer/
│   ├── docker-compose.yml    # Configuración de servicios (Airflow, Kafka, MinIO)
│   └── devcontainer.json     # Configuración del contenedor de desarrollo
├── dags/
│   └── fan_engagement_dag.py # DAG de Airflow para la ETL
├── data/
│   └── fan_engagement.jsonl  # Datos de entrada en formato JSONL
├── output/                   # Directorio de salida para archivos Avro
├── schemas/
│   └── fan_engagement_schema.json # Esquema Avro
├── src/
│   ├── beam/
│   │   └── fan_engagement_etl.py # Pipeline de Beam
│   └── settings.py           # Configuraciones centralizadas
├── logs/                     # Logs de Airflow
├── plugins/                  # Plugins de Airflow
└── requirements.txt          # Dependencias de Python
```

## Requisitos Previos

- Docker Desktop
- Visual Studio Code con extensión Dev Containers
- Al menos 4GB de RAM disponible

## Instrucciones de Ejecución

### 1. Iniciar el Workspace

1. Abrir el proyecto en Visual Studio Code
2. Cuando se solicite, hacer clic en "Reopen in Container"
3. Esperar a que se construya el contenedor de desarrollo

### 2. Verificar Servicios

Una vez que el contenedor esté listo, verificar que todos los servicios estén ejecutándose:

```bash
docker ps
```

Deberías ver los siguientes servicios:
- `ws-airflow-webserver` (puerto 8081)
- `ws-airflow-scheduler`
- `ws-kafka` (puerto 9092)
- `ws-kafka-ui` (puerto 8080)
- `ws-minio` (puertos 9000, 9001)

### 3. Acceder a Airflow

1. Abrir el navegador y ir a: `http://localhost:8081`
2. Credenciales por defecto:
   - Usuario: `airflow`
   - Contraseña: `airflow`

### 4. Ejecutar el DAG

1. En la interfaz de Airflow, buscar el DAG `fan_engagement_etl_dag`
2. Hacer clic en el botón "Play" para ejecutar manualmente
3. O esperar a que se ejecute automáticamente a las 2:00 AM diariamente

### 5. Monitorear la Ejecución

1. Hacer clic en el DAG para ver los detalles
2. Verificar que ambas tareas se completen exitosamente:
   - `run_beam_pipeline`: Ejecuta la ETL de Beam
   - `send_kafka_notification`: Envía notificación a Kafka

### 6. Verificar Resultados

1. **Archivos Avro**: Los archivos de salida se generan en `/workspaces/output/`
2. **Notificaciones Kafka**: Verificar en Kafka UI (`http://localhost:8080`) el tópico `fan_engagement_notifications`

## Configuración

### Variables de Entorno

Las configuraciones principales están en `src/settings.py`:

- **Kafka**: `ws-kafka:19092`
- **MinIO**: `ws-minio:9000`
- **Bucket**: `fan-engagement-bucket`

### Esquema Avro

El esquema de salida está definido en `schemas/fan_engagement_schema.json` e incluye:

- `FanID`: Identificador único del fan
- `RaceID`: Identificador de la carrera
- `Timestamp`: Timestamp original en formato string
- `Timestamp_unix`: Timestamp en milisegundos (nuevo campo calculado)
- `ViewerLocationCountry`: País del espectador
- `DeviceType`: Tipo de dispositivo
- `EngagementMetric_secondswatched`: Segundos vistos
- `PredictionClicked`: Si accedió a predicciones
- `MerchandisingClicked`: Si accedió a merchandising

## Solución de Problemas

### Error de Conexión a Kafka
- Verificar que el servicio `ws-kafka` esté ejecutándose
- Comprobar la configuración en `src/settings.py`

### Error en Pipeline de Beam
- Verificar que el archivo de entrada existe en `data/fan_engagement.jsonl`
- Comprobar que el esquema Avro es válido

### Airflow no Inicia
- Verificar que PostgreSQL esté ejecutándose
- Revisar logs en el directorio `logs/`

## Desarrollo Local

Para ejecutar el pipeline de Beam directamente sin Airflow:

```bash
cd /workspaces
python src/beam/fan_engagement_etl.py \
    --input "data/fan_engagement.jsonl" \
    --output "output/fan_engagement" \
    --schema "schemas/fan_engagement_schema.json"
```

## Notificaciones Kafka

Al finalizar el procesamiento, se envía automáticamente un mensaje JSON a Kafka con la siguiente estructura:

```json
{
  "event_type": "data_processing_completed",
  "data_entity": "FanEngagement",
  "status": "success",
  "location": "/opt/airflow/output/fan_engagement_*.avro",
  "processed_at": "2025-01-01 02:00:00",
  "source_system": "fan_engagement_etl_dag"
}
```

## Contacto

Para soporte técnico o preguntas sobre la implementación, contactar al equipo de desarrollo.
