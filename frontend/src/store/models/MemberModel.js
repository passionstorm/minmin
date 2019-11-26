import {BaseModel} from './BaseModel'

export const ENTITY = 'members';

export const module = {};

export default class MemberModel extends BaseModel {
  // This is the name used as module name of the Vuex Store.
  static entity = ENTITY;
  static fields() {
    return {
      _id: this.uid(),
      name: this.attr(''),
      email: this.attr(''),
      username: this.string('test'),
      password: this.string('123456'),
      phone: this.string(''),
      address: this.string(''),
      status: this.number('active'),
      hash: this.string(''),
      token: this.string('qqyzkzldrx'),
      role: this.attr('guest'),
      avatarUrl: this.string('/avatar/man_1.jpg'),
      permissions: this.attr([]),
    };
  }
}