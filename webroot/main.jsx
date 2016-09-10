import React from 'react';
import {render} from 'react-dom';
import Connection from './src/Connection.jsx';

class App extends React.Component {
  render () {
    return (
      <div>
        <Connection/>
      </div>
    );
  }
}

render(<App/>, document.getElementById('container'));
