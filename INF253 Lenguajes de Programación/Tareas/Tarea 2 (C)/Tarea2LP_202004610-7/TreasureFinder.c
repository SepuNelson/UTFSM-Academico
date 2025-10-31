#include <stdio.h>
#include <stdlib.h>
#include <stdbool.h>
#include <time.h>
#include "Tierra.h"
#include "Tablero.h"
#include "Bomba.h"

int main(){

    int tamanio; int turno = 0; bool flag = true; bool flag_2 = true; bool flag_3 = true;

    printf("\n\n\n¡Bienvenido a TreasureFinder!\n\n");
    printf("Indique el tamaño del tablero a jugar:\n");
    printf("1. 7 x 7\n");
    printf("2. 10 x 10\n");
    printf("3. 12 x 12\n\n");
    printf("Su input: ");
    scanf("%i", &tamanio);
    printf("\n");
    
    while (flag == true){
        if (tamanio == 1 || tamanio == 2 || tamanio == 3){
            flag = false;
            if(tamanio == 1){tamanio = 7;}
            if(tamanio == 2){tamanio = 10;}
            if(tamanio == 3){tamanio = 12;}
            }
        else{
            printf("ERROR Indique una de las opciones señaladas\n");
            printf("Su input: ");
            scanf("%i", &tamanio);
            printf("\n");
        }
    }
    
    printf("Empezando juego... ¡listo!\n\n\n");
    IniciarTablero(tamanio);
    IniciarTablero_2();

    while (flag_2 == true){

        int accion; turno++; 

        while(flag_3 == true){
            printf("Tablero (Turno %d)\n",turno);
            MostrarTablero();
            flag_3 = false;
        }


        printf("\nSeleccione una accion:\n");
        printf("0. Finalizar\n");
        printf("1. Colocar Bomba\n");
        printf("2. Mostrar Bombas\n");
        printf("3. Mostrar Tesoros\n\n");
        printf("Escoja una opcion: ");
        scanf("%i", &accion);
        printf("\n");

        while (accion != 0 && accion != 1 && accion != 2 && accion != 3){
            printf("ERROR Indique una de las opciones señaladas\n");
            printf("Escoja una opcion: ");
            scanf("%i", &accion);
            printf("\n");
        }

        if(accion == 0){
            flag_2 = false;

        }else if(accion == 1){

            int fila, columna;
            Bomba* bomba = (Bomba*)malloc(sizeof(Bomba));

            printf("Indique coordenadas de la Bomba\n");
            printf("Fila: ");
            scanf("%i", &fila);
            printf("Columna: ");
            scanf("%i", &columna);
            printf("\n");
            ColocarBomba(bomba,fila,columna);

            int accion_2;
            printf("Indique forma en que explota la bomba\n");
            printf("1.Punto\n");
            printf("2. X\n");
            printf("Su input: ");
            scanf("%i", &accion_2);
            printf("\n");

            if(accion_2 == 1){
                Bomba* bomba = (Bomba*)tablero[fila - 1][columna - 1];
                bomba->contador_turnos = 1;
                bomba->explotar = ExplosionPunto;

            }
            else if(accion_2 == 2){
                printf("\n\nPor Favor leer README\n\n");
                Bomba* bomba = (Bomba*)tablero[fila - 1][columna - 1];
                bomba->contador_turnos = 3;
                bomba->explotar = ExplosionX;
            }

            PasarTurno();
            flag_3 = true;

        }
        if(accion == 2){MostrarBombas();}
        if(accion == 3){VerTesoros();}

    }
    BorrarTablero();
    BorrarTablero_2();
    return 0;
}