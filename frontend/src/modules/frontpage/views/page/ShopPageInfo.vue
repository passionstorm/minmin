<template>
  <div class="content_header">
    <ul class="attr card-d">
      <li class="attr_i" v-for="attr in attrs">
        <dl>
          <dt>
            <icon size="m" name="eye"/>
            <span class="attr_name">{{attr.name}}:</span>
          </dt>
          <dd>
            <span class="attr_val">{{attr.val}}</span>
            <span class="attr_note" v-if="attr.note">({{attr.note}})</span>
          </dd>
        </dl>
      </li>
    </ul>
  </div>
</template>

<script>
  import Icon from '../../../../widgets/Icon';
  export default {
    name: 'ShopPageInfo',
    components:{
      Icon
    },
    data() {
      return {
        attrs: [
          {ic: '', name: 'Đánh giá', val: '4.7/5', note: '50 luot danh gia'},
          {ic: '', name: 'Đã tham gia', val: '8 tháng', note: ''},
          {ic: '', name: 'Bình luận', val: '100 lượt', note: ''},
          {ic: '', name: 'Xếp hạng', val: '100', note: 'Bình luận gần nhất cách 5 ngày'},
          {ic: '', name: 'Sản phẩm', val: '100k', note: ''},
        ],
      };
    },
  };
</script>

<style scoped>
  .attr{
    margin-top: 8px;
  }
  .attr_i{
    border-color: #f2f2f2;
    border-bottom: 1px solid #e5e5e5;
  }
  .attr_i dl{
    display: inline-flex;
  }

  .attr_val{
    font-size: 14px;
  }

  .attr_name{
    font-size: 14px;
  }

  .attr_note{
    font-size: 12px;
  }

  .attr_i dd{
    display: inline-flex;
    flex-direction: column;
    justify-content: center;
  }
  .attr_i dt{
    display: inline-flex;
    align-items: center;
    min-width: 8em;
  }


</style>