import react from 'react';
import {BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import InfoProductPage from '../pages/InfoProductPage';
import Home from '../pages/Home';
import Admin from '../pages/Admin'



const AppRouter = () => {
  return (
    <Router>
      <Routes>
        <Route exact path="/" element={<Home/>}></Route>
        <Route exact path="/products/:id" element={<InfoProductPage/>}></Route>
        <Route exact path="/admin" element={<Admin/>}></Route>
        <Route exact path="/admin/products" element={<Admin/>}></Route>
        <Route exact path="/admin/reklams" element={<Admin/>}></Route>
        <Route exact path="/admin/categories" element={<Admin/>}></Route>
      </Routes>
  </Router>
  )
}


export default AppRouter


