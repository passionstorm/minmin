const state = {
  item: {},
  items: [],
  errors: [],
};
const mutations = {
  ITEM(state, payload) {
    state.item = payload;
  },
  ITEMS(state, payload) {
    state.items = payload;
  },
  ERRORS(state, payload) {
    state.errors = payload;
  },
};

const actions = {
  get_all({commit}) {
    axios.get('api/post', {
      headers: {
        // Authorization: `Bearer ${
        //     context.getters.currentUser.token
        // }`,
      },
    }).then(res => commit('ITEMS', res.data)).catch(err => commit('ERRORS', err));
  },
  get({commit}, payload) {
    axios.get('api/post/edit/' + payload).then(res => commit('ITEM', res.data)).catch(err => commit('ERRORS', err));
  },
  create() {
  },
};

const namespaced = true;
export default {
  namespaced,
  state,
  actions,
  mutations,
};