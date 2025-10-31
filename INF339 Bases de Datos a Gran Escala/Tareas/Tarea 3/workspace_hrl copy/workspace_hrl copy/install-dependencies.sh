#!/bin/bash

echo "🔧 Iniciando instalación de dependencias..."
echo "📁 Directorio actual: $(pwd)"
echo "🐍 Versión de Python: $(python --version)"

# Actualizar pip
echo "⬆️  Actualizando pip..."
pip install --upgrade pip

# Verificar que requirements.txt existe
if [ -f "requirements.txt" ]; then
    echo "📦 Instalando dependencias desde requirements.txt..."
    echo "📋 Contenido de requirements.txt:"
    cat requirements.txt
    echo ""
    
    # Instalar dependencias
    pip install -r requirements.txt
    
    if [ $? -eq 0 ]; then
        echo "✅ Dependencias instaladas exitosamente!"
    else
        echo "❌ Error instalando dependencias"
        exit 1
    fi
else
    echo "⚠️  No se encontró requirements.txt"
    exit 1
fi

# Verificar instalación de Apache Beam
echo "🔍 Verificando instalación de Apache Beam..."
python -c "import apache_beam; print('✅ Apache Beam instalado correctamente')" 2>/dev/null || echo "❌ Apache Beam no está instalado"

# Verificar instalación de Airflow
echo "🔍 Verificando instalación de Airflow..."
python -c "import airflow; print('✅ Airflow instalado correctamente')" 2>/dev/null || echo "❌ Airflow no está instalado"

# Verificar instalación de Kafka
echo "🔍 Verificando instalación de Kafka..."
python -c "import kafka; print('✅ Kafka instalado correctamente')" 2>/dev/null || echo "❌ Kafka no está instalado"

echo "🎉 Instalación de dependencias completada!"
echo "🚀 El workspace está listo para usar." 