<template>
  <li v-if="isAlone()" class="nav-item">
    <router-link :to="{name: it.name}" class="nav-link">
      <icon :name="it.meta.icon ? it.meta.icon : 'circle'" class="nav-icon"/>
      <p>{{it.meta.title}}</p>
    </router-link>
  </li>
  <li class="nav-item treeview" v-else>
    <a href="javascript:;" class="nav-link" @click="toggleMenu">
      <icon :name="it.meta.icon ? it.meta.icon : 'circle'" class="nav-icon"/>
      <p>
        {{it.meta.title}}
        <icon :name="menuOpen ? 'arrow_down' : 'arrow_right'" class="right"/>
      </p>
    </a>
    <ul class="nav nav-treeview" v-show="menuOpen" :class="it.path">
      <slide-item v-for="(item,index) in childItems" :data="item" :key="index" :to="it.path"/>
    </ul>
  </li>
</template>

<script>
  import path from 'path';
  import {isExternal} from '../../utils/validate';
  import Icon from '../../widgets/Icon';

  export default {
    name: 'slide-item',
    components: {
      Icon,
    },
    props: {
      isHeader: {
        type: Boolean,
        default: false,
      },
      data: {
        type: Object,
        default: null,
      },
      badge: {
        type: Object,
        default() {
          return {};
        },
      },
      items: {
        type: Array,
        default() {
          return [];
        },
      },
      to: {
        type: String,
        default: '',
      }
    },
    data() {
      return {
        hasChild: this.items.length > 0,
        it: this.data,
        basePath: this.to,
        menuOpen: this.data.path === this.$route.matched[0].path,
        childItems: this.items.filter(item => !item.hidden),
      };
    },
    computed: {},
    methods: {
      toggleMenu() {
        this.menuOpen = !this.menuOpen;
      },
      resolvePath(routePath) {
        if (isExternal(routePath)) {
          return routePath;
        }
        if (isExternal(this.basePath)) {
          return this.basePath;
        }
        return path.resolve(this.basePath, routePath);
      },
      isAlone() {
        const children = this.items.filter(item => {
          return !item.hidden;
        });
        // When there is only one child router, the child router is displayed by default
        if (children.length === 1) {
          this.it = children[0];
          return true;
        }
        // Show parent if there are no child router to display
        return children.length === 0;
      },
    },
    created() {

    },
  };
</script>
<style scoped>
  .nav-sidebar .nav-link p {
    display: inline-block;
    margin: 0;
    transition: margin-right 1s linear, opacity 1s ease, visibility 1s ease;
  }

  .nav-sidebar > .nav-item .nav-icon {
    font-size: 1.2rem;
    margin-right: .2rem;
    text-align: center;
    width: 20px;
  }

  .nav-link {
    position: relative;
    display: block;
    padding: 0.5rem 1rem;
    white-space: nowrap;
  }

  .nav-icon {
    color: #fff;
  }

  .nav-link > p > .right {
    color: #fff;
    position: absolute;
    font-weight: 900;
    width: 1.5rem;
    right: 1rem;
    top: .7rem;
    transition: all 1s ease;
  }

  .nav-sidebar .nav-item > .nav-link {
    margin-bottom: .2rem;
  }

  .nav-sidebar .nav-treeview {
    list-style: none;
  }

  .sidebar_close .nav-sidebar .nav-link p {
    display: none;
  }

  .sidebar_close .nav-link {
    padding: 0.5rem 0.5rem;
  }

  .nav-sidebar .nav-link p {
    display: inline-block;
    color: #fff;
    margin: 0;
  }
</style>