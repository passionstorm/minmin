<template>
  <div class="post card-d">
    <h2 class="p_headline">{{item.title}}</h2>
    <div class="p_info">
      <div class="p_attr">
        <icon size="sm" name="eye"/>
        <span class="p_attr_v">200</span>
      </div>
      <div class="p_attr">
        <icon name="history" size="sm"/>
        <span class="p_attr_v">2 thang truoc</span>
      </div>
    </div>
    <div class="p_content">
      <p>{{item.content}}</p>
      <div class="img_area">
        <div class="ui_img" v-for="img in item.images">
          <img :src="img" alt="">
        </div>
      </div>
    </div>
    <div class="p_acts">
      <div class="act">
        <icon name="like"/>
        <span class="act_n">Thich</span>
      </div>
      <div class="act">
        <icon name="comment"/>
        <span class="act_n">Bình luận</span>
      </div>
    </div>
  </div>
</template>

<script>
  import Icon from '../../../../widgets/Icon';

  export default {
    name: 'ShopPagePost',
    props: {
      item: Object,
    },
    components: {
      Icon,
    },
  };
</script>

<style scoped>
  .post {
    padding: 10px;
    margin-bottom: 16px;
  }

  .p_info {
    border-bottom: 1px solid #e4e6e8;
  }

  .img_area {
    max-width: 512px;
  }

  .img_area {
    display: flex;
    justify-content: space-between;
    align-items: center;
  }

  .ui_img {
    width: 256px;
    display: flex;
    margin: 1px;
  }

  .p_acts {
    display: flex;
    border-top: 1px solid #e4e6e8;
    border-bottom: 1px solid #e4e6e8;
  }

  .act {
    padding: 4px 8px;
    display: flex;
    flex: 1 1 100%;
    cursor: pointer;
    user-select: none;
    justify-content: center;
    -webkit-user-select: none;
  }

  .act:first-child {
    border-right: 1px solid #e4e6e8;

  }

  .act_n {
    margin-left: 4px;
  }


  .p_attr {
    padding: 4px 8px;
    display: inline-flex;
    align-items: center;
  }

  .p_attr_v {
    font-size: 12px;
    margin-left: 4px;
  }

  .p_content {
    padding: 8px 0;
  }

  .p_headline {
    font-size: 1.3rem;
    font-weight: bold;
    overflow-wrap: break-word;
    word-wrap: break-word;
  }
</style>