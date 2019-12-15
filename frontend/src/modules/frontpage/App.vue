<template>
  <div>
    <section id="mainarea">
      <div class="left_bar"></div>
      <transition name="fade-in-right">
        <router-view/>
      </transition>
      <div></div>
    </section>
    <footer></footer>
  </div>
</template>

<script>
  import Icon from '../../widgets/Icon';

  export default {
    components: {
      Icon,
    },
  };
</script>

<style scoped>
  header {
  }




  .left_bar {
    width: 200px;
    flex-shrink: 0;
    position: relative !important;
    z-index: 1000;
    box-shadow: 0 0 0 rgba(12, 13, 14, 0.05);
    transition: box-shadow ease-in-out .1s, transform ease-in-out .1s;;
  }
</style>