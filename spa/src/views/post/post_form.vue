<template>
  <div>
    <form @submit.prevent="onSubmit">
      <div class="d-flex bd-highlight mb-2">
        <div class="mr-auto p-2 bd-highlight">
          <a href="javascript:;" @click="$router.go(-1)" class="btn btn-d">❮ Quay lai</a>
        </div>
        <div class="p-2 bd-highlight">
          <a href="javascript:;" class="btn btn-danger" @click="modal.reset = true">Reset</a>
          <a href="javascript:;" class="btn btn-d ">Lưu</a>
          <a href="javascript:;" @click="onSubmit" class="btn btn-primary">Đăng</a>
        </div>
      </div>
      <collapse aria-id="colaps.post" class="card card-primary" :open.sync="colaps.post.open">
        <div role="button" class="card-header" slot="trigger" aria-controls="colaps.post" slot-scope="props">
          <div class="card-header-title">Bài viết</div>
          <a class="card-header-icon">
            <icon :name="props.open ? 'arrow_down' : 'arrow_right'"></icon>
          </a>
        </div>
        <div class="card-body">
          <field class="col-sm-10" label="Tiêu đề" required="required">
            <v-input vid="title" rules="required" v-model="form.title" type="'text'" title="Tiêu đề*"
                     placeholder="Tiêu đề"/>
          </field>
          <field class="col-sm-10" label="Nội dung">
            <editor v-model="form.content"></editor>
          </field>
          <field>
            <upload v-model="dropFiles"
                    drag-drop>
              <section class="section">
                <div class="content has-text-centered">
                  <p>
                    <icon name="cloud_upload" size="l"/>
                  </p>
                  <p>Drop your files here or click to upload</p>
                </div>
              </section>
            </upload>
          </field>
        </div>
      </collapse>
      <!--      <collapse aria-id="colaps.seo" class="card card-primary" :open.sync="colaps.seo.open">-->
      <!--        <div slot="trigger" class="card-header" role="button" slot-scope="props" aria-controls="colaps.seo">-->
      <!--          <div class="card-header-title">-->
      <!--            Seo-->
      <!--          </div>-->
      <!--          <a class="card-header-icon">-->
      <!--            <icon :name="props.open ? 'arrow_down' : 'arrow_right'"/>-->
      <!--          </a>-->
      <!--        </div>-->
      <!--        <div class="card-body">-->
      <!--          <field class="col-sm-10" label="Tiêu đề">-->
      <!--            <v-input v-model="form.seo.title" type="textarea" maxlength="60"/>-->
      <!--          </field>-->
      <!--          <field class="col-sm-10" label="Mô tả ngắn">-->
      <!--            <v-input v-model="form.seo.description" type="textarea" maxlength="100"/>-->
      <!--          </field>-->
      <!--          <field label="Từ khoá tìm kiếm" class="col-sm-10">-->
      <!--            <taginput :allow-new="true" :data="filterSeoTags" autocomplete field="content" v-model="form.seo.tags"-->
      <!--                      icon="label" @add="onAddSeoTag" @typing="onTypeSeoTags"-->
      <!--                      placeholder="Nhập từ khoá"-->
      <!--            ></taginput>-->
      <!--          </field>-->
      <!--        </div>-->
      <!--      </collapse>-->
    </form>
    <modal v-model="modal.reset">
      <template slot="header">
        <h5 class="modal-title">Xác nhận</h5>
      </template>
      <template slot="body">
        <p>Bạn muốn reset nội dung đã nhập?</p>
      </template>
      <template slot="footer">
        <button type="button" class="btn btn-secondary" @click="modal.reset = false">NO</button>
        <button type="button" class="btn btn-primary" autofocus @click="reset">YES</button>
      </template>
    </modal>
  </div>
</template>

<script>
  import {Check, Collapse, Editor, Field, Icon, Modal, Pin, Taginput, Upload, VInput, VSelect} from '../../widgets';
  import {mapState} from 'vuex';
  import {htmlDecode} from '../../utils';
  import {msg} from '../../utils/msg';

  const defaultForm = {
    status: '1',
    seo: {},
    title: '',
    content: undefined,
    desc: '',
    source_uri: '',
    image_uri: '',
    display_time: undefined,
    id: undefined,
    comment_disabled: false,
    importance: 0,
  };
  export default {
    components: {
      VInput, Field, Pin, VSelect, Modal, Check, Taginput, Collapse, Icon, Upload, Editor,
    },
    props: {
      id: [Number, String],
    },
    data() {
      return {
        tags: [],
        colaps: {
          post: {open: true},
          seo: {open: true},
        },
        dataRes: {},
        editedItem: {},
        listCategory: [
          {
            id: 0,
            name: 'ban hang',
          }, {id: 1, name: 'duoc'}, {id: 3, name: 'nha cua'}],
        filterCategory: this.listCategory,
        statusOptions: {
          0: 'draft',
          1: 'publish',
        },
        modal: {
          reset: false,
        },
        filterSeoTags: [],
        listSeoTag: [],
        check: {
          test: false,
        },
        dropFiles: [],
        form: Object.assign({}, defaultForm),
      };
    },
    watch: {
      dropFiles(val) {
        console.log(val);
      },
    },
    computed: {
      ...mapState('post', ['item']),
    },
    methods: {
      onTypeCategoryTag(text) {
        this.filterCategory = this.listCategory.filter((i) => {
          return i.name.toString().toLowerCase().indexOf(text.toLowerCase()) >= 0;
        });
      },
      onTypeSeoTags(text) {
        console.log(this.form.seo.tags);
        this.filterSeoTags = this.listSeoTag.filter((i) => {
          return i.content.toString().toLowerCase().indexOf(text.toLowerCase()) >= 0;
        });
      },
      reset() {
        this.form = Object.assign({}, defaultForm);
        this.modal.reset = false;
      },
      onSubmit() {
        axios.post('post/create', this.form).then(rs => {
          if (rs.data.success) {
            msg.success(this.id ? 'Cập nhật thành công!' : 'Thêm mới thành công!');
            this.$router.push({name: 'post_index'});
          } else {
            msg.success(this.id ? 'Cập nhật thất !' : 'Thêm mới thất bại!');
          }
        }).catch(err => msg.fail(err));
      },
      async onAddSeoTag(text) {
      },
    },
    created() {
    },
    mounted() {
      if (this.id) {
        axios.get('post/edit/' + this.id).then(res => {
          console.log(res.data);
          this.form = res.data;
          this.form.content = htmlDecode(this.form.content);
          this.form.seo = {};
        });
      }
    },
    beforeDestroy() {
    },
  };
</script>
<style scoped>

</style>>
