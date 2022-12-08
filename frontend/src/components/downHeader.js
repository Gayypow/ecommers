import React from 'react'
import Properties from './properties';
import OrdersModal from './ordersModal';
import UsersModal from './usersModal';
import "../styles/header/downHeader.scss"

const DownHeader = (props) => {
  return (
    <div>
        <div className='downHeader'>
            <Properties/>
        </div>
        <OrdersModal rbool={props.rbool} toogleRegistry={props.toogleRegistry}/>
        <UsersModal  rbool={props.rbool}  toogleRegistry={props.toogleRegistry} toogleUserModal={props.toogleUserModal} ubool={props.ubool}/>
    </div>
  )
}

export default DownHeader
