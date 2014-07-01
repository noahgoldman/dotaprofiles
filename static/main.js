$(function() {
  // Hide the other sections
  $("#upload").show()
  $("#make_images").hide()

  $('#crop_target').Jcrop({
    onSelect: setCoords,
    aspectRatio: 1/5
  });

  $('#mainForm').submit(function(event) {
    event.preventDefault();
  });

  $('#upload_form').submit(function(event) {
    event.preventDefault()

    var formData = new FormData()
    formData.append("picture", $("#original_picture").get(0).files[0])

    $.ajax({
      type: 'POST',
      url: '/upload',
      data: formData,
      cache: false,
      contentType: false,
      processData: false,

      success: function(data, textStatus, jqXHR) {
        showCropView(data)
      },
    });
  });
});

function setCoords(c) {
  $('#x1').val(c.x);
  $('#y1').val(c.y);
  $('#x2').val(c.x2);
  $('#y2').val(c.y2);
}

function showCropView(url) {
  $("#crop_target").attr("src", url).load()

  $("#upload").hide()
  $("#make_images").show()
}
