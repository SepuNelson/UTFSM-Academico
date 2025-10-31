package tarea3.main;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.util.List;
import java.util.Scanner;
import java.util.concurrent.ForkJoinPool;

public class Main {
    public static void main(String[] args) throws IOException {

        // Leemos el archivo que se encuentra en la carpeta tarea3 -> laberinto
        List<String> lines = Files.readAllLines(Paths.get("laberinto/laberinto.txt"));

        // Leemos la dimensión del laberinto
        String strDim = lines.get(0).trim();
        String[] arrDim = strDim.split("x");
        int dimension = Integer.parseInt(arrDim[0].trim());

        // Leemos el punto de partida y Setteamos Coords X e Y
        String[] coords = lines.get(1).trim().replace("[", "").replace("]", "").split(",");
        int coordX = Integer.parseInt(coords[0].trim());
        int coordY = Integer.parseInt(coords[1].trim());

        // Lee el archivo .txt y pasa el laberinto a una matriz de caracteres de tamaño dimensión
        char[][] laberinto = new char[dimension][dimension];
        for (int i = 0; i < dimension; i++) {
            String line = lines.get(2 + i).replace(" ", "");
            laberinto[i] = line.toCharArray();
        }

        // Las 2 formas de resolver el problema
        System.out.println("Seleccione el método de resolución:");
        System.out.println("1. Usando Threads");
        System.out.println("2. Usando ForkJoin");

        // Pedimos la forma de resolver el problema
        Scanner scanner = new Scanner(System.in);
        int option = scanner.nextInt();

        if (option == 1) {
            // Opción de Threads
            long startTime = System.currentTimeMillis();
            Thread thread = new Thread(new Threads(laberinto, coordX, coordY));
            thread.start();

            // Esperamos a que termine el thread principal
            try {
                thread.join();
            } catch (InterruptedException e) {
                e.printStackTrace();
            }

            long endTime = System.currentTimeMillis();
            System.out.println("Tiempo de ejecución con Threads: " + (endTime - startTime) + " ms");

        } else if (option == 2) {
            // Opción de Fork/Join
            long startTime = System.currentTimeMillis();
            ForkJoinPool pool = null;
            try {
                pool = new ForkJoinPool();
                pool.invoke(new ForkJoin(laberinto, coordX, coordY));
            } finally {
                if (pool != null) {
                    pool.shutdown();
                }
            }
            long endTime = System.currentTimeMillis();
            System.out.println("Tiempo de ejecución con ForkJoin: " + (endTime - startTime) + " ms");
            
        } else {
            System.out.println("Opción no válida.");
        }

        scanner.close(); 

    }
}
