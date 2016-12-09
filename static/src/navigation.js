import React from 'react'
import { Link, IndexLink } from 'react-router'

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
                <IndexLink className="nav-tab" activeClassName='active' to='/'>Home</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/catalog'>Catalog</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/about'>About</IndexLink>
              </li>
            </ul>
            <ul className="nav navbar-nav navbar-right icon-pad">
              <li><a href="#">1</a></li>
              <li><a href="#">2</a></li>
            </ul>
          </nav>
        </div>
    </header>
  </div>
)

export default Nav