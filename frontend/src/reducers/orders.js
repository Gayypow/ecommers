

export default  (state = [], action) => {
    switch (action.type) {
      case 'ADD':
        return [...state,
          {id: action.id,name: action.name, price: action.price, amount: action.amount}
        ];
        case 'DELETE':
          return [state.filter((p) => {
            return p.id !== action.id
        })]
      default:
        return state;
    }
  };
