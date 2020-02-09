import {initializeApp} from 'firebase';

const config = {
  databaseURL: process.env.FIREBASE_DB_URL,
  storageBucket: process.env.FIREBASE_STORAGE_BUCKET,
};
const db = initializeApp(config).database();

const state = {};

const actions = {
  post({commit}, payload) {
    return new Promise((res, rej) => {
      let path = payload.path;
      let frm = payload.data;
      db.ref(path).push(frm).then(ref => {
        res(ref.key);
      }).catch(err => {
        rej(err);
      });
    });
  },
  get({commit}, payload) {

  },
};

const mutations = {};
const namespaced = true;
export default {
  namespaced,
  state,
  actions,
  mutations,
};