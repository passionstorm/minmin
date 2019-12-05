import Vue from 'vue';

const filter = function(text, length, clamp) {
  clamp = clamp || '...';
  let node = document.createElement('div');
  node.innerText  = text;
  let content = node.textContent;
  return content.length > length ? content.slice(0, length) + clamp : content;
};

Vue.filter('truncate', filter);