<template>
  <div>
    <div class="upload_area">
      <ul class="upload_list">
        <li class="upload_item" v-for="src in listImages">
          <img class="preview" :src="src"/>
        </li>
      </ul>
      <div class="upload_item">
        <div class="input_image">
          <div>
            <svg class="image-icon-drag" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
              <path fill="currentColor"
                    d="M512 144v288c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V144c0-26.5 21.5-48 48-48h88l12.3-32.9c7-18.7 24.9-31.1 44.9-31.1h125.5c20 0 37.9 12.4 44.9 31.1L376 96h88c26.5 0 48 21.5 48 48zM376 288c0-66.2-53.8-120-120-120s-120 53.8-120 120 53.8 120 120 120 120-53.8 120-120zm-32 0c0 48.5-39.5 88-88 88s-88-39.5-88-88 39.5-88 88-88 88 39.5 88 88z"></path>
            </svg>
          </div>
          <div class="image-input">
            <label for="image-upload" class="btn_upload"></label>
          </div>
        </div>
      </div>
      <input id="image-upload" type="file" accept="image/gif,image/jpeg,image/png,image/bmp,image/jpg"
             @change="uploadImage($event)">

    </div>

  </div>
</template>

<script>
  export default {
    name: 'UploadImage',
    data() {
      return {
        listImages: [],
      }
    },
    methods: {
      addImagePreview(e) {
        const file = e.target.files[0]
        this.listImages.push(URL.createObjectURL(file))
      },
      uploadImage(event) {
        const URL = 'https://fineuploader.com/server/success.html'
        let data = new FormData()
        data.append('name', 'my-picture')
        data.append('file', event.target.files[0])
        let config = {
          header: {
            'Content-Type': 'image/png',
          },
        }
        this.$http.get(
          URL,
          data,
          config,
        ).then(
          response => {
            this.addImagePreview(event)
          },
        )
      },
    },
  }
</script>

<style scoped>
  .upload_list {
    display: inline-flex;
  }

  .preview {
    width: 100%;
    height: 100%;
  }

  .upload_area {
    display: flex;
    flex-wrap: wrap;
  }

  .image-input {
    overflow: hidden;
    opacity: 0;
    top: 0;
    left: 0;
    position: absolute;
    bottom: 0;
    width: 100%;
    height: 100%;
  }

  .input_image {
    display: flex;
    position: relative;
    justify-content: center;
    align-items: center;
    background-color: #fbfdff;
    border: 1px dashed #c0ccda;
    border-radius: 6px;
    box-sizing: border-box;
    width: 148px;
    height: 148px;
    cursor: pointer;
    text-align: center;
    line-height: 146px;
    vertical-align: top;
  }

  .btn_upload {
    width: 100%;
    height: 100%;
    cursor: pointer;
    display: block;
  }

  .image-icon-drag {
    fill: #c9c8c8;
    height: 50px;
    width: 50px;
  }

  .upload_item {
    background-color: #fbfdff;
    border: 1px dashed #c0ccda;
    border-radius: 6px;
    box-sizing: border-box;
    width: 148px;
    height: 148px;
    margin: 0 8px 8px 0;
    position: relative;
    cursor: pointer;
    line-height: 146px;
    vertical-align: top;
    transition: all .5s cubic-bezier(.55, 0, .1, 1);
    list-style: none;
  }

  #image-upload {
    display: none
  }
</style>