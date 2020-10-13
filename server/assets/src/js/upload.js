import Dropzone from "dropzone"

Dropzone.autoDiscover = false

window.onload = () => {
  if (document.URL.match(/\/admin\/posts\/edit/)) {
    let elm = document.querySelector('textarea#image-dropzone')
    let imageDropzone = new Dropzone(elm, {
      paramName: 'image',
      url: '/admin/api/v1/images/',
      clickable: false,
      acceptedFiles: 'image/*'
    })
    imageDropzone.on('success', (f, response, _e) => {
      let sentence = elm.value
      let len = sentence.length
      let pos = elm.selectionStart
      let before = sentence.substr(0, pos)
      let word = `![${f.name}](${response.location})`
      let after = sentence.substr(pos, len)

      elm.value = before + word + after
    })
  }
}
