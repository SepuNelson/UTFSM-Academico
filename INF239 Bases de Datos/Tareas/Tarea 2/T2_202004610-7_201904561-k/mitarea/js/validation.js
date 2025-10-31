// Función para validar el email
function validateEmail(input) {
  var email = input.value;
  var emailError = input.parentNode.querySelector(".email-error");

  if (email.match(/^[A-Za-z0-9._%-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,4}$/)) {
    emailError.classList.remove("active");
  } else {
    emailError.classList.add("active");
  }
}

// Función para validar la contraseña
function validatePassword(input) {
  var password = input.value;
  var passwordError = input.parentNode.querySelector(".password-error");

  if (password.length >= 8) {
    passwordError.classList.remove("active");
  } else {
    passwordError.classList.add("active");
  }
}

// Función para confirmar contraseña
function validateConfirmPassword(input) {
  var confirmPassword = input.value;
  var password = document.getElementById("pswd1").value;
  var confirmPasswordError = input.parentNode.querySelector(".cPassword-error");

  if (confirmPassword === password) {
    confirmPasswordError.classList.remove("active");
  } else {
    confirmPasswordError.classList.add("active");
  }
}
