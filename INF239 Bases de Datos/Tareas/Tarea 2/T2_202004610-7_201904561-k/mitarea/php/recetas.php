<?php

include "../php_bd/funciones.php";
$con = conexion();
$sql = $con->query("SELECT foto, nombre, instrucciones, tiempoPrep, calificacion FROM recetas");
$sql->execute();
$resultado = $sql->fetchAll(PDO::FETCH_ASSOC);



?>


<!DOCTYPE html>
<html lang="en">
<head>
  <link rel="stylesheet" href="../css/styleFav.css">
</head>
<body>
  <?php
  include("barra.php"); // Incluye la barra lateral de home.php
  ?>

  
  

  <section class="home-section">
    <!-- Agrega contenido específico de user.php aquí -->
  <div class="text">RECETAS</div>
    <div class="cards-wrapper">
    <?php foreach($resultado as $row) { ?>
    <div class="card">
      <?php
      $foto = $row['foto'];
      $nombre = $row['nombre'];
      $instr = $row['instrucciones'];
      $tiempo = $row['tiempoPrep'];
      $cali = $row['calificacion'];
      ?>
      <img src = <?php echo $foto;?>  width="600" height="400" class="card-img-top">
        <div class="card-body">
          <h2 class="card-title"><?php echo $nombre; ?></h2>
          <p> Instrucciones: <br> <?php echo $instr; ?> </p>
          <p> <br> Tiempo de Preparación: <?php echo $tiempo; ?> minuntos. </p>
          <h4> <br> Calificación: <?php echo $cali; ?><?php include("calificacion.php");?></h4>
          <br>
          
        </div>
    </div>
    <br>
    <?php } ?>
  </div>
    
  </section>

  <!-- Script y cierre de HTML -->
  <script  src="../js/scriptfav.js"></script>
</body>
</html>