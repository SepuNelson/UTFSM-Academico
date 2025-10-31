import threading
import datetime
import random

# Función para el ingreso a la isla
def ingreso_isla(num_pirata, isla, numero_de_isla, hora_fila):
    try:
        # Pedir Lock
        with locks_islas[isla]:
            # Decrementar Semáforos
            puentes[isla].acquire()
            islas_semaforos[isla].acquire()

            # Escribir en el archivo de la isla
            hora_acceso = datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3]
            registros_islas[isla].write(f"Pirata {num_pirata}, {hora_fila}, {hora_acceso}, {numero_de_isla}\n")
        
        # Tiempo en la isla
        duracion = islas[isla]['duracion_busqueda']
        threading.Event().wait(duracion)
    
    except Exception as e:
        # Mostrar algún Error
        print(f"Error: {e}")
    
    finally:
        # Incrementar Semáforos
        islas_semaforos[isla].release()
        puentes[isla].release()

# Función que simula la llegada de piratas al BarOraculo
def pirata_llega(num_pirata):

    # Hora llegada
    hora = datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3]

    # Elige dos islas aleatorias para cada pirata
    isla1, isla2 = random.sample(list(islas.keys()), 2)

    # Trabajar en la Isla 1
    hora_isla_1 = datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3]
    ingreso_isla(num_pirata, isla1, 1, hora_isla_1)

    # Trabajar en la Isla 2
    hora_isla_2 = datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3]
    ingreso_isla(num_pirata, isla2, 2, hora_isla_2)

    # Registro en BarOraculo.txt
    with lock_bar:
        registro_bar.write(f"Pirata {num_pirata}, {hora}, {isla1}, {hora_isla_1}, {isla2}, {hora_isla_2}\n")
 
    # Registro de Salida.txt
    hora_salida = datetime.datetime.now().strftime("%H:%M:%S.%f")[:-3]
    with lock_bar:
        archivo_salida.write(f"Pirata {num_pirata}, {hora_salida}\n")

# Definición de las características de cada isla
islas = {
    'Isla_de_las_Sombras': {'capacidad_puente': 20, 'duracion_busqueda': 9, 'capacidad_isla': 10},
    'Isla_del_Dragón_Rojo': {'capacidad_puente': 8, 'duracion_busqueda': 5, 'capacidad_isla': 2},
    'Isla_de_las_Estrellas': {'capacidad_puente': 15, 'duracion_busqueda': 7, 'capacidad_isla': 5},
    'Isla_del_Bosque_Encantado': {'capacidad_puente': 6, 'duracion_busqueda': 4, 'capacidad_isla': 3},
    'Isla_de_los_Susurros': {'capacidad_puente': 6, 'duracion_busqueda': 1, 'capacidad_isla': 5},
    'Isla_del_Trueno': {'capacidad_puente': 9, 'duracion_busqueda': 4, 'capacidad_isla': 4},
    'Isla_del_Tesoro_Oculto': {'capacidad_puente': 7, 'duracion_busqueda': 5, 'capacidad_isla': 2}
}

# Semaforos y locks para sincronización
lock_bar = threading.Lock()
locks_islas = {isla: threading.Lock() for isla in islas}

# Semáforos para controlar acceso a puentes e islas
puentes = {isla: threading.Semaphore(islas[isla]['capacidad_puente']) for isla in islas}
islas_semaforos = {isla: threading.Semaphore(islas[isla]['capacidad_isla']) for isla in islas}

# Abre archivos de registro
registro_bar = open("BarOraculo.txt", "w")
registros_islas = {isla: open(f"{isla}.txt", "w") for isla in islas}
archivo_salida = open("Salida.txt", "w")

# Crear y lanzar hilos para los piratas
threads = []
for i in range(500):
    thread = threading.Thread(target=pirata_llega, args=(i + 1,))
    threads.append(thread)
    thread.start()

# Esperar a que todos los hilos terminen
i = 0
while i < len(threads):    
    thread.join()
    i += 1
