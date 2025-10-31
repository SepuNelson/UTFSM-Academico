    Nombres      |     Rol
Javiera Cortés   | 201904618-7
Nelson Sepúlveda | 202004610-5

Instrucciones:
    - Para ejecutar el Makefile se debe escribir "make" en la consola, con la dirección del archivo Makefile correcta, esto ejecutará los Docker iniciando los archivos
    cazarrecompensas.go, gobierno.go, marina.go y submundo.go
    - El programa funciona de la siguiente manera, una vez iniciados los archivos .go el gobierno lee el csv de los piratas, el cazarrecompensas pide la lista,
    busca por piradas en estado de BUSQUEDA y finaliza una vez todos los piratas de la lista están en estado de CAPTURADOS
    - Para ejecutar correctamente esta tarea se debe utilizar un archivo .csv con nombre "Piratas.csv" de otra forma se tendrá que editar el cpodigo para que lea el archivo
    csv con otro nombre 