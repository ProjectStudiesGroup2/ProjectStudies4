import React from 'react';
import ReactDOM from 'react-dom';
// import {ReactRouter, Router, Route} from 'react-router';
import MainP from './MainP';

ReactDOM.render(
  <ReactRouter.Router>
    <ReactRouter.Route path="./static" component={MainP}>

    </ReactRouter.Route>
  </ReactRouter.Router>,
  document.getElementById('mainP')
);
