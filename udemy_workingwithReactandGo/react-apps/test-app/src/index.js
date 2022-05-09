import React, { Component } from 'react';
import ReactDOM from 'react-dom';
import AppFooter from './AppFooter';

class App extends Component {

    //this is required for all react component
    render(){
      return (
        <div>
          <h1>Hello, world!</h1>
        <div>
        </div>
          <AppFooter />
        </div>
      )
    };

}

ReactDOM.render(<App />, document.getElementById('root') )