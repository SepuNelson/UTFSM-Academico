<?php
session_start();
?>

<!DOCTYPE html>
<html lang="en" >

<head>
  <meta charset="UTF-8">
  <title>Sabor USM</title>
  <link rel="icon" href="../imagenes/UTFSM.ico" type="image/x-icon">
  <link rel="stylesheet" href="../css/style.css">
  <link rel='stylesheet' href='https://unpkg.com/boxicons@2.1.4/css/boxicons.min.css'>

</head>

<body>
  <!-- partial:index.partial.html -->
  <div class="sidebar">
    <div class="logo-details">
      <a href="home.php">
        <i class="bx bx-home-alt-2 icon"></i>
      </a>
      <div class="logo_name">Sabor USM</div>
      <i class='bx bx-menu' id="btn"></i>
    </div>
    <ul class="nav-list"> 

      <!-- BUSCADOR -->

      <li>
        <a href="#">
          <i class='bx bx-search'></i>
        </a>
        <input type="text" placeholder="Search...">
        <span class="tooltip">Search</span>
      </li>

      <!-- FILTROS -->

      <li>
        <a href="filtros.php">
          <i class='bx bx-filter'></i>
          <span class="links_name">Filtros</span>
        </a>
        <span class="tooltip">Filtros</span>
      </li>

      <!-- USUARIO -->

      <li>
        <a href="user.php">
          <i class='bx bx-user'></i>
          <span class="links_name">Perfíl</span>
        </a>
        <span class="tooltip">Perfíl</span>
      </li>

      <!-- RECETAS -->
      
      <li>
        <a href="recetas.php">
          <i class='bx bx-fork'></i>
          <span class="links_name">Recetas</span>
        </a>
        <span class="tooltip">Recetas</span>
      </li>

      <!-- VOTACION -->

      <li>
        <a href="votacion.php">
          <i class='bx bx-bar-chart-alt-2'></i>
          <span class="links_name">Votación Semanal</span>
        </a>
        <span class="tooltip">Votación Semanal</span>
      </li>
  
      <li class="profile">
        <div class="profile-details">
          <img src="../imagenes/profile.png" alt="profileImg">
          <div class="name_job">
            <div class="name"><?php echo $_SESSION['nombre']; ?></div>
          </div>
        </div>
        <a href="../php/inicio.php">
          <i class='bx bx-log-out' id="log_out"></i>
        </a>
      </li>
    </ul>
  </div>

  <!-- partial -->
  <script  src="../js/script_menu.js"></script>
</body>
</html>