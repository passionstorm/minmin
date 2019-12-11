<template>
  <main>
    <div class="shop_page_header card-d">
      <div class="shop_name">
        <h2 class="shop_name_title">ABCXYZ.vn</h2>
        <p class="shop_name_account">ngutaihean</p>
      </div>
    </div>
    <div class="active_status">
      Đang online
    </div>
    <ShopPageMenu/>
  </main>
</template>

<script>
  import ShopPageInfo from './ShopPageInfo';
  import ShopPageMenu from './ShopPageMenu';

  export default {
    name: 'shop_page',
    components: {
      ShopPageInfo,
      ShopPageMenu,
    },
  };
</script>

<style scoped>
  .main {
  }

  .shop_page_header {
    display: flex;
    padding: 8px;
    background-color: #6c757d;
    color: #fff;
    flex-direction: column;
  }

  .active_status {
    background-color: #fff;
    border-bottom: 1px solid #ccd0d5;
    border-top: 1px solid #ccd0d5;
    padding: 4px 0;
    text-align: center;
  }

  .shop_name_account {
    font-size: 12px;
    margin-bottom: 0;
  }

  .shop_name_title {
    font-size: 24px;
    line-height: 28px;
  }

  .shop_name {

  }
</style>