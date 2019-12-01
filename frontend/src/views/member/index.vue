<template>
  <div>
    <div class="d-flex mb-2">
      <div class="mr-auto"></div>
      <div>
        <a v-show="checkedRows.length" @click="del" class="btn btn-danger" href="javascript:;"><b>Xoá</b></a>
        <router-link :to="{name: 'member_edit'}" class="btn btn-primary"><b>Thêm thành viên</b></router-link>
      </div>
    </div>
    <div class="card">
      <div class="card-body">
        <v-table
            class="table table-striped"
            :data="data"
            :paginated="true"
            :pagination-simple="true"
            pagination-position="bottom"
            :columns="columns"
            :current-page.sync="currentPage"
            :per-page="perPage"
            :checked-rows.sync="checkedRows"
            checkable
        >
          <template slot-scope="props">
            <v-table-column v-for="(column, index) in columns"
                            :key="index"
                            :label="column.label"
                            :field="column.field"
            >
              {{ props.row[column.field] }}
            </v-table-column>
            <v-table-column label="">
              <router-link :to="{name: 'member_edit', params:{id: props.row['_id']}}" class="btn">
                <icon name="pen"></icon>
              </router-link>

            </v-table-column>
          </template>
          <template slot="top-left">
              <div class="p-2">
                <b>Đã chọn</b>: {{ checkedRows.length }}
              </div>
              <div class="d-flex">
                <v-select class="is-success" v-model="perPage">
                  <option value="5">5 per page</option>
                  <option value="10">10 per page</option>
                  <option value="15">15 per page</option>
                  <option value="20">20 per page</option>
                </v-select>
              </div>
          </template>
        </v-table>
      </div>
    </div>
  </div>
</template>

<script>
  import {VTableColumn, Check, Icon, Modal, Radio, VTable, VSelect} from '../../widgets';

  export default {
    components: {Modal, Icon, Check, Radio, VTable, VTableColumn, VSelect},
    data() {
      return {
        currentPage: 1,
        perPage: 15,
        data: [],
        selected: {},
        columns: [
          {
            field: 'username',
            label: 'Username',
            searchable: true,

          },
          {
            field: 'name',
            label: 'Name',
            searchable: true,

          },
          {
            field: 'email',
            label: 'Email',
            centered: true,
            searchable: true,

          },
          {
            field: 'phone',
            label: 'phone',
          },
        ],
        checkedRows: [],
        entity: ENTITY,
        editModal: false,
      };
    },
    computed: {},
    created() {
    },
    methods: {
      edit(item) {
        // this.setEditedItem(item)
      },
      del() {
        this.$app.dialog.confirm({
          message: 'Bạn có chắc chắn muốn xoá không?',
          onConfirm: () => this.$app.toast.open('User confirmed')
        })
        // let ids = this.checkedRows.map(e => e._id);
        // ids.forEach(id => MemberModel.$delete(id));
        // this.checkedRows = [];
        // this.data = MemberModel.query().all();
      },
    },
    async beforeMount() {
    },
  };
</script>
<style scoped>
  .btn-icon {
    display: block;
    width: 48px;
    cursor: pointer;
  }
</style>
