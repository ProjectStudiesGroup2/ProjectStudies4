import React from 'react'
import SimpleSlider from './slider'
import ProdSlider from './productSlider'

const Home = () => (
    <div id="wrapper"> 
        <div className="content">
            <SimpleSlider />           
            <div className="container">            
                <h3>Most viewed </h3>
            </div> 
            <ProdSlider /> 
        </div>        
        <div id="footer">
            <div className="container">	
                <div className="row">
                    <div className="col-sm-7">
                        <p>Something here</p>
                    </div>
                </div>			
            </div>
        </div> 
    </div> 
)

export default Home;
