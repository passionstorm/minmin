const fs = require('fs');
const chalk = require('chalk');

class ManifestPlugin {
  constructor(options) {
    this.options = options;
  }

  apply(compiler) {
    const {outputPath, fileName = 'manifes.json'} = this.options;

    compiler.hooks.done.tap('Custom Manifest', stats => {
      const assetsManifest = [];
      const {assets, hash} = stats.compilation;
      const content = {
        hash: hash,
        assets: [],
      };
      Object.keys(assets).map(name => {
        let size = assets[name]['_cachedSize'];
        if (!isNaN(size)) {
          let s = size > 1000 ? Math.round(size / 1000) + ' KB' : size + ' B';
          content.assets.push({name, s});
        }
      });

      try {
        let filePath = outputPath + '/' + fileName;
        fs.writeFileSync(filePath, JSON.stringify(content, null, 4));

        console.log(chalk.green.bold('\n Manifest generated'));
      } catch (error) {
        console.log(chalk.bold.bgRed('Exception:'), chalk.bold.red(error.message));
      }
    });
  }
}

module.exports = ManifestPlugin;