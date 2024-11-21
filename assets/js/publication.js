$('#nova-publicacao').on('submit', createPublication);
$('.likes-publication').on('click', likesPublication);

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

function likesPublication(event) {
  event.preventDefault();
  const element = $(event.target);
  const publicationId = element.closest('div').data('publication-id');

  $.ajax({
    url: `publications/${publicationId}/likes`,
    method: 'POST',
  }).done(function() {
    alert("Publicação curtida");
  }).fail(function() {
    alert("Erro ao curtir publicação");
  });
}