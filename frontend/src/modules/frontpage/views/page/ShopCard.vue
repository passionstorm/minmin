<template>
  <div class="shop_i">
    <div class="layout_head bg-d">
      <div class="layout_head_title">
        <h2 class="main_title">
          <router-link :to="{name: 'shop_page'}" class="t">{{item.name}}</router-link>
        </h2>
      </div>
    </div>
    <div class="sh_card">
      <div class="im">
        <div class="im_thumb">
          <img alt="" :src="item.thumb"/>
        </div>
        <div class="im_chil">
          <img alt="" :src="item.thumb"/>
          <img alt="" :src="item.thumb"/>
          <img alt="" :src="item.thumb"/>
        </div>

      </div>
      <div class="ct">
        <!--        <div class="review">-->
        <!--          <span class="rv_rate">3.5</span>-->
        <!--          <span class="_star">-->
        <!--            <span style="width:52px"></span>-->
        <!--          </span>-->
        <!--          <span class="n_rate">({{item.count_review}})</span>-->
        <!--        </div>-->
        <div class="review_emoji">
          <div class="r_emo" v-for="r in raters" v-if="r.val">
            <span class="r_ic">{{r.ic}}</span>
            <span class="r_val">{{r.val}}</span>
          </div>
        </div>
        <ul class="attr">
          <li class="attr_l">
            <dl class="attr_i">
              <dt>
                <icon name="location" size="m"/>
              </dt>
              <dd>
                {{item.attr.loc}}
              </dd>
            </dl>
          </li>
          <li class="attr_l">
            <dl class="attr_i">
              <dt>
                <icon name="time" size="m"/>
              </dt>
              <dd>
                {{item.attr.time}}
                <span style="padding-left: 8px; color: green">Đang mở cửa</span>
              </dd>
            </dl>
          </li>
        </ul>
      </div>

    </div>
    <dl class="box_intro" v-if="item.desc">
      <dt>{{item.desc.title}}</dt>
      <dd>{{item.desc.content}}</dd>
    </dl>
  </div>
</template>

<script>
  import Icon from '../../../../widgets/Icon';

  export default {
    props: {
      item: Object,
    },
    data(){
      return{
        raters: [
          {ic: '😡', val: '50'},
          {ic: '🙂', val: '60'},
          {ic: '😃', val: '10'},
          {ic: '😂', val: ''},
          {ic: '🤩', val: '3'},
        ]
      }
    },
    components: {
      Icon,
    },
    name: 'ShopCard',
  };
</script>

<style scoped>
  .shop_i {
    margin: 8px 0;
    box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  }

  .review_emoji {
    display: flex;
    flex-wrap: wrap;
    margin-bottom: 8px;
  }

  .r_emo {
    display: flex;
    margin: 0 4px;
  }

  .r_ic {

  }

  .r_val {

  }

  .attr_i dt {
    min-width: 16px;
  }

  .attr_i dt .icon svg {
    width: 1.2rem !important;
    height: 1.2rem;
  }

  .attr_i dd {
    display: flex;
    align-items: center;
  }

  .layout_head {
    display: flex;
    padding: 10px;
    background-color: #E0E0E0;
  }

  .main_title a {
    color: #007bff;
    font-weight: normal;
    font-size: 1.25rem;
    line-height: 1.5rem;
    letter-spacing: 0;
  }

  .main_title a:hover {
    text-decoration: underline;
  }

  .box_intro {
    margin-top: 2px;
    background-color: #f5f5f5;
    flex-direction: column;
    padding: 10px;
  }

  .im_thumb {
    display: flex;
  }

  .im_chil {
    margin-top: 3px;
    display: flex;
    flex-direction: row;
    justify-content: space-between;
  }

  .sh_card .im .im_chil img {
    width: 48px;
    height: 48px;
  }

  .attr_i dt {
    flex-shrink: 0;
    font-weight: bold;
  }

  .sh_card {
    display: flex;
    overflow: hidden;
    margin-top: 2px;
    border: 1px solid rgba(0, 0, 0, .12);
    box-sizing: border-box;
  }

  .sh_card .im img {
    width: 150px;
    height: 150px;
  }

  .sh_card .ct {
    display: flex;
    flex-direction: column;
    flex: 1;
    background-color: #fff;
    padding: 10px 4px;
  }

  .review {
    font-size: 1rem;
  }

  .review .p_star {
    color: rgba(0, 0, 0, .54);

  }

  .sh_card .ct .t {
    color: #222;
  }

  .sh_card .im {
    display: flex;
    flex-direction: column;
  }

  ._star span {
    background-repeat: repeat-x;
    display: block;
    background-size: 20px 19px;
    height: 19px;
    top: 1px;
    width: 100px;
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACoAAAAmCAMAAACF3/kSAAAARVBMVEUAAADncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvQdmh2AAAAFnRSTlMAEDAg71CgQJDg8HBvYMDPv4/fr3+fs8YwdQAAANlJREFUeF6Nk9uOgCAMBS0g4hVve/7/UzdrE2NXqJ3neSmHad4QqDEyo7WqEdGqDhiMZgfA2dQEINnUHkBvMj3+8BZ1vNTRoi6XuhhMAkOWqZjZMhUTLVMxpcHagG8CfyC/4ov1fuYTOufjFrcr4i7/Dm2osZHxulAKgkrXrZXRppc5NTUcJE5LRdJpqUiSloqkV1OReC0VyailIlnUVCSkp4LDH2Awq6mE6TFc1FLJPJHLnI0y1XZH/nO5bW2q0P2PLlWmOkhOksuDOb5HchZ/15hd6YD9Mdgvr0crAhbtroEAAAAASUVORK5CYII=);
  }

  .rv_rate {
    cursor: pointer;
  }

  ._star {
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACoAAAAmCAQAAAAYCMGrAAAA+klEQVR4AcWV4cbtMBBFF0MIVUopoVSrhDDv/3gf/RFRpzdNOty1HiBO99mzeYWgCMZMKCPGrCgrxiSUhCkDeukxJKCXAUMiehkxw6FZhxEzmp0x4kCzByYISqlYdal0supS6WrVpdLEK0YSamJiJOPY0c/uOG4s6CcXfuKJaJcRzyNCQJsNiF1sRTR1hP11NNJ8RCrONOPRf+r7J+TZgQ5CNfMOYvW/2YxDqzqA/57+gVY9eiakrnyZEGXDsaE3p/4JScwPX3rtnZATDxnPWT7X16XAHaH8HWNrlxJD9TyGti5tCM84zpZe+RxNjeX9tZqLaGoMxN/P/wHP5Vw+8ZxnEQAAAABJRU5ErkJggg==);
    background-repeat: repeat-x;
    display: inline-block;
    overflow: hidden;
    position: relative;
    background-size: 20px 19px;
    height: 19px;
    top: 1px;
    width: 100px;
  }

  .attr {
    margin-top: 4px;
    font-size: 12px;
    line-height: 1.2;
  }

  .attr_l {
    color: #757575;
    overflow: hidden;
    word-break: break-word;
    margin: 0 0 8px;
  }
</style>