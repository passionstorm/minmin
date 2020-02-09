const t = [
  {
    path: '/',
    name: 'shop_list',
    index: 1,
    meta: {title: 'Nhân sự', icon: 'users'},
    component: () => import('./page'),
  },
  {
    path: '/shop/1',
    name: 'shop_page',
    index: 1,
    meta: {title: 'Nhân sự', icon: 'users'},
    component: () => import('./shop_page'),
    children: [
      {
        path: '/shop/1/',
        name: 'shop_page_info',
        index: 1,
        meta: {title: 'Nhân sự', icon: 'users'},
        component: () => import('./ShopPageInfo'),
      },
      {
        path: '/shop/1/product',
        name: 'shop_page_prod',
        index: 1,
        meta: {title: 'Nhân sự', icon: 'users'},
        component: () => import('./ShopProduct'),
      },
    ],
  },

];
export default t;