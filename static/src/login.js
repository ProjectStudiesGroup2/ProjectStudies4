import React from 'react';
import DocumentTitle from 'react-document-title';
 
export default class LoginPage extends React.Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event) {
    this.setState({value: event.target.value});
  }

  handleSubmit(event) {
    alert('A name was submitted: ' + this.state.value);
    event.preventDefault();
  }

  render() {
    return (
      <DocumentTitle title={`Login`} className="container">
        <div className="container">
          <form onSubmit={this.handleSubmit}>
            <h3 className="center">
              Sign in!
            </h3>
            <label><b>Username</b></label>
            <input type="text" placeholder="Username" name="uname" id="user" value={this.state.value} onChange={this.handleChange} />
            <label><b>Password</b></label>
            <input type="password" placeholder="Password" name="psw" id="pass" required />
                
            <button className="button" type="submit" value="Submit">Log in!</button>
            <input type="checkbox" /> Remember me
            <a className="psw"><span>Forgot password?</span></a>
          </form>
        </div>
      </DocumentTitle>
    );
  }
}
