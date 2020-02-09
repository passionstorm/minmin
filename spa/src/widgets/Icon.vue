<template>
  <i class="icon">
    <svg class="svg-icon" :class="iconClass">
      <use :xlink:href="'#' + name"></use>
    </svg>
  </i>
</template>

<script>
  import Icons from '../assets/icons'

  const defaultSize = {
    'sm': 'is-small',
    'm': 'is-medium',
    'l': 'is-large',
  }
  export default {
    name: 'icon',
    props: {
      name: {
        type: String,
        required: true,
        default: '',
        validator(val) {
          return Icons[val]
        },
      },
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
        }
      },
      classes() {
        return [
          defaultSize[this.size],
        ]
      },
      iconClass() {
        return 'icon-' + this.name
      },
    },
    created() {
      // console.log(this.loadIcon);

    },
  }
</script>

<style scoped>
  .svg-icon {
    width: 1.5rem;
    height: 1.5rem;
  }
</style>