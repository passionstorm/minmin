<template>
  <div>
    <div style="margin-bottom: 10px">
      <a href="javacript:;" class="btn btn_img btn-outline-secondary" v-if="listImages.length < 5">
        <input id="image-upload" type="file" accept="image/gif,image/jpeg,image/png,image/bmp,image/jpg"
               @change="uploadImage($event)">
        <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 512 512">
          <path fill="#c9c8c8"
                d="M512 144v288c0 26.5-21.5 48-48 48H48c-26.5 0-48-21.5-48-48V144c0-26.5 21.5-48 48-48h88l12.3-32.9c7-18.7 24.9-31.1 44.9-31.1h125.5c20 0 37.9 12.4 44.9 31.1L376 96h88c26.5 0 48 21.5 48 48zM376 288c0-66.2-53.8-120-120-120s-120 53.8-120 120 53.8 120 120 120 120-53.8 120-120zm-32 0c0 48.5-39.5 88-88 88s-88-39.5-88-88 39.5-88 88-88 88 39.5 88 88z"/>
        </svg>
        <span>Thêm ảnh </span>
      </a>
      <span>Thêm được {{5 - listImages.length}} ảnh</span>
    </div>
    <div class="upload_area">
      <transition-group tag="ul" name="list" class="upload_list">
        <li class="upload_item" v-for="(src,idx) in listImages" :key="src">
          <span class="btn_del" @click="removeImage(idx)">×</span>
          <img class="preview img-thumbnail" :src="src"/>
        </li>
      </transition-group>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'UploadImage',
    model: {
      prop: 'modelValue',
      event: 'change',
    },
    props: {
      modelValue: {
        type: [Array],
        default: () => [],
      },
    },
    data() {
      return {
        listImages: [...this.modelValue],
      }
    },
    methods: {
      addImagePreview(e) {
        const file = e.target.files[0]
        this.listImages.push(URL.createObjectURL(file))
        this.$emit('change', this.listImages)
      },
      removeImage(index){
        this.listImages.splice(index, 1);
        this.$emit('change', this.listImages)
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
        ).then(res => {
            this.addImagePreview(event)
          },
        )
      },
    },
  }
</script>

<style scoped>
  .btn_del {
    position: absolute;
    top: 2px;
    right: 2px;
    color: #fff;
    background: #23282d;
    border: 2px solid #fff;
    border-radius: 50%;
    text-align: center;
    width: 25px;
    height: 25px;
    opacity: 0.5;
    font-size: 1.15rem;
    line-height: 1.35rem;
  }
  .btn_del:hover{
    opacity: 1;
  }

  .upload_list {
    display: flex;
    flex-wrap: wrap;
    margin-block-start: 0;
    padding-inline-start: 0;
    -webkit-transition:  opacity .5s;
    -moz-transition: opacity .5s;
    -ms-transition: opacity .5s;
    -o-transition:  opacity .5s;
   transition: opacity .5s;
  }

  .list-enter-active,
  .list-leave-active,
  .list-move {
    transition: 500ms cubic-bezier(0.59, 0.12, 0.34, 0.95);
    transition-property: opacity, transform;
  }

  .list-enter {
    opacity: 0;
    transform: translateX(50px) scaleY(0.5);
  }

  .list-enter-to {
    opacity: 1;
    transform: translateX(0) scaleY(1);
  }

  .list-leave-active {
    position: absolute;
  }

  .list-leave-to {
    opacity: 0;
    transform: scaleY(0);
    transform-origin: center top;
  }

  .preview {
    width: 100%;
    height: 100%;
    padding: 2px;
    -webkit-user-drag: none;
    -khtml-user-drag: none;
    -moz-user-drag: none;
    -o-user-drag: none;
    user-drag: none;
  }

  .upload_area {
    width: 100%;
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
    vertical-align: top;
  }

  .btn_upload {
    width: 100%;
    height: 100%;
    cursor: pointer;
    display: block;
  }
  .btn_img{
    position: relative;
    cursor: pointer;
  }
  .btn_img svg{
    width: 24px;
    cursor: pointer;
    margin-right: 2px;
  }

  .btn_img:hover svg{
    fill: #dddddd;
  }

  .btn_img span{
    cursor: pointer;
    vertical-align: middle;
  }

  .image-icon-drag {
    fill: #c9c8c8;
    height: 50px;
    width: 50px;
  }

  .upload_item {
    border-radius: 6px;
    box-sizing: border-box;
    width: 148px;
    height: 148px;
    margin: 0 1px 1px 0;
    position: relative;
    user-select: none;
    -webkit-user-select: none;
    cursor: pointer;
    vertical-align: top;
    transition: all .5s cubic-bezier(.55, 0, .1, 1);
    list-style: none;
  }

  #image-upload {
    position: absolute;
    top: 0;
    right: 0;
    left: 0;
    height: 100%;
    opacity: 0;

  }

</style>