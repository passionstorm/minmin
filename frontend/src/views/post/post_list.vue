<template>
  <div>
    <div class="d-flex mb-2">
      <div class="mr-auto"></div>
      <router-link class="btn btn-primary" :to="{name: 'post_edit'}">Them bai viet</router-link>
    </div>
    <div class="list_p" v-if="items.length">
      <div class="i_p" v-for="(i,idx) in items">
        <div class="img">
          <img :src="i.image_thumb" alt="">
        </div>
        <div class="con">
          <router-link class="ti" :to="{name: 'post_edit', params:{id: i.id}}">{{i.title}}</router-link>
          <span class="desc">{{i.updated_at}}</span>
          <div class="mb-auto"></div>
          <div style="display: flex">
            <div class="mr-auto"></div>
            <div class="mn">
              <a class="btn btn-d btn-s">
                <icon name="eye"/>
              </a>
              <a class="btn btn-danger btn-s" @click="del(i.id, idx)">
                <icon name="delete"/>
              </a>
            </div>
          </div>
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
  import {msg} from '../../utils/msg';
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
    methods: {
      del(id, idx) {
        this.$app.dialog.confirm({
          title: 'Yêu cầu xoá',
          message: 'Bạn có chắc muốn xoá không ?',
          onConfirm: () => {
            axios.delete('post/' + id).then(rs => {
              this.items.splice(idx, 1);
              msg.success('Đã xoá');
            }).catch(err => {
              msg.fail('Không xoá được!');
            });
          },
        });
      },
    },
    watch: {
      items(val) {
      },
    },
    mounted() {
      this.$store.dispatch('post/get_all');
    },
  };
</script>

<style scoped>
  .list_p {
    display: flex;
    flex-direction: column;
  }

  .list_p .i_p {
    display: flex;
    flex-direction: row;
    padding: 8px 16px 8px 0;
    border-bottom: 1px solid rgba(34, 36, 38, .15);
  }

  .list_p .ti {
    display: flex;
    font-size: 18px;
    font-weight: bold;
  }

  .list_p .con {
    display: flex;
    flex: 100%;
    margin-left: 8px;
    flex-direction: column;
  }

  .list_p .con a {

  }

  .list_p .mn {
    display: flex;
  }

  .list_p .mn .btn {
    margin: 0 4px;
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