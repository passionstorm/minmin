<template>
  <div>
    <div class="table-responsive">
      <table class="table table-striped">
        <thead>
        <tr>
          <th v-for="col in _columns" scope="col">
            <template v-if="col.slot">
              <slot/>
            </template>
            <template v-else>
              {{col.title}}
            </template>
          </th>
        </tr>
        </thead>
        <tbody>
        <tr v-for="i in items">
          <td v-for="col in _columns">
            <pre>{{i[col.key]}}</pre>
          </td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
  const defaultColumn = {
    title: '',
    key: '',
  };
  export default {
    name: 'v-table',
    props: {
      columns: {
        type: Array,
        default: [],
      },
      items: {
        type: Array,
        data: [],
      },
    },
    data() {
      return {};
    },
    created() {
      console.log(this._columns);
    },
    computed: {
      _columns() {
        if (Array.isArray(this.columns)) {
          return this.columns.map(e => {
            if (typeof e == 'object') {
              return e;
            }

            return {title: e, key: e};
          });
        }

        return Object.assign(this.columns, {});
      },
      _items() {

      },
    },
  };
</script>

<style scoped>

</style>