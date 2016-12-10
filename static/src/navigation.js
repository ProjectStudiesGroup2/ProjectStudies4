import React from 'react'
import { IndexLink } from 'react-router'

const Nav = () => (
  <div>
    <header className="navbar navbar-static-top bar-style">
        <div className="container">
          <div className="navbar-header">
            <button type="button" className="navbar-toggle collapsed glyphicon glyphicon-align-justify" data-toggle="collapse" data-target="#bs-navbar" aria-expanded="false">
              <span className="sr-only">Toggle navigation</span>
              <span className="icon-bar"></span>
            </button>					
            <a href="#" className="navbar-brand logo">Web Store</a>
          </div>
          <nav id="bs-navbar" className="collapse navbar-collapse">          
            <ul className="nav navbar-nav nav-tabs">
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/'>Home</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/catalog'>Catalog</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/about'>About</IndexLink>
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/filters'>Filters*</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="activeTab" activeClassName='active' to='/search'>Search</IndexLink>&nbsp;
              </li>
            </ul>
            <ul className="nav navbar-nav navbar-right">
              <li><IndexLink to='/login'>Sign in</IndexLink></li>
              <li><IndexLink to='/register'>Register</IndexLink></li>
              <li><IndexLink to='/cart'>Cart</IndexLink></li>
            </ul>
          </nav>
        </div>
    </header>
  </div>
)

export default Nav
