import React, {useState} from 'react';
import HeaderNavigation from './HeaderNavigation';
import DownHeader from './downHeader';


const Header = (props) => {
    const [bool, setbool] = useState(false);
    const [rbool, setrbool] = useState(false)
    const [ubool, setubool] = useState(false)
    const toogleRegistry = () => {
        if(bool){
            setbool(!bool)
        }
        setrbool(!rbool)
    }
    const toogle = () =>{
        if(rbool){
            setrbool(!rbool)
        }
       setbool(!bool)
    } 
    const toogleUserModal = () =>{
        if(rbool){
            setrbool(!rbool)
        }
        if(bool){
            setbool(!bool)
        }
       setubool(!ubool)
    } 
    return (
        <div className='header'>
            <HeaderNavigation bool={bool} toogleRegistry={toogleRegistry} toogle={toogle} toogleUserModal={toogleUserModal}/>
            <DownHeader rbool={rbool} toogleRegistry={toogleRegistry} toogleUserModal={toogleUserModal} ubool={ubool}/>
        </div>
            )
}


export default Header;