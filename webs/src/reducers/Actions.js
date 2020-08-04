export const SET_TOAST = 'SET_TOAST'
export const SET_ITEM = 'SET_ITEM'
export const SET_BASKET = 'SET_BASKET'
export const SET_PAGE = 'SET_PAGE'
export const SET_DATA = 'SET_DATA'


export function setToast(title){
    return{
        type: SET_TOAST, title
    }
}

export function setItems(number){
    return{
        type: SET_ITEM, number
    }
}

export function setBasket(basket) {
    return{
        type: SET_BASKET, basket
    }
}

export function setPage(page) {
    return{
        type: SET_PAGE, page
    }
}

export function setData(data) {
    return{
        type: SET_DATA, data
    }
}

export const reducerMap = dispatch =>({
    setToast: title => dispatch(setToast(title)),
    setItem: number => dispatch(setItems(number)),
    setBasket: basket => dispatch(setBasket(basket)),
    setPage: page => dispatch(setPage(page)),
    setData: data => dispatch(setData(data)),
})