<?php
    include "funciones.php";
    session_start();
    
    $email = limpiar_cadena($_POST['Email']);

    $check_user = conexion();
    $check_user = $check_user->query("UPDATE usuario SET mail = '$email' WHERE nombre = '{$_SESSION['nombre']}'");
    $_SESSION['mail'] = $email;
    $save_user = null;

    header("Location: ../php/user.php");

?>