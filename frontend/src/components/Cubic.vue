<template>
  <div class="cube_modal" v-show="show">
    <div class="cube_bg"></div>
    <div class="cube">
      <div class="sides">
        <div class="top"></div>
        <div class="right"></div>
        <div class="bottom"></div>
        <div class="left"></div>
        <div class="front"></div>
        <div class="back"></div>
      </div>
      <div class="text">Đang tải trang...</div>
    </div>
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

  .cube_modal .text {
    margin-top: 20px;
    font-size: 1rem;
    color: #fff;
    width: 100%;
    font-weight: 200;
    display: flex;
    text-align: center;
  }

  .cube,
  .cube * {
    position: absolute;
    width: 6rem;
    height: 6rem;
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
    transform: rotateX(90deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes top-animation {
    0% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(100px);
    }
    20% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(100px);
    }
    100% {
      opacity: 1;
      transform: rotateX(90deg) translateZ(100px);
    }
  }

  .cube .sides .bottom {
    animation: bottom-animation 3s ease infinite;
    animation-delay: 0ms;
    transform: rotateX(-90deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes bottom-animation {
    0% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(100px);
    }
    20% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(100px);
    }
    100% {
      opacity: 1;
      transform: rotateX(-90deg) translateZ(100px);
    }
  }

  .cube .sides .front {
    animation: front-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(0deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes front-animation {
    0% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(96px);
    }
    20% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateY(0deg) translateZ(96px);
    }
    100% {
      opacity: 1;
    }
  }

  .cube .sides .back {
    animation: back-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(-180deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes back-animation {
    0% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(96px);
    }
    20% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(96px);
    }
    100% {
      opacity: 1;
      transform: rotateY(-180deg) translateZ(96px);
    }
  }

  .cube .sides .left {
    animation: left-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(-90deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes left-animation {
    0% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(96px);
    }
    20% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(96px);
    }
    100% {
      opacity: 1;
      transform: rotateY(-90deg) translateZ(96px);
    }
  }

  .cube .sides .right {
    animation: right-animation 3s ease infinite;
    animation-delay: 100ms;
    transform: rotateY(90deg) translateZ(96px);
    animation-fill-mode: forwards;
    transform-origin: 50% 50%;
  }

  @keyframes right-animation {
    0% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(96px);
    }
    20% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(48px);
    }
    70% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(48px);
    }
    90% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(96px);
    }
    100% {
      opacity: 1;
      transform: rotateY(90deg) translateZ(96px);
    }
  }

</style>