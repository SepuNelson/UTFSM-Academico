##########################################
#                                        #
#  Programe sus funciones aquí           #
#                                        #
##########################################

def disparo(tablero, barcos, fila, columna): #Arreglar la parte en donde cambia el O para que no sea solo el primero
    i = 0
    while i <len(barcos):
        largo_barco = barcos[i][0]
        orientacion_barcos = barcos[i][1]
        fila_barco = barcos[i][2]
        columna_barco = barcos[i][3]
        if orientacion_barcos == 1:
            iv = 0
            while iv < largo_barco:
                if fila_barco + iv == fila and columna_barco == columna :
                    tablero[fila][columna] = "O"
                    return
                iv += 1
        elif orientacion_barcos == 2:
            ih = 0
            while ih < largo_barco:
                if fila_barco == fila and columna_barco + ih == columna:
                    tablero[fila][columna] = "O"
                    return
                ih += 1
        i += 1
    tablero[fila][columna] = " "

def destruidos(tablero, barcos):
    a = 0
    c = 0
    while c < 5:
        largo_barco = barcos[c][0]
        orientacion_barcos = barcos[c][1]
        fila_barco = barcos[c][2]
        columna_barco = barcos[c][3]
        if orientacion_barcos == 1: #Vertical
            cv = 0
            contv = 0
            while cv < largo_barco:
                if tablero[fila_barco + cv][columna_barco] == "O":
                    contv += 1
                    if contv == largo_barco:
                        a += 1
                        cv1 = 0
                        while cv1 < largo_barco:
                            tablero[fila_barco + cv1][columna_barco] = "X"
                            cv1 += 1
                cv += 1
        elif orientacion_barcos == 2: #Horizontal
            ch = 0
            conth = 0
            while ch < largo_barco:
                if tablero[fila_barco][columna_barco + ch] == "O":
                    conth += 1
                    if conth == largo_barco:
                        a += 1
                        ch1 = 0
                        while ch1 < largo_barco:
                            tablero[fila_barco][columna_barco + ch1] = "X"
                            ch1 += 1
                ch += 1
        c += 1
    return a



# OPCIONAL:
# Cambie el valor de esta variable a 1 si desea ver
# la ubicación de los barcos antes de comenzar.
# Esto puede ser útil para probar sus funciones.
mostrar_solucion = 1




##################################################
#                                                #
#  NO MODIFIQUE EL CÓDIGO A PARTIR DE ESTE PUNTO #
#                                                #
##################################################

import random as rd

# Función que muestra el tablero con el formato deseado para la pantalla
def show(tablero):
    print("\n  123456789")
    for i in range(9):
        fila_texto = " "
        for j in tablero[i]:
            fila_texto += str(j)
        print(chr(65+i)+fila_texto)

# Creación del tablero (inicialmente únicamente con "-" en todas las posiciones)
tablero = []
board = []
for i in range(9):
    fila = []
    for j in range(9):
        fila.append("-")
    tablero.append(fila)
    board.append(list(fila))

# Creación de los barcos con orientación y posición aleatorias
barcos = []
length = [2,3,3,4,5]
for i in range(5):
    l = length[i]
    orientation = rd.randint(1,2)
    if orientation == 1:
        flag = True
        while flag:
            row = rd.randint(0,9-l)
            column = rd.randint(0,8)
            empty = True
            for j in range(l):
                empty = empty and board[row+j][column] != "X"
            if empty:
                flag = False
        for j in range(l): board[row+j][column] = "X"
    else:
        flag = True
        while flag:
            row = rd.randint(0,8)
            column = rd.randint(0,9-l)
            if "X" not in board[row][column:column+l]:
                flag = False
        for j in range(l): board[row][column+j] = "X"
    barcos.append([l,orientation,row,column])
print(barcos)
# Se muestra la solución al inicio en caso de que se haya seleccionado esta opción
if mostrar_solucion == 1:
    print("Solución:")
    show(board)
    print("\n\n")

# Ciclo principal del programa donde se reciben los disparos, se validan y se llama a la función disparo()
print("¡Bienvenido a Solitary Battleship!")
destroyed = 0
while destroyed < 5:
    not_valid = True
    while not_valid:
        turn = input("\n¿Que casilla desea disparar? (Ejemplo: A1): ")
        not_valid = False
        if len(turn) != 2:
            not_valid = True
            print("Ingrese una casilla válida por favor.")
        elif not("A" <= turn[0] and turn[0] <= "I"):
            not_valid = True
            print("Ingrese una casilla válida por favor.")
        elif not("1" <= turn[1] and turn[1] <= "9"):
            not_valid = True
            print("Ingrese una casilla válida por favor.")
        else:
            fila = "ABCDEFGHI".index(turn[0])
            columna = int(turn[1])-1
            if tablero[fila][columna] != "-":
                not_valid = True
                print("Ya ha disparado a esta casilla.")
    disparo(tablero, barcos, fila, columna)
    destroyed += destruidos(tablero, barcos)
    show(tablero)
    print("\n"+str(destroyed)+" barcos destruidos.")
    # Fin del juego
    if destroyed == 5:
        print("Felicidades, juego terminado.")
