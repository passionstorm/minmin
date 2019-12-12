<template>
  <div>
    <header>
    </header>
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
    <section id="mainarea">
      <div class="left_bar"></div>
      <transition name="fade-in-right">
        <router-view/>
      </transition>
      <div></div>
    </section>
    <footer></footer>
  </div>
</template>

<script>
  import Icon from '../../widgets/Icon';

  export default {
    components: {
      Icon,
    },
  };
</script>

<style scoped>
  header {
  }

  .a_search_ad {
    display: flex;
    align-items: center;
  }

  .head_area {
    margin: 20px 0 10px 0;
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


  .left_bar {
    width: 200px;
    flex-shrink: 0;
    position: relative !important;
    z-index: 1000;
    box-shadow: 0 0 0 rgba(12, 13, 14, 0.05);
    transition: box-shadow ease-in-out .1s, transform ease-in-out .1s;;
  }
</style>