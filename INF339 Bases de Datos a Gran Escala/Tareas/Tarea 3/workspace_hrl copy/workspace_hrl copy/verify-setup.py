#!/usr/bin/env python3
"""
Script de verificación para el workspace de ETL Fan Engagement
Verifica que todos los componentes estén funcionando correctamente
"""

import os
import sys
import json
import subprocess
import time
from datetime import datetime

def print_status(message, status="INFO"):
    """Imprime un mensaje con formato de estado"""
    timestamp = datetime.now().strftime("%H:%M:%S")
    if status == "SUCCESS":
        print(f"[{timestamp}] ✅ {message}")
    elif status == "ERROR":
        print(f"[{timestamp}] ❌ {message}")
    elif status == "WARNING":
        print(f"[{timestamp}] ⚠️  {message}")
    else:
        print(f"[{timestamp}] ℹ️  {message}")

def check_file_exists(file_path, description):
    """Verifica que un archivo exista"""
    if os.path.exists(file_path):
        print_status(f"{description}: {file_path}", "SUCCESS")
        return True
    else:
        print_status(f"{description}: {file_path} (no encontrado)", "ERROR")
        return False

def check_python_import(module_name, description):
    """Verifica que un módulo de Python se pueda importar"""
    try:
        __import__(module_name)
        print_status(f"{description}: {module_name}", "SUCCESS")
        return True
    except ImportError:
        print_status(f"{description}: {module_name} (no disponible)", "ERROR")
        return False

def check_docker_container(container_name, description):
    """Verifica que un contenedor Docker esté ejecutándose"""
    try:
        result = subprocess.run(
            ["docker", "ps", "--filter", f"name={container_name}", "--format", "{{.Status}}"],
            capture_output=True, text=True, timeout=10
        )
        if result.stdout.strip():
            print_status(f"{description}: {container_name}", "SUCCESS")
            return True
        else:
            print_status(f"{description}: {container_name} (no ejecutándose)", "WARNING")
            return False
    except (subprocess.TimeoutExpired, FileNotFoundError):
        print_status(f"{description}: Docker no disponible", "WARNING")
        return False

def check_service_health(url, description, timeout=5):
    """Verifica que un servicio web esté respondiendo"""
    try:
        import requests
        response = requests.get(url, timeout=timeout)
        if response.status_code == 200:
            print_status(f"{description}: {url}", "SUCCESS")
            return True
        else:
            print_status(f"{description}: {url} (status: {response.status_code})", "WARNING")
            return False
    except Exception as e:
        print_status(f"{description}: {url} (error: {str(e)})", "WARNING")
        return False

def main():
    print("🔍 Verificando configuración del workspace ETL Fan Engagement")
    print("=" * 60)
    
    # Contadores para el resumen
    total_checks = 0
    successful_checks = 0
    
    # Verificar archivos del proyecto
    print("\n📁 Verificando archivos del proyecto:")
    print("-" * 40)
    
    required_files = [
        ("requirements.txt", "Archivo de dependencias"),
        ("data/fan_engagement.jsonl", "Datos de entrada"),
        ("schemas/fan_engagement_schema.json", "Esquema Avro"),
        ("src/beam/fan_engagement_etl.py", "Pipeline de Beam"),
        ("dags/fan_engagement_dag.py", "DAG de Airflow"),
        ("src/settings.py", "Configuraciones"),
        ("test_pipeline.py", "Script de pruebas"),
        (".devcontainer/devcontainer.json", "Configuración DevContainer"),
        (".devcontainer/docker-compose.yml", "Compose de servicios")
    ]
    
    for file_path, description in required_files:
        total_checks += 1
        if check_file_exists(file_path, description):
            successful_checks += 1
    
    # Verificar módulos de Python
    print("\n🐍 Verificando módulos de Python:")
    print("-" * 40)
    
    required_modules = [
        ("apache_beam", "Apache Beam"),
        ("airflow", "Apache Airflow"),
        ("kafka", "Kafka Python"),
        ("fastavro", "FastAvro"),
        ("pandas", "Pandas"),
        ("pyarrow", "PyArrow"),
        ("minio", "MinIO Client"),
        ("boto3", "Boto3"),
        ("psycopg", "PostgreSQL")
    ]
    
    for module_name, description in required_modules:
        total_checks += 1
        if check_python_import(module_name, description):
            successful_checks += 1
    
    # Verificar contenedores Docker
    print("\n🐳 Verificando contenedores Docker:")
    print("-" * 40)
    
    required_containers = [
        ("ws-airflow-webserver", "Airflow WebServer"),
        ("ws-airflow-scheduler", "Airflow Scheduler"),
        ("ws-airflow-postgres", "PostgreSQL"),
        ("ws-kafka", "Kafka Broker"),
        ("ws-zookeeper", "Zookeeper"),
        ("ws-kafka-ui", "Kafka UI"),
        ("ws-minio", "MinIO")
    ]
    
    for container_name, description in required_containers:
        total_checks += 1
        if check_docker_container(container_name, description):
            successful_checks += 1
    
    # Verificar servicios web
    print("\n🌐 Verificando servicios web:")
    print("-" * 40)
    
    required_services = [
        ("http://localhost:8081", "Airflow Web UI"),
        ("http://localhost:8080", "Kafka UI"),
        ("http://localhost:9000", "MinIO API"),
        ("http://localhost:9001", "MinIO Console")
    ]
    
    for url, description in required_services:
        total_checks += 1
        if check_service_health(url, description):
            successful_checks += 1
    
    # Verificar configuración de Kafka
    print("\n📡 Verificando configuración de Kafka:")
    print("-" * 40)
    
    try:
        sys.path.append('src')
        from settings import KAFKA_BROKERS, KAFKA_TOPIC
        
        print_status(f"Kafka Brokers: {KAFKA_BROKERS}", "SUCCESS")
        print_status(f"Kafka Topic: {KAFKA_TOPIC}", "SUCCESS")
        successful_checks += 2
        total_checks += 2
    except Exception as e:
        print_status(f"Error cargando configuración de Kafka: {str(e)}", "ERROR")
        total_checks += 2
    
    # Resumen final
    print("\n" + "=" * 60)
    print(f"📊 RESUMEN DE VERIFICACIÓN")
    print("=" * 60)
    print(f"✅ Verificaciones exitosas: {successful_checks}/{total_checks}")
    print(f"📈 Porcentaje de éxito: {(successful_checks/total_checks)*100:.1f}%")
    
    if successful_checks == total_checks:
        print_status("🎉 ¡Todo está configurado correctamente!", "SUCCESS")
        print("\n🚀 Puedes comenzar a trabajar:")
        print("   • Airflow UI: http://localhost:8081 (airflow/airflow)")
        print("   • Kafka UI: http://localhost:8080")
        print("   • MinIO: http://localhost:9001 (minio-root-user/minio-root-password)")
    else:
        print_status("⚠️  Algunos componentes necesitan atención", "WARNING")
        print("\n🔧 Comandos útiles para solucionar problemas:")
        print("   • Ver contenedores: docker ps")
        print("   • Ver logs: docker logs <nombre-contenedor>")
        print("   • Reiniciar servicios: docker-compose -f .devcontainer/docker-compose.yml restart")
    
    print("\n" + "=" * 60)

if __name__ == "__main__":
    main() 