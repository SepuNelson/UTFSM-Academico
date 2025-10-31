#ifndef EQUIPO
#define EQUIPO


#include <iostream>
#include <sstream>
#include <string>
using namespace std;

struct Persona{                                    
    string nombre;
    bool capitan;
    int poder; 
};

class Equipo {                              //declaracion de la clase Equipo
        private:
            typedef struct nodo{
                Persona dato;
                struct nodo *siguiente; 
            }nodo;

            nodo *cabeza;          // puntero del nodo cabeza de la  lista
            nodo *cola;            // puntero del nodo cola de la lista
            nodo *curr;            // puntero del nodo actual de la lista
            unsigned int largo;    // largo de la lista
            unsigned int posicion; // posicion actual en la lista
        public:
                    // metodos de la clase
                Equipo(){          // declaracion o creacion de una lista vacia
                    cabeza = NULL;
                    cola = NULL;
                    curr = NULL;
                    largo = 0;
                    posicion = 0;
                }
                ~Equipo(){
                    //destructor
                }    
                void append(Persona item);          // funcion que agrega un integrante del equipo al final de la lista                  
                int length();                       // funcion que retorna el largo de la lista
                void moveToStart();                 // funcion que mueve el puntero curr o cursor hacia el inicio de la lista (justo antes de la cabeza en el caso de verlo como un cursor) 
                int agregar_companero(string name,bool captain, int power); 
                void imprimir_equipo();
                int calcular_poder();
                string printC();
};


#endif
