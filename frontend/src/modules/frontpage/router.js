import Vue from 'vue';
import Router from 'vue-router';
import {push} from '../../utils';

Vue.use(Router);
let _asyncRoutes = [];
const collectRoute = require.context('./views', true, /route\.js$/);
collectRoute.keys().forEach((r) => {
  let route = collectRoute(r).default;
  if (Array.isArray(route)) {
    route.forEach(subRoute => {
      _asyncRoutes = push(_asyncRoutes, subRoute, subRoute.index);
    });
  } else {
    _asyncRoutes = push(_asyncRoutes, route, route.index);
  }

});

const router = new Router({
  mode: 'history',
  routes: _asyncRoutes,
});
export default router;