import { combineReducers } from "redux";
import { connectRouter } from "connected-react-router";
import { 
  SET_TOAST,
  SET_ITEM, 
  SET_BASKET,
  SET_PAGE,
  SET_DATA,
} from "./Actions";

const initialState = {
  toastMessage:"",
  showToast: false,
  items: 0,
  basket: [],
  page: '/',
  data: [],
}

const rootReducer = (state = initialState, action) => {
    switch (action.type) {
      case SET_TOAST:
        return Object.assign({}, state, {
          toastMessage:action.title,
          showToast: true
        })
      case SET_ITEM: 
        return Object.assign({}, state, {
          items: state.items + action.number,
        }) 
      case SET_BASKET: 
        return Object.assign({}, state, {
          basket: action.basket,
        }) 
      case SET_PAGE: 
        return Object.assign({}, state, {
          page: action.page,
        })
      case SET_DATA:
        return Object.assign({}, state, {
          data: action.data,
        }) 
      default:
        return state
    }
  }

  
  
  export default (history) => combineReducers({
    router: connectRouter(history),
    rootReducer
  })