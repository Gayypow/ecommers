import React, {useState} from 'react';
import Icon from '@mdi/react'
import { mdiAccount } from '@mdi/js'
import { mdiCartOutline } from '@mdi/js';
import { mdiMagnify } from '@mdi/js';
import Properties from './properties';


const Header = (props) => {
    const [bool, setbool] = useState(false);
    const toogle = () =>{
       setbool(!bool)
    } 
    return (
        <div className='header'>
            <nav >
                <p>MERCI</p>
                <div className='icons'>
                {bool ? 
                <div className='searchbox'>
                <input type="text" placeholder='Haryt Ady Boyunca Gozle'/>
                <div className='inside_searchbox' onClick={toogle}><Icon path={mdiMagnify }></Icon></div>
                </div> : 
                <div className='icon search' onClick={toogle}><Icon path={mdiMagnify }></Icon></div>
                }      
                    <div className='icon'><Icon path={mdiCartOutline}></Icon></div>
                    <div className='icon'><Icon path={mdiAccount}></Icon></div>
                </div>
            </nav>
            <div className='downHeader'>
                <Properties/>
            </div>
        </div>
    )
}


export default Header;