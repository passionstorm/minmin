import Layout from '../../components/layouts/DefaultLayout';

const t = {
  path: '/post',
  component: Layout,
  type: 'bar',
  children: [
    {
      path: 'index',
      name: 'post',
      meta: {
        noCache: true,
        title: 'Bài viết', group: 'apps',
        roles: ['admin', 'editor'],
        icon: 'edit',
      },
      component: () => import('./PostForm'),
    },
  ],
};

export default t;