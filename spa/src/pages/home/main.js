import Vue from 'vue';
import Layout from './Layout';
// import router from './router';
// import store from './store';
import toast from '@widget/toast'
import dialog from '@widget/dialog'

Vue.use(toast);
Vue.use(dialog);


export const bus = new Vue();
new Vue({
  // router,
  // store,
  bus,
  render: h => h(Layout)
}).$mount('#app');