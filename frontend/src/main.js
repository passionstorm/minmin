import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
require('./bootstrap')
require('../dist/app.css')

export const bus = new Vue();
new Vue({
  render: h => h(App),
  router,
  store,
  bus,
}).$mount('#app');

