import React from 'react'
import {Link} from 'react-router-dom'


const Nav = () => {
  return (
    <nav className='nav'>
        <div className="logo"><p>Merci</p></div>
        <ul>
        <li><Link to={`/admin/products`} style={{textDecoration: 'none'}}><p>Products</p></Link></li>
        <li><Link to={`/admin/reklams`} style={{textDecoration: 'none'}}><p>reklams</p></Link></li>
        <li><Link to={`/admin/categories`} style={{textDecoration: 'none'}}><p>categories</p></Link></li>
        </ul>
  </nav>
  )
}


export default Nav;