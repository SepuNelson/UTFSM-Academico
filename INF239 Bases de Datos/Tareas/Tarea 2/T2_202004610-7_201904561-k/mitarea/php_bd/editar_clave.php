<?php
    include "funciones.php";
    session_start();
    
    $clave = limpiar_cadena($_POST['pswd1']);   

    $check_user = conexion();
    $check_user = $check_user->query("UPDATE usuario SET `password` = '$clave' WHERE nombre = '{$_SESSION['nombre']}'");

    $_SESSION['password'] = $clave;
    $save_user = null;

    header("Location: ../php/user.php");

?>