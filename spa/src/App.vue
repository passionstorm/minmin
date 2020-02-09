<template>

  <div class="app">
    <transition>
      <router-view/>
    </transition>
    <!-- setting drawer -->
    <!-- global snackbar -->
  </div>
</template>

<script>
  // import Cubic from './components/Cubic';

  export default {
    name: 'App',
    components: {
      // Cubic,
    },
    data() {
      return {
        rightDrawer: false,
        snackbar: {
          show: false,
          text: '',
          color: '',
        },
      };
    },
    computed: {

    },
    created() {

    },
    methods: {},
  };
</script>
<style lang="scss">
  .container {
    display: flex;
  }
</style>
