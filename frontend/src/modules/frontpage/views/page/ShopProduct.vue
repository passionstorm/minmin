<template>
  <div>
    <ShopProductItem :item="i" v-for="(i, idx) in items" :key="idx"/>
  </div>
</template>

<script>
  import ShopProductItem from './ShopProductItem';
  export default {
    name: 'ShopProduct',
    data() {
      return {
        items: [
          {},
        ],
      };
    },
    components: {
      ShopProductItem,
    },
  };
</script>

<style scoped>

</style>