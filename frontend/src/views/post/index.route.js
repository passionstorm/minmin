import Layout from '../../components/layouts/DefaultLayout';

const t = {
  path: '/post',
  component: Layout,
  type: 'bar',
  meta: {title: 'Bài Viết', icon: 'edit'},
  children: [
    {
      path: '',
      name: 'post_index',
      meta: {
        noCache: true,
        title: 'Danh sách', group: 'apps',
        roles: ['admin', 'editor'],
      },
      component: () => import('./index'),
    },
    {
      path: 'edit/:id?',
      name: 'post_edit',
      meta: {
        noCache: true,
        title: 'Thêm bài viết', group: 'apps',
        roles: ['admin', 'editor'],
        icon: 'edit',
      },
      component: () => import('./PostForm'),
    },
  ],
};

export default t;