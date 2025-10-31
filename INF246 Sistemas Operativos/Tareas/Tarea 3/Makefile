# Nombre del archivo JAR de salida
JAR_NAME = Main.jar

# Directorios de origen y compilación
SRC_DIR = main
BUILD_DIR = build

# Compilador Java y banderas
JAVAC = javac
JAVAC_FLAGS = -sourcepath $(SRC_DIR) -d $(BUILD_DIR)/classes

# Archivo de manifiesto para el JAR
MANIFEST = manifest.mf

# Objetivo predeterminado: Ejecuta el objetivo 'jar'
run: jar
	java -jar $(BUILD_DIR)/$(JAR_NAME)

# Objetivo para crear el archivo JAR
jar: classes
	jar cfm $(BUILD_DIR)/$(JAR_NAME) $(MANIFEST) -C $(BUILD_DIR)/classes .

# Objetivo para compilar las clases
classes: dir
	$(JAVAC) $(JAVAC_FLAGS) $(SRC_DIR)/*.java

# Objetivo para crear directorios si no existen
dir:
	if [ ! -d $(BUILD_DIR)/classes ]; then mkdir -p $(BUILD_DIR)/classes; fi

# Objetivo para limpiar los archivos generados
clean:
	rm -rf $(BUILD_DIR)
