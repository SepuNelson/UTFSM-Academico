import apache_beam as beam
import json
import datetime
from datetime import datetime
import argparse
import sys
import os

# Agregar el directorio src al path para importar settings
sys.path.append(os.path.join(os.path.dirname(__file__), '..'))
from settings import MINIO_ENDPOINT, MINIO_ACCESS_KEY, MINIO_SECRET_KEY, MINIO_SECURE, bucket_name

class ParseJsonAndTransform(beam.DoFn):
    """
    ParDo que parsea JSON y transforma el timestamp a Unix timestamp en milisegundos
    """
    def process(self, element):
        try:
            # Parsear el JSON
            data = json.loads(element)
            
            # Convertir el timestamp a Unix timestamp en milisegundos
            timestamp_str = data['Timestamp']
            dt = datetime.strptime(timestamp_str, '%Y-%m-%d %H:%M:%S')
            timestamp_unix = int(dt.timestamp() * 1000)
            
            # Crear el nuevo registro con el campo Timestamp_unix
            transformed_data = {
                'FanID': data['FanID'],
                'RaceID': data['RaceID'],
                'Timestamp': data['Timestamp'],
                'Timestamp_unix': timestamp_unix,
                'ViewerLocationCountry': data['ViewerLocationCountry'],
                'DeviceType': data['DeviceType'],
                'EngagementMetric_secondswatched': data['EngagementMetric_secondswatched'],
                'PredictionClicked': data['PredictionClicked'],
                'MerchandisingClicked': data['MerchandisingClicked']
            }
            
            yield transformed_data
            
        except Exception as e:
            # Log del error pero continuar procesando
            print(f"Error procesando línea: {element[:100]}... Error: {str(e)}")
            pass

def run_pipeline(input_path, output_path, schema_path):
    """
    Ejecuta el pipeline de Beam para convertir JSON a Avro
    """
    # Leer el esquema Avro
    with open(schema_path, 'r') as f:
        avro_schema = json.load(f)
    
    with beam.Pipeline() as pipeline:
        # Leer archivos JSON
        json_data = pipeline | 'Read JSON Files' >> beam.io.ReadFromText(input_path)
        
        # Transformar datos
        transformed_data = json_data | 'Parse and Transform' >> beam.ParDo(ParseJsonAndTransform())
        
        # Escribir en formato Avro
        transformed_data | 'Write to Avro' >> beam.io.WriteToAvro(
            output_path,
            schema=avro_schema,
            file_name_suffix='.avro'
        )

def main():
    parser = argparse.ArgumentParser(description='ETL Pipeline para convertir JSON a Avro')
    parser.add_argument('--input', required=True, help='Ruta de entrada (archivo JSON o patrón glob)')
    parser.add_argument('--output', required=True, help='Ruta de salida para archivos Avro')
    parser.add_argument('--schema', required=True, help='Ruta al archivo de esquema Avro')
    
    args = parser.parse_args()
    
    print(f"Iniciando pipeline ETL...")
    print(f"Input: {args.input}")
    print(f"Output: {args.output}")
    print(f"Schema: {args.schema}")
    
    run_pipeline(args.input, args.output, args.schema)
    
    print("Pipeline completado exitosamente!")

if __name__ == '__main__':
    main() 