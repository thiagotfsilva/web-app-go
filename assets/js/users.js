$('#unfollow').on('click', unfollow);
$('#follow').on('click', follow);
$('#edit-user').on('submit', editUser);

function unfollow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/unfollow`,
    method: 'POST'
  }).done(function() {
    window.location = `/users/${userId}`;
  }).fail(function() {
    alert("ops...erro ao parara de seguir")
    $('#unfollow').prop('disabled', false)
  })
}

function follow() {
  const userId = $(this).data('user-id');
  $(this).prop('disabled', true);

  $.ajax({
    url: `/users/${userId}/follow`,
    method: 'POST'
  }).done(function() {
    window.location = `/users/${userId}`;
  }).fail(function() {
    alert("ops...erro ao seguir")
    $('#follow').prop('disabled', false)
  })
}

function editUser(event) {
  event.preventDefault();
  $.ajax({
    url: "/edit-user",
    method: "PUT",
    data: {
      name: $("#name").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
    }
  }).done(function() {
    alert("Sucesso, usuário atualizado")
   /*  window.location = "/profile" */
  }).fail(function() {
    alert("Erro ao atualizar, usuário atualizado")
  })
}