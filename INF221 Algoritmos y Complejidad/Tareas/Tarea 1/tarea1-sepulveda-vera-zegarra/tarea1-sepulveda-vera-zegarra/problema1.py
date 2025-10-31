#!/usr/bin/env python3

import sys

def suma_minimo_gaps(L, l_ladrillos):
    def suma_por_niveles(L, l_ladrillos, index):
        if index == len(l_ladrillos):
            return 0

        suma = float('inf')
        fila_largo = 0
        for i in range(index, len(l_ladrillos)):
            fila_largo += l_ladrillos[i]
            if fila_largo <= L:
                gap = L - fila_largo
                gap_sum = gap * gap + suma_por_niveles(L, l_ladrillos, i + 1)
                suma = min(suma, gap_sum)
            else:
                break

        return suma

    return suma_por_niveles(L, l_ladrillos, 0)

def main():
    for line in sys.stdin:
        L, n = map(int, line.split())
        l_ladrillos = list(map(int, sys.stdin.readline().split()))
        resultado = suma_minimo_gaps(L, l_ladrillos)
        print(resultado)

if __name__ == "__main__":
    main()
