import Layout from '../../components/layouts/DefaultLayout';
const t = {
  path: '/assets',
  component: Layout,
  index: 1,
  type: 'bar',
  meta: {title: 'Phương tiện', icon: 'icon'},
  children: [
    {
      path: 'icon',
      name: 'icon',
      meta: {
        noCache: true,
        title: 'Biểu tượng',
        roles: ['admin', 'editor'],
        icon: 'grin',
      },
      component: () => import('./index'),
    },
    {
      path: 'photo',
      name: 'photo',
      meta: {
        noCache: true,
        title: 'Ảnh',
        roles: ['admin', 'editor'],
        icon: 'images',
      },
      component: () => import('./photo'),
    },
  ],
};

export default t;