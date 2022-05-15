import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';
import AppHeader from './AppHeader';
import AppContent from './AppContent';

import 'bootstrap/dist/css/bootstrap.min.css';
import './index.css';

class App extends Component {

    //this is required for all react component
    render(){      
      const myProps = {
        title: "My Cool App!",
        subject: "My subject",
        favourite_color: "red color",
      }
      return (
        <div className='app'>
          <AppHeader {...myProps} />
          <AppContent />
          <AppFooter />
        </div>
      )
    };

}

ReactDOM.render(<App />, document.getElementById('root') )