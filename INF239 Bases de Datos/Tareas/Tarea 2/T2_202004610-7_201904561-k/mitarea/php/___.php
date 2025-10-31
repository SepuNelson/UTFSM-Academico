<?php
  include("barra.php");?>

<!DOCTYPE html>
<html lang="en" >
<head>
  <meta charset="UTF-8">
  <title>Gato</title>
  <link rel="stylesheet" href="../css/styleCarrusel.css">
</head>

<body>
<?php


include "../php_bd/funciones.php";
$con = conexion();
$sql = $con->query("SELECT PP.nombre, PP.foto, SD.calificacionUsuario, SD.reseñaUsuario FROM
                          (SELECT idReceta, nombre, foto FROM recetas) PP,
                          (SELECT idReceta, calificacionUsuario, reseñaUsuario FROM almuerzo WHERE favorito = 1 AND idUsuario = 2) SD
                    WHERE PP.idReceta = SD.idReceta");
$sql->execute();
$resultado = $sql->fetchAll(PDO::FETCH_ASSOC);
?>
<section class="home-section">
  <div class="text">Favoritos</div>
  
  <?php
      foreach ($resultado as $row) { 
        $nombre = $row['nombre'];
        $foto = $row['foto'] ?>
    <main>
  
    
    <!-- Swiper -->
    <div class="swiper mySwiper">
      
      <div class="swiper-wrapper">
        
        
        
        <div class="swiper-slide">
          
          <div class="text"> <?php echo $nombre ?> </div>
          <img src=<?php echo $foto ?> width="500" height="350" class="card-img-top">
            <div class="content">
              <div class="title">
                <h1>
                </h1>
              </div>
            </div>
          </div>
         
        </div>
        
        
  <?php
  }
  ?>
        
        
      </div>
      
        
    </main>
   
</section>

<!-- partial -->
  <script  src="../js/scriptCarrusel.js"></script>

</body>
</html>
