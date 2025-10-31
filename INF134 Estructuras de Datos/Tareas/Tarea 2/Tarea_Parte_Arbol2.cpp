#include <iostream>
#include <sstream>
#include <fstream>
#include <string.h>
#include <string>
#include <stdio.h>
#include <math.h>
#include "Tarea_Parte_Arbol.hpp"


using namespace std;


bool batallar(Equipo &a, Equipo &b){        //funcion que compara poder de dos equipos y retorna un ganador 
    bool ret;                               //mediante un bool, true para equipo izq y false en caso contrario
    int podera = a.calcular_poder();
    int poderb = b.calcular_poder();
    if(podera >= poderb){ret = true;}
    if(podera < poderb){ret = false;}
    return ret;
}

Torneo::Torneo(){                           //constructor de clase Torneo
    raiz = NULL;
    num  = 0;
}
Torneo::~Torneo(){}                         //destructor de clase Torneo




void Torneo::crear_torneo(Equipo* equipos, int n){          //funcion que crea un arbol binario completo que representa
    int contador = 0;                                       //a un torneo con sus brackets correspondientes
    int alturaH = 1;                                        //y que ubica en un principio a todos los equipos
    num = n;                                                //en las hojas, ordenados por orden de entrada
    raiz = new Nodo();
    raiz->equipo = NULL;
    raiz->visitado = true;
    raiz->h = 0;
    raiz->izq = NULL;
    raiz->der = NULL;
    insertNodo(equipos, raiz, alturaH, n, &contador);       //utilizacion de funcion recursiva
}


void Torneo::insertNodo(Equipo* equipo, Nodo* raiz, int alturaH, int n, int* cont){ //funcion recursiva que inserta nodos y a su vez
                                                                                    //equipos en las hojas del mismo, dejando en un principio
    if ((raiz->izq == NULL || !raiz->izq->visitado)  && alturaH != log2(n)){        //los nodos internos y la raiz como vacios
        raiz->izq = new Nodo();
        raiz->izq->equipo = NULL;
        raiz->izq->visitado = true;
        raiz->izq->h = alturaH;
        raiz->izq->izq = NULL;
        raiz->izq->der = NULL;
        insertNodo(equipo, raiz->izq, alturaH + 1, n, cont);
    }
    if(( raiz->izq == NULL || !raiz->izq->visitado) && alturaH == log2(n)){ 
        raiz->izq = new Nodo();
        raiz->izq->equipo = &equipo[*cont];
        raiz->izq->visitado = true;
        raiz->izq->h = alturaH;
        raiz->izq->izq = NULL;
        raiz->izq->der = NULL;
        *cont= *cont + 1;
    }
    if(( raiz->der == NULL || !raiz->der->visitado) && alturaH == log2(n)){ 
        raiz->der = new Nodo();
        raiz->der->equipo = &equipo[*cont];
        raiz->der->visitado = true;
        raiz->der->h = alturaH;
        raiz->der->izq = NULL;
        raiz->der->der = NULL;
        *cont = *cont + 1;
    }
    if((raiz->der == NULL || !raiz->der->visitado) && alturaH != log2(n)){
        raiz->der = new Nodo();
        raiz->der->equipo = NULL;
        raiz->der->visitado = true;
        raiz->der->h = alturaH;
        raiz->der->izq = NULL;
        raiz->der->der = NULL;
        insertNodo(equipo, raiz->der, alturaH + 1, n, cont);
    }
    raiz->der->visitado = false;
    raiz->izq->visitado = false;
}

void Torneo::avanzar_ronda(){                       //funcion que modifica los nodos internos del arbol, de modo
    raiz->visitado = true;                          //que simula una batalla dentro del torneo, llenando el arbol con equipos ganadores
    int alturaH = 1;                                //desde las hojas hacia la raiz segun la ronda correspondiente
    if (raiz->izq->equipo !=NULL && raiz->der->equipo != NULL){
        if (batallar(*raiz->izq->equipo, *raiz->der->equipo)) {raiz->equipo = raiz->izq->equipo;}
        else {raiz->equipo = raiz->der->equipo;}
    }
    visitarNodoBatallar(raiz, alturaH, num);        //utilizacion de funcion recursiva
}


void Torneo::visitarNodoBatallar(Nodo* raiz, int alturaH, int n){   //funcion recursiva que recorre el arbol hasta la altura
    if ((!raiz->izq->visitado)  && alturaH != log2(n)){             //de los ultimos nodos vacios y que simula batalla entre nodo izq y der 
        raiz->izq->visitado = true;                                 //del nodo correspondiente
        visitarNodoBatallar(raiz->izq, alturaH + 1, n);    
    }
    if((!raiz->izq->visitado) && alturaH == log2(n)){
        raiz->izq->visitado = true;
        raiz->der->visitado = true;
        if (batallar(*raiz->izq->equipo, *raiz->der->equipo)) {raiz->equipo = raiz->izq->equipo;}
        else {raiz->equipo = raiz->der->equipo;}
    }
    if((!raiz->der->visitado) && alturaH != log2(n)){
        raiz->der->visitado = true;
        visitarNodoBatallar(raiz->der, alturaH + 1, n);
    }
    raiz->der->visitado = false;
    raiz->izq->visitado = false;
}



void Torneo::imprimir_bracket(){                                //funcion que recorre el arbol e imprime el estado del torneo
    raiz->visitado = true;                                      //en el instante en que es llamada
    int alturaH = 1;                            
    if (raiz->equipo == NULL) cout << " -- " << endl;
    else cout << raiz->equipo->printC() << endl;
    if (raiz->izq->equipo == NULL) cout << " -- vs -- " << endl;
    else cout << raiz->izq->equipo->printC() << " " << raiz->izq->equipo->calcular_poder() << " vs " << raiz->der->equipo->printC() << " " << raiz->der->equipo->calcular_poder() << endl;
    for (int i = 1; i <= log2(num); i++)
    {
        visitarNodoBracket(raiz, alturaH, num, &i);             //utilizacion de funcion recursiva que pasa por distintas alturas del arbol
        cout << endl;
    }
    
    
}

void Torneo::visitarNodoBracket(Nodo* raiz, int alturaH, int n, int* cont){         //funcion recursiva que visita todos los nodos del arbol
            if ((!raiz->izq->visitado)  && alturaH != log2(n)){                     //e imprime los equipos correspondientes en los nodos no vacios
                                                                                    //y en caso de ser nodos vacios imprime "--"
                raiz->izq->visitado = true;

                if (raiz->izq->equipo == NULL && raiz->h == *cont) {
                    cout << "| -- vs -- |";
                }
                else if (raiz->izq->equipo != NULL && raiz->h == *cont) 
                    cout << raiz->izq->equipo->printC() << " " << raiz->izq->equipo->calcular_poder() << " vs " << raiz->der->equipo->printC() << " " << raiz->der->equipo->calcular_poder() << " " << endl;

                visitarNodoBracket(raiz->izq, alturaH + 1, n, cont);
        }
            if((!raiz->izq->visitado) && alturaH == log2(n)){
                raiz->izq->visitado = true;
                raiz->der->visitado = true;

                if ( raiz->h == *cont){ 
                    cout << raiz->izq->equipo->printC() << " " << raiz->izq->equipo->calcular_poder() << " vs " << raiz->der->equipo->printC() << " " << raiz->der->equipo->calcular_poder() << " | ";
                }
        }
            if((!raiz->der->visitado) && alturaH != log2(n)){

                raiz->der->visitado = true;

                if (raiz->izq->equipo == NULL && raiz->h == *cont){
                    cout << "| -- vs -- |";

                    }
                else if (raiz->izq->equipo != NULL && raiz->h == *cont) 
                    cout << " " << raiz->izq->equipo->printC() << " " << raiz->izq->equipo->calcular_poder() << " vs " << raiz->der->equipo->printC() << " " << raiz->der->equipo->calcular_poder() << " |" << endl;

                visitarNodoBracket(raiz->der, alturaH + 1, n, cont);
        }
            raiz->der->visitado = false;
            raiz->izq->visitado = false;
    }
