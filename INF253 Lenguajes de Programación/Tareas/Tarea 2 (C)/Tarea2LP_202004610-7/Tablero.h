#ifndef TABLERO_H
#define TABLERO_H
#include "Bomba.h"

extern void*** tablero;
extern int dimension;
extern Bomba** tablero_2;

void IniciarTablero(int n);
void PasarTurno();
void ColocarBomba(Bomba* b, int fila, int columna);
void MostrarTablero();
void MostrarBombas();
void VerTesoros();
void BorrarTablero();
void IniciarTablero_2();
void BorrarTablero_2();

#endif