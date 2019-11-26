<template>
  <label
      class="b-checkbox checkbox"
      :class="[size, { 'is-disabled': disabled }]"
      ref="label"
      :disabled="disabled"
      @click="focus"
      @keydown.prevent.enter="$refs.label.click()">
    <input
        v-model="computedValue"
        :indeterminate.prop="indeterminate"
        type="checkbox"
        ref="input"
        @click.stop
        :disabled="disabled"
        :required="required"
        :name="name"
        :value="val"
        :true-value="trueValue"
        :false-value="falseValue">
    <span class="check" :class="color"></span>
    <span class="control-label"><slot/></span>
  </label>
</template>

<script>
  import CheckRadioMixin from './mixins/checkradio.mixin'
  export default {
    name: 'VCheck',
    mixins: [CheckRadioMixin],
    props: {
      indeterminate: Boolean,
      trueValue: {
        type: [String, Number, Boolean, Function, Object, Array],
        default: true
      },
      falseValue: {
        type: [String, Number, Boolean, Function, Object, Array],
        default: false
      }
    }
  }
</script>
