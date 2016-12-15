import React from 'react'
import SimpleSlider from './slider'
import ProdSlider from './productSlider'

const Home = () => (
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
                        <p>{this.state.name}</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
)

export default Home;
