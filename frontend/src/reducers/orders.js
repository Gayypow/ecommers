

export default  (state = [], action) => {
    switch (action.type) {
      case 'ADD':
        return [...state,
          {id: action.id, amount: action.amount}
        ];
      default:
        return state;
    }
  };
