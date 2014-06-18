$(function() {
    $('#crop_target').Jcrop({
      onSelect: setCoords,
      aspectRatio: 1/5
    });
});

function setCoords(c) {
  $('#x1').val(c.x);
  $('#y1').val(c.y);
  $('#x2').val(c.x2);
  $('#y2').val(c.y2);
}

$('#mainForm').submit(function(event) {
  event.preventDefault();
});  
