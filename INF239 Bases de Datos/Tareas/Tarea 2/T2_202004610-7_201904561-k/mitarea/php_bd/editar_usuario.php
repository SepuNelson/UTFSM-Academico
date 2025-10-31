<?php
    include "funciones.php";
    session_start();
    
    $usuario = limpiar_cadena($_POST['Username']);    

    $check_user = conexion();
    $check_user = $check_user->query("UPDATE usuario SET nombre = '$usuario' WHERE nombre = '{$_SESSION['nombre']}'");
    $_SESSION['nombre'] = $usuario;
    $save_user = null;

    header("Location: ../php/user.php");

?>