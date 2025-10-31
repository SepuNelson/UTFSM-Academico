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
        <form action="../php_bd/editar_usuario.php" method="post">
          <div class="input-with-button">
            <input type="text" placeholder="Username" name="Username" style="margin-bottom: 20px;">
            <button class="button-2 primary" formaction="../php_bd/editar_usuario.php">Editar</button>
          </div>
        </form>

        <form action="../php_bd/editar_email.php" method="post">
          <div class="input-with-button">
            <input type="email" placeholder="Email" name="Email" onblur="validateEmail(this)" style="margin-bottom: 20px;" />
            <button class="button-2 primary" formaction="../php_bd/editar_email.php">Editar</button>
          </div>
        </form>

        <form action="../php_bd/editar_clave.php" method="post">
          <div class="input-with-button">
            <input type="password" placeholder="Password" name="pswd1" onblur="validatePassword(this)" style="margin-bottom: 20px;" />
            <button class="button-2 primary" formaction="../php_bd/editar_clave.php">Editar</button>
          </div>
        </form>

        <form action="../php_bd/borrar_usuario.php">
          <div class="input-with-button">
            <button class="button-2 primary" formaction="../php_bd/borrar_usuario.php">Borrar</button>
          </div>
        </form>
        </div>
      </div>
    </div>
  </section>
</body>
</html>