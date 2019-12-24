import Vue from 'vue'
import App from './App'
import router from './router'
import toast from '../../widgets/toast'
import dialog from '../../widgets/dialog'
import axios from 'axios'
// import store from './store';
require('./css/frontpage.css')
require('../../config')
require('../../../dist/app.css')
require('../../../dist/style.css')
require('../../helper_filter')
Vue.use(toast)
Vue.use(dialog)

const base = axios.create({
  baseURL: 'https://mintoot-minh.firebaseio.com/',
})
Vue.prototype.$http = base

export const bus = new Vue()
new Vue({
  render: h => h(App),
  router,
  // store,
  bus,
}).$mount('#app')

