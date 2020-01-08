<template>
  <div>
    <table class="form-table">
      <tbody>
        <tr>
          <th>
            <label>Nhập tên sản phẩm</label>
          </th>
          <td class="d">
            <input class="form-control" type="text">
          </td>
        </tr>
        <tr>
          <th>
            <label for="">Nhập tên sản phẩm ssssss</label>
          </th>
          <td class="d">
            <input class="form-control" type="text">
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import Editor from '../../../../widgets/Editor'

  export default {
    name: 'EditorForm',
    components: {Editor},
    data() {
      return {
        frm:{}
      }
    },

  }
</script>

<style scoped>
  .form-table{

  }
  dl {
    display: flex;
    align-items: center;
  }

  dt {
    width: 100px;
  }
</style>