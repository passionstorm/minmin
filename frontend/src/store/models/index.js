const files = require.context('.', false, /\.js$/);
const models = {};
files.keys().forEach(fileName => {
  if (fileName === './index.js') return;
  const f = files(fileName);
  const entity = f.default.entity;
  models[entity] = {
    model: files(fileName).default,
    module: f.module,
  };
});

export default models;