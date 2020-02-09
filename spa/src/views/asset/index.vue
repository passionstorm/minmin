<template>
  <div>
    <div class="d-flex bd-highlight mb-2 justify-content-end">
      <a href="javascript:;" class="btn btn-d">Thêm Icon</a>
    </div>
    <div class="card">
      <div class="card-header">
        <div class="card-header-title">Icons</div>
      </div>
      <div class="card-body">
        <div class="d-flex flex-wrap">
          <div class="icon_item" v-for="i in icons">
            <icon :name="i"/>
            <span>{{i}}</span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
  import Icon from '../../widgets/Icon';
  export default {
    components: {
      Icon
    },
    data() {
      return {
        icons: ICON_LIST,
      };
    },
    created() {

    },
    mounted() {
    },
  };
</script>

<style scoped>
  .icon_item{
    text-align: center;
    width: 110px;
    margin: 15px 8px;
    display: flex;
    align-items: center;
    flex-direction: column;
  }

  .icon_item i {
    width: 30px;
    height: 40px;
    display: block;
    text-align: center;
    margin: 0 auto;
  }
</style>