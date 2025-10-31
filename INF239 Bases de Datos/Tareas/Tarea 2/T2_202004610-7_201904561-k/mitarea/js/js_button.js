
// Obtén una referencia a los botones
const favoritosButton = document.getElementById("favoritosButton");
const editarButton = document.getElementById("editarButton");

// Agrega un controlador de clic al botón "Favoritos"
favoritosButton.addEventListener("click", function() {
  // Redirige a favoritos.php
  window.location.href = "favoritos.php";
});

// Agrega un controlador de clic al botón "Editar"
editarButton.addEventListener("click", function() {
  // Redirige a user_edit.php
  window.location.href = "../php/user_edit.php";
});

