import React from 'react';
import Slider from 'react-slick';
import "../styles/reklam/react-slick/slick.scss"; 
import "../styles/reklam/react-slick/slick-theme.scss";


const Reklam = () =>{
    const settings = {
        dots: true,
        infinite: true,
        speed: 500,
        slidesToShow: 3,
        slidesToScroll: 1,
        arrows: false
      };
    return(
      <div className='Reklam'>
        <Slider {...settings} >
            <div className='Reklam-wrap'>
                <img src={require('../sources/Rectangle 1.png')}/>
            </div>
            <div className='Reklam-wrap'>
                <img src={require('../sources/Rectangle 2.png')}/>
            </div>
            <div className='Reklam-wrap'>
                <img src={require('../sources/Rectangle 3.png')}/>
            </div>
            <div className='Reklam-wrap'>
                <img src={require('../sources/Rectangle 1.png')}/>
            </div>
        </Slider>
      </div>
    )
}

export default Reklam; 