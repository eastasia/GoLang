import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';
import AppContent from './AppContent';

import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';

class App extends Component {

    //this is required for all react component
    render(){      
      return (
        <div className='app'>
          <h1>Hello, world!</h1>
          <AppContent />
        <div>
        </div>
          <AppFooter />
        </div>
      )
    };

}

ReactDOM.render(<App />, document.getElementById('root') )