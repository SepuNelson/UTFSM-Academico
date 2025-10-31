Nombre: Nelson Antonio Sepúlveda Valenzuela
Rol: 202004610-7


Comentarios:
Eliminé int argc, char const *argv[] del int main en TreasureFinder, ya que me generaba warnings por no usarlos.

Creé 2 funciones extras en Tablero.c y Tablero.h para identificar si en el tablero se posiciona una Bomba o una Tierra.
void IniciarTablero_2();
void BorrarTablero_2();

Creé una función extra en Bomba.c y Bomba.h para quitar vida en ExplosiónX sin tener un código extremadamente largo.
void BajarVida(int fila, int columna);

Tengo problemas al borrar una Bomba con ExplosiónX por lo que habrá leak de memoria, espero
que en el descuento solo cuenten 1 (o ninguno, mejor, no?) y no hagan nxn Bombas, ya que si me descuentan por muchas
bombas me saldría mucho mejor no haber mandado el archivo con la función ExplotarX.

Instrucciones:
Para compilar los archivos uso:
    make -f Makefile.txt

Para usar valgrind ejecuto:
    valgrind --leak-check=full ./TreasureFinder