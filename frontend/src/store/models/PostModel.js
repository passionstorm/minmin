import {make} from 'vuex-pathify';
import {BaseModel} from './BaseModel';
import User from './UserModel';

export const ENTITY = 'posts';

const state = {
  name: ENTITY,
  items: [],
  defaultArticle: {},
};
export const module = {
  namespaced: true,
  state,
  mutations: {
    ...make.mutations(state),
  },
  actions: {
    ...make.actions(state),
  },
  getters: {
    ...make.getters(state),
  },
};

export default class PostModel extends BaseModel {
  static entity = ENTITY;
  static fields() {
    return {
      title: this.string(''),
      summary: this.string(''),
      body: this.string(''),
      status: this.number(0),
      author: this.belongsTo(User, 'user_id'),
      tags: this.attr('dev'),
      star: this.attr('10'),
      like: this.attr('24'),
      published: this.attr(''),
      updateAt: this.attr('2019-07-01'),
      href: this.attr('http://localhost'),
    };
  }
  static apiConfig = {
    baseURL: 'http://localhost:9000/api',
    actions: {
      fetch: {
        method: 'get',
        url: '/post/create'
      },
      // post: {
      //   method: 'post',
      //   url: '/post/create'
      // }
    }
  }
}
