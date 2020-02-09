<template>
  <label
      class="b-radio radio"
      ref="label"
      :class="[size, { 'is-disabled': disabled }]"
      :disabled="disabled"
      @click="focus"
      @keydown.prevent.enter="$refs.label.click()">
    <input
        v-model="computedValue"
        type="radio"
        ref="input"
        @click.stop
        :disabled="disabled"
        :required="required"
        :name="name"
        :value="val">
    <span class="check" :class="color"/>
    <span class="control-label"><slot/></span>
  </label>
</template>

<script>
  import CheckRadioMixin from './mixins/checkradio.mixin'

  export default {
    name: 'VRadio',
    mixins: [CheckRadioMixin],
    components: {},
  };
</script>
<style scoped>

</style>