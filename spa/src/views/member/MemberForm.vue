<template>
  <form ref="form" @submit.prevent="onSubmit">
    <field required="true" label="Username">
      <v-input name="username" rules="required|min:3|max:30|username" v-model="editedItem.username" maxlength="30"/>
    </field>
    <field required="true" label="Chức vụ">
      <div class="block">
        <radio color="is-success" v-for="role in roleSelect" :key="role" v-model="editedItem.role" name="name"
               :val="role"> {{role}}
        </radio>
      </div>
    </field>
    <field required="true" label="Mật khẩu">
      <v-input type="password" name="Mật khẩu" v-model="editedItem.password" rules="required|min:6|max:20"
               vid="password"/>
    </field>
    <field required="true" label="Xác nhận mật khẩu" >
      <v-input name="Xác nhận mật khẩu" type="password" v-model="editedItem.rePassword"
               rules="required|confirmed:password"/>
    </field>
    <field required="true" label="Tên">
      <v-input name="Tên" title="Tên" v-model="editedItem.name" rules="required"/>
    </field>
    <field required="true" label="Email">
      <v-input name="Email" v-model="editedItem.email" rules="required|email"/>
    </field>
    <field label="Số điện thoại">
      <v-input rules="phone" v-model="editedItem.phone"/>
    </field>
    <div class="col-sm-10"
        style="display:flex;
           box-shadow: 0 4px 8px 0 rgba(0,0,0,0.6);
            padding: 0.5em 2em 0.5em 0.1em;
            bottom: 0;
            margin-right: -20px;
            position: fixed;
            background: rgba(0, 0, 0, 0.4);">
      <div class="col-sm-12" style="display: flex">
        <router-link :to="{name: 'member_list'}" class="btn btn-d">Quay lai</router-link>
        <div class="mr-auto"></div>
        <button type="submit" class="btn btn-primary"><b>Submit</b></button>
      </div>
    </div>
  </form>
</template>

<script>
  import {ROLE_TYPES} from '../../utils/constant';
  import {Collapse, Field, Icon, Radio, VInput} from '../../widgets';

  export default {
    components: {VInput,Field, Collapse,Icon,Radio, },
    props: {},
    data() {
      return {
        editedItem: {
          role: 'admin',
        },
        id: this.$route.params.id,
        roleSelect: Object.values(ROLE_TYPES),
      };
    },
    watch: {},
    mounted() {
    },
    methods: {
      async onSubmit() {
      },
      resetForm() {
        this.$refs.form.reset();
        this.editedItem = Object.assign({}, {});
      },
    },
    beforeMount() {
      this.$route.meta.title = this.id ? 'Cập nhật thành viên' : 'Thêm thành viên';
    },
    created() {
    },
    computed: {},
  };
</script>

<style scoped>

</style>