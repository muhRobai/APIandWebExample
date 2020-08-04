import React, { Component } from 'react';
import { connect } from 'react-redux';
import '../Assets/styles/checkout.css';
import {setBasket, setData, setItems, setPage} from '../reducers/Actions';
import store from '../store';
import Utils from '../utils';

class Checkout extends Component {
	constructor(props) {
		super(props);
		this.state = {
				basket: this.props.basket,
				item: this.props.item,
			}
		}

		onHandleClick = () => {
			var empty = []
			
			store.dispatch(setBasket(empty))
			store.dispatch(setData(empty))
			store.dispatch(setItems(-this.state.item))
			store.dispatch(setPage('/'))
			Utils.navigateTo('/')()
		}	

		render(){
		return(
			<div className='item'>
				{this.state.basket.map((item, index) => {
					return (
						<div key={item.id} className='content'>
								<div>
									nama: {item.title}
								</div>
								<div>
									jumlah : {item.item}
								</div>
						</div>
					)
				})}
				<div>
				<button className='btn' 
					onClick={() => {
						this.onHandleClick()
					}}
				>
					Beli
				</button>
				</div>
			</div>
		)
	}
}

const propsMapper = (state) => {
	return {
		basket: state.rootReducer.basket,
		item: state.rootReducer.items,
	};
};

const checkouts = connect(propsMapper)(Checkout);
export default checkouts;