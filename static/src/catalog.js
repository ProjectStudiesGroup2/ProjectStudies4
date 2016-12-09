import React from 'react'
import { Link, IndexLink } from 'react-router'

const Catalog = (props) => 
<div className="container-fluid">
  <Link to='/catalog'>Twitter Feed</Link>&nbsp;
  <Link to='/catalog/instagram'>Instagram Feed</Link>
  <h1>We are located at 555 Jackson St.</h1>
  {props.children}
</div>

export default Catalog