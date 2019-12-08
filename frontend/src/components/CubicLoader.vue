<template>
  <div class="loading_container" v-show="show">
    <div class="cube">
      <div class="sides">
        <div class="top"></div>
        <div class="right"></div>
        <div class="bottom"></div>
        <div class="left"></div>
        <div class="front"></div>
        <div class="back"></div>
      </div>
    </div>
    <div class="text">Đang đồng bộ dữ liệu...</div>
  </div>
</template>

<script>
  export default {
    name: 'cubic',
    props: ['show'],
    serverCacheKey: props => props.show,
  };
</script>

<style scoped>
  .loading_container{
    position: relative;
    justify-content: center;
    align-items: center;
    margin-top: 8px;
    display: flex;
  }
  .cube_bg {
    background-color: rgba(10, 10, 10, 0.86);
    bottom: 0;
    left: 0;
    position: absolute;
    right: 0;
    top: 0;
  }

  .cube_modal {
    align-items: center;
    display: flex;
    bottom: 0;
    left: 0;
    right: 0;
    top: 0;
    flex-direction: column;
    justify-content: center;
    overflow: hidden;
    position: fixed;
    z-index: 50;
  }

  .loading_container .text {
    font-size: 1rem;
    width: 100%;
    padding-left: 20px;
    display: flex;
  }

  .cube,
  .cube * {
    position: absolute;
    width: 1rem;
    height: 1rem;
  }

  .cube .sides * {
    box-sizing: border-box;
    border: 0.12rem solid #007bff;
    border-radius: 0.25rem;
    background: rgba(255, 255, 255, 0.1);
  }

  @keyframes rotate {
    0% {
      transform: rotateX(-37.5deg) rotateY(45deg);
    }
    50% {
      transform: rotateX(-37.5deg) rotateY(405deg);
    }
    100% {
      transform: rotateX(-37.5deg) rotateY(405deg);
    }
  }

  .sides {
    animation: rotate 3s ease infinite;
    animation-delay: 0.8s;
    transform-style: preserve-3d;
    transform: rotateX(-37.5deg) rotateY(45deg);
  }

  .cube .sides .top {
    animation: top-animation 3s ease infinite;
    animation-delay: 0ms;
    transform: rotateX(90deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes top-animation {
    0% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(16px);
    }
    100% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(16px);
    }
  }

  .cube .sides .bottom {
    animation: bottom-animation 3s ease infinite;
    animation-delay: 0ms;
    transform: rotateX(-90deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes bottom-animation {
    0% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(16px);
    }
    100% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(16px);
    }
  }

  .cube .sides .front {
    animation: front-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(0deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes front-animation {
    0% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(16px);
    }
    100% {
      opacity: 1;
    }
  }

  .cube .sides .back {
    animation: back-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(-180deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes back-animation {
    0% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(16px);
    }
    100% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(16px);
    }
  }

  .cube .sides .left {
    animation: left-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(-90deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes left-animation {
    0% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(16px);
    }
    100% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(16px);
    }
  }

  .cube .sides .right {
    animation: right-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(90deg) translateZ(16px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes right-animation {
    0% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(16px);
    }
    20% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(8px);
    }
    70% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(8px);
    }
    90% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(16px);
    }
    100% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(16px);
    }
  }

</style>