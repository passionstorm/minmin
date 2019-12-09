<template>
  <div>

    <div class="sh_l">
      <ShopCard v-for="(i, idx) in listShop" :key="idx" :item="i"/>
    </div>
  </div>

</template>

<script>
  import ShopCard from './ShopCard';

  export default {
    components: {ShopCard},
    name: 'ShopListCard',
    data() {
      return {
        listShop: [
          {
            id: 0,
            name: 'Dana Pearl Hotel',
            addr: 'K1B/54 Le Dinh Duong',
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            attr: [{ic: '', name: ''}],
          },
          {
            id: 0,
            name: 'Dana Pearl Hotel',
            addr: 'K1B/54 Le Dinh Duong',
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            attr: [{ic: '', name: ''}],
          },
          {
            id: 0,
            name: 'Dana Pearl Hotel',
            addr: 'K1B/54 Le Dinh Duong',
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            attr: [{ic: '', name: ''}],
          },
        ],
      };
    },
  };
</script>

<style scoped>
.sh_l {
  display: flex;
  flex-wrap: wrap;
  flex-direction: column;
}
</style>