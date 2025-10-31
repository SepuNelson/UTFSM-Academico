#ifndef TAREA_PARTE_ARBOL
#define TAREA_PARTE_ARBOL


#include <iostream>
#include <sstream>
#include <fstream>
#include <string.h>
#include <string>
#include <stdio.h>
#include <math.h>
#include "Equipo.hpp"



using namespace std;


typedef Equipo tElem;

struct Nodo{
    Equipo* equipo;
    bool visitado;
    int h;
    Nodo* izq;
    Nodo* der;
};

class Torneo{
    private:
    Nodo* raiz;
    int num;

    public:
    Torneo();
    ~Torneo();

    void crear_torneo(Equipo* equipos, int n);
    void insertNodo(Equipo* equipo, Nodo* raiz, int alturaH, int n, int* cont);
    void avanzar_ronda();
    void visitarNodoBatallar(Nodo* raiz, int alturaH,int n);
    void imprimir_bracket();
    void imprimirAltura(Nodo* raiz, int alturaH, int n, int cont);
    void visitarNodoBracket(Nodo* raiz, int alturaH, int n, int* cont);
    
};


#endif