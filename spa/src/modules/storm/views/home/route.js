export default [
  {
    path: '/',
    name: 'home',
    index: 1,
    meta: {title: 'Nhân sự', icon: 'users'},
    component: () => import('./home'),
  },
]