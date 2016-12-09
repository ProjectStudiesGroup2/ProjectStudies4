import React from 'react'
import { Link, IndexLink } from 'react-router'
import App from './app'

var Nav = React.createClass ({
  render: function() {
    return (
      <div>
        <IndexLink activeClassName='active' to='/'>Home</IndexLink>&nbsp;
        <IndexLink activeClassName='active' to='/filters'>Filters</IndexLink>&nbsp;
        <IndexLink activeClassName='active' to='/search'>Search</IndexLink>&nbsp;
        <IndexLink activeClassName='active' to='/about'>About</IndexLink>
      </div>
    );
  }
});

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
export default Nav
