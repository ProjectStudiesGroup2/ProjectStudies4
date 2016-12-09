import React from 'react'

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
        <Content onClick={this.handleViewSidebar}/>
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

var Content = React.createClass({
  render: function() {
    // var contentClass = this.props.isOpen ? 'content open' : 'content';
    return (
      <div className="container-fluid">
        <button className='button button-red' onClick={this.props.onClick}>Filters</button>
      </div>
    );
  }
});

export default Side
