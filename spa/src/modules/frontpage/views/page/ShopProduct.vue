<template>
  <div>
    <a href="javascript:;" class="btn btn-d" @click="isCreate = !isCreate">Create product</a>
    <div v-show="isCreate">
      <CreateProduct />
    </div>
    <ShopProductItem :item="i" v-for="(i, idx) in items" :key="idx"/>
  </div>
</template>

<script>
  import CreateProduct from '../../pages/CreateProduct';
  import ShopProductItem from './ShopProductItem';
  export default {
    name: 'ShopProduct',
    data() {
      return {
        isCreate: true,
        items: [
          {},
        ],
      };
    },
    components: {
      ShopProductItem,CreateProduct,
    },
  };
</script>

<style scoped>

</style>