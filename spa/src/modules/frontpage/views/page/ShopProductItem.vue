<template>
  <div>
    <div class="i_img">
      <div class="i_img_wrap">
        <div class="badge_img"></div>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'ShopProductItem',
    props: {
      item: [String, Object],
    },
  };
</script>

<style scoped>
  .i_img {
    margin-right: .75rem;
    width: 6.875rem;
    height: 6.875rem;
    background: rgba(0, 0, 0, .04);
  }
  .i_img_wrap{

  }

  .badge_img{
    top: 0;
    right: 0;
    position: absolute;
  }
</style>