import React from 'react'
import SimpleSlider from './slider'
import ProdSlider from './productSlider'
import DocumentTitle from 'react-document-title'

const Home = () => (
    <DocumentTitle title={`Web Store - Home page`} className="container">
        <div id="wrapper"> 
            <div className="content">
                <SimpleSlider />         
                <div className="container margin-top">                        
                    <h4>Most viewed </h4>
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
    </DocumentTitle>
)


export default Home;
