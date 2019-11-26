<template>
    <div class="d-flex justify-content-center login_form">
        <card>
            <form ref="form">
                <div class="form-group">
                    <input type="text" class="form-control" placeholder="Username" v-model="loginForm.username"/>
                </div>
                <div class="form-group">
                    <input type="password" class="form-control" placeholder="Password" v-model="loginForm.password"/>
                </div>
                <div class="form-group">
                    <a href="javascript:;" class="btn btn-primary" @click="handleLogin">Login</a>
                </div>
                <div class="form-group">
                    <a href="javascript:;" class="ForgetPwd">Forget Password?</a>
                </div>
            </form>
        </card>
    </div>
</template>

<script>
  import Card from '../../widgets/Card';

  export default {
    components: {
      Card,
    },
    data() {
      return {
        loginForm: {
          username: 'admin',
          password: '111111',
        },
        loading: true,
      };
    },
    methods: {
      handleLogin() {
        this.loading = true;
        this.$store.dispatch('user/login', this.loginForm).then(() => {
          this.$router.push({path: this.redirect || '/', query: this.otherQuery});
          this.loading = false;
        }).catch((err) => {
          console.log(err);
          this.loading = false;
        });
      },
    },
  };
</script>
<style scoped>
    .login_form {
        padding-top: 20px;
    }
</style>