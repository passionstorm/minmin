<template>
  <div>
    <h2>Review</h2>
    <pre>
      {{frm}}
    </pre>
    <table class="form-table">
      <tbody>
      <tr>
        <th>
          Danh mục đánh giá
        </th>
        <td>
          <ASelect v-model="frm.cat" :options="data.cats" option-text="name" option-value="id"
          />
        </td>
      </tr>
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
      </tr>
      </tbody>
    </table>
  </div>
</template>

<script>
  import Editor from '../../../../widgets/Editor'
  import UploadImage from '../../../../widgets/UploadImage'
  import ASelect from '../../../../widgets/ASelect'
  export default {
    name: 'EditorForm',
    components: {Editor, UploadImage,ASelect},
    data() {
      return {
        data: {
          cats: [
            {id: 1, name: 'Món ăn'},
            {id: 2, name: 'Quán ăn'},
            {id: 3, name: 'Địa điểm du lịch'},
            {id: 4, name: 'Địa điểm du lịch'},
            {id: 5, name: 'Địa điểm du lịch'},
            {id: 6, name: 'Địa điểm du lịch'},
            {id: 7, name: 'Địa điểm du lịch'},
            {id: 8, name: 'Địa điểm du lịch'},
            {id: 9, name: 'Địa điểm du lịch'},
            {id: 10, name: 'Địa điểm du lịch'},
            {id: 11, name: 'Địa điểm du lịch'},
            {id: 12, name: 'Địa điểm du lịch'},
            {id: 13, name: 'Địa điểm du lịch'},
          ]
        },
        frm: {
        },
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