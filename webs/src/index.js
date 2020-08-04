import React from 'react';
import { render } from 'react-dom';
import { Provider } from 'react-redux';
import { ConnectedRouter } from 'connected-react-router';
import AppRouter from './AppRouter';
import store, { history } from './store';
import * as serviceWorker from './serviceWorker';

window.nav = history
const target = document.querySelector('#root')

render(
  <Provider store={store}>
    <ConnectedRouter history={history}>
        <AppRouter history={history}/>
    </ConnectedRouter>
  </Provider>,
  target
)
serviceWorker.unregister();