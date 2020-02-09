<template>
  <div class="container">
    <div id="mainarea">
      <transition name="fade-in-right">
        <router-view/>
      </transition>
    </div>
    <footer></footer>
  </div>
</template>

<script>
  export default {
  }
</script>

<style scoped>
</style>