const CracoLessPlugin = require('craco-less');

module.exports = {
  plugins: [
    {
      plugin: CracoLessPlugin,
      options: {
        lessLoaderOptions: {
          lessOptions: {
            modifyVars: {
              '@primary-color': '#78BA4D',
              '@menu-horizontal-line-height': '60px'
            },
            javascriptEnabled: true,
          },
        },
      },
    },
  ],
};
