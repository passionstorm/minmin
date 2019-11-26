import Vue from 'vue';
import App from './App.vue';
import router from './router';
import store from './store';
import toast from './widgets/toast';
import notification from './widgets/notification';
import {loadLocale} from './validator';
import VueFroala from 'vue-froala-wysiwyg';

Vue.use(toast);
Vue.use(notification);
Vue.use(VueFroala);
let locale = 'vi';
loadLocale(locale);

export const bus = new Vue();
new Vue({
  render: h => h(App),
  router,
  store,
  bus,
  locale: locale,
}).$mount('#app');

