import React, { Component } from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import LastPage from './components/layout/lastPage'
import PersonalInfo from './components/basicInfo/PersonalInfo'
import instructions from './components/instructions/instructions'
import samtest from './components/samtest/samtest';
import CountDownTimer from './components/coding/CountDownTimer';
import Q1 from './components/coding/Q1';
import Q2 from './components/coding/Q2';
import Q3 from './components/coding/Q3';
import Q4 from './components/coding/Q4';
import Q5 from './components/coding/Q5';
import Q6 from './components/coding/Q6';

class App extends Component {
  render() {
    const hoursMinSecs = {hours:1, minutes: 30, seconds: 0}
    return (
      <BrowserRouter>
        <div className="App">
          <CountDownTimer hoursMinSecs={hoursMinSecs} />
          <Switch>
            <Route exact path='/' component={PersonalInfo} />
            <Route path='/last' component={LastPage} />
            <Route path='/instructions' component={instructions} />
            <Route path='/samtest' component={samtest} />
            <Route path='/Q1' component={Q1} />
            <Route path='/Q2' component={Q2} />
            <Route path='/Q3' component={Q3} />
            <Route path='/Q4' component={Q4} />
            <Route path='/Q5' component={Q5} />
            <Route path='/Q6' component={Q6} />
          </Switch>
        </div>
      </BrowserRouter>
    );
  }
}

// Video 17

export default App;