$('#nova-publicacao').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.deslike-publication', deslikePublication);

$('#edit-publication').on('click', editPublication);
$('.delete-publication').on('click', deletePublication)

function createPublication(event) {
  event.preventDefault();
  $.ajax({
    url: "/publications",
    method: "POST",
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    window.location = "/home";
  }).fail(function() {
    alert("Erro ao criar uma publicação")
  });
}

function deslikePublication(event) {
  event.preventDefault();
  const element = $(event.target);
  const publicationId = element.closest('div').data('publication-id');

  $.ajax({
    url: `publications/${publicationId}/dislike`,
    method: 'POST',
  }).done(function() {
    const likesCount = element.next("span");
    const likes = parseInt(likesCount.text());

    likesCount.text(likes - 1);

    element.removeClass('deslike-publication');
    element.removeClass('text-danger');
    element.addClass('like-publication')
  }).fail(function() {
    alert("Erro ao curtir publicação");
  }).always(function() {
    element.prop('disabled', false);
  });

}

function likePublication(event) {
  event.preventDefault();
  const element = $(event.target);
  const publicationId = element.closest('div').data('publication-id');
  element.prop('disabled', true);
  $.ajax({
    url: `publications/${publicationId}/like`,
    method: 'POST',
  }).done(function() {
    const likesCount = element.next("span");
    const likes = parseInt(likesCount.text());

    likesCount.text(likes + 1);

    element.addClass('deslike-publication');
    element.addClass('text-danger');
    element.removeClass('like-publication')
  }).fail(function() {
    alert("Erro ao curtir publicação");
  }).always(function() {
    element.prop('disabled', false);
  });
}

function editPublication() {
  $(this).prop('disabled', true);
  const publicationId = $(this).data('publication-id');
  console.log(publicationId)
  $.ajax({
    url: `/publications/${publicationId}`,
    method: 'PUT',
    data: {
      title: $('#title').val(),
      content: $('#content').val(),
    }
  }).done(function() {
    alert("sucesso ao editar publicação")
  }).fail(function() {
    alert("Erro ao editar publicação")
  }).always(function() {
    $('#edit-publication').prop('disabled', false)
  });
}

function deletePublication(event) {
  event.preventDefault();
  const element = $(event.target);
  const publication = element.closest('div');
  const publicationId = publication.data('publication-id');
  element.prop('disabled', true);

  $.ajax({
    url: `publications/${publicationId}`,
    method: 'DELETE'
  }).done(function() {
    publication.fadeOut('slow', function() {
      $(this).remove();
    });
  }).fail(function() {
    alert("Erro ao excluir a publicação");
  });
}