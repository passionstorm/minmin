import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

const routes = [
  {path: '/foo', component: Foo},
  {path: '/bar', component: Bar},
]

const router = new Router({
  mode: 'history',
  routes: _asyncRoutes,
})
export default router