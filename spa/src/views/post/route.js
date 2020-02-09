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
      component: () => import('./post_list'),
    },
    {
      path: 'edit/:id?',
      name: 'post_edit',
      props: true,
      meta: {
        noCache: true,
        title: 'Thêm bài viết', group: 'apps',
        roles: ['admin', 'editor'],
        icon: 'edit',
      },
      component: () => import('./post_form'),
    },
  ],
};

export default t;