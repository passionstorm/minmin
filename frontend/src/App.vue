<template>
    <div class="app">
        <transition>
            <router-view />
        </transition>
        <!-- setting drawer -->
        <!-- global snackbar -->
    </div>
</template>

<script>
  export default {
    name: 'App',
    components: {},
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
    created() {
    },
    methods: {},
  };
</script>
<style>
    .container {
        display: flex;
    }
</style>
