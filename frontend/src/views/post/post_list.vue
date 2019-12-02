<template>
  <div>
    <div class="list" v-if="items.length">
      <div class="list-item" v-for="i in items">
        <div class="content">
          <router-link :to="{name: 'post_edit', params:{id: i.id}}">{{i.title}}</router-link>
          <span class="desc">{{i.updated_at}}</span>
        </div>
        <div class="mr-auto">
        </div>
        <div class="img">
          <img :src="i.image_thumb" alt="">
          <icon name="more_vert"></icon>
        </div>
      </div>
    </div>
    <div v-else>
      <h2>
        Hiện tại không có bài viết nào
      </h2>
      <router-link class="btn btn-primary" :to="{name: 'post_edit'}">Thêm bài viết</router-link>
    </div>
  </div>
</template>

<script>
  import {mapState} from 'vuex';
  import {Check, Icon, Modal, Radio, VSelect, VTable, VTableColumn} from '../../widgets';

  export default {
    components: {Modal, Icon, Check, Radio, VTable, VTableColumn, VSelect},
    data() {
      return {
        columns: [
          {field: 'title', label: 'Title', searchable: true},
        ],
      };
    },
    computed: {
      ...mapState('post', ['items']),
    },
    watch:{
      items(val){

        // console.log(val);
      }
    },
    mounted() {
      this.$store.dispatch('post/get_all');
    },
  };
</script>

<style scoped>
  .list-item {
    display: flex;
    flex-direction: row;
  }

  .list-item .content {
    display: flex;
    flex-direction: column;
  }

  .list-item .content a {
    font-size: 24px;
    font-weight: bold;
  }

  .img {
    position: relative;
  }

  .img .icon {
    position: absolute;
    top: 5px;
    right: 0;
  }
</style>