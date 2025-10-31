package tarea3.main;

import java.util.concurrent.atomic.AtomicInteger;

public class Threads implements Runnable {
    private char[][] laberinto;
    private int x, y;
    private int idProceso;
    private AtomicInteger contadorProceso;
    private static volatile boolean foundExit = false;

    // Constructor principal
    public Threads(char[][] laberinto, int x, int y) {
        this(laberinto, x, y, new AtomicInteger(0));
    }

    // Constructor privado con idProceso y contadorProceso
    private Threads(char[][] laberinto, int x, int y, AtomicInteger contadorProceso) {
        this.laberinto = laberinto;
        this.x = x;
        this.y = y;
        this.contadorProceso = contadorProceso;
        this.idProceso = contadorProceso.getAndIncrement();
    }

    @Override
    public void run() {

        if (foundExit) {
            System.exit(0);;
        }

        // Si no es 1 no es necesario seguir
        if (x < 0 || x >= laberinto.length || y < 0 || y >= laberinto[0].length || laberinto[x][y] == '0' || laberinto[x][y] == 'V') {
            return;
        }

        // Si encuentra una Salida
        if (laberinto[x][y] == 'S') {
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "] - Salida");
            foundExit = true;
            System.exit(0);
            return;
        }

        // Marcar la posición actual como visitada
        laberinto[x][y] = 'V';

        // Contador para contar las opciones de movimiento disponibles
        // Si esCamino llama nuevamente a la clase Threads
        int options = 0;
        Threads taskIzquierda = null, taskDerecha = null, taskArriba = null, taskAbajo = null;
        if (esCamino(x, y - 1)) {
            options++;
            taskIzquierda = new Threads(laberinto, x, y - 1, contadorProceso);
        }
        if (esCamino(x, y + 1)) {
            options++;
            taskDerecha = new Threads(laberinto, x, y + 1, contadorProceso);
        }
        if (esCamino(x - 1, y)) {
            options++;
            taskArriba = new Threads(laberinto, x - 1, y, contadorProceso);
        }
        if (esCamino(x + 1, y)) {
            options++;
            taskAbajo = new Threads(laberinto, x + 1, y, contadorProceso);
        }

        if (options > 1) {
            // Muestra por pantalla la ubicación del proceso antes de crear los otros cuando hay una división.
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "]");

            // Si hay más de una opción de movimiento, crear threads
            if (taskIzquierda != null) new Thread(taskIzquierda).start();
            if (taskDerecha != null) new Thread(taskDerecha).start();
            if (taskArriba != null) new Thread(taskArriba).start();
            if (taskAbajo != null) new Thread(taskAbajo).start();
        } else if (options == 1) {

            // Si solo hay una opción de movimiento continua recursivamente sin crear threads
            if (taskIzquierda != null) taskIzquierda.run();
            if (taskDerecha != null) taskDerecha.run();
            if (taskArriba != null) taskArriba.run();
            if (taskAbajo != null) taskAbajo.run();
        } else {
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "]");
        }
    }

    // Método para ver si es Camino
    private boolean esCamino(int x, int y) {
        return x >= 0 && x < laberinto.length && y >= 0 && y < laberinto[0].length && laberinto[x][y] != '0' && laberinto[x][y] != 'V';
    }
}