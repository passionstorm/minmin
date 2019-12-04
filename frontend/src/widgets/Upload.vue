<template>
  <div ref="myVueDropzone" id="dropzone" :options="dropzoneOptions"></div>
</template>

<script>
  // import vue2Dropzone from 'vue2-dropzone'
  // import 'vue2-dropzone/dist/vue2Dropzone.min.css'
  export default {
    name: 'app',
    components: {
      // vueDropzone: vue2Dropzone
    },
    data: function () {
      return {
        dropzoneOptions: {
          // url: 'http://127.0.0.1:9000/api/upload',
          thumbnailWidth: 150,
          maxFilesize: 0.5,
          headers: { "My-Awesome-Header": "header value" }
        }
      }
    }
  }
</script>