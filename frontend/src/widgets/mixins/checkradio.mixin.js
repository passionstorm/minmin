export default {
  props: {
    value: [String, Number, Boolean, Function, Object, Array],
    val: [String, Number, Boolean, Function, Object, Array],
    color: String,
    disabled: Boolean,
    required: Boolean,
    name: String,
    size: String,
  },
  data() {
    return {
      newValue: this.value
    }
  },
  computed: {
    computedValue: {
      get() {
        return this.newValue
      },
      set(value) {
        this.newValue = value
        this.$emit('input', value)
      }
    }
  },
  watch: {
    value(value) {
      this.newValue = value
    }
  },
  methods: {
    focus() {
      this.$refs.input.focus()
    }
  }
}