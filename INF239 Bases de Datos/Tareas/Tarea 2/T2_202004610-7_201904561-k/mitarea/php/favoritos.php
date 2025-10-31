<!DOCTYPE html>
<html lang="en">
<head>
  <link rel="stylesheet" href="../css/style_card.css" />
</head>
<body>
  <?php
  include("barra.php");

include "../php_bd/funciones.php";
$con = conexion();
$sql = $con->query("SELECT PP.nombre, PP.foto, SD.calificacionUsuario, SD.reseñaUsuario FROM
                          (SELECT idReceta, nombre, foto FROM recetas) PP,
                          (SELECT idReceta, calificacionUsuario, reseñaUsuario FROM almuerzo WHERE favorito = 1 AND idUsuario = '{$_SESSION['id']}') SD
                    WHERE PP.idReceta = SD.idReceta");
$sql->execute();
$resultado = $sql->fetchAll(PDO::FETCH_ASSOC);

?>


  <section class="home-section">
    <!-- Agrega contenido específico de user.php aquí -->
    <div class="text">FAVORITOS</div>
    <div class="cards-wrapper">
        <?php
        foreach ($resultado as $row) { 
        $nombre = $row['nombre'];
        $foto = $row['foto'] ?>

        <h3> <?php echo $nombre;?> <br></h3>
        <img src = <?php echo $foto;?>  width="500" height="350" class="card-img-top">

        <?php
        }
        ?>
    </div>  
  </section>
  <script src="../js/js_button.js"></script>
</body>
</html>

