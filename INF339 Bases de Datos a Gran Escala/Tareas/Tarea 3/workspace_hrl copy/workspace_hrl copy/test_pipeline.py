#!/usr/bin/env python3
"""
Script de prueba para verificar que el pipeline de Beam funciona correctamente
"""

import os
import sys
import json
import tempfile
import shutil

# Agregar el directorio src al path
sys.path.append(os.path.join(os.path.dirname(__file__), 'src'))

def test_pipeline():
    """
    Prueba el pipeline de Beam con datos de ejemplo
    """
    print("🧪 Iniciando prueba del pipeline de Beam...")
    
    # Verificar que los archivos necesarios existen
    required_files = [
        'data/fan_engagement.jsonl',
        'schemas/fan_engagement_schema.json',
        'src/beam/fan_engagement_etl.py'
    ]
    
    for file_path in required_files:
        if not os.path.exists(file_path):
            print(f"❌ Error: No se encuentra el archivo {file_path}")
            return False
        print(f"✅ Archivo encontrado: {file_path}")
    
    # Crear directorio temporal para la salida
    temp_output_dir = tempfile.mkdtemp()
    output_path = os.path.join(temp_output_dir, "test_output")
    
    try:
        # Importar y ejecutar el pipeline
        from src.beam.fan_engagement_etl import run_pipeline
        
        print("🔄 Ejecutando pipeline...")
        run_pipeline(
            input_path='data/fan_engagement.jsonl',
            output_path=output_path,
            schema_path='schemas/fan_engagement_schema.json'
        )
        
        # Verificar que se generaron archivos de salida
        output_files = [f for f in os.listdir(temp_output_dir) if f.endswith('.avro')]
        
        if output_files:
            print(f"✅ Pipeline ejecutado exitosamente!")
            print(f"📁 Archivos generados: {len(output_files)}")
            for file in output_files:
                file_size = os.path.getsize(os.path.join(temp_output_dir, file))
                print(f"   - {file} ({file_size} bytes)")
        else:
            print("❌ No se generaron archivos de salida")
            return False
            
    except Exception as e:
        print(f"❌ Error ejecutando pipeline: {str(e)}")
        return False
    finally:
        # Limpiar archivos temporales
        shutil.rmtree(temp_output_dir)
    
    return True

def test_schema():
    """
    Prueba que el esquema Avro es válido
    """
    print("\n🔍 Verificando esquema Avro...")
    
    try:
        with open('schemas/fan_engagement_schema.json', 'r') as f:
            schema = json.load(f)
        
        # Verificar campos requeridos
        required_fields = [
            'FanID', 'RaceID', 'Timestamp', 'Timestamp_unix',
            'ViewerLocationCountry', 'DeviceType', 'EngagementMetric_secondswatched',
            'PredictionClicked', 'MerchandisingClicked'
        ]
        
        schema_fields = [field['name'] for field in schema['fields']]
        
        for field in required_fields:
            if field not in schema_fields:
                print(f"❌ Campo requerido no encontrado: {field}")
                return False
        
        print("✅ Esquema Avro válido")
        return True
        
    except Exception as e:
        print(f"❌ Error verificando esquema: {str(e)}")
        return False

def test_data_sample():
    """
    Prueba que los datos de entrada son válidos
    """
    print("\n📊 Verificando datos de entrada...")
    
    try:
        with open('data/fan_engagement.jsonl', 'r') as f:
            lines = f.readlines()
        
        if not lines:
            print("❌ Archivo de datos vacío")
            return False
        
        # Verificar primera línea
        first_line = json.loads(lines[0])
        required_fields = [
            'FanID', 'RaceID', 'Timestamp', 'ViewerLocationCountry',
            'DeviceType', 'EngagementMetric_secondswatched',
            'PredictionClicked', 'MerchandisingClicked'
        ]
        
        for field in required_fields:
            if field not in first_line:
                print(f"❌ Campo requerido no encontrado en datos: {field}")
                return False
        
        print(f"✅ Datos de entrada válidos ({len(lines)} registros)")
        print(f"   Ejemplo: {first_line['FanID']} - {first_line['RaceID']} - {first_line['Timestamp']}")
        return True
        
    except Exception as e:
        print(f"❌ Error verificando datos: {str(e)}")
        return False

if __name__ == '__main__':
    print("🚀 Iniciando pruebas del workspace Tarea 3...\n")
    
    tests = [
        test_schema,
        test_data_sample,
        test_pipeline
    ]
    
    passed = 0
    total = len(tests)
    
    for test in tests:
        if test():
            passed += 1
        print()
    
    print(f"📊 Resultados: {passed}/{total} pruebas pasaron")
    
    if passed == total:
        print("🎉 ¡Todas las pruebas pasaron! El workspace está listo para usar.")
    else:
        print("⚠️  Algunas pruebas fallaron. Revisar los errores antes de continuar.")
    
    sys.exit(0 if passed == total else 1) 