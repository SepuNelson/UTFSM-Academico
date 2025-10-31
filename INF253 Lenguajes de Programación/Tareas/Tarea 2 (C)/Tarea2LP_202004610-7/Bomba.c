#include "Tierra.h"
#include "Tablero.h"
#include "Bomba.h"
#include <stddef.h>
#include <stdlib.h>
#include <stdio.h>

void TryExplotar(int fila, int columna){
    /*Esta función lo que hace es descontar el contador de la bomba entregada en 1 y diferencia el tipo
    de explosión para llamar a la función respectiva*/
    Bomba* bomba = (Bomba*)tablero[fila][columna];
    bomba->contador_turnos -= 1;
    if(bomba->contador_turnos == 0){
        if(bomba->explotar == ExplosionPunto){ExplosionPunto(fila,columna);}
        else if(bomba->explotar == ExplosionX){ExplosionX(fila,columna);}
    }
}

void BorrarBomba(int fila, int columna){
    /*Cuando el contador de una Bomba llega a 0 se llama a esta función y lo que hace es borrar la bomba
    que es reemplazada por la tierra debajo*/
    Bomba* bomba = (Bomba*)tablero[fila][columna];
    Tierra* tierra = (Tierra*)bomba->tierra_debajo;
    free(bomba);
    tablero[fila][columna] = tierra;
    tablero_2[fila][columna].contador_turnos = -1;
}

void ExplosionPunto(int fila, int columna){
    /*Se usa para explotar la bomba en la misma celda cuando se llame a la función*/
    Bomba* bomba = (Bomba*)tablero[fila][columna];
    bomba->tierra_debajo->vida -= 3;
    if (bomba->tierra_debajo->vida < 0){bomba->tierra_debajo->vida = 0;}
    if(bomba->tierra_debajo->vida == 0){BorrarBomba(fila,columna);}
}

void ExplosionX(int fila, int columna){
    /*Se usa para explotar la bomba en forma de X celda cuando se llame a la función*/
    Bomba* bomba = (Bomba*)tablero[fila][columna];
    bomba->tierra_debajo->vida -= 1;
    if (bomba->tierra_debajo->vida < 0){bomba->tierra_debajo->vida = 0;}

    fila -=1; columna -=1; //SUPERIOR IZQUIERDO DE LA X
    if ((fila < 0) && (columna < 0)){ //ARRIBA E IZQUIERDA
        fila = dimension - 1; columna = dimension - 1;
        BajarVida(fila,columna);

    } else if ((fila < 0) && ((0 < columna) < (dimension - 1))){ //ARRIBA
        fila = dimension - 1;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && (columna < 0)){ //IZQUIERDO
        columna = dimension - 1;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && ((0 < columna) < (dimension - 1))){ //EOC
        BajarVida(fila,columna);

    }
    fila +=1; columna +=1;

    fila -=1; columna +=1; //SUPERIOR DERECHO DE LA X
    if ((fila < 0) && (columna > (dimension - 1))){ //ARRIBA Y DERECHA
        fila = dimension - 1; columna = 0;
        BajarVida(fila,columna);

    } else if ((fila < 0) && ((0 < columna) < (dimension - 1))){ //ARRIBA
        fila = dimension - 1;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && (columna > (dimension - 1))){ //DERECHO
        columna = 0;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && ((0 < columna) < (dimension - 1))){ //EOC
        BajarVida(fila,columna);

    }
    fila +=1; columna -=1;

    fila +=1; columna -=1; //INFERIOR IZQUIERDO DE LA X
    if ((fila > (dimension - 1)) && (columna < 0)){ //ABAJO E IZQUIERDA
        fila = 0; columna = dimension - 1;
        BajarVida(fila,columna);

    } else if ((fila > (dimension - 1)) && ((0 < columna) < (dimension - 1))){ //ABAJO
        fila = 0;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && (columna < 0)){ //IZQUIERDO
        columna = dimension - 1;
        BajarVida(fila,columna);

    } else if (((0 < fila) < (dimension - 1)) && ((0 < columna) < (dimension - 1))){ //EOC
        BajarVida(fila,columna);

    }
    fila -=1; columna +=1;

    fila +=1; columna +=1; //INFERIOR DERECHO DE LA X
    if ((fila > dimension) && (columna > dimension)){ //ABAJO Y DERECHA
        fila = 0; columna = 0;
        BajarVida(fila,columna);

    } else if ((fila > dimension) && ((0 < columna) < dimension)){ //ABAJO
        fila = 0;
        BajarVida(fila,columna);

    } else if (((0 < fila) < dimension) && (columna > dimension)){ //DERECHO
        columna = 0;
        BajarVida(fila,columna);

    } else if (((0 < fila) < dimension) && ((0 < columna) < dimension)){ //EOC
        BajarVida(fila,columna);

    }

}

void BajarVida(int fila, int columna){
    /*Esta función lo que hace es disminuir en 1 la vida de la tierra y verificar que la vida
    es inferior a 0, en ese caso la iguala a 0*/
    if(tablero_2[fila][columna].contador_turnos == 0){
        Bomba* bomba = (Bomba*)tablero[fila][columna];
        bomba->tierra_debajo->vida -= 1;
        if (bomba->tierra_debajo->vida < 0){bomba->tierra_debajo->vida = 0;}
    } else {
        Tierra* tierra = (Tierra*)tablero[fila][columna];
        tierra->vida -= 1;
        if (tierra->vida < 0){tierra->vida = 0;}
    }
}
