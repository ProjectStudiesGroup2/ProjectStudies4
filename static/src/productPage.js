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
        </div>
      </DocumentTitle>
    );
  }
}