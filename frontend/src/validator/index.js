import {extend, localize} from 'vee-validate';
import * as rules from 'vee-validate/dist/rules';

for (let rule in rules) {
  extend(rule, rules[rule]);
}
extend('phone', {
  validate: (value, {other}) => {
    return value.match(/^[\+]?[(]?[0-9]{3}[)]?[-\s\.]?[0-9]{3}[-\s\.]?[0-9]{4,6}$/im);
  },
  // message: 'Số điện thoại chưa đúng định dạng',
  // params: [{name: 'other', isTarget: true}],
});
// extend('username', (value, {other}) => {
//   if (!value.match(/^[a-zA-Z0-9]+$/)) return 'username chưa đúng';
//
//   return true;
// });

export function loadLocale(code) {
  return import(`vee-validate/dist/locale/${code}.json`).then(defaultLocale => {
    localize(code, defaultLocale);
    return import(`./${code}.json`).then(locale => localize(code, locale));
  });
}

