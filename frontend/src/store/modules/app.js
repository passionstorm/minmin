import {push} from '../../utils';

const types = {
  TOGGLE_SIDEBAR: 'TOGGLE_SIDEBAR',
};

let _asyncRoutes = [];
const state = {
  sidebarOpen: true,
  isLoading: false,
};

const loadState = () => {
  try {
    var serializedState = localStorage.getItem('app');
    if (serializedState === null || serializedState == '[]') {
      return undefined;
    }
    return JSON.parse(serializedState);
  } catch (err) {
    return undefined;
  }
};

const saveState = (state) => {
  try {
    const serializedState = JSON.stringify(state);
    localStorage.setItem('app', serializedState);
  } catch (err) {
    console.error(`Something went wrong: ${err}`);
  }
};

const actions = {
  toggleSidebar({commit, state}) {
    commit(types.TOGGLE_SIDEBAR);
  },
};

const mutations = {
  [types.TOGGLE_SIDEBAR](state) {
    state.sidebarOpen = !state.sidebarOpen;
  },
  loading(state, data){
    state.isLoading = data
  }
};

export default {
  state,
  actions,
  mutations,
};