import React from 'react';
import './App.css';
import Button from '@material-ui/core/Button';
import TextField from '@material-ui/core/TextField';

  class App extends React.Component {
    constructor(props) {
      super(props);
      this.state = { 
        fibnumber: '',
        output: ''
      };
      this.handleChange = this.handleChange.bind(this);
      this.handleClick = this.handleClick.bind(this);
    }
    handleChange(event){
      this.setState({ fibnumber: event.target.value });
    }
    handleClick(event){
      this.componentDidMount();
    }
    componentDidMount() {
      var url = 'http://localhost:8080/api/fibonacci/' + this.state.fibnumber;
      fetch(url)
      .then(res => res.text())
      .then(response => 
        this.setState({output: response}))
      .catch(error => console.error('Error:', error));
    }

  render(){
    return (
      <div className="App">
        <header className="App-header">
          <h2>
            Fibonacci Sequence Generator
          </h2>
          <p class="normal">
            Please enter an integer to generate your Fibonacci sequence
          </p>
          <form noValidate autoComplete="off">
            <TextField 
              id="fibsequence" 
              label="Enter Sequence Length" 
              variant="outlined" 
              fullWidth
              onChange={this.handleChange}
            />
            <Button 
              variant="contained"
              onClick={this.handleClick}
            >Generate</Button>
          </form>
          <h4>
              {this.state.output}
          </h4>
        </header>
      </div>
    );
    }
};

export default App;
