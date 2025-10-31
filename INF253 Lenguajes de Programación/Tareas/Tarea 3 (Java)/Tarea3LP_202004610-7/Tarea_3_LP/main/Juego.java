import java.util.Scanner;

import zona.Zona;
import zona.Pildora;
import zona.Pieza;
import zona.Muralla;
import zona.Enemigo;

import pikinim.Pikinim;
import pikinim.Amarillo;
import pikinim.Cyan;
import pikinim.Magenta;



public class Juego {
    public static void main(String[] args) {

        System.out.println("====================================================");
        System.out.println(".             Bienvenido a Pikinim                 .");
        System.out.println("====================================================");
        Scanner reader = new Scanner(System.in);


        // Generar Mapa
        Zona[] Mapa = new Zona[11];
        Mapa[0] = new Pieza(true,50);
        Mapa[1] = new Enemigo(true, 130, 20, 25);
        Mapa[2] = new Enemigo(true, 50, 10, 15);
        Mapa[3] = new Pildora(true, 25);
        Mapa[4] = new Muralla(true, 50);
        Mapa[5] = new Pieza(true,100);
        Mapa[6] = new Enemigo(true, 45, 8, 10);
        Mapa[7] = new Pieza(true,150);
        Mapa[8] = new Pildora(true, 15);
        Mapa[9] = new Enemigo(true, 75, 15, 20);
        Mapa[10] = new Muralla(true, 150);

        System.out.println("\nA continuación se mostrará el Tablero");
        System.out.println(".........................................................................................................");
        System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
        System.out.println("........................................................................................................."); 

        // Inicializar los Pikinims
        Pikinim P_cyan = new Cyan(10);
        Pikinim P_magenta = new Magenta(10);
        Pikinim P_amarillo = new Amarillo(10);


        // Turnos del Juego
        int turno = 1;
        int pos = 5;
        while(turno != 31){
            System.out.println("\nTurno " + turno + "( Cyan - " + P_cyan.get_cantidad() + ", Amarillo - " + P_amarillo.get_cantidad() + ", Magenta - " + P_magenta.get_cantidad() + " )\n");
            boolean flag = true;


            // Loop para ingresar un movimiento
            while(flag == true){
                int izquierda = pos -1 ;
                int derecha = pos +1;
                int aqui = pos;
                if(izquierda < 0){izquierda = 10;}
                if(derecha > 10){derecha = 0;}

                System.out.println("Ingrese una de las opciones de movimiento");
                System.out.println("1. Ir a derecha ( " + Mapa[derecha].get_nombre() + " ) | 2. Ir a la izquierda ( " + Mapa[izquierda].get_nombre() + " ) | 3. Quedarse aquí ( " + Mapa[aqui].get_nombre() + " ) ");

                // Operaciones para variable pos
                int accion = reader.nextInt();
                if (accion == 1) {pos += 1;flag = false;}
                else if (accion == 2) {pos -= 1;flag = false;}
                else if (accion == 3) {flag = false;}
                else {System.out.println("\n\nIngrese una opción correcta\n\n");}

                // Circularidad del Tablero
                if(pos < 0){pos = 10;}
                else if(pos > 10){pos = 0;}
            }


            // Juego del Tablero
            if(pos == 0){
                // Pieza

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("........................................................................................................");
                System.out.println("| Aquí | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("........................................................................................................");

                Pieza pieza = (Pieza)Mapa[0];
                int peso_de_pieza = pieza.get_peso();
                System.out.println("\nZona Actual: Pieza (peso - " + peso_de_pieza + ")");

            }else if(pos == 1){
                // Enemigo

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Aquí | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                Enemigo enemigo = (Enemigo)Mapa[1];
                int vida_de_enemigo = enemigo.get_vida();
                System.out.println("\nZona Actual: Enemigo (vida - " + vida_de_enemigo + ")");

            }else if(pos == 2){
                // Enemigo

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Aquí | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                Enemigo enemigo = (Enemigo)Mapa[2];
                int vida_de_enemigo = enemigo.get_vida();
                System.out.println("\nZona Actual: Enemigo (vida - " + vida_de_enemigo + ")");
                
            }else if(pos == 3){
                // Pildora

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Aquí | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                if(Mapa[3].get_usado() == false){
                    System.out.println("\nPíldora Ya Encontrada y Usada");
                }else{
                    Pildora pildora = (Pildora)Mapa[3];
                    int cantidad_de_pildora = pildora.get_cantidad();
                    System.out.println("\nZona Actual: Píldora (cantidad - " + cantidad_de_pildora + ")");

                    double cantidad_de_pikinim = P_cyan.get_cantidad() + P_magenta.get_cantidad() + P_amarillo.get_cantidad();
                    System.out.println("\nLomiar llegó a un lugar lleno de unas figuras con forma de píldoras, los " + cantidad_de_pikinim + " Pikinim se llevan las píldoras.");

                    boolean flag_2 = true;
                    while(flag_2 ==true){
                        System.out.println("Qué color de pikinim desea que se multiplique? (cantidad a multiplicar: 10)");
                        System.out.println("1.Cyan");
                        System.out.println("2.Magenta");
                        System.out.println("3.Amarillo");
                        int accion_2 = reader.nextInt();
                        if(accion_2 == 1){
                            flag_2 = false;
                            P_cyan.multiplicar(P_cyan.get_cantidad());
                            double can = P_cyan.get_cantidad();
                            System.out.println("\nLos Pikinim Cyan han aumentado su cantidad en " + can + ".");
                            
                        }
                        else if(accion_2 == 2){
                            flag_2 = false;
                            P_magenta.multiplicar(P_magenta.get_cantidad());
                            double can = P_magenta.get_cantidad();
                            System.out.println("\nLos Pikinim Magenta han aumentado su cantidad en " + can + ".");
                        }
                        else if(accion_2 == 3){
                            flag_2 = false;
                            P_amarillo.multiplicar(P_amarillo.get_cantidad());
                            double can = P_amarillo.get_cantidad();
                            System.out.println("\nLos Pikinim Amarillos han aumentado su cantidad en " + can + ".");
                        }
                        else{System.out.println("\n\nIngrese una opción correcta\n\n");}
                    }
                    Mapa[3].set_usado(false);
                }
            }else if(pos == 4){
                // Muralla

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Aquí | Pieza | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                if(Mapa[4].get_usado() == false){
                    System.out.println("\nMuralla Destruida");
                    System.out.println("\nVida de Muralla = 0");
                }else{
                    Muralla muralla = (Muralla) Mapa[4];
                    muralla.TryRomper(P_cyan, P_magenta, P_amarillo);
                    System.out.println("\nVida restante de la Muralla es: " + muralla.get_vida());
                    if(muralla.get_vida() == 0){Mapa[4].set_usado(false);}
                }
                
            }else if(pos == 5){
                // Pieza

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("........................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Aquí | Enemigo | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("........................................................................................................");

                Pieza pieza = (Pieza)Mapa[5];
                int peso_de_pieza = pieza.get_peso();
                System.out.println("\nZona Actual: Pieza (peso - " + peso_de_pieza + ")");
                
            }else if(pos == 6){
                // Enemigo

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Aquí | Pieza | Pildora | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                Enemigo enemigo = (Enemigo)Mapa[6];
                int vida_de_enemigo = enemigo.get_vida();
                System.out.println("\nZona Actual: Enemigo (vida - " + vida_de_enemigo + ")");

            }else if(pos == 7){
                // Pieza

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("........................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Aquí | Pildora | Enemigo | Muralla |");
                System.out.println("........................................................................................................");

                Pieza pieza = (Pieza)Mapa[7];
                int peso_de_pieza = pieza.get_peso();
                System.out.println("\nZona Actual: Pieza (peso - " + peso_de_pieza + ")");

            }else if(pos == 8){
                // Píldora

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Aquí | Enemigo | Muralla |");
                System.out.println("......................................................................................................");

                if(Mapa[8].get_usado() == false){
                    System.out.println("\nPíldora Ya Encontrada y Usada");
                }else{
                    Pildora pildora = (Pildora)Mapa[8];
                    int cantidad_de_pildora = pildora.get_cantidad();
                    System.out.println("\nZona Actual: Píldora (cantidad - " + cantidad_de_pildora + ")");

                    double cantidad_de_pikinim = P_cyan.get_cantidad() + P_magenta.get_cantidad() + P_amarillo.get_cantidad();
                    System.out.println("\nLomiar llegó a un lugar lleno de unas figuras con forma de píldoras, los " + cantidad_de_pikinim + " Pikinim se llevan las píldoras.");

                    boolean flag_2 = true;
                    while(flag_2 ==true){
                        System.out.println("Qué color de pikinim desea que se multiplique? (cantidad a multiplicar: 10)");
                        System.out.println("1.Cyan");
                        System.out.println("2.Magenta");
                        System.out.println("3.Amarillo");
                        int accion_2 = reader.nextInt();
                        if(accion_2 == 1){
                            flag_2 = false;
                            P_cyan.multiplicar(P_cyan.get_cantidad());
                            double can = P_cyan.get_cantidad();
                            System.out.println("\nLos Pikinim Cyan han aumentado su cantidad en " + can + ".");
                            
                        }
                        else if(accion_2 == 2){
                            flag_2 = false;
                            P_magenta.multiplicar(P_magenta.get_cantidad());
                            double can = P_magenta.get_cantidad();
                            System.out.println("\nLos Pikinim Magenta han aumentado su cantidad en " + can + ".");
                        }
                        else if(accion_2 == 3){
                            flag_2 = false;
                            P_amarillo.multiplicar(P_amarillo.get_cantidad());
                            double can = P_amarillo.get_cantidad();
                            System.out.println("\nLos Pikinim Amarillos han aumentado su cantidad en " + can + ".");
                        }
                        else{System.out.println("\n\nIngrese una opción correcta\n\n");}
                    }
                    Mapa[8].set_usado(false);
                }
                
            }else if(pos == 9){
                // Enemigo

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Aquí | Muralla |");
                System.out.println("......................................................................................................");

                Enemigo enemigo = (Enemigo)Mapa[9];
                int vida_de_enemigo = enemigo.get_vida();
                System.out.println("\nZona Actual: Enemigo (vida - " + vida_de_enemigo + ")");
                
            }else if(pos == 10){
                // Muralla

                System.out.println("\nUsted se encuentra Aquí");
                System.out.println("......................................................................................................");
                System.out.println("| Pieza | Enemigo | Enemigo | Pildora | Muralla | Pieza | Enemigo | Pieza | Pildora | Enemigo | Aquí |");
                System.out.println("......................................................................................................");

                if(Mapa[10].get_usado() == false){
                    System.out.println("\nMuralla Destruida");
                    System.out.println("\nVida de Muralla = 0");
                }else{
                    Muralla muralla = (Muralla) Mapa[10];
                    muralla.TryRomper(P_cyan, P_magenta, P_amarillo);
                    System.out.println("\nVida restante de la Muralla es: " + muralla.get_vida());
                    if(muralla.get_vida() == 0){Mapa[10].set_usado(false);}
                }
            }
            turno += 1;
        }
        reader.close();
    }  
}

// INDICAR EN QUE POS SE ENCUENTRA AL COMIENZO DEL JUEGO