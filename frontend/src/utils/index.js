// export function camel (str) {
//   const camel = (str || '').replace(/-([^-])/g, g => g[1].toUpperCase());

//   return capitalize(camel);
// }

export const lazySorter = (sortKey) => (order) => (data) => {
  return data.slice().sort(compareObjectValues(sortKey)(order));
};

export function removeElement(el) {
  if (typeof el.remove !== 'undefined') {
    el.remove()
  } else if (typeof el.parentNode !== 'undefined') {
    el.parentNode.removeChild(el)
  }
}

export const isMobile = {
  Android: function () {
    return (
        typeof window !== 'undefined' &&
        window.navigator.userAgent.match(/Android/i)
    )
  },
  BlackBerry: function () {
    return (
        typeof window !== 'undefined' &&
        window.navigator.userAgent.match(/BlackBerry/i)
    )
  },
  iOS: function () {
    return (
        typeof window !== 'undefined' &&
        window.navigator.userAgent.match(/iPhone|iPad|iPod/i)
    )
  },
  Opera: function () {
    return (
        typeof window !== 'undefined' &&
        window.navigator.userAgent.match(/Opera Mini/i)
    )
  },
  Windows: function () {
    return (
        typeof window !== 'undefined' &&
        window.navigator.userAgent.match(/IEMobile/i)
    )
  },
  any: function () {
    return (
        isMobile.Android() ||
        isMobile.BlackBerry() ||
        isMobile.iOS() ||
        isMobile.Opera() ||
        isMobile.Windows()
    )
  }
}

const mergeFn = (target, source, deep = false) => {
  if (deep || !Object.assign) {
    const isDeep = (prop) =>
        isObject(source[prop]) && target.hasOwnProperty(prop) && isObject(target[prop])
    const replaced = Object.getOwnPropertyNames(source)
    .map((prop) =>
        ({ [prop]: isDeep(prop) ? mergeFn(target[prop], source[prop]) : source[prop] }))
    .reduce((a, b) => ({ ...a, ...b }), {})

    return {
      ...target,
      ...replaced
    }
  } else {
    return Object.assign(target, source)
  }
}
export const merge = mergeFn

export function escapeRegExpChars(value) {
  if (!value) return value

  // eslint-disable-next-line
  return value.replace(/[\-\[\]\/\{\}\(\)\*\+\?\.\\\^\$\|]/g, '\\$&')
}


export function getValueByPath(obj, path) {
  const value = path.split('.').reduce((o, i) => o ? o[i] : null, obj)
  return value
}

/**
 * Extension of indexOf method by equality function if specified
 */
export function indexOf(array, obj, fn) {
  if (!array) return -1

  if (!fn || typeof fn !== 'function') return array.indexOf(obj)

  for (let i = 0; i < array.length; i++) {
    if (fn(array[i], obj)) {
      return i
    }
  }

  return -1
}

const isObject = (item) => typeof item === 'object' && !Array.isArray(item)

export const compareObjectValues = (key) => (order) => {
  return (a, b) => comparePairs(a[key])(b[key])(order);
};

export const comparePairs = a => b => (order) => {
  const notEqualCompare = a > b ? 1 : -1;
  const equalCompare = a === b ? 0 : notEqualCompare;
  return equalCompare * order;
};

export const push = (arr, newItem, index = -1) => {
  if (index !== -1) return [
    ...arr.slice(0, index),
    newItem,
    ...arr.slice(index),
  ];

  return [...arr, newItem];
};

export const lazyFilter = (filter) => (data) => {
  return data.reduce((list, item) => {
    Object.keys(item).some(key => {
      if (checkStringMatch(item[key])(filter)) {
        list.push(item);
        return true;
      }
    });
    return list;
  }, []);
};

export const checkStringMatch = (test) => (filter) => {
  return String(test).toLowerCase().indexOf(filter) > -1;
};

export const baseFilter = ({sort, search}, items = []) => {
  let filter = search && search.toLowerCase();
  let order = 1;
  let data = items;
  data = lazyFilter(filter)(data);
  data = lazySorter(sort)(order)(data);
  return data;
};

// export function camelActual (str) {
//   return (str || '').replace(/-(\w)/g, (_, c) => (c ? c.toUpperCase() : ''));
// }

// export function kebab (str) {
//   return (str || '').replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase();
// }

// export function capitalize (str) {
//   str = str || '';

//   return `${str.substr(0, 1).toUpperCase()}${str.slice(1)}`;
// }

// export function findProduct (store, id) {
//   return store.state.store.products.find(p => p.id === id);
// }

// export function isOnSale (variants) {
//   return variants.some(variant => {
//     return parseFloat(variant.price) < parseFloat(variant.compareAtPrice);
//   });
// }

// export function randomNumber (min, max) {
//   return Math.floor(Math.random() * max) + min;
// }

// export function randomString (length = 5) {
//   let text = '';
//   const possible = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';

//   for (let i = 0; i < length; i++) {
//     text += possible.charAt(Math.floor(Math.random() * possible.length));
//   }

//   return text;
// }

const randomElement = (arr) => {
  return arr[Math.floor(Math.random() * arr.length)];
};

const kebab = str => {
  return (str || '').replace(/([a-z])([A-Z])/g, '$1-$2').toLowerCase();
};

const toggleFullScreen = () => {
  let doc = window.document;
  let docEl = doc.documentElement;

  let requestFullScreen =
      docEl.requestFullscreen || docEl.mozRequestFullScreen ||
      docEl.webkitRequestFullScreen || docEl.msRequestFullscreen;
  let cancelFullScreen =
      doc.exitFullscreen || doc.mozCancelFullScreen || doc.webkitExitFullscreen ||
      doc.msExitFullscreen;

  if (!doc.fullscreenElement && !doc.mozFullScreenElement &&
      !doc.webkitFullscreenElement && !doc.msFullscreenElement) {
    requestFullScreen.call(docEl);
  } else {
    cancelFullScreen.call(doc);
  }
};

export default {
  randomElement,
  toggleFullScreen,
  kebab,
};