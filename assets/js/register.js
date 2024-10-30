$('#form-cadastro').on('submit', createUser);

function createUser(event) {
  event.preventDefault();
  
  if ($('#password').val() != $('#confirmPassword').val()) {
    alert("As senhas não coincidem!")
    return;
  }

  $.ajax({
    url: '/users',
    method: "POST",
    data: {
      name: $("#name").val(),
      nick: $("#nick").val(),
      email: $("#email").val(),
      password: $("#password").val(),
    }
  })
  .done(function(){
    alert("user registered success")
  })
  .fail(function(error){
    console.log(error)
    alert("user not registered")
  })
}