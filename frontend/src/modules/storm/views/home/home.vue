<template>
  <div class="container">
    <header>
      <div class="main_title">
        <img class="" src="../../img/logo.png" alt="">
        <h1>MinTooT</h1>
      </div>
    </header>
    <section>
      <editor-form/>
    </section>
  </div>
</template>

<script>
  import EditorForm from './ReviewForm'
  export default {
    name: 'home',
    components:{
      EditorForm
    },
    data() {
      return {}
    },
    methods: {
      onCreate() {

      },
    },
  }
</script>

<style scoped>
  .main_title {
    display: flex;
    flex-direction: column;
    align-items: center;
    margin: 1em 0 3em 0;
  }

  img {
    width: 125px;
    max-width: 100%;
    height: auto;
  }
</style>