#include <stdio.h>
#include <sys/stat.h>
#include <sys/types.h>
#include <stdlib.h>
#include <dirent.h>
#include <string.h>
#include <ctype.h>
#include <unistd.h>
#include <stdbool.h>
#include <time.h>

void crear_abd (){
    char* nombre = "Alfabetico";
    char* abc = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
    char ruta[100];

    mkdir(nombre, 0777); //Crea carpeta padre
    for (int i = 0; abc[i] != '\0'; i++){
        sprintf(ruta, "%s/%c", nombre, abc[i]);
        mkdir(ruta, 0777); //Crea todas las carpetas con nombre de todo el abecedario
    }
}

void crear_gen(){
    char* nombre = "Generacion";
    char* abd[] = {"I", "II", "III", "IV"};
    char ruta[100];

    mkdir(nombre, 0777); // Crea carpeta padre
    for (int i = 0; i < 4; i++){
        sprintf(ruta, "%s/%s", nombre, abd[i]);
        mkdir(ruta, 0777); //Crea las carpetas de las generaciones
    }
}

int main() {

    clock_t inicio, fin;
    double tiempo_transcurrido;
    inicio = clock();

    crear_abd();
    crear_gen();

    // Abre el directorio Sprites
    // Ciclo para la lectura de las imagenes
    struct dirent *entrada;     
    DIR *dir_1 = opendir("Sprites");
    while ((entrada = readdir(dir_1)) != NULL) {

        // Excluye los directorios "." y ".."
        if (strcmp(entrada->d_name, ".") != 0 && strcmp(entrada->d_name, "..") != 0) {
            
            // Extraer Datos
            char png[256]; strcpy(png, entrada->d_name);    // Nombre del archivo .png
            char *pokemon = strtok(entrada->d_name, "_.");  // Nombre del Pokemón
            char num[256]; strcpy(num, strtok(NULL, "_.")); // Id del Pokemón
            
            // Ir a la Carpeta correspondiente del Alfabeto
            char letra = toupper(pokemon[0]); char carpeta_1[256];
            sprintf(carpeta_1, "/Alfabetico/%c", letra);
            DIR *dir_2 = opendir(carpeta_1);
        
            // Mover Archivos .png
            char sprites[516];
            char alfa[516];
            sprintf(sprites, "Sprites/%s", png);
            sprintf(alfa, "Alfabetico/%c/%s", letra, png);
            rename(sprites, alfa);
            closedir(dir_2);
        }   
    }
    closedir(dir_1);

    // Abre el directorio Alfabetico
    // Ciclo para la lectura de las imagenes
    struct dirent *entrada_6;     
    DIR *dir_8 = opendir("Alfabetico");
    while ((entrada_6 = readdir(dir_8)) != NULL) {

        // Excluye los directorios "." y ".."
        if (strcmp(entrada_6->d_name, ".") != 0 && strcmp(entrada_6->d_name, "..") != 0) {

            struct dirent *entrada_7;
            
            char subdir_path[516];
            snprintf(subdir_path, sizeof(subdir_path), "%s/%s", "Alfabetico", entrada_6->d_name);

            DIR *dir_9 = opendir(subdir_path);
            while ((entrada_7 = readdir(dir_9)) != NULL) {
                if (strcmp(entrada_7->d_name, ".") != 0 && strcmp(entrada_7->d_name, "..") != 0) {
    
                    // Extraer Datos
                    char png[256]; strcpy(png, entrada_7->d_name);   // Nombre del archivo .png
                    char *pokemon = strtok(entrada_7->d_name, "_."); // Nombre del Pokemón
                    char num[256]; strcpy(num, strtok(NULL, "_."));  // Id del Pokemón



                    // Verificar la Gen del Pokemón
                    int id = atoi(num); char *gen;
                    if(1 <= id && id <= 151){gen = "I";}
                    else if(152 <= id && id <= 251){gen = "II";}
                    else if(252 <= id && id <= 386){gen = "III";}
                    else if(387 <= id && id <= 493){gen = "IV";}

                    // Ir a la Carpeta correspondiente de la Generación
                    char carpeta_2[256];
                    sprintf(carpeta_2, "/Generacion/%s", gen);
                    DIR *dir_3 = opendir(carpeta_2);

                    // Ctrl + C y Ctrl + V
                    char comando_2[1000];
                    sprintf(comando_2, "cp %s/%s Generacion/%s/", subdir_path, png, gen);
                    system(comando_2);
                    closedir(dir_3);
                    
                }
            }
            closedir(dir_9);
        }
    }
    closedir(dir_8);

    // Escrituraa en el archivo .txt
    // Abrir archivo .txt
    FILE *arch;
    arch = fopen("RegistroPokemon.txt", "w");

    // Conteo y Escritura ded la carpeta Generaación
    fprintf(arch, "Generación\n");

    struct dirent *entrada_2;     
    DIR *dir_4 = opendir("Generacion");
    while ((entrada_2 = readdir(dir_4)) != NULL) {
        if (strcmp(entrada_2->d_name, ".") != 0 && strcmp(entrada_2->d_name, "..") != 0) {

            int contador_1 = 0;
            struct dirent *entrada_3;
            
            char subdir_path[516];
            snprintf(subdir_path, sizeof(subdir_path), "%s/%s", "Generacion", entrada_2->d_name);

            DIR *dir_5 = opendir(subdir_path);
            while ((entrada_3 = readdir(dir_5)) != NULL) {
                if (strcmp(entrada_3->d_name, ".") != 0 && strcmp(entrada_3->d_name, "..") != 0) {
                    contador_1 += 1;
                }
            }
            char char_num_1[256];
            sprintf(char_num_1, "%d", contador_1);
            fprintf(arch, "%s - %s\n", entrada_2->d_name, char_num_1); // Escibe en el .txt
            closedir(dir_5);
        }
    }
    closedir(dir_4);

    // Conteo y Escritura ded la carpeta Alfabetico
    fprintf(arch, "Alfabético\n");

    struct dirent *entrada_4;     
    DIR *dir_6 = opendir("Alfabetico");
    while ((entrada_4 = readdir(dir_6)) != NULL) {
        if (strcmp(entrada_4->d_name, ".") != 0 && strcmp(entrada_4->d_name, "..") != 0) {

            int contador_2 = 0;
            struct dirent *entrada_5;
            
            char subdir_path[516];
            snprintf(subdir_path, sizeof(subdir_path), "%s/%s", "Alfabetico", entrada_4->d_name);

            DIR *dir_7 = opendir(subdir_path);
            while ((entrada_5 = readdir(dir_7)) != NULL) {
                if (strcmp(entrada_5->d_name, ".") != 0 && strcmp(entrada_5->d_name, "..") != 0) {
                    contador_2 += 1;
                }
            }
            char char_num_2[256];
            sprintf(char_num_2, "%d", contador_2);
            fprintf(arch, "%s - %s\n", entrada_4->d_name, char_num_2); // Escibe en el .txt
            closedir(dir_7);
        }
    }
    closedir(dir_6);

    fclose(arch);;

    fin = clock();
    tiempo_transcurrido = ((double) (fin - inicio)) / CLOCKS_PER_SEC;
    printf("Tiempo transcurrido: %.2f segundos\n", tiempo_transcurrido);
    return 0;
}
