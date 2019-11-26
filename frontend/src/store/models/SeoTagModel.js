import {BaseModel} from './BaseModel'

export const ENTITY = 'seo_tag';

export const module = {};

export default class SeoTagModel extends BaseModel {
  // This is the name used as module name of the Vuex Store.
  static entity = ENTITY;
  static fields() {
    return {
      _id: this.uid(),
      content: this.string(''),
      shop_id: this.number(),
    };
  }
}