<template>
  <div class="shop_i">
    <div class="layout_head">
      <div class="layout_head_title">
        <h2 class="main_title">
          <a class="t">{{item.name}}</a>
        </h2>
      </div>
    </div>
    <div class="sh_card">
      <div class="im">
        <div class="im_thumb">
          <img :src="item.thumb"/>
        </div>
        <div class="im_chil">
          <img :src="item.thumb"/>
          <img :src="item.thumb"/>
          <img :src="item.thumb"/>
        </div>

      </div>
      <div class="ct">
        <div class="review">
          <span class="rv_rate">3.5</span>
          <span class="_star">
          <span style="width:52px"></span>
        </span>
          <span class="n_rate">({{item.count_review}})</span>
        </div>
        <div class="addr">
          {{item.addr}}
        </div>
        <ul class="attr">
          <li class="attr_l">
            <div class="attr_i">
              <div class="attr_i_name">
                Đỗ xe miễn phí
              </div>
            </div>
          </li>
          <li class="attr_l">
            <dl class="attr_i">
              <dt>
                <i class="fa fa-subway" aria-hidden="true"></i><span>Giờ làm việc</span>
              </dt>
              <dd>
                Đỗ xe miễn phí
              </dd>
            </dl>
          </li>
        </ul>
      </div>
    </div>
  </div>
</template>

<script>
  export default {
    props: {
      item: Object,
    },
    name: 'ShopCard',
  };
</script>

<style scoped>
  .shop_i {
    margin: 8px 0;
  }

  .layout_head {
    display: flex;
    padding: 10px;
    background-color: #E0E0E0;
  }

  .main_title {
    color: #007bff;
    font-weight: bold;
    font-size: 1.5rem;
    line-height: 2rem;
    letter-spacing: 0;
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

  .sh_card {
    display: flex;
    border-radius: 8px;
    overflow: hidden;
    border: 1px solid rgba(0, 0, 0, .12);
    box-sizing: border-box;
    margin: 4px;
  }

  .sh_card .im img {
    width: 150px;
    height: 150px;
  }

  .sh_card .ct {
    display: flex;
    flex-direction: column;
    padding: 10px;
  }

  .review {
    font-size: 1.5rem;
  }

  .review .p_star {
    color: rgba(0, 0, 0, .54);

  }

  .sh_card .ct .t {
    color: #222;
    font-size: 16px;
    line-height: 20px;
  }

  .sh_card .im {
    display: flex;
    flex-direction: column;
  }

  ._star span {
    background-repeat: repeat-x;
    display: block;
    background-size: 24px 23px;
    height: 23px;
    top: 1px;
    width: 120px;
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
    background-size: 24px 23px;
    height: 23px;
    top: 1px;
    width: 120px;
  }

  .attr {
    font-size: 1.2rem;
    line-height: 1.6rem;
  }

  .attr_l {
    margin: 0 0 5px;
  }

  .n_rate {

  }

  .attr_l {
    color: #757575;
    overflow: hidden;
    word-break: break-word;
    font-size: 12px;
    line-height: 16px;
  }
</style>