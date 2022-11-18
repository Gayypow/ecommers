import React from 'react';
import Haryt from './Haryt'
import { useState, useEffect } from 'react';
import axios from 'axios'



const infoHaryt = [{
    id: 1,
    imgSrc: '../sources/image 2.png',
    maglumat: 'Chorek 150gr',
    bahasy: '5 manat'
},
{
    id: 2,
    imgSrc: '../sources/image 2.png',
    maglumat: 'Chorek 150gr',
    bahasy: '5 manat'
}
]






const Harytlar = (props) => {
    useEffect(() => {

    async function getMyAPI() {
        const products = await axios.get(`http://192.168.1.106:9000/api/products`)
        
   
        getProduct({state: products.data.products })
      }
  
      getMyAPI()
    }, [])
    
    const [product, getProduct] = useState({state: []})
    console.log(product.state)
    
    return(
        <div>
        <div className='harytlar'>
        {
            product.state.map((e) => (<Haryt key={e.id} maglumat={e}/>))
        }
        </div>
        </div>
    )
} 


export default Harytlar

