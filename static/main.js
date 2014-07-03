$(function() {
  // Hide the other sections
  $("#make_images").hide()
  $("#display_images").hide()

  $('#crop_form').submit(function(event) {
    event.preventDefault();

    $.ajax({
      type: 'POST',
      url: '/make_images',
      data: $(this).serialize(),
      success: showOutputImages
    });      
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

      success: showCropView
    });
  });
});

function setCoords(c) {
  $('#x1').val(c.x);
  $('#y1').val(c.y);
  $('#x2').val(c.x2);
  $('#y2').val(c.y2);
}

function showCropView(data, textStatus, jqXHR) {
  $("#crop_target").attr("src", data.url);
  $("#ps_id").val(data.id)

  $('#crop_target').Jcrop({
    onSelect: setCoords,
    aspectRatio: 1/5
  });

  $("#upload").hide();
  $("#make_images").show();
  $("#display_images").hide();
}

function showOutputImages(data, textStatus, jqXHR) {
  $(".output_images").each(function(index, img) {
    $(img).attr("src", data[index]);
  });

  $("#upload").hide();
  $("#make_images").hide();
  $("#display_images").show();
}
