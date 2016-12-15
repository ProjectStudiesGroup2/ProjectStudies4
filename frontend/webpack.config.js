path = require('path');
const webpack = require('webpack');

module.exports = {
  devtool: 'eval',
  entry: [
    './src/index.js',
  ],
  output: {
    path: path.join(__dirname, 'public'),
    filename: 'bundle.js',
    publicPath: '/static/',
  },
  module: {
    loaders: [{
        exclude: /node_modules/,
        test: /\.(js|jsx)$/,
        loader: 'babel-loader',
        query: {
          presets: ['es2015', 'react']
        }
      }
    ],
  },
};
