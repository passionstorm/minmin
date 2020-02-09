<template>
  <div>
    <div class="head_area">
      <div class="a_search">
        <icon class="ic_search" name="search"/>
        <input class="">
      </div>
      <div class="a_search_ad">
        <select class="select">
          <option value="Đà nẵng"></option>
        </select>
      </div>
    </div>
    <div class="sh_l">
      <ShopCard v-for="(i, idx) in listShop" :key="idx" :item="i"/>
    </div>
  </div>
</template>

<script>
  import Icon from '../../../../widgets/Icon';
  import ShopCard from './ShopCard';

  export default {
    components: {ShopCard, Icon,},
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
            desc: {
              title: 'Nếu bạn đang tìm kiếm 1 thứ gì đó, Mintoot.com sẽ giúp bạn tìm thứ đó!',
              content: 'Số lượng cửa hàng và thông tin tài sản hàng đầu trên toàn quốc! Ngoài ra còn có một loạt các tài sản với các điều kiện đặc biệt như vật nuôi, dụng cụ và nhà thiết kế! Dịch vụ Wi-Fi miễn phí!',
            },
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
  .a_search_ad {
    display: flex;
    align-items: center;
  }

  .head_area {
    padding: 80px 0 10px 0;
    justify-content: center;
  }

  .a_search {
    display: flex;
    background: #fff;
    border: 1px solid #dfe1e5;
    box-shadow: none;
    border-radius: 24px;
    z-index: 3;
    height: 46px;
    margin: 0 auto;
    max-width: 550px;
    padding: 4px 8px 4px 0;
    align-items: center;

  }

  .head_area {
    padding: 10px;
    background-color: #ffff;
    border-bottom: 1px solid #d5d5d5;
  }

  .ic_search {
    margin: 10px;
    transition: 0.5s;
  }

  .ic_search svg {
    width: 20px !important;
    height: 20px !important;
  }

  .a_search:hover {
    box-shadow: 0 4px 4px 0 rgba(0, 0, 0, 0.2);
  }

  .a_search input {
    background-color: transparent;
    border: none;
    color: rgba(0, 0, 0, .87);
    word-wrap: break-word;
    outline: none;
    display: flex;
    flex: 1;
    -webkit-tap-highlight-color: transparent;
    font: 16px arial, sans-serif;
    line-height: 46px;
    height: 46px;
  }



  .sh_l {
    display: flex;
    flex-wrap: wrap;
    flex-direction: column;
  }
</style>