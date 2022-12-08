import React, { useState } from 'react';
import Modal from 'react-modal';
import {connect} from 'react-redux';
import Icon from '@mdi/react';
import { mdiClose, mdiDelete, mdiPlus } from '@mdi/js';
import '../styles/modal/OrdersModal.scss';

const OrdersModal = (props) => {
  const orders = props.orders
  const [count, setCount] = useState(0)
  // const [orders, setOrders] = useState(ordersClicked)
  return (
    <Modal
        isOpen={props.rbool}
        contentLabel="Selected Option"
        onRequestClose={props.toogleRegistry}
        closeTimeoutMS={200}
        className="modal"
        ariaHideApp={false}
        style={{
          overlay: {
            marginTop: '60px'
          }
        }}
    >
    {orders.map((order) => {
      return(
        <div className='orderItem' key={Math.random()}>
          <div className='information'>
            <h3>{order.name}</h3>
            <p>{order.price}</p>
          </div>
          <div className='actions'>
            <Icon path={mdiDelete} onClick={(e) => {props.dispatch({type: 'DELETE', id: e.id})}}></Icon>
            <div className='price'><p>{count}</p></div>
            <Icon path={mdiPlus} onClick={() => {setCount(count+1)}}></Icon>
          </div>
      </div> 
      )
    })} 
      <div className='orderInfo'>
        <div className='order_number'>
          <p>Haryt sany</p>
          <p className='number'>0</p>
        </div>
        <div className='order_number'>
          <p>Umumy bahasy</p>
          <p className='number'>0tmt  </p>
        </div>
      </div>  
      <div className='order_button'>
          <button>Sargyt etmek</button>
      </div>  
      <div className="close" onClick={props.toogleRegistry}><Icon path={mdiClose}></Icon></div>
    </Modal>
  )
}

const connectStore = (state) => {
  return {
    orders: state.orders
  }
}

export default connect(connectStore)(OrdersModal)