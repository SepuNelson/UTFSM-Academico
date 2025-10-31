package tarea3.main;

import java.util.concurrent.RecursiveTask;
import java.util.concurrent.atomic.AtomicInteger;

public class ForkJoin extends RecursiveTask<Boolean> {
    private static final long serialVersionUID = 1L;
    private char[][] laberinto;
    private int x, y;
    private int idProceso;
    private AtomicInteger contadorProceso;
    private static volatile boolean foundExit = false;

    // Constructor principal
    ForkJoin(char[][] laberinto, int x, int y) {
        this(laberinto, x, y, new AtomicInteger(0));
    }

    // Constructor privado con idProceso y contadorProceso
    private ForkJoin(char[][] laberinto, int x, int y, AtomicInteger contadorProceso) {
        this.laberinto = laberinto;
        this.x = x;
        this.y = y;
        this.contadorProceso = contadorProceso;
        this.idProceso = contadorProceso.getAndIncrement(); // Incrementar idProceso
    }

    // Ejecución Principal del Ejercicio con los forks
    @Override
    protected Boolean compute() {

        if (foundExit) {
            return false;
        }

        // Si no es 1 no es necesario seguir
        if (x < 0 || x >= laberinto.length || y < 0 || y >= laberinto[0].length || laberinto[x][y] == '0' || laberinto[x][y] == 'V') {
            return false;
        }

        // Si encuentra una Salida
        if (laberinto[x][y] == 'S') {
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "] - Salida");
            foundExit = true;
            return true;
        }

        // Marcar la posición actual como visitada
        laberinto[x][y] = 'V';

        // Contador para contar las opciones de movimiento disponibles
        // Si esCamino llama nuevamente a la clase ForkJoin
        int options = 0;
        ForkJoin taskIzquierda = null, taskDerecha = null, taskArriba = null, taskAbajo = null;
        if (esCamino(x, y - 1)) {
            options++;
            taskIzquierda = new ForkJoin(laberinto, x, y - 1, contadorProceso);
        }
        if (esCamino(x, y + 1)) {
            options++;
            taskDerecha = new ForkJoin(laberinto, x, y + 1, contadorProceso);
        }
        if (esCamino(x - 1, y)) {
            options++;
            taskArriba = new ForkJoin(laberinto, x - 1, y, contadorProceso);
        }
        if (esCamino(x + 1, y)) {
            options++;
            taskAbajo = new ForkJoin(laberinto, x + 1, y, contadorProceso);
        }

        if (options > 1) {
            // Muestra por pantalla la ubicación del proceso antes de crear los otros cuando hay una división.
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "]");

            // Si hay más de una opción de movimiento, crear forks
            if (taskIzquierda != null) taskIzquierda.fork();
            if (taskDerecha != null) taskDerecha.fork();
            if (taskArriba != null) taskArriba.fork();
            if (taskAbajo != null) taskAbajo.fork();

            // Esperar a que todas las tareas finalicen
            boolean found = false;
            if (taskIzquierda != null) found = taskIzquierda.join() || found;
            if (taskDerecha != null) found = taskDerecha.join() || found;
            if (taskArriba != null) found = taskArriba.join() || found;
            if (taskAbajo != null) found = taskAbajo.join() || found;

            return found;

        } else if (options == 1) {

            // Si solo hay una opción de movimiento continua recursivamente sin crear forks
            if (taskIzquierda != null) return taskIzquierda.compute();
            if (taskDerecha != null) return taskDerecha.compute();
            if (taskArriba != null) return taskArriba.compute();
            if (taskAbajo != null) return taskAbajo.compute();
        } else {
            System.out.println("P" + idProceso + "- [" + y + ", " + x + "]");
        }

        return false;
    }

    // Método para ver si es Camino
    private boolean esCamino(int x, int y) {
        return x >= 0 && x < laberinto.length && y >= 0 && y < laberinto[0].length && laberinto[x][y] != '0' && laberinto[x][y] != 'V';
    }
}
