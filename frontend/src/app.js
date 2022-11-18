import React from 'react'
import AppRouter from "./routers/AppRouter";
import configureStore from './store/configureStore';
import { Provider } from 'react-redux';
import { useEffect } from 'react';




const store = configureStore();
// store.dispatch({type: 'ADD', id: 1, amount: 1})
const state = store.getState();
store.subscribe(() => {console.log(store.getState())})
console.log(state)  


const App = () => {
  return (
    <Provider store={store}>
      <AppRouter />
    </Provider>
  )
}



export  default App;