import { createStore, combineReducers } from 'redux';
import ordersReducer from '../reducers/orders'

export default () => {
  const store = createStore(
    combineReducers({
      orders: ordersReducer
    })
  );

  return store;
};
