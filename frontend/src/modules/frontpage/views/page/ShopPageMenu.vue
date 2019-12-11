<template>
  <div>
    <div class="tab_router">
      <router-link :to="{name: 'shop_page_info'}">Trang Chủ</router-link>
      <router-link :to="{name: 'shop_page_info'}">Sản phẩm</router-link>
    </div>
    <router-view />
  </div>
</template>

<script>
  import {TabItem, Tabs} from '../../../../widgets';

  export default {
    name: 'ShopPageMenu',
    components: {
      Tabs, TabItem,
    },
    data() {
      return {
        activeTab: 0,
      };
    },
  };
</script>

<style scoped>

</style>