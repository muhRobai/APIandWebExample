import React, { Component } from 'react';
import Settings from '../services/settings';
import '../Assets/styles/home.css';
import lamp from '../Assets/images/lampu.png';
import plus from '../Assets/images/plus-circle-solid.svg';
import minus from '../Assets/images/minus-circle-solid.svg';
import { setItems, setBasket, setData } from '../reducers/Actions';
import store from '../store';
import { connect } from 'react-redux';


const service = new Settings();

class Homepage extends Component {
  constructor(props) {
      super(props);
      this.state = {
          data: [],
      }
  }

  componentDidMount() {
      this.fetchData()
  }

  fetchData = () => {
      service.getData().then((v) => {
          this.setState({
              data: v.data,
          });
      })
  }

  onAddItem = (payload) => {
    var data = this.state.data
    if (data[payload.id-1] && !data[payload.id-1].item) {
      data[payload.id-1].item = 0  
    }

    data[payload.id-1].item += 1

    this.setState({
      data: data,
    })
    store.dispatch(setData(data))
    store.dispatch(setItems(1))
    this.onManageBaasket(payload.item)
  }

  onDeleteItem = (payload) => {
    var data = this.state.data
    if (data[payload.id-1] && data[payload.id-1].item && data[payload.id-1].item > 0 ) {
      data[payload.id-1].item -= 1  
      store.dispatch(setItems(-1))
    }  

    this.setState({
      data: data,
    })
    store.dispatch(setData(data))
    this.onManageBaasket(payload.item)
  }

  onManageBaasket = (payload) => {
    payload.isDelete = false
    var basket = this.props.basket
    var isAble = false
    
    if (payload.item === 0) {
      payload.isDelete = true
    }

    basket.map((item) => {
      if (item && item.id === payload.id) {
        item.item = payload.item
        item.isDelete = payload.isDelete
        isAble = true
      }
    })

    if (!isAble) {
      basket.push(payload)
    }

    store.dispatch(setBasket(basket))
  }

  render(){
    var data = this.props.data.length > 0 ? this.props.data : this.state.data
    return (
      <div className='card'>
        {data.map((item) => {
          return (
            <div className='card-item' key={item.id}>
              <img src={lamp} alt ='lamp'></img>
              <div className='card-text'>
                {item.title}
              </div>
              <div className='card-footer'>
                <img 
                  src={plus} alt='plus' 
                  style={{cursor: 'pointer'}}
                  onClick={() => {
                    var payload = {
                      id: item.id,
                      item: item,
                    }
                    this.onAddItem(payload)
                  }}></img>
                <div className='text'>{item.item !== undefined ? item.item : "0" }</div>
                <img 
                  src={minus} alt='minus'
                  style={{cursor: 'pointer'}}
                  onClick={() => {
                    var payload = {
                      id: item.id,
                      item: item,
                    }

                    this.onDeleteItem(payload)
                  }}></img>
              </div>
            </div>
          )
        })}
      </div>
    )
  }
}

const propsMapper = (state) => {
  return {
    basket: state.rootReducer.basket,
    data: state.rootReducer.data,
  };
};

const HomePages = connect(propsMapper)(Homepage);
export default HomePages;