
import {bus} from '../main';
const EVT_MSG = 'msg';

export const msg = {
  fail: (msg) =>{
    bus.$emit(EVT_MSG, {
      msg: msg,
      type: 'is-error'
    })
  },
  success: (msg) => {
    bus.$emit(EVT_MSG, {
      msg: msg,
      type: 'is-success',
    })
  }
};