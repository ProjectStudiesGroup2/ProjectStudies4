import React from 'react';
import DocumentTitle from 'react-document-title';
import { LoginForm } from 'react-stormpath';
 
export default class LoginPage extends React.Component {
  render() {
    return (
      <DocumentTitle title={`Login`} className="container">
        <div>
          <div className="row">
            <div className="col-xs-12">
              <h3 className="center">Login</h3>
              <hr />
            </div>
          </div>
          <LoginForm />
        </div>
      </DocumentTitle>
    );
  }
}