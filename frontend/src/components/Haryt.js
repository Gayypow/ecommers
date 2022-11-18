import React, {useState} from 'react';
import {Link} from 'react-router-dom';
import { connect } from 'react-redux';


const Haryt = (props) => {
    const [count, setCount] = useState(0)
    
    return(
        <div className='harytMaglumat'>
            <Link to={`/products/${props.maglumat.id}`} className='picture'>
                <img src={require('../sources/image 2.png')}/>
            </Link>
            <p className='maglumat'>{props.maglumat.name.tk}</p>
            <p className='Bahasy'>{props.maglumat.price}</p>
            <button onClick={() => {console.log('added')}}>Sebede gosh</button>
        </div>
    )
} 



export default connect()(Haryt);

