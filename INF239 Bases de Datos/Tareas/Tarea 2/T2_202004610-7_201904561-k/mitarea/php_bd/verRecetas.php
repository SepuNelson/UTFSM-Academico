<?php

    include "funciones.php";
    session_start();

    //===================================//
    //      V E R I F I C A C I Ó N      //
    //===================================//

    $check_receta = conexion();
    $check_receta = $check_receta->query("SELECT * FROM recetas WHERE nombre = 'Arroz con Pollo'");

 
        $check_receta = $check_receta->fetch();
                    
            $_SESSION['nombreR'] = $check_receta['nombre'];
            $_SESSION['fotoR'] = $check_receta['foto'];


            header("Location: ../php/home.php");
            exit();
            

    

    $check_receta = null;
?>