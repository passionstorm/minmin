<template>
  <div class="comment_l">
    <div class="comment_i" v-for="item in items">
      <div class="comm_img">
        <a class="avtar">
          <img src='https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraightStrand&accessoriesType=Blank&hairColor=Red&facialHairType=BeardMagestic&facialHairColor=Red&clotheType=ShirtScoopNeck&clotheColor=PastelBlue&eyeType=Side&eyebrowType=Angry&mouthType=Tongue&skinColor=Brown'
          />
          <div class="rank">🐣</div>
        </a>
      </div>
      <div class="comm_area">
        <div class="b_cmn">
          <div class="usr_bg">
            <div class="h_cmn">
              <div class="usr">{{item.username}}</div>
              <div class="tip">2 tieng truoc</div>
            </div>
            <div class="cmn_tx">{{item.cmt}}</div>
          </div>
          <div class="stas">
            <span>👍 10</span>
          </div>
        </div>
        <div class="file_atth">
          <div class="img_ath" v-for="i in item.images">
            <img :src="i" alt="">
          </div>
        </div>
        <div class="cmn_btn">
        </div>
      </div>
    </div>
    <div class="u_cmt">
      <div class="tool"></div>
      <textarea rows="3" placeholder="Nhập tin nhắn" class="input u_in"/>
      <a href="javascript:;" class="btn btn-d">Gửi</a>
    </div>
  </div>
</template>

<script>
  export default {
    name: 'Comment',
    data() {
      return {
        items: [
          {
            username: 'Trần Hưng',
            cmt: 'Thực sự ban đầu k ấn tượng lắm . Nhưng nhận và sau 1 ngày sử dụng. Thích cực. Máy nhẹ . Pin ổn.chụp hình bao net bao đẹp. Tai nghe trong . Hoàn hảo cho giá này. Thanks shop.mà vẫn chưa biết kích hoạt bảo hành',
            avatar: 'https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraightStrand&accessoriesType=Blank&hairColor=Red&facialHairType=BeardMagestic&facialHairColor=Red&clotheType=ShirtScoopNeck&clotheColor=PastelBlue&eyeType=Side&eyebrowType=Angry&mouthType=Tongue&skinColor=Brown',
            images: [
              'https://via.placeholder.com/75',
              'https://via.placeholder.com/75',
              'https://via.placeholder.com/75',
            ],
          },
          {
            username: 'Trần Hưng',
            cmt: 'Thực sự ban đầu k ấn tượng lắm . Nhưng nhận và sau 1 ngày sử dụng. Thích cực. Máy nhẹ . Pin ổn.chụp hình bao net bao đẹp. Tai nghe trong . Hoàn hảo cho giá này. Thanks shop.mà vẫn chưa biết kích hoạt bảo hành',
            avatar: 'https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraightStrand&accessoriesType=Blank&hairColor=Red&facialHairType=BeardMagestic&facialHairColor=Red&clotheType=ShirtScoopNeck&clotheColor=PastelBlue&eyeType=Side&eyebrowType=Angry&mouthType=Tongue&skinColor=Brown',
          },
          {
            username: 'Trần Hưng',
            cmt: 'Thực sự ban đầu k ấn tượng lắm . Nhưng nhận và sau 1 ngày sử dụng. Thích cực. Máy nhẹ . Pin ổn.chụp hình bao net bao đẹp. Tai nghe trong . Hoàn hảo cho giá này. Thanks shop.mà vẫn chưa biết kích hoạt bảo hành',
            avatar: 'https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraightStrand&accessoriesType=Blank&hairColor=Red&facialHairType=BeardMagestic&facialHairColor=Red&clotheType=ShirtScoopNeck&clotheColor=PastelBlue&eyeType=Side&eyebrowType=Angry&mouthType=Tongue&skinColor=Brown',
          },
          {
            username: 'Trần Hưng',
            cmt: 'Thực sự ban đầu k ấn tượng lắm . Nhưng nhận và sau 1 ngày sử dụng. Thích cực. Máy nhẹ . Pin ổn.chụp hình bao net bao đẹp. Tai nghe trong . Hoàn hảo cho giá này. Thanks shop.mà vẫn chưa biết kích hoạt bảo hành',
            avatar: 'https://avataaars.io/?avatarStyle=Circle&topType=LongHairStraightStrand&accessoriesType=Blank&hairColor=Red&facialHairType=BeardMagestic&facialHairColor=Red&clotheType=ShirtScoopNeck&clotheColor=PastelBlue&eyeType=Side&eyebrowType=Angry&mouthType=Tongue&skinColor=Brown',
          },
        ],
      };
    },
  };
</script>

<style scoped>
  .u_cmt{
    display: flex;
    align-items: flex-start;
    margin-top: 8px;
  }
  .comment_l {

  }

  .rank {
    text-align: center;
  }

  .comm_area {
    display: flex;
    flex: 0 100%;
    flex-direction: column;
  }

  .comm_img {
    margin-right: 4px;
  }

  .usr_bg {
    background-color: #ebedf0;
    border-radius: 18px;
    line-height: 18px;
    margin-top: 0;
    min-width: 0;
    overflow: hidden;
    padding: 6px 12px 7px 12px
  }

  .h_cmn {
    display: flex;
    line-height: 14px;
    padding: 0 0 4px 0;
  }

  .u_in {
    height: auto;
    line-height: 16px;
    white-space: pre-wrap;
    word-wrap: break-word;
    margin-right: 8px;
  }

  .tip {
    border-left: 1px solid #ccc;
    padding-left: 4px;
    font-size: 12px;
    margin-left: 4px;
  }

  .usr {
    font-weight: bold;
    font-size: 12px;
  }

  .b_cmn {
    display: flex;
  }

  .stas {
    margin-top: 4px;
    display: flex;
    line-height: 16px;
    color: #8d949e;
    font-size: 11px;
    text-align: center;
    flex-direction: row;
  }

  .stas span {
    margin-left: 2px;
  }

  .file_atth {
    margin-top: 4px;
    display: flex;
  }

  .file_atth img {
    margin-left: 4px;
    width: 75px;
    height: 75px;
  }

  .cmn_tx {
  }

  .comment_i {
    display: flex;
    margin-top: 8px;
  }

  .avtar img {
    background-size: 100% 100%;
    width: 36px;
    height: 36px;
  }

  .cmn_act {

  }
</style>