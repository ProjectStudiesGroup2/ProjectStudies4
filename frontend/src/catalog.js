import React from 'react'
import { Link } from 'react-router'

var Side = React.createClass({
  getInitialState: function(){
    return {sidebarOpen: false};
  },
  handleViewSidebar: function(){
    this.setState({sidebarOpen: !this.state.sidebarOpen});
  },
  render: function(){
    return (
      <div>
        <Catalog onClick={this.handleViewSidebar}/>
        <SideBar isOpen={this.state.sidebarOpen}/>
      </div>
    );
  }
});

var SideBar = React.createClass({
  render: function() {
    var sidebarClass = this.props.isOpen ? 'sidebar open' : 'sidebar';
    return (
      <div className={sidebarClass}>
        <div>Contents</div>
      </div>
    );
  }
});

var Catalog = React.createClass({
  render: function() {
    return (
  <div className="container-fluid">
    <h3>Hello from catalog!</h3>
    <button className='button button-red' onClick={this.props.onClick}>Tab of filters</button>
  </div> )
  }
});

export default Side
