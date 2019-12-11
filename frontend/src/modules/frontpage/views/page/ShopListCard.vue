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
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            desc: {title: 'Nếu bạn đang tìm kiếm 1 thứ gì đó, Mintoot.com sẽ giúp bạn tìm thứ đó!', content: 'Số lượng cửa hàng và thông tin tài sản hàng đầu trên toàn quốc! Ngoài ra còn có một loạt các tài sản với các điều kiện đặc biệt như vật nuôi, dụng cụ và nhà thiết kế! Dịch vụ Wi-Fi miễn phí!'},
            attr: {loc: 'K1B/54 Le Dinh Duong ', time: '8:30 - 24:00'},

          },
          {
            id: 0,
            name: 'Dana Pearl Hotel',
            addr: 'K1B/54 Le Dinh Duong',
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            attr: {loc: '293 Nguyễn lương ', time: '8:30 - 24:00'},
          },
          {
            id: 0,
            name: 'Dana Pearl Hotel',
            addr: 'K1B/54 Le Dinh Duong',
            star: '',
            count_review: 10,
            thumb: 'https://via.placeholder.com/200',
            attr: {loc: '293 Nguyễn lương ', time: '8:30 - 24:00'},
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