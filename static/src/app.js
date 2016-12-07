import React from 'react'
import { Router, Route, Link, IndexRoute, hashHistory, IndexLink } from 'react-router'

class App extends React.Component {
  render () {
    return (
      <Router history={hashHistory}>
        <Route path='/' component={Container}>
          <IndexRoute component={Home} />
          <Route path='/address' component={Address}>
            <IndexRoute component={Categ1} />
            <Route path='cat2' component={Categ2} />
          </Route>
          <Route path='/search' component={Search} />
          <Route path='/about' component={About} />
          <Route path='*' component={NotFound} />
        </Route>
      </Router>
    )
  }
}

const Nav = () => (
  <div>
    <IndexLink activeClassName='active' to='/'>Home</IndexLink>&nbsp;
    <IndexLink activeClassName='active' to='/address'>Categories</IndexLink>&nbsp;
    <IndexLink activeClassName='active' to='/search'>Search</IndexLink>&nbsp;
    <IndexLink activeClassName='active' to='/about'>About</IndexLink>
  </div>
)

const Container = (props) => <div><Nav />{props.children}</div>

const Home = () => <h1>Hello from Home!</h1>

const Address = (props) => <div>
  <br />
  <Link to='/address'>Categories</Link>&nbsp;
  <Link to='/address/cat2'>More categories</Link>
  <h1>Category page:</h1>
  {props.children}
</div>

const Categ1 = () => <h3>Some categories here</h3>
const Categ2 = () => <h3>Some more categories here</h3>

const Search = () => <h2>Coming soon</h2>

const About = () => <h3>Welcome to the About Page</h3>

const NotFound = () => <h1>404.. This page is not found!</h1>

export default App
