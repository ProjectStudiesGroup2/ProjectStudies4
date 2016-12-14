import React from 'react'
import DocumentTitle from 'react-document-title'

export default class Product extends React.Component {
  render() {
    return (
      <DocumentTitle title={`Product page`}>
        <div className="container-fluid">
        <p className="margin-top"> <a>Category</a> > <a>Type</a> > <a>Product</a></p>
          <div className="row margin-top">
            <div className="col-xs-5">
              <img alt="prod" src="img/product/discription/mercg65.jpg" className="product_img_big"/>
            </div>
            <div className="col-xs-7">
              <p>Or kind rest bred with am shed then. In raptures building an bringing be. Elderly is detract 
              tedious assured private so to visited. Do travelling companions contrasted it. Mistress strongly 
              remember up to. Ham him compass you proceed calling detract. Better of always missed we person mr. 
              September smallness northward situation few her certainty something.</p>
            </div>
          </div>
          <div className="row margin-top">
            <div className="col-xs-12">
              <Tabs>
                <Pane label="Tab 1">
                  <div>This is my tab 1 contents!</div>
                </Pane>
                <Pane label="Tab 2">
                  <div>This is my tab 2 contents!</div>
                </Pane>
                <Pane label="Tab 3">
                  <div>This is my tab 3 contents!</div>
                </Pane>
              </Tabs>
            </div>
          </div>
        </div>
      </DocumentTitle>
    );
  }
}

const Tabs = React.createClass({
  displayName: 'Tabs',
  propTypes: {
    selected: React.PropTypes.number,
    children: React.PropTypes.oneOfType([
      React.PropTypes.array,
      React.PropTypes.element
    ]).isRequired
  },
  getDefaultProps() {
    return {
      selected: 0
    };
  },
  getInitialState() {
    return {
      selected: this.props.selected
    };
  },
  handleClick(index, event) {
    event.preventDefault();
    this.setState({
      selected: index
    });
  },
  _renderTitles() {
    function labels(child, index) {
      let activeClass = (this.state.selected === index ? 'active' : '');
      return (
        <li key={index}>
          <a href="#" 
            className={activeClass}
            onClick={this.handleClick.bind(this, index)}>
            {child.props.label}
          </a>
        </li>
      );
    }
    return (
      <ul className="tabs__labels">
        {this.props.children.map(labels.bind(this))}
      </ul>
    );
  },
  _renderContent() {
    return (
      <div className="tabs__content">
        {this.props.children[this.state.selected]}
      </div>
    );
  },
  render() {
    return (
      <div className="tabs">
        {this._renderTitles()}
        {this._renderContent()}
      </div>
    );
  }
});

const Pane = React.createClass({
  displayName: 'Pane',
  propTypes: {
    label: React.PropTypes.string.isRequired,
    children: React.PropTypes.element.isRequired
  },
  render() {
    return (
      <div>
        {this.props.children}
      </div>
    );
  }
});
 