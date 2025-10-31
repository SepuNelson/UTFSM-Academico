<?php
    include "funciones.php";
    
    $usuario = limpiar_cadena($_POST['Username']);
    $email = limpiar_cadena($_POST['Email']);
    $clave_1 = limpiar_cadena($_POST['pswd1']);
    $clave_2 = limpiar_cadena($_POST['pswd2']);

    /* VERIFICAR USUARIO */

    if($usuario == "" || $email == "" || $clave_1 == "" || $clave_2 == ""){
        echo '
        <div class="notification is-danger is-light">
            <strong>¡Ocurrio un error inesperado!</strong><br>
            No has llenado todos los campos que son obligatorios
        </div>';
        exit();
    }

    $check_user = conexion();
    $check_user = $check_user->query("SELECT nombre FROM usuario WHERE nombre = '$usuario'");

    if($check_user->rowCount()>0){
        echo '
        <div class="notification is-danger is-light">
            <strong>¡Ocurrio un error inesperado!</strong><br>
            El usuario ingresado ya esta registrado, por favor digite uno distinto.
        </div>';
        exit(); 
    }

    $check_user = null;

    /* Verificacion que las contraseñas coincidan */

    if($clave_1 != $clave_2){
        echo '
        <div class="notification is-danger is-light">
            <strong>¡Ocurrio un error inesperado!</strong><br>
            Las claves ingresadas no coinciden.
        </div>';
        exit(); 
    }else{ 
        $clave = $clave_1;
    }

    /* GUARDAR USUARIO */

    $save_user = conexion();

    $save_user = $save_user->query("INSERT INTO usuario (`nombre`, `mail`, `password`, `almuerzos`, `lastLog`) VALUES ('$usuario', '$email', '$clave', 0, NOW())");


    if($save_user->rowCount() == 1){
        $mensaje = "Te registraste exitosamente";
        header("Location: ../php/inicio.php?mensaje=" . urlencode($mensaje));
        exit();
    }else{
        echo '
        <div class="notification is-danger is-light">
            <strong>¡Ocurrio un error inesperado!</strong><br>
            No se pudo registrar el usuario, intente nuevamente.
        </div>';
    }
    $save_user = null;

?>