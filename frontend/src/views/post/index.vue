<template>
  <div>
    <div class="list">
      <div class="list-item" v-for="i in data">
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
  </div>
</template>

<script>
  import {Check, Icon, Modal, Radio, VSelect, VTable, VTableColumn} from '../../widgets';
  import {default as PostModel, ENTITY} from '../../store/models/PostModel';

  export default {
    components: {Modal, Icon, Check, Radio, VTable, VTableColumn, VSelect},
    data() {
      return {
        currentPage: 1,
        perPage: 15,
        data: [],
        selected: {},
        columns: [
          {field: 'title', label: 'Title', searchable: true},
        ],
        checkedRows: [],
        entity: ENTITY,
        editModal: false,
      };
    },
    async mounted() {
      let res = await PostModel.api().fetch();
      this.data = res.response.data;
      console.log(res.response.data);
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