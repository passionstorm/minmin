import Firebase from 'firebase'

const config = {}
let app = Firebase.initializeApp(config)
let db = app.database()

const state = {
  sidebarOpen: true,
  isLoading: false,
}

const actions = {
  post({commit}, payload) {
    let table = payload.path
    let frm = payload.form
    db.ref(path).push(frm).then()
  },
  get({commit}, payload) {

  },
}

const mutations = {}

export default {
  state,
  actions,
  mutations,
}