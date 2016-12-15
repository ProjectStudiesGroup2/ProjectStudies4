import React from 'react'
import { Router, Route, IndexRoute, hashHistory } from 'react-router'
import Container from './container'
import axios from 'axios'
import Home from './home'
import Catalog from './catalog'
import Filters from './filters'
import Search from './search'
import About from './about'
import LoginPage from './login'
import RegistrationPage from './register'
// import ProfilePage from './profile'
import Product from './productPage'

var TestComponent = React.createClass({
  getInitialState: function() {
    return { name: [] }
  },

  componentDidMount: function() {
    var instance = axios.create({
    });
    instance.get('https://g2p4.herokuapp.com/api/categories').then(function(response) {
      this.setState({ name: response })
    }.bind(this));
  },

  render: function() {
    return <div>{this.state.name}</div>
  }
});


// ReactDOM.render(<TestComponent />, document.getElementById('app'));

//*********//
// React.createClass({
//   getInitialState: function() {
//     return {
//       name: ''
//     }
//   },
//
//   componentDidMount: function(){
//     var instance = axios.create({
//       baseURL: 'https://g2p4.herokuapp.com/api'
//     });
//
//     instance.get('https://g2p4.herokuapp.com/api/categories')
//     .then(function (response){
//       this.setState({ name: response.data });
//     }.bind(this));
//   },
//
//   render: function() {
//     return (
//       <div>{this.state.name}</div>
//     )
//   }
// });

//******//
// var instance = axios.create({
//   baseURL: 'https://g2p4.herokuapp.com/api',
//   headers: {'Access-Control-Allow-Origin': 'http://localhost:3000'}
// });
// instance.get('/categories')
// .then(function (response){
//   console.log(response)
// });

const NotFound = () => (
  <div className="container-fluid">
    <h1>404.. This page is not found!</h1>
  </div>
)

const Instagram = () => <h3>Instagram Feed</h3>
const TwitterFeed = () => <h3>Twitter Feed</h3>

class App extends React.Component {
  render () {
    return (
      <Router history={hashHistory}>
        <Route path='/' component={Container}>
          <IndexRoute component={Home} />
          <Route path='/catalog' component={Catalog}>
            <IndexRoute component={TwitterFeed} />
            <Route path='instagram' component={Instagram} />
          </Route>
          <Route path='/filters' component={Filters} />
          <Route path='/search' component={Search} />
          <Route path='/about' component={About} />
          <Route path='/login' component={LoginPage} />
          <Route path='/register' component={RegistrationPage} />
          {/*<Route path='/profile' component={ProfilePage} />*/}
          <Route path='/product' component={Product} />
          <Route path='*' component={NotFound} />
        </Route>
      </Router>
    )
  }
}

export default App
