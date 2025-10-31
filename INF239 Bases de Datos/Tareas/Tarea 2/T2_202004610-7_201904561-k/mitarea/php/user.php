<!DOCTYPE html>
<html lang="en">
<head>
  <link rel="icon" href="../imagenes/UTFSM.ico" type="image/x-icon">
  <link rel="stylesheet" href="../css/style_card.css" />
</head>
<body>
  <?php
  include("barra.php");
  ?>

  <section class="home-section">
    <div class="text"></div>
    <div class="container">
      <div class="card">
        <img src="../imagenes/profile.png" />
        <div>
          <h2><?php echo $_SESSION['nombre']; ?></h2>
          <h3><?php echo $_SESSION['mail']; ?></h3>
          <ul class="stats">
            <li>
              <var><?php echo $_SESSION['almuerzos']; ?></var>
              <label>Almuerzos</label>
            </li><br>
            <li>
              <var><?php echo $_SESSION['lastLog']; ?></var>
              <label>Último Inicio de Sesión</label>
            </li>
          </ul>
          <nav class="buttons">
            <button id="favoritosButton" class="primary">Favoritos</button>
            <button id="editarButton">Editar</button>
          </nav>
        </div>
      </div>
    </div>
  </section>
  <script src="../js/js_button.js"></script>
</body>
</html>