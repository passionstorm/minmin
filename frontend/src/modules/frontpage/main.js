import Vue from 'vue';
import App from './App';
import router from './router';
// import store from './store';
require('../../config');
require('../../../dist/app.css');
require('../../../dist/style.css');
require('../../helper_filter');
import toast from '../../widgets/toast';
import dialog from '../../widgets/dialog'
Vue.use(toast);
Vue.use(dialog);


export const bus = new Vue();
new Vue({
  render: h => h(App),
  router,
  // store,
  bus,
}).$mount('#app');

