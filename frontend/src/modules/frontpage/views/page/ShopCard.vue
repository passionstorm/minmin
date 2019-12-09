<template>
  <div class="sh_card">
    <div class="im">
      <div class="im_thumb">
        <img :src="item.thumb"/>
      </div>
      <div class="im_chil">
        <img :src="item.thumb"/>
        <img :src="item.thumb"/>
      </div>

    </div>
    <div class="ct">
      <div class="t">{{item.name}}</div>
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
      <div class="attr">
        <div class="attr_l">
          <div class="attr_i">
            <div class="attr_i_name">
              Đỗ xe miễn phí
            </div>
          </div>
        </div>
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
  .sh_card {
  }

  .im_thumb {
    display: flex;
  }
  .im_chil {
    display: flex;
    flex-direction: row;
  }

  .sh_card .im .im_chil img {
    width: 100px;
    height: 100px;
  }

  .sh_card {
    display: flex;
    border-radius: 8px;
    overflow: hidden;
    border: 1px solid rgba(0, 0, 0, .12);
    box-sizing: border-box;
    margin: 8px;
  }

  .sh_card .im img {
    width: 100%;
    height: auto;
  }

  .sh_card .ct {
    display: flex;
    flex-direction: column;
    padding: 10px;
  }

  .review .p_star {
    color: rgba(0, 0, 0, .54);
    font-size: 12px;
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
    background-size: 14px 13px;
    height: 13px;
    top: 1px;
    width: 69px;
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACoAAAAmCAMAAACF3/kSAAAARVBMVEUAAADncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvncRvQdmh2AAAAFnRSTlMAEDAg71CgQJDg8HBvYMDPv4/fr3+fs8YwdQAAANlJREFUeF6Nk9uOgCAMBS0g4hVve/7/UzdrE2NXqJ3neSmHad4QqDEyo7WqEdGqDhiMZgfA2dQEINnUHkBvMj3+8BZ1vNTRoi6XuhhMAkOWqZjZMhUTLVMxpcHagG8CfyC/4ov1fuYTOufjFrcr4i7/Dm2osZHxulAKgkrXrZXRppc5NTUcJE5LRdJpqUiSloqkV1OReC0VyailIlnUVCSkp4LDH2Awq6mE6TFc1FLJPJHLnI0y1XZH/nO5bW2q0P2PLlWmOkhOksuDOb5HchZ/15hd6YD9Mdgvr0crAhbtroEAAAAASUVORK5CYII=);
  }

  .rv_rate {
    font-size: 12px;
    line-height: 20px;
    cursor: pointer;
  }

  ._star {
    background-image: url(data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAACoAAAAmCAQAAAAYCMGrAAAA+klEQVR4AcWV4cbtMBBFF0MIVUopoVSrhDDv/3gf/RFRpzdNOty1HiBO99mzeYWgCMZMKCPGrCgrxiSUhCkDeukxJKCXAUMiehkxw6FZhxEzmp0x4kCzByYISqlYdal0supS6WrVpdLEK0YSamJiJOPY0c/uOG4s6CcXfuKJaJcRzyNCQJsNiF1sRTR1hP11NNJ8RCrONOPRf+r7J+TZgQ5CNfMOYvW/2YxDqzqA/57+gVY9eiakrnyZEGXDsaE3p/4JScwPX3rtnZATDxnPWT7X16XAHaH8HWNrlxJD9TyGti5tCM84zpZe+RxNjeX9tZqLaGoMxN/P/wHP5Vw+8ZxnEQAAAABJRU5ErkJggg==);
    background-repeat: repeat-x;
    display: inline-block;
    overflow: hidden;
    position: relative;
    background-size: 14px 13px;
    height: 13px;
    top: 1px;
    width: 69px;
  }

  .attr {
    display: flex;
    flex: 1 0 auto;
    flex-wrap: wrap;
    align-items: center;
  }

  .n_rate {
    font-size: 12px;
  }

  .attr_l {
    color: #757575;
    overflow: hidden;
    word-break: break-word;
    font-size: 12px;
    line-height: 16px;
  }
</style>