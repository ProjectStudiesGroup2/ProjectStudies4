import React from 'react'
import { Link, IndexLink } from 'react-router'
import Nav from './navigation'

const Container = (props) => (
  <div> 
    <div>
      <Nav />
    </div>
    <div>
      {props.children} 
    </div>
  </div>
)

export default Container
  