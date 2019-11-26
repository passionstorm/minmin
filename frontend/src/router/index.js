import Vue from 'vue';
import Router from 'vue-router';
import Layout from '../components/layouts/DefaultLayout';
import store from '../store';
import {push} from '../utils';
import {getToken} from '../utils/auth';

Vue.use(Router);

export const constantRoutes = [
  {
    path: '/login',
    meta: {title: 'login'},
    component: () => import('../views/login/index'),
    hidden: true,
  },
  {
    path: '/',
    component: Layout,
    type: 'bar',
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        name: 'dashboard',
        component: () => import('../views/dashboard'),
        meta: {
          title: 'Home', group: 'apps',
          roles: ['admin', 'editor'],
          icon: 'dashboard',
        },
      },
    ],
  },
];

/**
 * Note: sub-menu only appear when route children.length >= 1
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */
let _asyncRoutes = [
  // {
  //   path: '/goods',
  //   meta: {title: 'Sản phẩm', icon: 'cart'},
  //   component: Layout,
  //   type: 'bar',
  //   children: [
  //     {
  //       path: 'menu1',
  //       name: 'goods_edit',
  //       meta: {
  //         noCache: true,
  //         title: 'menu1',
  //         roles: ['admin', 'editor'],
  //       },
  //       component: () => import('../views/post/post'),
  //     },
  //   ],
  // },
];

const collectRoute = require.context('../views', true, /index.route\.js$/);
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

export const asyncRoutes = _asyncRoutes;

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({y: 0}),
  routes: constantRoutes,
});

const router = createRouter();

// console.log("ahncsbhncsdndjn");
export function resetRouter() {
  const newRouter = createRouter();
  router.matcher = newRouter.matcher; // reset router
}

const whiteList = ['/login', '/auth-redirect'];

router.beforeEach(async (to, from, next) => {
  // start progress bar
  // document.title = getPageTitle(to.meta.title)
  // determine whether the user has logged in
  const hasToken = getToken();
  if (hasToken) {
    if (to.path === '/login') {
      // if is logged in, redirect to the home page
      next({path: '/'});
      // NProgress.done()
    } else {
      // determine whether the user has obtained his permission roles through getInfo
      const hasRoles = store.getters.roles && store.getters.roles.length > 0;
      if (hasRoles) {
        next();
      } else {
        try {
          // get user info
          // note: roles must be a object array! such as: ['admin'] or ,['developer','editor']
          const {roles} = await store.dispatch('user/getInfo');
          // generate accessible routes map based on roles
          const accessRoutes = await store.dispatch('permission/generateRoutes', roles);

          // dynamically add accessible routes
          router.addRoutes(accessRoutes);
          // console.log(router.);
          // set the replace: true, so the navigation will not leave a history record
          next({...to, replace: true});
        } catch (error) {
          // remove token and go to login page to re-login
          await store.dispatch('user/resetToken');
          // Message.error(error || 'Has Error')
          next(`/login?redirect=${to.path}`);
          // NProgress.done()
        }
      }
    }
  } else {
    if (whiteList.indexOf(to.path) !== -1) {
      // in the free login whitelist, go directly
      next();
    } else {
      // other pages that do not have permission to access are redirected to the login page.
      next(`/login?redirect=${to.path}`);
    }
  }
});
export default router;
