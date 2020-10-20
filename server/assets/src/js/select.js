import $ from "jquery"
import "select2"

$(document).ready(() => {
  $('.select2').select2({
    tags: true,
  });
});
