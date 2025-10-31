#include <iostream>
#include <fstream>
#include <string.h>

using namespace std;


// Estructura 'Registro' para leer el archivo binario
struct Registro {
    int dia;
    int mes;
    int anio;
    char rol[12];
    int ppm;
    float precision;
};

// Struct creado para almacenar los datos del estudiante con mejor precision en la fecha requerida (precision del estudiante y rol)
struct EstudiantePrecision {    
    float precisionEst = 0;
    char* rolPre;
};

// Struct creado para almacenar los datos del estudiante con mejor ppm en la fecha requerida (ppm del estudiante y rol)
struct EstudiantePpm {          
    int ppmEst = 0;
    char* rolPpm;
};



int main(){ 

    ifstream fp; //Abrimos el archivo binario para lectura
    
    fp.open("registros.dat", ios::binary);

    //Sentencia que actua si el archivo no es abierto correctamente
    if (!fp.is_open())    
    {
        cerr << " Error el abrir el archivo " << endl;
        return -1; // error
    }

    // lee el primer dato de tipo entero, el cual indica en número de registros
    
    int nRegistros;
    fp.read((char *)&nRegistros, sizeof(int));
    Registro registrosArray[nRegistros]; // declaramos el array para guardar la información de los registros
    

    Registro registros;
    int j = 0;
    while (fp.read((char *)&registros, sizeof(Registro))) // Leemos el archivo binario
    {
        registrosArray[j] = registros; // Guardamos los structs del Archivo Binario en un arreglo
        j++;
    }
    
    fp.close(); // Cerramos el archivo binario
    
    

    int q;   //corresponde a la cantidad de veces que se preguntará por los datos, que vendrá dada por el input de la linea 66

    int t, d, m, a; //corresponden a Tipo de dato, Dia, Mes y Anio, respectivamente. Todos designados por el input de las lineas 70-73
    cin >> q;

    while (q>0) //en este ciclo se trabaja todo el programa por cada input que se solicita 
    {
        cin >> t;
        cin >> d;
        cin >> m;
        cin >> a;

      EstudiantePrecision maxPre;    //declaramos los structs que almacenarán los datos requeridos en los inputs anteriores
      EstudiantePpm maxPpm;


      if(t == 0){ // si se requiere precision
        if(d == -1){ // no se cuenta el dia
          if(m == -1){ // solo se compara año
            for(int i = 0; i < nRegistros; i++){ // empezamos el ciclo para comparar y encontrar la fecha requerida
              if(a == registrosArray[i].anio){ 
                if(registrosArray[i].precision > maxPre.precisionEst){ // si la precision de la coincidencia encontrada en el archivo binario es mayor que la precision almacenada en el struct, entonces pasa a ser la nueva precision y nuevo rol 
                    maxPre.precisionEst = registrosArray[i].precision; //actualizamos maxima precision 


                    
                    maxPre.rolPre = registrosArray[i].rol; //actualizamos rol de maxima precision
                }
              } 
            }
          }else{ // se compara mes y año
            for(int i = 0; i < nRegistros; i++){  // iniciamos el ciclo para comparar y encontrar la fecha requerida
              if(m == registrosArray[i].mes && a == registrosArray[i].anio){ //aqui es necesario comparar solamente mes y anio
                if(registrosArray[i].precision > maxPre.precisionEst){
                  maxPre.precisionEst = registrosArray[i].precision;//actualizamos maxima precision 


                  maxPre.rolPre = registrosArray[i].rol;//actualizamos rol de maxima precision
                }
              }
            }
          }
          
        }else{ // se compara todo
          for(int i = 0; i < nRegistros; i++){  // iniciamos el ciclo para comparar y encontrar la fecha requerida
              if(m == registrosArray[i].mes && a == registrosArray[i].anio && d == registrosArray[i].dia){
               if(registrosArray[i].precision > maxPre.precisionEst){
                  maxPre.precisionEst = registrosArray[i].precision;//actualizamos maxima precision 


                  maxPre.rolPre = registrosArray[i].rol;//actualizamos rol de maxima precision
                 }
              }
            }
          
        }
          
        
      }else{ // si se requiere ppm
        if(d == -1){ // no se cuenta el dia
          if(m == -1){ // solo se compara año
            for(int i = 0; i < nRegistros; i++){  // iniciamos el ciclo para comparar y encontrar la fecha requerida
              if(a == registrosArray[i].anio){
                if(registrosArray[i].ppm > maxPpm.ppmEst){
                    maxPpm.ppmEst = registrosArray[i].ppm;//actualizamos maxima ppm

                  
                    maxPpm.rolPpm = registrosArray[i].rol;//actualizamo maximo rol de ppm
                }
              } 
            }
          }else{ // se compara mes y año
            for(int i = 0; i < nRegistros; i++){  // iniciamos el ciclo para comparar y encontrar la fecha requerida
              if(m == registrosArray[i].mes && a == registrosArray[i].anio){  
                if(registrosArray[i].ppm > maxPpm.ppmEst){
                  maxPpm.ppmEst = registrosArray[i].ppm;//actualizamos maxima ppm

                  
                  maxPpm.rolPpm = registrosArray[i].rol;//actualizamo maximo rol de ppm
                }
              }
            }
          }
          
        }else{ // se compara todo
          for(int i = 0; i < nRegistros; i++){  // iniciamos el ciclo para comparar y encontrar la fecha requerida
              if(m == registrosArray[i].mes && a == registrosArray[i].anio && d == registrosArray[i].dia){
               if(registrosArray[i].ppm > maxPpm.ppmEst){
                  maxPpm.ppmEst = registrosArray[i].ppm;//actualizamos maxima ppm

                
                  maxPpm.rolPpm = registrosArray[i].rol;//actualizamo maximo rol de ppm
                 }
              }
            }
          
        }
      }

     
        ifstream fp2;
        fp2.open("estudiantes.txt", ios::in);

        // verifica que el archivo se abrió
        if (!fp2.is_open())
        {
            cerr << " Error el abrir el archivo " << endl;
            return -1; // error
        }

        string rol, nombre, apellido, paralelo;

        while(!fp2.eof()){

          int z = 0;

          while(z<4){
            if(z==0){
              fp2 >> rol;
              z++;
            }
            if(z==1){
              fp2 >> nombre;
              z++;
            }
            if(z==2){
              fp2 >> apellido;
              z++;
            }
            if(z==3){
              fp2 >> paralelo;
              z++;
            }
          }

          if (t == 0){
            if(maxPre.rolPre == rol){
              cout << nombre << " " << apellido << endl;   
            }
          }
          if (t == 1){
            if(maxPpm.rolPpm == rol){
              cout << nombre << " " << apellido << endl;
            }
          }
        
        

        
        }

        fp2.close();        
        q--;
    }

} 