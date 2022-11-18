import React from 'react'
import Slider from 'react-slick';
import "../styles/header/Slick Menu slider/slick.scss"; 
import "../styles/header/Slick Menu slider/slick-theme.scss"; 
import "../styles/header/slickMenu.scss" 


const Properties = () => {
    const settings = {
        dots: false,
        infinite: true,
        speed: 500,
        slidesToShow: 8,
        slidesToScroll: 1,
        arrows: false
      };
    return(
      <div className='container'>
        <Slider {...settings} >
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 15.png')}/>
                </div>
                <div className='text'>
                    <p>One chykanlar</p>
                </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 16.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 17.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 18.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 19.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 20.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 21.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
            <div className='menu'>
                <div className='menuPicture'>
                    <img src={require('../sources/Rectangle 22.png')}/>
                </div>
                <div className='text'>
                    <p>Arzanladyshlar</p>
                 </div>
            </div>
        </Slider>
      </div>
  )
}

export default Properties

