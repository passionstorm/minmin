import Vue from 'vue';
import Dialog from './Dialog';

import {merge} from '../../utils';
import {registerComponentProgrammatic} from '../../utils/plugins';

function open(propsData) {
  const vm = typeof window !== 'undefined' && window.Vue ? window.Vue : Vue;
  const DialogComponent = vm.extend(Dialog);
  return new DialogComponent({
    el: document.createElement('div'),
    propsData,
  });
}

const DialogProgrammatic = {
  alert(params) {
    if (typeof params === 'string') {
      params = {
        message: params,
      };
    }
    const defaultParam = {
      canCancel: false,
    };
    const propsData = merge(defaultParam, params);
    return open(propsData);
  },
  confirm(params) {
    const defaultParam = {};
    const propsData = merge(defaultParam, params);
    return open(propsData);
  },
  prompt(params) {
    const defaultParam = {
      hasInput: true,
      confirmText: 'Done',
    };
    const propsData = merge(defaultParam, params);
    return open(propsData);
  },
};

const Plugin = {
  install(Vue) {
    registerComponentProgrammatic(Vue, 'dialog', DialogProgrammatic);
  },
};

export default Plugin;

export {
  DialogProgrammatic,
};