import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';

require('./config');
require('../dist/app.css');
require('./helper_filter');



export const bus = new Vue();
new Vue({
  render: h => h(App),
  router,
  store,
  bus,
}).$mount('#app');

