<template>
  <v-notification
      v-bind="$options.propsData"
      @close="close"/>
</template>

<script>
  import config from '../../utils/config_element'
  import NoticeMixin from '../mixins/notice.mixin'

  export default {
    name: 'VNotificationNotice',
    mixins: [NoticeMixin],
    props: {
      indefinite: {
        type: Boolean,
        default: false
      }
    },
    data() {
      return {
        newDuration: this.duration || config.defaultNotificationDuration
      }
    }
  }
</script>