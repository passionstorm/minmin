<template>
  <div class="card">
    <div class="card-content">
      <form @submit.prevent="onSubmit" method="post" action="/">
        <v-input placeholder="What is name of thing your sell" v-model.lazy="frm.name"/>
        <input type="submit"/>
      </form>
    </div>
    <div class="card-footer">
      <pre>{{frm}}</pre>
    </div>
  </div>
</template>

<script>
  import VInput from '../../../widgets/VInput'

  export default {
    name: 'CreateProduct',
    data() {
      return {
        frm: {},
      }
    },
    methods: {
      onSubmit(e) {
        this.$http.post('post', this.frm).then(res => {
          this.frm.res = res
        }).catch(err => {
          console.log(err)
          this.frm.err = err
        })
        return true
      },
      onReset() {
        this.frm = {}
      },
    },
    components: {
      VInput,
    },
  }
</script>

<style scoped>

</style>