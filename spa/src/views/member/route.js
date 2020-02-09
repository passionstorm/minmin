import Layout from '../../components/layouts/DefaultLayout';

const t = {
  path: '/member',
  component: Layout,
  type: 'bar',
  index: 1,
  meta: {title: 'Nhân sự', icon: 'users'},
  children: [
    {
      path: '',
      name: 'member_list',
      meta: {
        noCache: true,
        title: 'Danh sách',
        roles: ['editor'],
        icon: 'address_book',
      },
      component: () => import('./index'),
    },
    {
      path: 'edit/:id?',
      name: 'member_edit',
      hidden: true,
      meta: {
        noCache: true,
        title: 'Danh sach user',
        roles: ['admin', 'editor'],
        icon: 'address_book',
      },
      component: () => import('./MemberForm'),
    },
    {
      path: 'permission',
      name: 'post',
      meta: {
        noCache: true,
        title: 'Phân quyền',
        roles: ['admin', 'editor'],
        icon: 'user_role',
      },
      component: () => import('./MemberForm'),
    },
  ],
};
export default t;