       Nombre        |     Rol     | Paralelo 
Nelson Sepúlveda     | 202004610-7 |   201
Gonzalo Bahamondes   | 202073503-4 |   201 

Supuestos:
1. La función convert_base solo sirve para la tarea

Observaciones:
- No implementamos la letra D (overflow).
- Cambiar nombre del archivo de ejemplo en la linea que se abre por primera vez para leer.
- Respecto a la función "convert_base(a,b,c)" : En algunos casos de cambio de representación de números con base mayor a 10
(hacia esta base mayor a 10) se presenta una ambigüedad. Por ejemplo, los números 8463 y 687 en base 10 serán representados 
como 2AF en la base hexagecimal, cuando esta representación solo corresponde a 687, pues la representación de 8463 en 
hexagecimal es 210F.  Pero al realizar la conversión de 2AF y 210F  desde hexagecimal hacia decimal, la conversión es
realizada exitosamente.

Orden General del Código:
Funciones.
Ciclo general para que se pida un número.
Condicional en caso de que sea 1-32 el numero pedido.
	Leer archivo código.txt.
	Manipular variables A, B, C, D (En su mayoría usando las funciones).
	Escribir en archivo resultados.txt.
Condicional en caso que sea 0 el numero pedido.
	Verificar que la suma de "(B + C + D)" sea mayor a "A" para finalizar código.
Condicional en caso que sea menor a 0 o mayor a 32 el numero pedido.
	Printear por pantalla que el número fue mal ingresado.

