const getters = {
  sidebarOpen: state => state.app.sidebarOpen,
  permission_routes: state => state.permission.routes,
  visitedViews: state => state.tag_view.visitedViews,
  cachedViews: state => state.tag_view.cachedViews,
  token: state => state.user.token,
  avatar: state => state.user.avatar,
  name: state => state.user.name,
  roles: state => state.user.roles,
};
export default getters;