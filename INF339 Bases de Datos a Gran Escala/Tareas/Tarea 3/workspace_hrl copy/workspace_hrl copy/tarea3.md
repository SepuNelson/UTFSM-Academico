````markdown
# Tarea 3

**Bases de Datos a gran escala**

**mail:** claudio.torresf@usm.cl

## 1. Objetivo

[cite_start]El objetivo de esta tarea es aplicar tus conocimientos en Apache Beam y Apache Airflow para construir una ETL que permita convertir los datos de participación de fans de la Liga de Carreras de Helicópteros (HRL) de Json a Avro, para su posterior uso en el entrenamiento de modelos de Machine Learning. [cite: 1]

## 2. Descripción

[cite_start]La Liga de Carreras de Helicópteros (HRL) recolecta información de la participación de sus fans en sus aplicaciones durante sus transmisiones. [cite: 2] [cite_start]El objetivo es entrenar modelos de Machine Learning que permitan predecir si un fan utilizará el servicio de predicciones y/o comprará artículos de Merchandising con el fin de ofrecer banners de publicidad personalizados. [cite: 3] [cite_start]Los datos se recolectan en formato Json, con la siguiente estructura: [cite: 4]

* [cite_start]**FanID:** identificador único del fan. [cite: 4]
* [cite_start]**RaceID:** identificador de la carrera visualizada por el fan. [cite: 5]
* [cite_start]**Timestamp:** momento en el que el fan inicio la transmisión de la carrera, en formato `%Y-%m-%d %H:%M:%S`. [cite: 6]
* [cite_start]**ViewerLocationCountry:** país desde donde visualiza la carrera. [cite: 7]
* [cite_start]**DeviceType:** tipo de dispositivo desde el cual visualiza la carrera. [cite: 7]
* [cite_start]**EngagementMetric_secondswatched:** cantidad de segundos que un fan ha visualizado la carrera. [cite: 8]
* [cite_start]**PredictionClicked:** True si accedió a la sección de predicciones, False en caso contrario. [cite: 9]
* [cite_start]**MerchandisingClicked:** True si accedió a la sección de compra de Merchandising, False en caso contrario. [cite: 10]

[cite_start]Los ingenieros de ML le han pedido convertir la columna Timestamp a Unix timestamp en milisegundos, por lo que se ha agregado una nueva columna: Timestamp_unix. [cite: 11] [cite_start]La plataforma de ML soporta datos en formato Avro, los datos de salida deben utilizar el siguiente esquema Avro: [cite: 12]

```json
{
  "type": "record",
  "name": "FanEngagement",
  "fields": [
    {"name": "FanID", "type": "string"},
    {"name": "RaceID", "type": "string"},
    {"name": "Timestamp", "type": "string"},
    {
      "name": "Timestamp_unix",
      "type": {
        "type": "long",
        "logicalType": "timestamp-millis"
      }
    },
    {"name": "ViewerLocationCountry", "type": "string"},
    {"name": "DeviceType", "type": "string"},
    {"name": "EngagementMetric_secondswatched", "type": "int"},
    {"name": "PredictionClicked", "type": "boolean"},
    {"name": "MerchandisingClicked", "type": "boolean"}
  ]
}
````

Se incluye un fichero JsonL con datos de prueba.

## 3\. Actividades

[cite\_start]Para los requerimientos mencionados, deberás (100 puntos): [cite: 14]

1.  [cite\_start]Construir una ETL en Apache Beam que realice la conversión de los archivos Json a Avro con el esquema entregado. [cite: 14] [cite\_start]Su ETL deberá calcular el campo `Timestamp_unix`. [cite: 15] [cite\_start]Debe tener en cuenta que los datos de entrada y de salida pueden estar en un bucket y/o en el sistema de ficheros local. [cite: 15] [cite\_start](40 puntos) [cite: 16]

2.  [cite\_start]Construir un Dag en Airflow que permita ejecutar en forma diaria el procesamiento de los archivos Json de entrada. [cite: 16] [cite\_start](40 puntos) [cite: 17] [cite\_start]Al finalizar el job de Apache Beam deberá notificar a un tópico de Kafka que lo datos están disponibles para consumo, con el siguiente mensaje Json: [cite: 17]

    ```json
    {
      "event_type": "data_processing_completed",
      "data_entity": "FanEngagement",
      "status": "success",
      "location": "path_or_bucket",
      "processed_at": "processed_timestamp",
      "source_system": "pipeline_name"
    }
    ```

    Con:

      * [cite\_start]**path\_or\_bucket:** path o URI de los datos procesados. [cite: 18]
      * [cite\_start]**processed\_timestamp:** momento en el que se terminó el procesamiento de los datos y se envió el mensaje, en formato `%Y-%m-%d %H:%M:%S`. [cite: 18]
      * [cite\_start]**pipeline\_name:** nombre de su Dag. [cite: 19]

3.  [cite\_start]Construya un workspace usando Visual Studio Code y devcontainers en el cual se pueda ejecutar su código en local, incluyendo los servicios necesarios para su correcto funcionamiento, debe agregar un archivo Readme.md con las indicaciones para ejecutar su Dag. [cite: 19] [cite\_start](10 puntos) [cite: 20]

4.  [cite\_start]**Conclusiones:** desarrolle un pequeño informe con los principales desafios a los que se enfrentó para resolver esta actividad. [cite: 20] [cite\_start](10 puntos) [cite: 21]

## 4\. Consideraciones

  * [cite\_start]La tarea puede ser desarrollada en grupo de máximo 3 estudiantes. [cite: 21]
  * [cite\_start]El informe debe poseer el nombre de todos los integrantes del equipo de trabajo. [cite: 22]
  * [cite\_start]Este debe ser en formato PDF. [cite: 23]
  * [cite\_start]Debe entregar su código fuente. [cite: 23]
  * [cite\_start]Se debe enviar un informe por equipo en la plataforma AULA. [cite: 24]
  * [cite\_start]Está permitido utilizar Inteligencias Artificiales para el desarrollo de la tarea, pero debe agregar el prompt utilizado en el informe. [cite: 25]

[cite\_start]**Fecha entrega:** 29 Junio 2025 [cite: 26]

```
```