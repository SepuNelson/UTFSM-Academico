<?php
    include "funciones.php";
    session_start();

    $check_user = conexion();
    $check_user = $check_user->query("DELETE FROM usuario WHERE nombre = '{$_SESSION['nombre']}'");
    $save_user = null;

    header("Location: ../php/inicio.php");

?>