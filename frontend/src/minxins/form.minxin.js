import Request from '../utils/request.vuex.orm';
import {rules} from '../utils/validate';
import {baseFilter} from '../utils';
import {HTTP_METHOD} from '../utils/constant';

const req = new Request();

const t = {
  data() {
    return {
      editedItem: {}, // currently item to be edited
      editedIndex: -1, // when -1, create, else update or delete
      model: null,
      fields: [],
      defaultItem: {},
      filter: {
        search: '',
        sort: '_id',
      },
      rules,

    };
  },

  computed: {
    formTitle() {
      return this.editedIndex === -1 ? 'Thêm' : 'Sửa';
    },
    editing() {
      return this.editedIndex !== -1; // is in edit state
    },
    items() {
      let search = this.filter.search;
      let sort = this.filter.sort;
      let data = this.model ? this.model.query().all() : [];
      if (search === '') return data;
      return baseFilter({sort, search}, data);
    },
  },

  watch: {
    editedIndex(val) {
      console.log(`Editing item ${val}`);
    },
  },
  created() {
  },
  methods: {
    async fetch() {
      const res = await req.request({
        method: HTTP_METHOD.GET,
        entity: this.entity,
        params: {},
      });

      if (!res) return;

      this.fields = res.result.fields;
      this.model = res.result.model;
      console.log("reslsls",res);
      this.setEditedItem(new this.model());
    },
    async deleteItem(data) {
      await req.request({
        method: HTTP_METHOD.DELETE,
        entity: this.entity,
        data: data,
      });
    },

    async updateItem(data) {
      await req.request({
        method: HTTP_METHOD.PUT,
        entity: this.entity,
        data: data,
      });
    },
    async createItem(data) {
      console.log('Creating', data);
      const res = await req.request({
        method: HTTP_METHOD.POST,
        entity: this.entity,
        data: data,
      });
      console.log(res);
    },
    async saveItem(data) {
      this.setEditedItem(data);
      if (this.editedIndex > -1) {
        await this.updateItem(this.editedItem);
      } else {
        await this.createItem(this.editedItem);
      }
    },
    setEditedItem(data) {
      this.editedIndex = data._id ? data._id : -1;
      this.editedItem = Object.assign({}, data);
    },
    reset() {
      this.setEditedItem(new this.model());
    },
  },
};

export default t;