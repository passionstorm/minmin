import Vue from 'vue'
import Cookies from 'js-cookie'
const lang = Cookies.get('lang') || 'en'
Vue.config.lang = lang

import VueMeta from 'vue-meta'
Vue.use(VueMeta)

import Icon from './widgets/Icon'
Vue.component('icon', Icon)