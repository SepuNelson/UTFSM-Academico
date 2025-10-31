#include <iostream>
#include <sstream>
#include <fstream>
#include <string.h>
#include <string>
#include <stdio.h>
#include "Equipo.hpp"
#include "Tarea_Parte_Arbol.hpp"
#include "Equipo2.cpp"
#include "Tarea_Parte_Arbol2.cpp"

using namespace std;


int main(){
    ifstream fp; // Abrimos el archivo en modo lectura
    fp.open("Equipos.txt");

    // Verificamos si el archivo se abre bien
    if (!fp.is_open())    
    {
        cerr << " Error el abrir el archivo " << endl;
        return -1; 
    }

    // Calculamos el largo del archivo
    // Creamos variables
    int largo = 0;
    int m = 0;
    int abb4 = 0;
    string linea;
    // Creamos un ciclo que se ejecute hasta que el archivo se termine
    while(!fp.eof()){
        getline(fp,linea);
        
        if (m == 0){ // Extraemos la cantidad de equipos
            abb4 = stoi(linea);
            m++;
        }
        largo++;
    }
    
    fp.close(); // Cerramos el archivo

    Equipo *arr = new Equipo[abb4]; // Creamos un arreglo dinamico 
    fp.open("Equipos.txt"); // Nuevamente abrimos el archivo 

    // Verificamos si el archivo se abre bien
    if (!fp.is_open())    
    {
        cerr << " Error el abrir el archivo " << endl;
        return -1; 
    }

    // Leemos el archivo txt para extraer los datos
    // Asignamos variables
    int x = 0;
    int abb2, abb3;
    string integrantes, abb, jugador, capitan;
    // Un ciclo para x el cual sirve como comparador al avanzar en la lectura del txt
    while(x < largo){ 
        if(x == 0){ // Saca la cantidad de equipos para crear el ABB
            getline(fp,abb); // Extraemos la cantidad de equipos
            abb2 = stoi(abb); // Cambiamos el tipo de dato de abb a int y lo guardamos en abb2
            abb3 = 0;
            x++;
        }
        if(x != 0){ // Creamos un formato de extracción de equipos, 1er getline = cantidad de miembros, 2do getline = ciclo para sacar datos de jugadores, 3er getline = saca el capitán
            getline(fp,integrantes); // Extraemos la cantidad de integrantes del Equipo
            int integrantes2 = stoi(integrantes); // La variable integrante la dejamos en entero y la guardamos en integrantes2
            int var2 = x + integrantes2; // Creamos una variable var2 que ayuda en el ciclo de los integrantes del Equipo
            string *Array = new string[integrantes2]; // Creamos un arreglo de tamaño integrantes2
            int n = 0;
            while( x < var2 ){ // Ciclo para agregar cada integrante a un arreglo de su equipo
                getline(fp,jugador);
                Array[n] = jugador;
                n++;              
                x++;  
            }
            x = x+2;
            getline(fp,capitan); // Extraemos el capitán
            string capitan2;
            int largo2 = capitan.size(); // A la variable capitán le extraemos el tamaño para guardarlo en una variable largo2
            if( x != largo){ // Condición en caso de que la variable x sea distinto al largo del archivo
                for(int a = 0; a < (largo2 - 1 ); a++){ // Se recorta el valor de capitan[a] pq con el getline saca el espacio o salto de linea que no nos sirve
                capitan2 = capitan2 + capitan[a];
                }
            }
            if ( x == largo){ // Condición en caso de que la variable x sea igual al largo del archivo
                capitan2 = capitan; // En este caso es el ultimo getline que se pide del archivo por lo que no hay espacio o salto de linea y guarda la variable como la extrae 
            }

            
            Equipo capitan; // Creamos la lista que tendrá el nombre del capitán de tipo Equipo
            bool cap;
            for(int i = 0; i < integrantes2; i++){ // Ciclo para separar los datos de jugador en nombre y poder 
                jugador = Array[i];
                stringstream input_strinstream(jugador);
                char delimitador = ' ';
                string nombre, poder;
                getline(input_strinstream, nombre, delimitador);
                getline(input_strinstream, poder, delimitador);
                int poder2 = stoi(poder);
                if(nombre == capitan2){ // Si el nombre es igual al valor de capitán 2 cambiamos el valor de cap a true
                    cap = true;  
                } 
                capitan.agregar_companero(nombre,cap,poder2); // Agregamos los datos a la lista capitán 
                cap = false;
                        
            }

            if (abb3 < abb2){ // Añadimos al arreglo arr la lista capitán 
                arr[abb3] = capitan;
                abb3++;
            }
            
        }
    }
    fp.close(); // Cerramos el archivo

    Torneo tournament; 
    tournament.crear_torneo(arr, abb4); // Creamos el Torneo

    // Ciclo que funciona hasta que el torneo termine
    int aviso = 0;
    int t, k;
    while(aviso < log2(abb4)){
        cin >> t; // Pedimos datos
        if(t == 1){
            // Avanzamos una ronda
            tournament.avanzar_ronda();
           
            if (aviso==(log2(abb4) - 1)){ // Este es el ultimo caso en que se pueda avanzar
                tournament.imprimir_bracket(); // Imprime el bracket
            } 
            aviso++;

        }
        if(t == 2){
            // Imprimimos estado actual del bracket
            tournament.imprimir_bracket();
            
        }
        if(t == 3){
    
            // Pedimos otro dato
            cin >> k; // Pedimos k
            //  Comparar poderes de cada equipo con k
            for(int i = 0; i < abb4; i++){
                if(arr[i].calcular_poder() >= k){
                    arr[i].imprimir_equipo();
                }
            }
            
        }
    }
    
    return 0;
}
