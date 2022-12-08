import React, {useState} from 'react'
import Icon from '@mdi/react'
import { mdiDelete, mdiPlus, mdiClose } from '@mdi/js';


const Table = () => {
    const [active, setActive] = useState(true)
    const [name, setName] = useState('')
    const [price, setPrice] = useState('')
    const [products, setProducts] = useState([])
   
    const Submit = (e) => {
        e.preventDefault()
        console.log(name, price)
        setName('')
        setPrice('')
        setProducts([...products, {name, price}])
        setActive(true)
    }

    const Delete = (e) => {
        console.log(e)
    }
  return (
    <div className='t_container'>
        <div className='bigPlus' onClick={() => {setActive(false)}}>
            <Icon path={mdiPlus}  className="Plus"></Icon>
        </div>
        {
            active ? (
                    <table>
                        <thead>
                            <tr>
                                <th>Id:</th>
                                <th>Name:</th>
                                <th>Price:</th>
                                <th>Action:</th>
                     </tr>
                        </thead>
                        <tbody>
                           {
                            products.map((product) => (
                                <tr key={product.name}>
                                    <td>1</td>
                                    <td>{product.name}</td>
                                    <td>{product.price}</td>
                                    <td className='actions'><Icon path={mdiPlus}  className="Plus" onClick={() => {console.log('Clicked plus')}}></Icon><Icon path={mdiDelete} className="Delete" onClick={() => {setProducts(products.filter((p) => {
                                        return p.name !== product.name
                                    }))}} ></Icon>
                                    </td>
                                </tr>
                            ))
                           }
                        </tbody>
                    </table>
                    ): 
                    (
                <div className='AddProduct'>
                    <Icon path={mdiClose} onClick={() => {setActive(true)}}></Icon> 
                    <input type="text" placeholder='Name' value={name} onChange={(e) => {setName(e.target.value)}}/>
                    <input type="text" placeholder='Price'  value={price} onChange={(e) => {setPrice(e.target.value)}}/>
                    <button onClick={Submit}>Add product</button>
                </div>
            )
        }
    </div>
  )
}

export default Table
