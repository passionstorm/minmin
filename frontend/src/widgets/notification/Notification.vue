<template>
  <transition name="fade">
    <article
        v-show="isActive"
        class="notification"
        :class="[type, position]">
      <button
          v-if="closable"
          class="delete"
          type="button"
          @click="close"
          :aria-label="ariaCloseLabel"
      />
      <div class="media">
        <!--        <div v-if="icon && hasIcon" class="media-left">-->
        <!--          <b-icon-->
        <!--              :icon="icon"-->
        <!--              :pack="iconPack"-->
        <!--              both-->
        <!--              size="is-large"-->
        <!--              aria-hidden/>-->
        <!--        </div>-->
        <div class="media-content">
          <p
              class="text"
              v-if="message"
              v-html="message"/>
          <slot v-else/>
        </div>
      </div>
    </article>
  </transition>
</template>

<script>
  import MessageMixin from '../mixins/message.mixin'

  export default {
    name: 'VNotification',
    mixins: [MessageMixin],
    props: {
      position: String,
      ariaCloseLabel: String
    }
  }
</script>