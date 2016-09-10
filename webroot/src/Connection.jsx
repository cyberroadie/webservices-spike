import React from 'react';
import Websocket from 'react-websocket';

class Connection extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            text: 'Haaai'
        };
    }

    handleData(data) {
        let result = JSON.parse(data);
        this.setState({text: result.text});
    }

    render() {
        return (
        <div>
          Text: <strong>{this.state.text}</strong>
          <Websocket url='ws://localhost:8081/entry' onMessage={this.handleData.bind(this)}/>
        </div>
      );
    }
}

export default Connection;