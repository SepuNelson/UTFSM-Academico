

#include <iostream>
#include <sstream>
#include <string>
#include "Equipo.hpp"
using namespace std;


int Equipo::length(){                   // retorna el largo de la lista en el instante en que se llama la funcion  
    return largo;
}

void Equipo::moveToStart(){         // mueve el 'cursor' hacia el inicio de la lista      
    curr = cabeza;
    posicion = 0;
}

void Equipo::append(Persona item){      //agrega un integrante nuevo al equipo, que sera agregado en la ultima posicion 
    nodo *aux = new nodo;               //se crea el nodo donde se almacena el nuevo integrante y luego es introducido en la lista
    aux->dato = item;       
    aux->siguiente = NULL;
    if (largo == 0){                    //esta condicion es verdadera cuando se inserta el primer integrante del equipo          
        cabeza = aux;
        cola = aux;
        curr = cola;
        largo++;
    }
    else{                               //esta condicion es verdadera cuando hay 1 o mas integrantes en el equipo
        cola-> siguiente = aux;
        cola = aux; 
        curr = cola;
        largo++;
        posicion++;
    }
}

int Equipo::agregar_companero(string name,bool captain, int power){     //agrega integrantes al equipo utilizando la funcion append
    Persona* nuInt = new Persona;
    nuInt->nombre = name;
    nuInt->poder = power;
    nuInt->capitan = captain;
    append(*nuInt);
    return (length()-1);
}

int Equipo::calcular_poder(){               //calcula el poder total del equipo
    int sumPo = 0;
    moveToStart();
    while (posicion < largo){
        sumPo += curr->dato.poder;
        posicion++;
        curr = curr->siguiente;
    }
    return sumPo; 
}

void Equipo::imprimir_equipo(){             //muestra por pantalla los integrantes del equipo, su poder, y si es el capitan o no del equipo
    moveToStart();
    cout << "Equipo:" << endl;
    for (posicion = 0; posicion < largo; posicion++){
        cout << "Persona " << posicion << ": "<< curr->dato.nombre << " " << boolalpha << curr->dato.capitan << " " << curr->dato.poder << endl;
        curr = curr->siguiente;
    }
}

string Equipo::printC(){                    //funcion que busca al capitan del equipo e imprime su nombre
    moveToStart();                          //para representar al equipo en el arbol
    string retorno;
    while(posicion < largo ){
        if (curr->dato.capitan){
            retorno= curr->dato.nombre;
            break;
        } 
        if (!curr->dato.capitan){
            curr = curr->siguiente;
            posicion++;
        }
    }
    return retorno;
}
