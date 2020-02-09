<template>
  <div class="">
    <div class="sh_page_info_basic">
      <div class="navb bg-d navb_headl">Thông tin cơ bản</div>
      <ul class="attr card-d">
        <li class="attr_i" v-for="attr in attrs">
          <dl>
            <dt>
              <icon size="m" :name="attr.ic"/>
              <span class="attr_name">{{attr.name}}:</span>
            </dt>
            <dd>
              <span class="attr_val">{{attr.val}}</span>
              <span class="attr_note" v-if="attr.note">({{attr.note}})</span>
            </dd>
          </dl>
        </li>
        <template v-if="isShowMore">
          <li class="attr_i" v-for="attr in attrsMore">
            <dl>
              <dt>
                <icon size="m" :name="attr.ic"/>
                <span class="attr_name">{{attr.name}}:</span>
              </dt>
              <dd>
                <span class="attr_val">{{attr.val}}</span>
                <span class="attr_note" v-if="attr.note">({{attr.note}})</span>
              </dd>
            </dl>
          </li>
        </template>
        <template v-else>
          <li class="has-text-link show_more" @click="isShowMore = true">Hiển thị thêm</li>
        </template>
      </ul>
    </div>
    <ShopPagePostList/>
  </div>
</template>

<script>
  import ShopPagePostList from './ShopPagePostList';
  import Icon from '../../../../widgets/Icon';

  export default {
    name: 'ShopPageInfo',
    components: {
      Icon, ShopPagePostList,
    },
    data() {
      return {
        isShowMore: false,
        attrs: [
          {ic: 'eye', name: 'Đánh giá', val: '4.7/5', note: '50 luot danh gia'},
          {ic: 'eye', name: 'Xếp hạng', val: 'Bạch kim'},
          {ic: 'time', name: 'Hôm nay', val: '9: 30-12: 30, 15: 30-20: 00'},
        ],
        attrsMore: [
          {ic: 'eye', name: 'Đã tham gia', val: '8 tháng', note: ''},
          {ic: 'reply_all', name: 'Tương tác', val: '100 lượt', note: ''},
          {ic: 'cart', name: 'Sản phẩm', val: '100k', note: ''},
        ],
      };
    },
  };
</script>

<style scoped>
  .attr {
    margin: 0 4px;
  }
  .sh_page_info_basic{
    margin: 8px 0;
  }

  .show_more {
    padding: 10px 0;
    cursor: pointer;
    text-align: center;
  }

  .attr_i {
    border-color: #f2f2f2;
    border-bottom: 1px solid #e5e5e5;
  }

  .attr_i dl {
    display: inline-flex;
  }

  .attr_val {
    font-size: 14px;
  }

  .attr_name {
    font-size: 14px;
  }

  .attr_note {
    font-size: 12px;
  }

  .attr_i dd {
    display: inline-flex;
    flex-direction: column;
    justify-content: center;
  }

  .attr_i dt {
    display: inline-flex;
    align-items: center;
    min-width: 8em;
  }


</style>