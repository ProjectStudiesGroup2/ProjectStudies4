import React from 'react'
import DocumentTitle from 'react-document-title'
 
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
    if (document.getElementById('user').value=="admin" && document.getElementById('pass').value=="admin" ||
        document.getElementById('user').value=="roman" && document.getElementById('pass').value=="12345") {
      alert('You are loged in! Hello, ' + this.state.value + '!');
      event.preventDefault();
    }
    else if (document.getElementById('user').value=="admin" || document.getElementById('user').value=="roman"){
      alert('Wrong password!');
    }
    else {
      alert('No such usrname or password!');
    }
  }

  render() {
    return (
      <DocumentTitle title={`Login page`} className="container">
        <div className="container">
          <h3 className="center">
            Sign in!
          </h3>
          <form onSubmit={this.handleSubmit} className="form">            
            <label><b>Username</b></label>
            <input type="text" placeholder="Username" name="uname" id="user" value={this.state.value} onChange={this.handleChange} /><br />
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
