<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <title>Sabor USM</title>
  <link rel="icon" href="../imagenes/UTFSM.ico" type="image/x-icon">
  <link rel="stylesheet" href="../css/style2.css">
  <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.4.2/css/all.min.css"
    integrity="sha512-z3gLpd7yknf1YoNbCzqRKc4qyor8gaKU1qmn+CShxbuBusANI9QpRohGBreCFkKxLhei6S9CQXFEbbKuqLg0DA=="
    crossorigin="anonymous" referrerpolicy="no-referrer" />
</head>

<body>
  <div class="container" id="container">
    <div class="form-container sign-up-container">
      <form action="../php_bd/signup.php" method="post">
        <h1 style="margin-bottom: 20px;">Crear una Cuenta</h1>
        <input type="text" placeholder="Username" name="Username" />
        <input type="email" placeholder="Email" name="Email" onblur="validateEmail(this)" />
        <span class="error email-error">Error: Ingresa un correo válido</span>
        <input type="password" placeholder="Password" name="pswd1" onblur="validatePassword(this)" />
        <span class="error password-error">Error: La contraseña debe tener al menos 8 caracteres</span>
        <input type="password" placeholder="Confirm Password" name="pswd2" onblur="validateConfirmPassword(this)" />
        <span class="error cPassword-error">Error: Las contraseñas no coinciden</span>
        <button type="submit">Regístrate</button>
      </form>
    </div>
    <div class="form-container sign-in-container">
      <form action="../php_bd/login.php" method="post">
        <h1 style="margin-bottom: 20px;">Inicio Sesión</h1>
        <input type="text" placeholder="Username" name="Username" />
        <input type="password" placeholder="Password" name="Password" style="margin-bottom: 20px;" />
        <button type="submit" style="margin-top: 20px;">Iniciar Sesión</button>
      </form>
    </div>


    <div class="overlay-container">
      <div class="overlay">
        <div class="overlay-panel overlay-left">
          <h1>Bienvenido de Vuelta</h1>
          <p>Para mantenerte conectado con nosotros ingresa tus datos</p>
          <button class="ghost" id="signIn">Iniciar Sesión</button>
        </div>
        <div class="overlay-panel overlay-right">
          <h1>Hola, Amig@!</h1>
          <p>Registra tus datos para ingresar en nuestra página</p>
          <button class="ghost" id="signUp">Registrate</button>
        </div>
      </div>
    </div>
    <!-- Resto de tu código... -->
  </div>
  <script src="../js/app.js"></script>
  <script src="../js/validation.js"></script>
</body>
</html>