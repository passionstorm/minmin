<template>
  <div>
    <main class="main-container">
      <transition name="fade-in-right">
        <router-view/>
      </transition>
    </main>
    <footer>
    </footer>
  </div>
</template>

<script>
  export default {};
</script>

<style scoped>

</style>