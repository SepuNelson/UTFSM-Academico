# Makefile para ejecutar los programas problema1.py y problema2.py

# Definición de los objetivos
PROBLEMA1 = problema1.py
PROBLEMA2 = problema2.py

# Definición de los comandos
RUN_PROBLEMA1 = ./$(PROBLEMA1) < input-1.dat
RUN_PROBLEMA2 = ./$(PROBLEMA2) < input-1.dat

# Regla para ejecutar el primer problema
problema1:
	chmod +x $(PROBLEMA1)
	$(RUN_PROBLEMA1)

# Regla para ejecutar el segundo problema
problema2:
	chmod +x $(PROBLEMA2)
	$(RUN_PROBLEMA2)

# Regla por defecto
all: problema1 problema2
