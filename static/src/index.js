import React from 'react'
import ReactDOM from 'react-dom'
import ReactStormpath from 'react-stormpath'
import App from './app'

ReactStormpath.init();
ReactDOM.render(
  <App />,
  document.getElementById('app')
);
