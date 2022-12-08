import React, {useState} from 'react'
import Modal from 'react-modal';
import '../styles/modal/UsersModal.scss'

const UsersModal = (props) => {

    const [number, setNumber] = useState('')
    const [code, setCode] = useState('')

    const submitData = (e) => {
        e.preventDefault();
        console.log(number, code)
        setNumber('')
        setCode('')
        props.toogleUserModal()

    } 
  return (
    <Modal
        isOpen={props.ubool}
        contentLabel="Selected Option"
        onRequestClose={props.toogleUserModal}
        closeTimeoutMS={200}
        className="userModal"
        ariaHideApp={false}
    >
    <form className='um_container'>
        <h1 className='um_text'>Ulgama Gir</h1>
        <div className='phone_number'>
            <p>Telefon belginiz</p>
            <input type="text" placeholder='+99365266124' value={number} onChange={(e) => {setNumber(e.target.value)}}/>
        </div>
        <div className='phone_number code'>
            <div className='code_p'>
                <p>Acar soz</p>
                <p className='not'>Yatda yokmy?</p>
            </div>
            <input type="text" value={code} onChange={(e) => {setCode(e.target.value)}}/>
        </div>
        <div className='checkbox'>
            <input type="checkbox"/>
            <p>Maglumaty 30 günlik ýatda sakla</p>
        </div>
        <button onClick={submitData} >Sign in</button>
        <div className='checkbox bottom'>
            <p className='question'>Ulgama girmedinmi?</p>
            <p className='registration'>Registrasiya</p>
        </div>
    </form>
    </Modal>
  )
}

export default UsersModal
