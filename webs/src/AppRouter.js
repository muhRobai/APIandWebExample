import React, { Component } from 'react';
import { Route } from 'react-router-dom';
import Homepages from './Page/homepage';
import App from './Page/App';
import Checkout from './Page/checkout';
import './Assets/styles/index.css';
import { connect } from 'react-redux';
import Utils from './utils';
import basket from './Assets/images/shopping-basket-solid.svg';
import arrow from './Assets/images/arrow-circle-left-solid.svg';
import store from './store';
import { setPage } from './reducers/Actions';

class AppRouter extends Component {
	constructor(props) {
		super(props)
	}
	render() {
		return (
			<React.Fragment>
				<div className='header'>
					<div className='checkout'>
						<img src={this.props.page === '/' 
							? basket : arrow} 
							alt='basket'
							style={{cursor:'pointer'}}
							onClick={() => {
								var url = this.props.page === "/" ? '/checkout' : '/'
								store.dispatch(setPage(url))
								Utils.navigateTo(url)();
							}}
						></img>
					</div>
					<div className='item'>
						{this.props.item}
					</div>
				</div>
				<div className='body'>
					<div>&nbsp;</div>
					<Route exact path='/' component={Homepages}/>
					<Route exact path='/app' component={App}/>
					<Route exact path='/checkout' component={Checkout}/>
				</div>
			</React.Fragment>
		)
	}
}

const propsMapper = (state) => {
  return {
		item: state.rootReducer.items,
		page: state.rootReducer.page,
  };
};

const AppRouters = connect(propsMapper)(AppRouter);
export default AppRouters;