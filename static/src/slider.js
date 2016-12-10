var React = require('react');
var Slider = require('react-slick');

var SimpleSlider = React.createClass({
  render: function () {
    var settings = {
      arrows: false,
      dots: true,
      infinite: true,
      speed: 1500,
      autoplay: true,
      fade: true,
      draggable: false,
      slidesToShow: 1,
      slidesToScroll: 1
    };
    return (
      <Slider {...settings} className="slider">
        <div><img alt="car$plane" className="sliderImg" src={'img/car&plane.jpg'} /></div>
        <div><img alt="car$plane" className="sliderImg" src={'img/car&plane2.jpg'} /></div>
        <div><img alt="car$plane" className="sliderImg" src={'img/car&plane.jpg'} /></div>
        <div><img alt="car$plane" className="sliderImg" src={'img/car&plane2.jpg'} /></div>
      </Slider>
    );
  }
});

export default SimpleSlider