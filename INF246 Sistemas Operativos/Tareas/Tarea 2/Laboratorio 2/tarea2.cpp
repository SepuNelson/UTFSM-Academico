//............LIBRERÍAS...........//
#include <iostream>
#include <stdlib.h>
#include <unistd.h>
#include <time.h>
#include <sys/types.h>
#include <sys/wait.h>
using namespace std;

//.............STRUCTS............//
/* Struct de los Jugadores */
struct Personaje {
    int vida;
    int ataque;
    int defensa;
    int evasion;
};

//............FUNCIONES...........//
/* Función que entrega un número aleatorio
entre un mínimo y un máximo incluyendolos*/
int random(int min, int max) {
    return min + rand() % (max - min + 1);
}

/* Función que crea un personaje para cada proceso */
Personaje crear_personaje(){
    struct Personaje jugador;
    jugador.vida = 100;
    jugador.ataque = random(30, 40);
    jugador.defensa = random(10, 25);
    jugador.evasion = 60 - jugador.defensa;
    return jugador;
}

/* Función que calcula el daño al jugador atacado */
int calculo_ataque(int arrJugadores[], Personaje arrPersonaje[]){
    //............VARIABLES............//
    int NumR, J1, J2;
    J1 = arrJugadores[0];
    J2 = arrJugadores[1];

    cout << "\nJugador " << J1 + 1 << " ataca al Jugador " << J2 + 1 << endl;

    Personaje jugador1, jugador2;
    jugador1 = arrPersonaje[J1];
    jugador2 = arrPersonaje[J2];

    //..........CÁLCULO.DAÑO..........//
    int Damage;
    NumR = random(0,100);
    /* if -> jugador es atacado
       else -> jugador esquiva */
    if(NumR > jugador2.evasion){
        cout << "Jugador " << J2 + 1 << " recibe " << jugador1.ataque - jugador2.defensa << " de daño" << endl;
        Damage = jugador1.ataque - jugador2.defensa;
        return Damage;
    }else{
        cout << "Jugador " << J2 + 1 << " esquivó el ataque!\n" << endl;
        Damage = 0;
        return Damage;
    }
}

//..............MAIN..............//
int main() {

    cout << "\nComienza el Juego \n" << endl;

    srand(time(NULL)); //Semilla para los número aleatorios

    /* Arreglo del Estado de los Jugadores */
    int arrEstado[4];
    for(int i = 0; i < 4; i++){
        arrEstado[i] = 1;
    }

    //............PID.PADRE............//
    pid_t pidPadre = getpid();

    //..............PIPES..............//
    int pipes[4][2];  // Darle a los hijos un personaje
    int pipes2[4][2]; // Mostrar la Info
    int pipes3[4][2]; // Pedirle a los hijos el personaje
    int pipes4[4][2]; // Editar el Personaje de los hijos
    int pipes5[4][2]; // Pipe para detener los hijos 2, 3, 4 antes del cin del hijo 1
    int pipes6[4][2]; // Señal de padre a Hijos para nivelar velocidad de ejecución
    int pipes7[4][2]; // Info del arrEstado
    int pipes8[4][2]; // Actualizar flag
    for(int i = 0; i < 4; i++){
        pipe(pipes[i]);
        pipe(pipes2[i]);
        pipe(pipes3[i]);
        pipe(pipes4[i]);
        pipe(pipes5[i]);
        pipe(pipes6[i]);
        pipe(pipes7[i]);
        pipe(pipes8[i]);
    }


    //...............HIJOS...............//
    Personaje jugador;
    pid_t pid[4];
    for(int i = 0; i < 4; i++){
        pid[i] = fork();
        if(pid[i] == 0){
            close(pipes[i][1]);
            read(pipes[i][0], &jugador, sizeof(jugador));
            break;
        }
    }

    //............CREACIÓN.DE.PERSONAJES............//
    if(getpid() == pidPadre){
        for (int i = 0; i < 4; ++i) {
            Personaje jugador = crear_personaje();
            close(pipes[i][0]);
            write(pipes[i][1], &jugador, sizeof(jugador));  
        }
    }

    //............COMIENZA.EL.JUEGO............//
    /* Ciclo de Turnos */
    int flag = 0;
    for(int turno = 1; true; turno++){

        /* Todavía no hay 3 muertos */
        if(flag == 0){
            if(getpid() == pidPadre){
                /* Cuenta la cantidad de Jugadores Muertos*/
                int muertos = 0;
                for(int i = 0; i < 4; i++){
                    if(arrEstado[i] == 0){
                        /* 0 para Muertos*/
                        muertos++;
                    }
                }
                /* Detiene los procesos si hay 3 Jugadores Muertos */
                if(muertos == 3){
                    flag = 1;
                    for(int x = 0; x < 4; x++){
                        if(arrEstado[x] == 1){
                            cout << "\n\n\nFELICIDADES!!!!!" << endl;
                            cout << "El Ganador del Juego es el Jugador " << x + 1 << endl;
                            cout << "Fin del Juego\n" << endl;
                        }
                    }
                }else if(muertos == 4){
                    flag = 1;
                    cout << "\n\n\nTodos los Jugadores murieron" << endl;
                    cout << "Es un Empate!!!!\n" << endl;
                }

                /* Actualiza la flag que determina si los procesos deben continuar */
                for(int i = 0; i < 4; i++){
                    close(pipes8[i][0]);
                    write(pipes8[i][1], &flag, sizeof(flag));
                }


                if(flag == 0){
                    /* Mandar con Pipes la info del arreglo de estados */
                    for(int i = 0; i < 4; i++){
                        close(pipes7[i][0]);
                        write(pipes7[i][1], &arrEstado, sizeof(arrEstado));
                    }

                    /* Muestra la Ronda Actual  */
                    cout << "\nRonda Actual => " << turno << "\n" << endl;
                        
                    //............IMPRIMIR.INFORMACIÓN............//
                    Personaje arrPersonaje[4];
                    for(int i = 0; i < 4; i++){
                        close(pipes2[i][1]);
                        read(pipes2[i][0], &jugador, sizeof(jugador));
                        arrPersonaje[i] = jugador;
                        cout << "Jugador " << i + 1 << " => "
                            << " Vida: " << jugador.vida 
                            << " Ataque: " << jugador.ataque
                            << " Defensa: " << jugador.defensa
                            << " Evasión: " << jugador.evasion << endl;
                    }
                    
                    /* Espera a que se ingrese un valor de J2 en Jugador 1 */
                    for(int i = 0; i < 4; i++){
                        char Signal;
                        close(pipes6[i][0]);
                        write(pipes6[i][1], &Signal, sizeof(Signal));
                    }

                    
                    //............CÁCULOS.DE.DAÑOS............//
                    int arrVidas[4];
                    for(int i = 0; i < 4; i++){
                        int HP = arrPersonaje[i].vida;
                        arrVidas[i] = HP;
                    }

                    int arrJugadores[2];
                    for(int i = 0; i < 4; i++){

                        /* Espera el arreglo que contiene al jugaddor Atacante y al Atacado */
                        close(pipes3[i][1]);
                        read(pipes3[i][0], &arrJugadores, sizeof(arrJugadores));
                        if(arrJugadores[0] == i){
                            /* Solo atacan quienes tengan vida mayor a 0 */
                            if(arrPersonaje[i].vida > 0){
                                /* Calcula el Daño */
                                int Damage = calculo_ataque(arrJugadores, arrPersonaje);

                                /* Guarda la Vida Actualizada en el arrVidas */
                                arrVidas[arrJugadores[1]] = arrVidas[arrJugadores[1]] - Damage;
                            }
                        }
                    }

                    /* Manda la info de la nueva Vida a los Hijos */
                    for(int i = 0; i < 4; i++){
                        int Vida = arrVidas[i];
                        if(Vida <= 0){
                            arrEstado[i] = 0;
                        }
                        close(pipes4[i][0]);
                        write(pipes4[i][1], &Vida, sizeof(Vida));
                    }
                }
            }else if(pid[0] == 0){

                /* Recibe información de la flag que le permite continuar o no el proceso */
                close(pipes8[0][1]);
                read(pipes8[0][0], &flag, sizeof(flag));

                if(flag == 0){
                    /* Recibir información del padre de los estados actualizados */
                    int arrEstados[4];
                    close(pipes7[0][1]);
                    read(pipes7[0][0], &arrEstados, sizeof(arrEstados));

                    /* Pipes que mandan la información de
                    su Personaje a "IMPRIMIR INFORMACIÓN" */
                    close(pipes2[0][0]);
                    write(pipes2[0][1], &jugador, sizeof(jugador));
                    sleep(1);

                    int arrJugadores[2], J2, Vida;
                    if(jugador.vida > 0){

                        /* Pide el número del Jugador que quiere atacar */
                        cout << "\nEs tu Turno" << endl;
                        cout << "A qué jugador quiere atacar? (2,3,4) : ";
                        cin >> J2;
                        J2--;
                        while(J2 == 0){
                            cout << "No puedes atacarte a ti mismo!!" << endl;
                            cout << "A qué jugador quiere atacar? (2,3,4) : ";
                            cin >> J2;
                            J2--;
                        }
                        while(arrEstados[J2] == 0){
                            cout << "Ese Jugador no tiene vida, escoje otro para atacar: ";
                            cin >> J2;
                            J2--;
                        }
                        
                    }else{
                        J2 = 0;
                    }
                
                    /* Manda señal a sus Hermanos */
                    for(int i = 0; i < 4; i++){
                        char Signal;
                        close(pipes5[i][0]);
                        write(pipes5[i][1], &Signal, sizeof(Signal));
                    }

                    /* Manda la señal al Padre */
                    close(pipes6[0][0]);
                    char SignalPadre;
                    write(pipes6[0][1], &SignalPadre, sizeof(SignalPadre));

                    /* Manda la info del Atacante y el Atacado */
                    arrJugadores[0] = 0;
                    arrJugadores[1] = J2;
                    close(pipes3[0][0]);
                    write(pipes3[0][1], &arrJugadores, sizeof(arrJugadores));

                    /* Espera la info de la nueva Vida */
                    close(pipes4[0][1]);
                    read(pipes4[0][0], &Vida, sizeof(Vida));

                    /* Actualizar vida */
                    if(Vida <= 0){
                        Vida = 0;
                        jugador.vida = Vida;
                    }else{
                        jugador.vida = Vida;
                    }
                }
            }else{

                /* Indice que indica que hijo es */
                int i = getpid() - pid[0];

                /* Recibe información de la flag que le permite continuar o no el proceso */
                close(pipes8[i][1]);
                read(pipes8[i][0], &flag, sizeof(flag));
                
                if(flag == 0){
                    /* Recibir información del padre de los estados actualizados */
                    int arrEstados[4];
                    close(pipes7[i][1]);
                    read(pipes7[i][0], &arrEstados, sizeof(arrEstados));

                    /* Pipes que mandan la información de
                    su Personaje a "IMPRIMIR INFORMACIÓN" */
                    close(pipes2[i][0]);
                    write(pipes2[i][1], &jugador, sizeof(jugador));
                    sleep(1);

                    /* Espera a que el Jugador 1 ingrese el J2 */
                    close(pipes5[i][1]);
                    char Signal;
                    read(pipes5[i][0], &Signal, sizeof(Signal));

                    /* Manda la señal al Padre */
                    close(pipes6[i][0]);
                    char SignalPadre;
                    write(pipes6[i][1], &SignalPadre, sizeof(SignalPadre));

                    /* Define variables y saca el J2 como número random */
                    int arrJugadores[2], NumR, Vida;
                    NumR = random(0,3);
                    while(NumR == i){
                        NumR = random(0,3);
                    }

                    while(arrEstados[NumR] == 0){
                        NumR = random(0,3);
                        while(NumR == i){
                            NumR = random(0,3);
                        }
                    }

                    /* Manda la info del Atacante y el Atacado */
                    arrJugadores[0] = i;
                    arrJugadores[1] = NumR;
                    close(pipes3[i][0]);
                    write(pipes3[i][1], &arrJugadores, sizeof(arrJugadores));

                    /* Espera la info de la nueva Vida */
                    close(pipes4[i][1]);
                    read(pipes4[i][0], &Vida, sizeof(Vida));
                    
                    /* Actualizar vida */
                    if(Vida < 0){
                        Vida = 0;
                        jugador.vida = Vida;
                    }else{
                        jugador.vida = Vida;
                    }
                }
            }
        }
        /* Hay 3 muertos */
        else{
            break;
        }
    }
    return 0;
}
