import React from 'react'
import { Router, Route, Link, IndexRoute, hashHistory, IndexLink } from 'react-router'
import Filters from './filters'
import Home from './home'
import Search from './search'
import About from './about'
import Container from './container'
import Nav from './navigation'

class App extends React.Component {
  render () {
    return (
      <Router history={hashHistory}>
        <Route path='/' component={Container}>
          <IndexRoute component={Home} />
          <Route path='/filters' component={Filters} />
          <Route path='/search' component={Search} />
          <Route path='/about' component={About} />
          <Route path='*' component={NotFound} />
        </Route>
      </Router>
    )
  }
}

const NotFound = () => <h1>404.. This page is not found!</h1>

export default App
