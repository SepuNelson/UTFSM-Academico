.PHONY: docker-network docker-gobierno docker-marina docker-submundo docker-cazarrecompensas docker-all clean

# Crear red Docker compartida (si no existe)
docker-network:
	docker network create lab1-final-net || true

# Crear el servicio Gobierno
docker-gobierno: docker-network
	docker build -f gobierno/Dockerfile -t lab1/gobierno .
	docker run -d --rm --name gobierno --network lab1-final-net -p 50051:50051 lab1/gobierno

# Crear el servicio Marina
docker-marina: docker-network
	docker build -f marina/Dockerfile -t lab1/marina .
	docker run -d --rm --name marina --network lab1-final-net -p 50052:50052 lab1/marina

# Crear el servicio Submundo
docker-submundo: docker-network
	docker build -f submundo/Dockerfile -t lab1/submundo .
	docker run -d --rm --name submundo --network lab1-final-net -p 50053:50053 lab1/submundo

# Crear el servicio Cazarrecompensas
docker-cazarrecompensas: docker-network
	docker build -f cazarrecompensas/Dockerfile -t lab1/cazarrecompensas .
	docker run -it --rm --name cazarrecompensas --network lab1-final-net -p 50054:50054 lab1/cazarrecompensas

# Levantar todos los servicios
docker-all: docker-gobierno docker-marina docker-submundo docker-cazarrecompensas

# Limpiar contenedores y eliminar imágenes
clean:
	-docker rm -f gobierno marina submundo cazarrecompensas
	-docker rmi lab1/gobierno lab1/marina lab1/submundo lab1/cazarrecompensas
