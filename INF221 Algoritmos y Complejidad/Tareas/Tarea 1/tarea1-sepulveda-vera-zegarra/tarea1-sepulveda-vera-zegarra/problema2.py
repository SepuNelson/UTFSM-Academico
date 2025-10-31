#!/usr/bin/env python3

import sys

def suma_minimo_gaps_pd(L, l_ladrillos):
    n = len(l_ladrillos)
    lista_pd = [float('inf')] * (n + 1) #  Lista que cada elemento es la suma de los cuadrados de cada nivel
    lista_pd[0] = 0 # partimos sin ladrillos

    for i in range(1, n + 1): # Ciclo
        fila_largo = 0
        for j in range(i, 0, -1):
            fila_largo += l_ladrillos[j - 1]
            if fila_largo <= L:
                gap = L - fila_largo
                lista_pd[i] = min(lista_pd[i], gap * gap + lista_pd[j - 1])
            else:
                break

    return lista_pd[n]

def main():
    for line in sys.stdin:
        L, n = map(int, line.split())
        l_ladrillos = list(map(int, sys.stdin.readline().split()))
        resultado = suma_minimo_gaps_pd(L, l_ladrillos)
        print(resultado)

if __name__ == "__main__":
    main()
