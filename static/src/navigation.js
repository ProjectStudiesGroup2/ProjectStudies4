import React from 'react'
import { IndexLink } from 'react-router'

const Nav = () => (
  <div>
    <header className="navbar navbar-static-top bar-style">
        <div className="container">
          <div className="navbar-header">
            <button type="button" className="navbar-toggle collapsed glyphicon glyphicon-align-justify" data-toggle="collapse" data-target="#bs-navbar" aria-expanded="false">
              <span className="sr-only">Toggle navigation</span>
            </button>					
            <a href="#" className="navbar-brand logo">Web Store</a>
          </div>
          <nav id="bs-navbar" className="collapse navbar-collapse">          
            <ul className="nav navbar-nav">
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/'>Home</IndexLink>
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/catalog'>Catalog</IndexLink>
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/about'>About</IndexLink>
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/filters'>Filters*</IndexLink>
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/search'>Search</IndexLink>
              </li>
            </ul>
            <ul className="nav navbar-nav navbar-right">
              <li><IndexLink to='/profile'>Profile</IndexLink></li>
              <li><IndexLink to='/login'>Login</IndexLink></li>
              <li><IndexLink to='/register'>Register</IndexLink></li>
              <li><IndexLink to='/cart'><i className="glyphicon glyphicon-shopping-cart"></i></IndexLink></li>
            </ul>
          </nav>
        </div>
    </header>
  </div>
)

export default Nav
