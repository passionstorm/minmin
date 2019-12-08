<template>
  <nav class="navbar_menu">
    <div class="navbar_list">
      <div class="navbar_item">
        <div class="navbar_icon" @click="sidebarClick">
          <icon name="menu"/>
        </div>
      </div>
      <div class="navbar_item">
        <div class="navbar_link">{{title}}</div>
      </div>
      <div class="navbar_item">
        <div class="navbar_link">

        </div>
      </div>
    </div>
  </nav>
</template>

<script>
  import Icon from '../widgets/Icon';

  export default {
    components: {
      Icon
    },
    data() {
      return {
        breadcrumbs: [],
      };
    },
    computed: {
      title() {
        return this.$route.meta.title.toUpperCase();
      },
      sidebarOpen() {
        return this.$store.state.app.sidebarOpen;
      },
    },
    watch: {
      '$route.path': function(newVal) {
        this.computeBreadcrumbs();
      },
    },
    created() {

    },
    methods: {
      sidebarClick() {
        this.$store.dispatch('toggleSidebar');
      },
      computeBreadcrumbs() {
        let breadcrumbs = [
          {
            text: 'Home',
            href: '/',
            disabled: false,
          },
        ];
        let appends = [];
        // appends = this.$route.matched.map(item => {
        //   return {
        //     text: item.meta.title || '',
        //     href: item.path || '/',
        //     disabled: item.path === this.$route.path,
        //   };
        // });
        this.breadcrumbs = breadcrumbs.concat(appends);
      },
    },
    created() {
      this.computeBreadcrumbs();
    },
  };
</script>
<style scoped>

  .navbar_menu {
    border-bottom: 1px solid #dee2e6;
    background-color: #fff;
    display: -ms-flexbox;
    display: flex;
    justify-content: center;
    -ms-flex-direction: column;
    flex-direction: column;
    padding-left: 0;
    margin-bottom: 0;
    list-style: none;
  }

  .navbar_list {
    display: -ms-flexbox;
    display: flex;
    align-items: center;
    justify-content: space-between;
  }

  .navbar_icon {
    position: relative;
    display: block;
    padding: .5rem 1rem;
    cursor: pointer;
  }

  .navbar_icon i {
    width: 24px;
    height: 24px;
  }

  .navbar_link {
    padding: 0 1rem;
    text-align: center;
    cursor: pointer;
  }
</style>
