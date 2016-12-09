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
                <IndexLink className="nav-tab" activeClassName='active' to='/'>Home</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/catalog'>Catalog</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/about'>About</IndexLink>
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/filters'>Filters*</IndexLink>&nbsp;
              </li>
              <li>
                <IndexLink className="nav-tab" activeClassName='active' to='/search'>Search</IndexLink>&nbsp;
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

// var Nav = React.createClass ({
//   render: function() {
//     return (
//       <div>
//         <IndexLink activeClassName='active' to='/'>Home</IndexLink>&nbsp;
//         <IndexLink activeClassName='active' to='/filters'>Filters</IndexLink>&nbsp;
//         <IndexLink activeClassName='active' to='/search'>Search</IndexLink>&nbsp;
//         <IndexLink activeClassName='active' to='/about'>About</IndexLink>
//       </div>
//     );
//   }
// });

// var SideBar = React.createClass({
//   render: function() {
//     var sidebarClass = this.props.isOpen ? 'sidebar open' : 'sidebar';
//     return (
//       <div className={sidebarClass}>
//         <div>Contents</div>
//       </div>
//     );
//   }
// });
