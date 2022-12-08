import React from 'react'
import Icon from '@mdi/react'
import { mdiAccount } from '@mdi/js'
import { mdiCartOutline } from '@mdi/js';
import { mdiMagnify } from '@mdi/js';
import "../styles/header/HeaderNavigation.scss"


const HeaderNavigation = (props) => {
  return (
        <nav >
        <p>MERCI</p>
            <div className='icons'>
            {props.bool ? 
            <div className='searchbox'>
            <input type="text" placeholder='Haryt Ady Boyunca Gozle'/>
            <div className='inside_searchbox' onClick={props.toogle}><Icon path={mdiMagnify }></Icon></div>
            </div> : 
            <div className='icon search' onClick={props.toogle}><Icon path={mdiMagnify }></Icon></div>
            }      
                <div className='icon' onClick={props.toogleRegistry}><Icon path={mdiCartOutline}></Icon></div>
                <div className='icon' onClick={props.toogleUserModal}><Icon path={mdiAccount}></Icon></div>
            </div>
        </nav>
  )
}

export default HeaderNavigation
