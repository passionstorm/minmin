const types = {
  TOGGLE_SIDEBAR: 'TOGGLE_SIDEBAR',
};

const state = {
  sidebarOpen: true,
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
};

export default {
  state,
  actions,
  mutations,
};