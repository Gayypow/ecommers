import React from 'react';
import '../styles/infoProduct/InfoProduct.scss'
import { connect } from 'react-redux';

const InfoProduct = (props) => {
    return (
        <div className='InfoProduct'>
            <div className='img'>
                <img src={require('../sources/Rectangle 23.png')}/>
            </div>
            <div className='exactInfo'>
                <h1>Varisocks</h1>
                <p className="name">Çagalar üçin kitap "Tebigaty goralyň" (12 sahypa)</p>
                <div className="codes">
                    <div className='productCode'> 
                        <h3>Haryt kody</h3>
                        <h3>Barkod</h3>
                    </div>
                    <div className='barCode'>
                        <p>YNMKOK02</p>
                        <p>555555</p>
                    </div>
                </div>
                <button onClick={() => {props.dispatch({type: 'ADD', id: 1, amount: 1})}}>Sebede Gosh</button>
            </div>
        </div>
    )
}

export default connect()(InfoProduct);