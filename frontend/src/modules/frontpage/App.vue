<template>
  <div>
    <header>

    </header>
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
  export default {};
</script>

<style scoped>
  .left_bar {
    width: 164px;
    flex-shrink: 0;
    position: relative !important;
    z-index: 1000;
    box-shadow: 0 0 0 rgba(12, 13, 14, 0.05);
    transition: box-shadow ease-in-out .1s, transform ease-in-out .1s;;
  }
</style>