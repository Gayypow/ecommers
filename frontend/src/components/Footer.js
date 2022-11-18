import React from 'react';

const Footer = () => {
    return(
        <div>
        <div className='footer'>
            <div className='merci'>merci</div>
            <div className='market'>
                <ul>
                    <li><h3>Market</h3></li>
                    <li>Biz barada</li>
                    <li>Eltip bermek we toleg tertibi</li>
                    <li>Aragatnashyk</li>
                    <li>c</li>
                </ul>                                                       
            </div>
            <div className='habarlashmak'>
                <ul>
                    <li><h3>Habarlashmak ucin</h3></li>
                    <li>Telefon: +99365266124</li>
                    <li>IMO: +99361286522</li>
                    <li>E-mail: info@merci.com</li>
                    <li>@merci_.com </li>
                </ul>
            </div>
            <div className='mobileprogram'>
                <h3>Mobile Programmalar</h3>
                <img src={require('../sources/image 7.png')} />
                <img src={require('../sources/image 8.png')} />
                <img src={require('../sources/image 9.png')} />
            </div>
        </div>
        <div className='bottom'>
        <p>© 2019-2022 merci.com. Ähli hukuklary goraglydyr.</p>
        </div>        
        </div>
    )
} 


export default Footer