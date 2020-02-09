import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import './global.js'

require('./config');
// require('../dist/app.css');
// require('../dist/style.css');

require('./helper_filter');
import toast from './widgets/toast';
import dialog from './widgets/dialog'
Vue.use(toast);
Vue.use(dialog);


export const bus = new Vue();
new Vue({
  router,
  store,
  bus,
  render: h => h(App)
}).$mount('#app');

