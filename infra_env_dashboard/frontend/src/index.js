// src/index.js

import React from 'react';
import ReactDOM from 'react-dom';
import './App.css'; // Global styles, if any
import App from './App';

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById('root') // This matches the <div id="root"></div> in index.html
);
