<template>
  <div>
    <div class="navb_headl navb bg-d">
        Bài viết mới
    </div>
    <div class="post_list">
      <ShopPagePost :key="idx" v-for="(p,idx) in posts" :item="p" />
    </div>
  </div>

</template>

<script>
  import ShopPagePost from './ShopPagePost';
  export default {
    name: 'ShopPagePostList',
    components:{
      ShopPagePost
    },
    data(){
      return {
        posts: [{
          title: 'jadior sz 18 hàng super',
          content: 'Mẫu jadior sz 18 hàng super e còn hai màu đỏ và vàng.. ce e nào hốt ko ạ ?\n' +
              'Xinh vãi 🥰',
          images:['https://via.placeholder.com/512','https://via.placeholder.com/512']
        },{
          title: 'jadior sz 18 hàng super',
          content: 'Mẫu jadior sz 18 hàng super e còn hai màu đỏ và vàng.. ce e nào hốt ko ạ ?\n' +
              'Xinh vãi 🥰',
          images:['https://via.placeholder.com/512','https://via.placeholder.com/512']
        }]
      }
    }
  };
</script>

<style scoped>
 .post_list_title{
   display: flex;
   font-weight: bold;
   color: #0d64c6;
   padding: 10px;
   font-size: 20px;
   line-height: 24px;
   box-shadow: 0 2px 2px 0 rgba(4, 0, 0, 0.15);
 }
  .post_list{
    margin: 0 4px;
  }
</style>