#!/bin/bash

echo "🚀 Configurando workspace Tarea 3 - ETL Fan Engagement HRL"
echo "=================================================="

# Función para mostrar el progreso
show_progress() {
    echo "📋 $1"
}

# Función para verificar si un comando existe
command_exists() {
    command -v "$1" >/dev/null 2>&1
}

# Función para esperar a que un servicio esté disponible
wait_for_service() {
    local service_name=$1
    local service_url=$2
    local max_attempts=30
    local attempt=1
    
    show_progress "Esperando a que $service_name esté disponible..."
    
    while [ $attempt -le $max_attempts ]; do
        if curl -s -f "$service_url" >/dev/null 2>&1; then
            echo "✅ $service_name está disponible"
            return 0
        fi
        
        echo "⏳ Intento $attempt/$max_attempts - $service_name no está listo aún..."
        sleep 2
        attempt=$((attempt + 1))
    done
    
    echo "❌ $service_name no está disponible después de $max_attempts intentos"
    return 1
}

# Verificar que estamos en el directorio correcto
show_progress "Verificando directorio de trabajo..."
if [ ! -f "requirements.txt" ] || [ ! -f "tarea3.md" ]; then
    echo "❌ Error: No se encontraron archivos del proyecto. Asegúrate de estar en el directorio correcto."
    exit 1
fi
echo "✅ Directorio correcto"

# Verificar Python
show_progress "Verificando Python..."
if command_exists python; then
    echo "✅ Python encontrado: $(python --version)"
else
    echo "❌ Python no encontrado"
    exit 1
fi

# Verificar pip
show_progress "Verificando pip..."
if command_exists pip; then
    echo "✅ pip encontrado"
else
    echo "❌ pip no encontrado"
    exit 1
fi

# Instalar dependencias
show_progress "Instalando dependencias de Python..."
echo "📦 Actualizando pip..."
pip install --upgrade pip

echo "📦 Instalando dependencias desde requirements.txt..."
pip install -r requirements.txt

if [ $? -eq 0 ]; then
    echo "✅ Dependencias instaladas exitosamente!"
else
    echo "❌ Error instalando dependencias"
    exit 1
fi

# Verificar instalaciones
show_progress "Verificando instalaciones..."

echo "🔍 Verificando Apache Beam..."
python -c "import apache_beam; print('✅ Apache Beam instalado')" 2>/dev/null || echo "❌ Apache Beam no instalado"

echo "🔍 Verificando Airflow..."
python -c "import airflow; print('✅ Airflow instalado')" 2>/dev/null || echo "❌ Airflow no instalado"

echo "🔍 Verificando Kafka..."
python -c "import kafka; print('✅ Kafka instalado')" 2>/dev/null || echo "❌ Kafka no instalado"

echo "🔍 Verificando FastAvro..."
python -c "import fastavro; print('✅ FastAvro instalado')" 2>/dev/null || echo "❌ FastAvro no instalado"

# Verificar archivos del proyecto
show_progress "Verificando archivos del proyecto..."

required_files=(
    "data/fan_engagement.jsonl"
    "schemas/fan_engagement_schema.json"
    "src/beam/fan_engagement_etl.py"
    "dags/fan_engagement_dag.py"
    "src/settings.py"
)

for file in "${required_files[@]}"; do
    if [ -f "$file" ]; then
        echo "✅ $file"
    else
        echo "❌ $file (no encontrado)"
    fi
done

# Verificar servicios Docker si están disponibles
show_progress "Verificando servicios Docker..."

# Verificar si Docker está disponible
if command_exists docker; then
    echo "✅ Docker encontrado"
    
    # Verificar si los contenedores están ejecutándose
    if docker ps --format "table {{.Names}}\t{{.Status}}" | grep -q "ws-airflow-webserver"; then
        echo "✅ Contenedor ws-airflow-webserver está ejecutándose"
    else
        echo "⚠️  Contenedor ws-airflow-webserver no está ejecutándose"
    fi
    
    if docker ps --format "table {{.Names}}\t{{.Status}}" | grep -q "ws-kafka"; then
        echo "✅ Contenedor ws-kafka está ejecutándose"
    else
        echo "⚠️  Contenedor ws-kafka no está ejecutándose"
    fi
    
    if docker ps --format "table {{.Names}}\t{{.Status}}" | grep -q "ws-minio"; then
        echo "✅ Contenedor ws-minio está ejecutándose"
    else
        echo "⚠️  Contenedor ws-minio no está ejecutándose"
    fi
else
    echo "⚠️  Docker no está disponible en este entorno"
fi

# Ejecutar pruebas
show_progress "Ejecutando pruebas del pipeline..."
python test_pipeline.py

# Ejecutar verificación completa
show_progress "Ejecutando verificación completa del workspace..."
python verify-setup.py

# Mostrar información de servicios
echo ""
echo "🌐 Información de servicios disponibles:"
echo "========================================"
echo "🌐 Airflow Web UI: http://localhost:8081"
echo "   Usuario: airflow"
echo "   Contraseña: airflow"
echo ""
echo "🌐 Kafka UI: http://localhost:8080"
echo "🌐 MinIO: http://localhost:9000"
echo "   Usuario: minio-root-user"
echo "   Contraseña: minio-root-password"
echo "🌐 MinIO Console: http://localhost:9001"
echo ""
echo "📡 Kafka Broker: localhost:9092"
echo "📡 Kafka Broker (Docker): localhost:29092"

# Mostrar comandos útiles
echo ""
echo "📚 Comandos útiles:"
echo "=================="
echo "  • Ejecutar pipeline manualmente:"
echo "    python src/beam/fan_engagement_etl.py --input data/fan_engagement.jsonl --output output/fan_engagement --schema schemas/fan_engagement_schema.json"
echo ""
echo "  • Ejecutar pruebas:"
echo "    python test_pipeline.py"
echo ""
echo "  • Verificar configuración:"
echo "    python verify-setup.py"
echo ""
echo "  • Ver logs de Airflow:"
echo "    docker logs ws-airflow-webserver"
echo "    docker logs ws-airflow-scheduler"
echo ""
echo "  • Ver logs de Kafka:"
echo "    docker logs ws-kafka"
echo ""
echo "  • Verificar estado de contenedores:"
echo "    docker ps"
echo ""
echo "  • Reiniciar servicios:"
echo "    docker-compose -f .devcontainer/docker-compose.yml restart"

echo ""
echo "🎉 ¡Workspace configurado exitosamente!"
echo "🚀 Puedes comenzar a trabajar en tu proyecto."
echo ""
echo "💡 Consejo: Si los servicios no están disponibles inmediatamente,"
echo "   espera unos minutos a que todos los contenedores se inicialicen completamente." 