<template>
   <span class="icon" :class="classes" v-if="name" :style="styles">
        <i v-html="loadIcon"></i>
   </span>
</template>

<script>
  const defaultSize = {
    'sm': 'is-small',
    'm': 'is-medium',
    'l': 'is-large',
  };
  export default {
    name: 'icon',
    props: {
      name: String,
      size: {
        default: 'normal',
      },
      modifier: {
        default: null,
      },
      fixAlign: {
        default: true,
      },
      color: String,
    },
    components: {},
    computed: {
      styles() {
        return {
          color: this.color,
        };
      },
      classes() {
        return [
          defaultSize[this.size],
        ];
      },
      loadIcon() {
        return require(`../assets/images/icons/${this.name}.svg`);
      },
    },
    created() {
      // console.log(this.loadIcon);

    },
  };
</script>

<style scoped>
  .icon {
  }

  .icon svg {
    width: 100%;
    height: 100%;
  }

  .icon.is-large i svg {
    width: 3rem;
    height: 3rem;
  }
</style>