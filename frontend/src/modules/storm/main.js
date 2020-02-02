import Vue from 'vue'
import App from './App'
import router from './router'
import toast from '../../widgets/toast'
import dialog from '../../widgets/dialog'
import axios from 'axios'
import store from './store'
import * as VueGoogleMaps from "vue2-google-maps";
Vue.use(VueGoogleMaps, {
  load: {
    key: "AIzaSyAknIaoGLtaIUVfJ11cLhoedvf2pbhW0Tg ",
    libraries: "places"
  }
});

require('../../config')
require('./css/common.css')
require('../../helper_filter')
Vue.use(toast)
Vue.use(dialog)
const base = axios.create({
  baseURL: 'http://api.localhost:9000/',
})
Vue.prototype.$http = base

export const bus = new Vue()
new Vue({
  render: h => h(App),
  router,
  store,
  bus,
}).$mount('#app')

