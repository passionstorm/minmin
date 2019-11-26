<template>
  <div class="editor">
    <froala id="edit" @input="$emit('input', model)"
            tag="textarea"
            :config="config"
            v-model="model"
    />
  </div>
</template>

<script>
  import Icon from './Icon';
  import 'froala-editor/css/froala_editor.pkgd.min.css';

  export default {
    name: 'editor',
    components: {
      Icon,
    },
    props: {},
    computed: {},
    created() {
    },
    data() {
      return {
        model: '',
        config: {
          placeholderText: 'Edit Your Content Here!',
          charCounterCount: true,
        },
      };
    },
    methods: {},
    mounted() {
    },

    beforeDestroy() {
    },
  };
</script>

<style scoped>
  #logo{
    display: none !important;
  }
</style>