<template>
  <div>
    <h2>Review</h2>
    <pre>{{frm}}</pre>
    <table class="form-table">
      <tbody>
      <tr>
        <th>
          <label>Tên</label>
        </th>
        <td class="d">
          <input class="form-control" type="text" placeholder="Nhập tên cái cần review">
        </td>
      </tr>
      <tr>
        <th>
          <label>URL</label>
        </th>
        <td class="d">
          <input class="form-control" type="text" placeholder="Nhập url của cái cần review">
        </td>
      </tr>
      <tr>
        <th><label>Nội dung review</label></th>
        <td><textarea class="form-control" type="text" placeholder="Nhập nội dung review"/></td>
      </tr>
      <tr>
        <th><label>Hình ảnh</label></th>
        <td>
          <upload-image v-model="frm.images"/>
        </td>
      </tr>
      <tr>
        <th>Vị trí</th>
        <td>
          <QEditor />
        </td>
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import Editor from '../../../../widgets/Editor'
  import UploadImage from '../../../../widgets/UploadImage'
  import GMap from '../../../../widgets/GMap'
  import QEditor from '../../../../widgets/QEditor'
  export default {
    name: 'EditorForm',
    components: {Editor, UploadImage, GMap, QEditor},
    data() {
      return {
        frm: {},
      }
    },

  }
</script>

<style scoped>
  .form-table {

  }

  .form-table th {
    font-size: 16px;
    min-width: 150px;
  }
</style>