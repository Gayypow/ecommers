import React from 'react'
import Header from '../components/Header'
import Reklam from '../components/Reklam';
import Harytlar from '../components/Harytlar';
import Footer from '../components/Footer';

const Home = () => {
  return (
    <div>
        <Header/>
        <Reklam/>
        <Harytlar/>
        <Footer/>
    </div>
  )
}

export default Home;