<template>
  <div
      class="app_container"
      :class="{'sidebar_close': !sidebarOpen}"
  >
    <app-sidebar/>
    <div class="main-container">
      <app-header>
        <component :is="slotHeader"></component>
      </app-header>

      <section class="content app_main">
        <cubic-loader :show="isLoading"/>
        <main class="container-fluid">
          <transition name="fade-in-right">
            <router-view/>
          </transition>
        </main>
      </section>
      <app-footer>
        <component :is="slotFooter"></component>
      </app-footer>
    </div>
    <div
        class="sidebar_overlay"
        @click="sidebarOverlayClick"
    ></div>
  </div>
</template>

<script>
  import AppSidebar from '../AppSidebar';
  import AppHeader from '../AppHeader';
  import AppFooter from '../AppFooter';
  import AppMain from '../AppMain';
  import {bus} from '../../main';
  import CubicLoader from '../CubicLoader';

  export default {
    components: {
      AppSidebar,
      AppHeader,
      AppFooter,
      AppMain,
      CubicLoader,
    },
    data() {
      return {
        showDrawer: true,
        slotFooter: null,
        slotHeader: null,
        msg: null,
        routeName: null,
      };
    },
    computed: {
      isLoading() {
        return this.$store.state.app.isLoading;
      },
      sidebarOpen() {
        return this.$store.state.app.sidebarOpen;
      },

    },
    watch: {
      $route(to, from) {
        this.hideLoading()
      },
    },
    methods: {
      showLoading() {
        this.$store.commit('loading', true);
      },
      hideLoading() {
        this.$store.commit('loading', false);
      },
      sidebarOverlayClick() {
        this.$store.dispatch('toggleSidebar');
      },
      createLoading() {
        this.routeName = this.$router.currentRoute.name;
        axios.interceptors.request.use((config) => {
          this.showLoading();
          return config;
        }, (error) => {
          this.hideLoading();
          return Promise.reject(error);
        });
        axios.interceptors.response.use((response) => {
          this.hideLoading();
          return response;
        }, (error) => {
          this.hideLoading();
          return Promise.reject(error);
        });
      },
    },
    created() {
      this.createLoading();
      bus.$on('msg', (data) => {
        this.$app.toast.open({
          message: data.msg,
          type: data.type,
          position: 'is-bottom',
        });
      });
    },
    beforeDestroy() {
      bus.$off('msg');
    },
  };
</script>
<style scoped>
</style>
