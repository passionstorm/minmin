import {make} from 'vuex-pathify';
import {BaseModel} from './BaseModel';
import User from './UserModel';

export const ENTITY = 'items';

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

export default class ItemModel extends BaseModel {
  static entity = ENTITY;
  static fields() {
    return {
      _id: this.uid(),
      assetName: this.string('asset'),
      itemCode: this.string('itemCode'),
      itemName: this.string('itemName'),
      assetCategory: this.string('assetCategory'),
      assetOwner: this.string('assetOwner'),
      assetOwnerCompany: this.string('assetOwnerCompany'),
      image: this.string('image'),
      brand: this.string('brand'),
      model: this.string('model'),
      price: this.string('price'),
      producedDate: this.string('2019-03-08'),
      purchasedDate: this.string('2019-03-08'),
      descriptions: this.string('descriptions'),
      barCode: this.string('barCode'),
      supplier: this.string('supplier'),
      location: this.string('location'),
      custodian: this.string('custodian'),
    };
  }
}
