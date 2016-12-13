var React = require('react');
var Slider = require('react-slick')
import { IndexLink } from 'react-router'

var ProdSlider = React.createClass({
  render: function () {
    var settings2 = {
      arrows: true,
      dots: false,
      infinite: true,
      draggable: true,
      speed: 500,
      slidesToShow: 3,
      slidesToScroll: 3
    };
    return (
    <div className="container product_slider_cont">
      <Slider {...settings2} className="product_slider">
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
      </Slider>
      <br />
      <Slider {...settings2} className="product_slider">
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/mercg65s.png" /></div>
                <div id="product_header">Mercedes-AMG G65</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
        <div>
            <IndexLink to='/product/:id'><div id="product_box">
                <div id="product_img"><img alt="prod" src="img/product/shortcut/drone.png" /></div>
                <div id="product_header">Passenger Drone EHANG184</div>
                <div id="product_price">100000$</div>
            </div></IndexLink>
        </div>
      </Slider>
      <br />
    </div>
    );
  }
});

export default ProdSlider
