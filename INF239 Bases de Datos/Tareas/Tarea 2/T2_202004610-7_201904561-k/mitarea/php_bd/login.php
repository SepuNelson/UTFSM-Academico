<?php

    include "funciones.php";
    session_start();

    //===================================//
    //         V A R I A B L E S         //
    //===================================//

    $usuario = limpiar_cadena($_POST['Username']);
    $clave = limpiar_cadena($_POST['Password']);

    //===================================//
    //      V E R I F I C A C I Ó N      //
    //===================================//

    $check_user2 = conexion();
    $check_user = $check_user2->query("SELECT * FROM usuario WHERE nombre = '$usuario'");

    if($check_user->rowCount()==1){
        $check_user = $check_user->fetch();
    
        if($check_user['nombre'] == $usuario && $check_user['password'] == $clave){

            $check_user2->query("UPDATE usuario SET lastLog = NOW() WHERE nombre = '" . $check_user['nombre'] . "'");
                
            $_SESSION['id'] = $check_user['idUsuario']; 
            $_SESSION['nombre'] = $check_user['nombre'];
            $_SESSION['mail'] = $check_user['mail'];
            $_SESSION['password'] = $check_user['password'];
            $_SESSION['almuerzos'] =  $check_user['almuerzos'];

            $check_user2->query("UPDATE usuario SET lastLog = NOW() WHERE nombre = '" . $check_user['nombre'] . "'");
            $check_user2 = $check_user2->query("SELECT * FROM usuario WHERE nombre = '$usuario'");
            $check_user2 = $check_user2->fetch();

            $_SESSION['lastLog'] =  $check_user2['lastLog'];


            header("Location: ../php/home.php");
            exit();
            
        }else{
            header("Location: ../php/inicio.php");
            exit();
        }
    
    } 
    else{
        header("Location: ../php/inicio.php");
            exit();
    }
    $check_user = null;
?>