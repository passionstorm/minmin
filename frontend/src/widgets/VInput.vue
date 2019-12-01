<template>
    <div class="control" :class="rootClasses">
      <input
          v-if="type !== 'textarea'"
          ref="input"
          class="input"
          :class="[inputClasses, customClass]"
          :type="newType"
          :autocomplete="newAutocomplete"
          :maxlength="maxlength"
          :value="computedValue"
          v-bind="$attrs"
          @input="onInput"
          @blur="onBlur"
          @focus="onFocus">
      <transition
          v-else
          enter-active-class="fadeInDown"
          leave-active-class="fadeOut">
        <textarea
            :rows="textAreaRow"
            ref="textarea"
            class="textarea"
            :class="[inputClasses, customClass]"
            :maxlength="maxlength"
            :value="computedValue"
            v-bind="$attrs"
            @input="onInput"
            @blur="onBlur"
            @focus="onFocus"></textarea>
      </transition>
      <icon class="is-left" v-if="icon" :name="icon"/>
      <small
          v-if="maxlength && hasCounter && type !== 'number'"
          class="help counter"
          :class="{ 'is-invisible': !isFocused }">
        {{ valueLength }} / {{ maxlength }}
      </small>
    </div>
</template>

<script>
  import Icon from './Icon'
  import config from '../utils/config_element'
  import FormElementMixin from './mixins/form.mixin'
  export default {
    name: 'VInput',
    components: {Icon,},
    mixins: [FormElementMixin],
    inheritAttrs: false,
    props: {
      value: [Number, String],
      type: {
        type: String,
        default: 'text'
      },
      vid: String,
      name: String,
      rules: String,
      icon: String,
      validMode: {
        type: String,
        default: 'lazy'
      },
      passwordReveal: Boolean,
      hasCounter: {
        type: Boolean,
        default: () => config.defaultInputHasCounter
      },
      customClass: {
        type: String,
        default: ''
      }
    },
    data() {
      return {
        textAreaRow: 1,
        newValue: this.value,
        newType: this.type,
        newAutocomplete: this.autocomplete || config.defaultInputAutocomplete,
        isPasswordVisible: false,
        _elementRef: this.type === 'textarea'
            ? 'textarea'
            : 'input'
      }
    },
    created() {
      this.$emit('required', true);
    },
    computed: {
      computedValue: {
        get() {
          return this.newValue
        },
        set(value) {
          this.newValue = value
          this.$emit('input', value)
          !this.isValid && this.checkHtml5Validity()
        }
      },
      rootClasses() {
        return [
          this.iconPosition,
          this.size,
          {
            'is-expanded': this.expanded,
            'is-loading': this.loading,
            'is-clearfix': !this.hasMessage
          }
        ]
      },
      inputClasses() {
        return [
          this.statusType,
          this.size,
          {'is-rounded': this.rounded}
        ]
      },
      hasIconRight() {
        return this.passwordReveal || this.loading || this.statusTypeIcon
      },
      /**
       * Position of the icon or if it's both sides.
       */
      iconPosition() {
        if (this.icon) {
          if (this.hasIconRight) return 'has-icons-left has-icons-right'

          return 'has-icons-left'
        }
        // if(this.rules) return 'has-icons-right'
        return 'has-icons-right'
        // if(this.hasIconRight)  return 'has-icons-right'

      },
      /**
       * Icon name (MDI) based on the type.
       */
      statusTypeIcon() {
        switch (this.statusType) {
          case 'is-success':
            return 'check'
          case 'is-danger':
            return 'alert-circle'
          case 'is-info':
            return 'information'
          case 'is-warning':
            return 'alert'
        }
      },
      /**
       * Check if have any message prop from parent if it's a Field.
       */
      hasMessage() {
        return !!this.statusMessage
      },
      /**
       * Current password-reveal icon name.
       */
      passwordVisibleIcon() {
        return !this.isPasswordVisible ? 'eye' : 'eye-off'
      },
      /**
       * Get value length
       */
      valueLength() {
        if (typeof this.computedValue === 'string') {
          return this.computedValue.length
        } else if (typeof this.computedValue === 'number') {
          return this.computedValue.toString().length
        }
        return 0
      }
    },
    watch: {
      /**
       * When v-model is changed:
       *   1. Set internal value.
       */
      value(value) {
        if (this.type === 'textarea') {
          this.textAreaRow = (value.match(/\n/g) || []).length + 1;
        }

        this.newValue = value
      },
      isFocused(val) {
        if (this.type !== 'textarea') return;
        if (this.computedValue === undefined) return;
        const numRow = (this.computedValue.match(/\n/g) || []).length;
        if (val) {
          this.textAreaRow = numRow + 1;
          return;
        }
        this.textAreaRow = numRow > 3 ? 3 : numRow + 1;
      }
    },
    methods: {
      togglePasswordVisibility() {
        this.isPasswordVisible = !this.isPasswordVisible
        this.newType = this.isPasswordVisible ? 'text' : 'password'
        this.$nextTick(() => {
          this.$refs.input.focus()
        })
      },
      /**
       * Input's 'input' event listener, 'nextTick' is used to prevent event firing
       * before ui update, helps when using masks (Cleavejs and potentially others).
       */
      onInput(event) {
        this.$nextTick(() => {
          if (event.target) {
            this.computedValue = event.target.value
          }
        })
      }

    }
  }
</script>
<style scoped>
  .icon {
    transition: opacity 1s ease-in;
  }

  .icon-checker svg {
    width: 1rem !important;
    height: 1rem !important;
  }

  textarea {
    resize: none;
    overflow-y: hidden;
  }
</style>