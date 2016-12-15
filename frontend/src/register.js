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
    if (document.getElementById('email').value == "" && document.getElementById('user').value == "" && 
        document.getElementById('pass').value == "" && document.getElementById('repass').value == "" ) {
      alert('Some important fields should not be empty!');
      event.preventDefault();
    }
    else if (document.getElementById('pass').value != document.getElementById('repass').value){
      alert('Passwords do not match!');
    }
    else {
      alert('You are now registered as a ' + this.state.value + '!');
    }
  }

  render() {
    return (
      <DocumentTitle title={`Registration page`} className="container">
        <div className="container">
          <h3 className="center">
            Register new account!
          </h3>
          <form onSubmit={this.handleSubmit} className="form">    
            <label><b>Email *</b></label>
            <input type="text" placeholder="Username" name="uname" id="email" /><br />        
            <label><b>Username *</b></label>
            <input type="text" placeholder="Username" name="uname" id="user" value={this.state.value} onChange={this.handleChange} /><br />
            <label><b>Password *</b></label>
            <input type="password" placeholder="Password" name="psw" id="pass" required />
            <label><b>Repeat password *</b></label>
            <input type="password" placeholder="Password" name="psw" id="repass" required />
                
            <button className="button" type="submit" value="Submit">Save info!</button>
            <input type="checkbox" /> Remember me
            <a className="psw"><span>Forgot password?</span></a>
          </form>
        </div>
      </DocumentTitle>
    );
  }
}
