<template>
  <div>
    <pre>
      {{dt}}
    </pre>
    <h4>Collections</h4>
    <table class="table">
      <thead>
      <tr>
        <th>name</th>
      </tr>
      </thead>
      <tbody>
      <tr v-for="i in dt">
        <td>
          {{i}}
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  export default {
    name: 'index.vue',
    data() {
      return {
        dt: {},
      }
    },
    computed:{
      
    },
    created() {
      let t = this
      axios.get('/db').then(rs => {
        t.dt = rs.data
      })
    },
  }
</script>

<style scoped>

</style>