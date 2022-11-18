import React from 'react'
import Nav from '../components/admin/nav'
import Table from '../components/admin/table'
import '../styles/admin/Admin.scss'

const Admin = () => {
  return (
    <div className='body'>
      <Nav/>
      <div className='second_part'>
        <header className='admin_header'></header>
        <Table/>
      </div>
    </div>
  )
}

export default Admin;
