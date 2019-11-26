import {CODE_MESSAGE} from './constant';
import models from '../store/models';

const t = function() {
  t.method = '';
  t.namespace = '';
  t.params = '';
  t.model = null;
  t.fields = '';
  t.response = {};
  t.data = null;
  t.result = {
    model: t.model,
    fields: t.fields,
    namespace: t.namespace,
    reqParams: t.params,
    reqMethod: t.method,
    reqData: t.method,
  };
  t.resultNotFound = {
    status: 404,
    message: CODE_MESSAGE[404],
  };
  t.execute = {
    POST: async function() {
      let createdItems = await t.model.$create({data: t.data});
      if (!createdItems) return t.resultNotFound;
      return {
        result: {
          ...t.result,
          data: createdItems,
        },
        status: 200,
        message: CODE_MESSAGE[200],
      };
    },
    UPDATE: async function() {
      let updatedItems = await t.model.$update({data: t.data});
      if (!updatedItems) return t.resultNotFound;

      return {
        result: {
          ...result,
          data: updatedItems,
        },
        status: 200,
        message: CODE_MESSAGE[200],
      };
    },
    DELETE: async function() {
      t.data.map(async item => {
        await t.model.$delete(item._id);
      });
      return {
        result: {
          ...t.result,
          data: [],
        },
        status: 200,
        message: CODE_MESSAGE[200],
      };
    },
    GET: async function() {
      let allItems = await t.model.$fetch();
      if (!allItems) return t.resultNotFound;
      return {
        result: {
          ...t.result,
          data: allItems,
        },
        status: 200,
        message: CODE_MESSAGE[200],
      };
    },
  };
};

t.prototype.request = async ({method, entity, data, params = {}}) => {
  if (models[entity] === undefined) {
    throw new TypeError(`Not found ${entity} entity`);
  }

  t.method = method;
  t.namespace = entity;
  t.params = params;
  t.model = models[entity].model;
  t.fields = models[entity].model.fieldsKeys();
  if(t.model === undefined){
    throw new TypeError(`Not found model of ${entity} entity`);
  }
  console.log(t.method, t.model, entity, t.data);

  return await t.execute[method]();
};

export default t;