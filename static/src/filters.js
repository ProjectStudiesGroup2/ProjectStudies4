import React from 'react'
import styles from '/sideStyle.css'
import {Nav, NavItem, Navbar, NavDropdown, MenuItem, Glyphicon} from 'react-bootstrap';

var Side = React.createClass({
  render: function() {
    return <div id="sidebar-menu" className={styles.sideBarMenuContainer}>
      <Navbar fluid className={styles.sidebar} inverse>
        <Navbar.Collapse>
          <Nav>
            <NavDropdown eventKey={1} title="Item 1">
              <MenuItem eventKey={1.1} href="#">Item 1.1</MenuItem>
            </NavDropdown>
            <NavItem eventKey={2}>Item 2</NavItem>
            <NavItem eventKey={3}>Item 3</NavItem>
          </Nav>
        </Navbar.Collapse>

      </Navbar>
    </div>;
  }
});

// var Side = React.createClass({
//   getInitialState: function(){
//     return {sidebarOpen: false};
//   },
//   handleViewSidebar: function(){
//     this.setState({sidebarOpen: !this.state.sidebarOpen});
//   },
//   render: function(){
//     return (
//       <div>
//         <Content onClick={this.handleViewSidebar}/>
//         <SideBar isOpen={this.state.sidebarOpen}/>
//       </div>
//     );
//   }
// });
//
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
//
// var Content = React.createClass({
//   render: function() {
//     // var contentClass = this.props.isOpen ? 'content open' : 'content';
//     return (
//       <div className="container-fluid">
//         <button className='button button-red' onClick={this.props.onClick}>Filters</button>
//       </div>
//     );
//   }
// });

export default Side
