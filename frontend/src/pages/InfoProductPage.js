import React, {Component} from 'react';
import Header from '../components/Header';
import Footer from '../components/Footer';
import Harytlar from '../components/Harytlar';
import InfoProduct from '../components/InfoProduct';
import SameProducts from '../components/SameProducts';
import '../styles/App.scss'


export class InfoProductPage extends Component {
  render() {
    return (
      <div>
        <Header/>
        <InfoProduct/>
        <SameProducts/>
        <Footer/>
      </div>
    )
  }
}

export default InfoProductPage