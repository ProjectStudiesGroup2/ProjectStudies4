import React from 'react'
import Nav from './navigation'

const Container = (props) => (
  <div> 
    <div>
      <div>
        <Nav />
      </div>
      <div>
        {props.children} 
      </div>
    </div>
  </div>
)

export default Container
