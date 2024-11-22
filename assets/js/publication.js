$('#nova-publicacao').on('submit', createPublication);

$(document).on('click', '.like-publication', likePublication);
$(document).on('click', '.deslike-publication', deslikePublication);

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