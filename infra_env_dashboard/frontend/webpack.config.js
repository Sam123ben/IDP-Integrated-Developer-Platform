// webpack.config.js

const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const Dotenv = require('dotenv-webpack');

module.exports = {
  entry: './src/index.tsx',  // Updated to use index.tsx
  output: {
    path: path.resolve(__dirname, 'dist'),
    filename: 'bundle.js',
    clean: true // Clean the output directory before emit
  },
  module: {
    rules: [
      {
        test: /\.(ts|tsx)$/, // Add TypeScript loader for .ts and .tsx files
        exclude: /node_modules/,
        use: {
          loader: 'ts-loader'
        }
      },
      {
        test: /\.(js|jsx)$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader'
        }
      },
      {
        test: /\.css$/, // Add CSS loader configuration here
        use: ['style-loader', 'css-loader']
      }
    ]
  },
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'] // Add TypeScript extensions here
  },
  plugins: [
    new HtmlWebpackPlugin({
      template: './src/index.html', // Ensure this path is correct
      filename: 'index.html'
    }),
    new Dotenv() // Load environment variables from a .env file
  ],
  devServer: {
    static: './dist',
    hot: true,
    port: 3000,
    historyApiFallback: true
  }
};