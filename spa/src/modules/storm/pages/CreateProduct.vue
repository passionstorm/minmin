<template>
  <div class="card">
    <div class="card-content">
      <form @submit.prevent="onSubmit" method="post" action="/">
        <field label="Tên sản phẩm">
          <v-input placeholder="What is name of thing your sell" v-model.lazy="frm.name"/>
        </field>
        <field label="Giá sản phẩm" :grouped="true">
          <v-input placeholder="Nhập giá thấp nhất" v-model.lazy="frm.priceFrom" />
          <v-input placeholder="Nhập giá cao nhất" v-model.lazy="frm.priceTo" />
        </field>
        <input type="submit"/>
      </form>
    </div>

    <div class="card-footer">
      <pre>{{log}}</pre>
    </div>
  </div>
</template>

<script>
  import {VInput, Field} from '../../../widgets'

  export default {
    name: 'CreateProduct',
    components: {
      VInput,Field
    },
    data() {
      return {
        frm: {},
        log: {}
      }
    },
    methods: {
      async onSubmit(e) {
        let res = await this.$store.dispatch('api/post', {
          path: 'products',
          data: this.frm,
        })
        this.log = res
        return true
      },
      onReset() {
        this.frm = {}
      },
    },

  }
</script>

<style scoped>

</style>