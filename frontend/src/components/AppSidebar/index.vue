<template>
  <aside class="sidenav" @mouseover="mouseOver" :class="hover">
    <ul class="nav nav-pills nav-sidebar flex-column">
      <slide-item v-for="(item,index) in listSlides" :data="item" :key="index" :items="item.children" :to="item.path"></slide-item>
    </ul>
  </aside>
</template>

<script>
  import {mapGetters} from 'vuex';
  import SlideItem from './SlideItem';

  export default {
    components: {
      SlideItem,
    },
    data() {
      return {
        basePath: this.$route.path,
        listSlides: [],
        hover: '',
      };
    },
    computed: {
      ...mapGetters([
        'sidebarOpen',
        'permission_routes',
      ]),
    },
    methods: {
      mouseOver() {

      },
    },
    created() {
      // console.log(this.permission_routes.filter(e => e.type === 'bar'))
      this.listSlides = this.permission_routes.filter(e => e.type === 'bar');
    },
  };
</script>

<style scoped>

</style>
