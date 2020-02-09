
<template>
  <transition
      :enter-active-class="transition.enter"
      :leave-active-class="transition.leave">
    <div
        v-show="isActive"
        class="toast"
        :class="[type, position]"
        :aria-hidden="!isActive"
        role="alert">
      <div v-html="message"/>
    </div>
  </transition>
</template>

<script>
  import config from '../../utils/config_element'
  import NoticeMixin from '../mixins/notice.mixin'
  export default {
    name: 'Toast',
    mixins: [NoticeMixin],
    data() {
      return {
        newDuration: this.duration || config.defaultToastDuration
      }
    }
  }
</script>
<style scoped>
  .notices{
    padding: 0 2em;
  }
</style>