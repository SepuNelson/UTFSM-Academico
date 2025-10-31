#include <iostream>
#include <string>
#include <fstream>
#include <sstream>
#define VACIA 0;
#define UTILIZADA -1;
//#include "struct_tienda.cpp"
using namespace std;



struct Carta {
    int id;
    string nombre;
    int ataque;
    int defensa;
    int precio;
};

struct Sobre {
    int id;
    Carta cartas[10];
};

struct RanuraC{
    int clave;
    Carta carta;
    int cantidad;
};




int hash1(int clave, int celdas){                       //funcion de hash para el primer intento de insercion
    int posicion = ((clave%283)*29573)%celdas;
    return posicion;
};

int hashCo(int id, int rep){                          // Funcion de hash para resolucion de colisiones
    int pos = (id%37)*35803*rep;
    return pos;
};

int insertar(RanuraC arr[], int id,Carta cAux, int celdas){             //funcion que recibe la tabla hash de cartas, la id a insertar,
    int inicio, i;                                                      //la carta a insertar y la cantidad de celdas que contiene la tabla
    int pos = inicio = hash1(id, celdas);                               //y que inserta la carta requerida usando la funcion de hash
    for(i = 1; arr[pos].clave != 0 && arr[pos].clave != id; i++){
        pos = (inicio + hashCo(id, i)) % celdas;
    }
    if (arr[pos].clave == id){
        arr[pos].cantidad++;
        return 0;
    }
    else {
        arr[pos].clave = id;
        arr[pos].carta = cAux;
        arr[pos].cantidad++;
        return 1;
    }

};

int insertarSobre(Sobre arr[], int id,Sobre sAux, int celdas){  //Funcion que agrega un sobre a la tabla hash
    int inicio, i, x;                                           //Recibe una tabla hash de Sobres, el id a insertar, 
    int pos = inicio = hash1(id, celdas);                       //cantidad de celdas de la tabla hash (sobre)
    for(i = 1; arr[pos].id != 0 && arr[pos].id != id; i++){
        pos = (inicio + hashCo(id, i)) % celdas;
    }
    arr[pos].id = sAux.id;
    for(x = 0; x < 10; x++){
        arr[pos].cartas[x] = sAux.cartas[x];
    }
    return 0;
};

int tengo_la_carta(RanuraC arr[], int id, int celdas){          //funcion que retorna la cantidad de cartas en caso de que exista
    int inicio, i;                                              //la carta dentro de la tabla de hash
    int pos = inicio = hash1(id, celdas);
    for(i = 1; arr[pos].clave != 0 && arr[pos].clave != id && arr[pos].clave != -1; i++){
        pos = (inicio + hashCo(id, i)) % celdas;
    }
    if (arr[pos].clave == id){
        return arr[pos].cantidad;
    }
    else {
        
        return 0;
    }
};


int buscar_pos(RanuraC arr[], int id, int celdas){              //funcion que busca la posicion dentro de la tabla de hash de cartas
    int inicio, i;                                              //de una carta especifica
    int pos = inicio = hash1(id, celdas);
    for(i = 1; arr[pos].clave != 0 && arr[pos].clave != id && arr[pos].clave != -1; i++){
        pos = (inicio + hashCo(id, i)) % celdas;
    }
    if (arr[pos].clave == id){
        return pos;
    }
    else {
        return -5;
    }
};

void hashmostrar_cartas(int M, RanuraC array[]){                //funcion que imprime por pantalla las cartas disponibles en la tienda
        for (int z = 0; z<M; z++){
            if (array[z].clave != 0 && array[z].clave != -1){
                cout << array[z].clave << " " << array[z].carta.nombre << " " << array[z].carta.ataque << " " << array[z].carta.defensa << " " << array[z].carta.precio << " " << array[z].cantidad << endl;
            }
        }

};

void hashmostrar_sobres(int MS, Sobre array[]){             //funcion que imprime por pantalla los sobres disponibles en la tienda
        for (int z = 0; z<MS; z++){
            if (array[z].id != 0 && array[z].id != -1){
                cout << array[z].id << endl;
            }
        }

};

void vender_carta(int id, RanuraC arr[], int M, int &dinero){      //funcion que vende una carta especifica, disminuyendo la cantidad
    int posicion = buscar_pos(arr, id, M);                          //de esta dentro de la tienda y aumentando el dinero que posee                                         //la tienda luego de vender
    if (posicion < -1){
        cout << "Ese id no se encuentra!" << endl;
        return;
    }
    if(arr[posicion].cantidad > 0){
        arr[posicion].cantidad--;
        dinero = dinero + arr[posicion].carta.precio;
        cout << "Vendida la carta!" << endl;
        if (arr[posicion].cantidad == 0){
            arr[posicion].clave = UTILIZADA;
        }
        return;
    }
};

int buscar_sobre(Sobre arr[], int id, int M){                       //funcion que busca la posicion de un sobre dentro de la tabla de hash
    int inicio, i;                                                  //de los sobres
    int pos = inicio = hash1(id,M);
    for(i = 1; arr[pos].id != 0 && arr[pos].id != id && arr[pos].id != -1; i++){
        pos = (inicio + hashCo(id, i)) % M;
    }
    if (arr[pos].id == id){
        return pos;
    }
    else {
        return -5;
    }
};


void vender_sobre(int id, Sobre arr[], int M, int &dinero){        //funcion que vende y elimina sobre especifico y que aumenta dinero a la tienda
    int posicion = buscar_sobre(arr, id, M);                        //luego de haber vendido el sobre
    if(posicion < 0){
        cout << "Ese id no se encuentra!" << endl;
        return;
    }
    dinero = dinero + 1000;
    for(int n = 0; n < 10; n++){
        cout << arr[posicion].cartas[n].id << " " << arr[posicion].cartas[n].nombre << " " <<  arr[posicion].cartas[n].ataque << " " <<  arr[posicion].cartas[n].defensa << " " <<  arr[posicion].cartas[n].precio << endl;
    }
    cout << "Vendido el sobre!" << endl;
    arr[posicion].id = -1;
    return;

};


int main(){
    fstream file;
    string tamanoCarta, tamanoSobre, lectC, lectS, linea;
    int cont, contLin, largo, tamanoTcartas, tamanoTsobres, tamanoIn;
    contLin = cont = largo = tamanoTcartas = tamanoTsobres = tamanoIn = 0;
    int n1 = 0; int MS;
    Carta carta_temp, carta_temps; 
    Sobre sobre_temp;
    file.open("Tienda.txt");
    
    while(!file.eof()){
        getline(file,linea);
        n1 = stoi(linea);
        largo++;
    }
        
    file.close();
    
    file.open("Tienda.txt");
    string *arr = new string[n1];
    for(int i = 0; i < n1; i++){ //dejar en 0 todo el arreglo
        arr[i] = "a";   
    }
    if (!file.is_open()){
        cout << "Error al abrir el archivo" << endl;
        return 0;
    }
    int contLin1 = 0; int tarr = 0; int tarr1 = 0; int cartasunicas = 0; 
    string lectCarta;
    while(contLin1 < largo){
        
        
        if(contLin1 == 0){
            getline(file,linea);
            tarr = stoi(linea);
            tarr1 = tarr;
        } 
        else{
            
            if (contLin1 < tarr + 1 && tarr1 > 0){
                tarr1 = tarr1 - 1;
                getline(file, lectCarta);               
                stringstream input_strinstream(lectCarta);
                char delimitador = ' ';
                string id, nombre, ataque, defensa, precio;
                getline(input_strinstream, id, delimitador);
                getline(input_strinstream, nombre, delimitador);
                getline(input_strinstream, ataque, delimitador);
                getline(input_strinstream, defensa, delimitador);
                getline(input_strinstream, precio, delimitador);
                for(int i = 0; i < tarr; i++){ //Leer la cantidad de cartas distintas
                    for (int c = 0; c <= i; c++){
                        
                        if(arr[c] == id){
                            break;  
                        }
                        if(arr[c] == "a" && arr[c] != id){
                            arr[i] = id;
                        }
                        
                    }
                }  
            } 

            if (contLin1 == tarr + 1){
                getline(file, linea);
                MS = stoi(linea);
            } 
        }
        contLin1++;
    } 
    file.close();
    for(int i = 0; i < tarr; i++){          
        if(arr[i] != "a"){
            cartasunicas++;
        }
    }
    int M = cartasunicas/0.6;

    delete[] arr;

    RanuraC *tablaCarta = new RanuraC[M];
    for(int r = 0; r < M; r++){
        tablaCarta[r].clave = VACIA;
        tablaCarta[r].cantidad = 0;
    }

    Sobre *tablaSobre = new Sobre[MS];
    for(int r = 0; r < MS; r++){
        tablaSobre[r].id = VACIA;
    }


    file.open("Tienda.txt");
    if (!file.is_open()){
        cout << "Error al abrir el archivo" << endl;
        return 0;
    }
    while(contLin < largo){
        if(contLin == 0){
                getline(file, tamanoCarta);
                tamanoTcartas = stoi(tamanoCarta);
                tamanoIn = tamanoTcartas;
        }
        else{
            if (contLin < tamanoTcartas+1){
                getline(file,lectC);
                //cout << lectC << endl;
                stringstream input_strinstream(lectC);
                char delimitador = ' ';
                string id, nombre, ataque, defensa, precio;
                getline(input_strinstream, id, delimitador);
                getline(input_strinstream, nombre, delimitador);
                getline(input_strinstream, ataque, delimitador);
                getline(input_strinstream, defensa, delimitador);
                getline(input_strinstream, precio, delimitador);
                
                carta_temp.id = stoi(id);
                carta_temp.nombre = nombre;
                carta_temp.ataque = stoi(ataque);
                carta_temp.defensa = stoi(defensa);
                carta_temp.precio = stoi(precio);
                insertar(tablaCarta, stoi(id), carta_temp, M);
                                          
            }
            if (contLin == tamanoTcartas+1){            //se obtiene la cantidad de sobres de la tienda
                getline(file, tamanoSobre);
                //cout << tamanoTsobres << endl;
            }
            if(contLin > tamanoTcartas+1){              //lectura de texto correspondiente a sobres
                getline(file,lectS);
                if(contLin % (tamanoIn+2) == 0){        //se obtiene el ID de los sobres
                    tamanoIn = tamanoIn + 11;
                    sobre_temp.id = stoi(lectS);
                    cont = 0;
                }
                else{
                    stringstream input_strinstream(lectS);
                    char delimitador = ' ';
                    string id, nombre, ataque, defensa, precio;
                    getline(input_strinstream, id, delimitador);
                    getline(input_strinstream, nombre, delimitador);
                    getline(input_strinstream, ataque, delimitador);
                    getline(input_strinstream, defensa, delimitador);
                    getline(input_strinstream, precio, delimitador);
                    sobre_temp.cartas[cont].id = stoi(id);
                    sobre_temp.cartas[cont].nombre = nombre;
                    sobre_temp.cartas[cont].ataque = stoi(ataque);
                    sobre_temp.cartas[cont].defensa = stoi(defensa);
                    sobre_temp.cartas[cont].precio = stoi(precio);
                    cont++;
                    if (cont == 10){
                        insertarSobre(tablaSobre, sobre_temp.id, sobre_temp, MS);
                    }            
                }

            }

        }
        contLin++;
    }
    //PARTE INTERACTIVA DEL PROGRAMA
    int o; int cond = 0; int dinero = 0; int id_comprar; 
    while(cond == 0){
        cin >> o;
        if(o == 1){
            cout << dinero << endl;
        }
        if(o == 2){
            hashmostrar_cartas(M, tablaCarta);
        }
        if(o == 3){
            hashmostrar_sobres  (MS, tablaSobre);
            
        }
        if(o == 4){
            cin >> id_comprar;
            vender_carta(id_comprar, tablaCarta, M, dinero);
            // Vender cartas
        }
        if(o == 5){
            cin >> id_comprar;
            vender_sobre(id_comprar, tablaSobre, M, dinero);
            
            
        }
        if(o == 6){
            delete[] tablaCarta;
            delete[] tablaSobre;
            cond++;
        }
    }
    return 0;
}