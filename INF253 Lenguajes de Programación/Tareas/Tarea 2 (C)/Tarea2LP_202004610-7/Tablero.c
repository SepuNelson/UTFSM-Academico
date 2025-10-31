#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>
#include <time.h>
#include "Tierra.h"
#include "Tablero.h"
#include "Bomba.h"

void*** tablero;
int dimension;
Bomba** tablero_2;

void IniciarTablero(int n){
    /*La función recibe un parametro entero que será el tamaño del tablero, luego a medida que se genera
    el tablero se van colocando datos de tierra con los valores de los atributos al azar.*/
    dimension = n;
    tablero = (void***)malloc(n * sizeof(void**));
    srand(time(NULL));
    for (int i = 0; i < n; i++) {
        tablero[i] = (void**)malloc(n * sizeof(void*));
        for (int j = 0; j < n; j++) {
            tablero[i][j] = (void*)malloc(sizeof(Tierra));
            Tierra* tierra = (Tierra*)tablero[i][j];
            tierra->vida = (rand() % 3) + 1;
            tierra->es_tesoro = (rand() % 20) < 1;
        }
    }
}

void PasarTurno(){
    /*Esta función recorre el tablero auxiliar buscando si hay una bomba, cuando encuentra una bomba
    llama a la función TryExplotar, para explotar la bomba*/
    for (int i = 0; i < dimension; i++) {
        for (int j = 0; j < dimension; j++) {
            if (tablero_2[i][j].contador_turnos == 0) {TryExplotar(i,j);}
        }
    }
}

void ColocarBomba(Bomba* b, int fila, int columna){
    /*Esta función lo que hace es colocar una Bomba en la pos fila - 1 columna - 1, ya que el tablero
    parte del 0,0 y los usuarios suelen usar el 0,0 como 1,1*/
    fila -= 1; columna -= 1;
    Tierra* tierra_temporal = (Tierra*)tablero[fila][columna];
    tablero[fila][columna] = b;
    b->tierra_debajo = tierra_temporal;
    tablero_2[fila][columna].contador_turnos = 0;
}

void MostrarTablero(){
    /*Esta función lo que hace es buscar en el tablero auxiliar si hay una Bomba, si no la encuentra
    transforma la celda del tablero principal en una tipo Tierra y printea la vida, si la vida es 0 y
    es tesoro printea un "*", si la celda en el tablero aux es Bomba, printea una "o"*/
    for (int i = 0; i < dimension; i++) {
        for (int j = 0; j < dimension; j++) {
            Tierra* tierra = (Tierra*)tablero[i][j];
            if (tablero_2[i][j].contador_turnos == -1) {
                if (tierra->vida == 0 && tierra->es_tesoro == 1){printf(" * ");} 
                else{printf("%2d ", tierra->vida);}
            } else {
                if (tierra->vida == 0){printf(" 0 ");}
                else {printf(" o ");}
            }
            if (j != dimension - 1) {printf("|");}
        }
        printf("\n");
    }
}

void MostrarBombas(){
    /*Esta función busca en el tablero aux las Bombas, y printea los datos de la bomba, como cuantos turnos
    falta para que explote, la posición, el tipo de forma de la explosión y la vida de la tierra debajo*/     
    printf("Bombas\n");
    for (int i = 0; i < dimension; i++) {
        for (int j = 0; j < dimension; j++) {
            if (tablero_2[i][j].contador_turnos == 0) {
                Bomba* bomba = (Bomba*)tablero[i][j];
                printf("Turnos para explotar: %d\n", bomba->contador_turnos);
                printf("Coordenadas: %d %d\n", i + 1, j + 1);
                if(bomba->explotar == ExplosionPunto){printf("Forma de la Explosión: Explosión Punto\n");}
                else if(bomba->explotar == ExplosionX){printf("Forma de la Explosión: Explosión X\n");}
                printf("Vida de Tierra Debajo: %d\n", bomba->tierra_debajo->vida);
                printf("\n\n");
            }
        }
    }
}

void VerTesoros(){
    /*Esta función busca en el tablero aux si hay una Bomba o no, para luego revisar si la tierra debajo es tesoro,
    cuando no es Bomba revisa si la tierra en el tablero es tesoro*/
    printf("Tesoros\n");
    for (int i = 0; i < dimension; i++) {
        for (int j = 0; j < dimension; j++) {
            if (tablero_2[i][j].contador_turnos == -1) {
                Tierra* tierra = (Tierra*)tablero[i][j];
                if (tierra->es_tesoro == 1) {printf(" * ");}
                else {printf("%2d ", tierra->vida);}
            } else {
                Bomba* bomba = (Bomba*)tablero[i][j];
                if(bomba->tierra_debajo->es_tesoro == 1){printf(" * ");}
                else {printf("%2d ", bomba->tierra_debajo->vida);}
            }
            if (j != dimension - 1) {printf("|");}
        }
        printf("\n");
    }
}

void BorrarTablero(){
    /*Esta función limpia con Free todo el tablero principal y lo vuelve nulo*/
    for (int i = 0; i < dimension; i++) {
        for (int j = 0; j < dimension; j++) {
            if(tablero_2[i][j].contador_turnos == 0){printf("hay bomba");BorrarBomba(i,j);} 
            else { free(tablero[i][j]);}
            }
        free(tablero[i]);
    }
    free(tablero);
    tablero = NULL;
}

void IniciarTablero_2(){
    /*Esta es la función auxiliar de tablero, la que hace la diferencia si una casilla en el tablero original es
    Bomba o no, ayuda en gran parte al programa*/
    tablero_2 = (Bomba**)malloc(dimension * sizeof(Bomba*));
    for (int i = 0; i < dimension; i++) {
        tablero_2[i] = (Bomba*)malloc(dimension * sizeof(Bomba));
        for (int j = 0; j < dimension; j++) {
            tablero_2[i][j].contador_turnos = -1;
        }
    }
}

void BorrarTablero_2() {
    /*Esta función limpia el tablero auxiliar con los Free en cada celda y lo vuelve nulo*/
    for (int i = 0; i < dimension; i++) {free(tablero_2[i]);}
    free(tablero_2);
    tablero_2 = NULL;
}