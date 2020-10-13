import Dropzone from "dropzone"

Dropzone.autoDiscover = false

window.onload = () => {
  if (document.URL.match(/\/admin\/posts\/edit/)) {
    let imageDropzone = new Dropzone(document.querySelector('textarea#image-dropzone'), {
      paramName: 'image',
      url: '/admin/images/',
      clickable: false,
      acceptedFiles: 'image/*'
    })
  }
}
