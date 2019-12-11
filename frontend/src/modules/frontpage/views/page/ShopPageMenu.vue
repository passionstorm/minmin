<template>
  <div>
    <ul class="tab_router">
      <li>
        <router-link class="tab_item" :to="{name: 'shop_page_info'}">Trang Chủ</router-link>
      </li>
      <li>
        <router-link class="tab_item" :to="{name: 'shop_page'}">Sản phẩm</router-link>
      </li>
    </ul>
    <router-view/>
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
  .tab_router {
    display: flex;
    font-weight: normal;
    background-color: #f5f5f5;
    box-shadow: 0 2px 2px 0 rgba(4, 0, 0, 0.15);
  }

  .tab_item {
    display: inline-flex;
    cursor: pointer;
    color: #333;
    padding: 12px 15px 8px;
    font-size: 1.1rem;
    line-height: 2.2rem;
  }

  .tab_item.router-link-exact-active {
    border-bottom: 3px solid #3578e5;
    color: #3578e5;
  }
</style>