import React, { Component, useEffect, useState } from "react";
import "./questions.css";
import imgQ1 from "./imgQ1.png";
import IdleTimer from "react-idle-timer";

export class Q1 extends Component {
  constructor(props) {
    super(props);
    this.idleTimer = null;
    this.handleOnIdle = this.handleOnIdle.bind(this);
    this.keyDown = null;
    this.handleKeyDown = this.handleKeyDown.bind(this);
    this.state = {
      input: localStorage.getItem("input")||"",
      output: "",
      language_id:localStorage.getItem("language_Id")|| 2,
      user_input: "",
      count: 0,
      intervalId: 0
    };
  }
  input = (event) => {
    event.preventDefault();
    this.setState({ input: event.target.value });
    localStorage.setItem("input", event.target.value);
  };

  userInput = (event) => {
    event.preventDefault();
    this.setState({ user_input: event.target.value });
  };

  language = (event) => {
    event.preventDefault();
    this.setState({ language_id: event.target.value });
    localStorage.setItem("language_Id", event.target.value);
  };

  run = async (e) => {
    e.preventDefault();
    this.setState({button:e.target.name, countClick:this.state.countClick+1});
    console.log("Click Run");
    const outputText = document.getElementById("output");
    outputText.innerHTML = "";
    outputText.innerHTML += "Creating Submission ...\n";
    const response = await fetch(
      "https://judge0-ce.p.rapidapi.com/submissions",
      {
        method: "POST",
        headers: {
          "x-rapidapi-host": "judge0-ce.p.rapidapi.com",
          "x-rapidapi-key": "6e061ff853mshd93ec34cf96e638p1d4e48jsnbe918de0e2bb", // Get yours for free at https://rapidapi.com/judge0-official/api/judge0-ce/
          "content-type": "application/json",
          accept: "application/json"
        },
        body: JSON.stringify({
          source_code: this.state.input,
          stdin: this.state.user_input,
          language_id: this.state.language_id
        })
      }
    );

    outputText.innerHTML += "Submission Created ...\n";
    const jsonResponse = await response.json();
    let jsonGetSolution = {
      status: { description: "Queue" },
      stderr: null,
      compile_output: null
    };
    while (
      jsonGetSolution.status.description !== "Accepted" &&
      jsonGetSolution.stderr == null &&
      jsonGetSolution.compile_output == null
    ) {
      outputText.innerHTML = `Creating Submission ... \nSubmission Created ...\nChecking Submission Status\nstatus : ${jsonGetSolution.status.description}`;
      if (jsonResponse.token) {
        const url = `https://judge0-ce.p.rapidapi.com/submissions/${jsonResponse.token}?base64_encoded=true`;
        const getSolution = await fetch(url, {
          method: "GET",
          headers: {
            "x-rapidapi-host": "judge0-ce.p.rapidapi.com",
            "x-rapidapi-key": "6e061ff853mshd93ec34cf96e638p1d4e48jsnbe918de0e2bb", // Get yours for free at https://rapidapi.com/judge0-official/api/judge0-ce/
            "content-type": "application/json"
          }
        });
        jsonGetSolution = await getSolution.json();
      }
    }

    if (jsonGetSolution.stdout) {
      const output = atob(jsonGetSolution.stdout);
      outputText.innerHTML = "";
      outputText.innerHTML += `Results :\n${output}\nExecution Time : ${jsonGetSolution.time} Secs\nMemory used : ${jsonGetSolution.memory} bytes`;
    } else if (jsonGetSolution.stderr) {
      const error = atob(jsonGetSolution.stderr);
      outputText.innerHTML = "";
      outputText.innerHTML += `\n Error :${error}`;
    } else {
      const compilation_error = atob(jsonGetSolution.compile_output);
      outputText.innerHTML = "";
      outputText.innerHTML += `\n Error :${compilation_error}`;
    }
  };

  handleKeyDown = (event) => {
    console.log(event.key);

    const currentText = event.target.value;
    const characterCount = currentText.length;
    console.log("character " + characterCount);

    const wordcount = currentText.split(" ").length;
    console.log("word " + wordcount);

    let delCount = 0;
    if (event.code==="Backspace") {
      delCount++;
      console.log(delCount);
      }

    //var delCount = 0;
    //var keycode = (event.keycode ? event.keycode : event.which);
    //if (keycode=='8') {
    //  delCount++;
    //  console.log(delCount);
    //}
    //console.log("total active time ", this.idleTimer.getTotalActiveTime()/1000 + "s")

    const timestamp = Date.now()/1000;
    const date = new Date(timestamp * 1000);
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];
    const day = date.getDate();
    const month = months[date.getMonth()];
    const year = date.getFullYear();
    const hour = date.getHours();
    const min = "0" + date.getMinutes();
    const sec = "0" + date.getSeconds();
    const currentDay = "Date " + day + " : " + month + " : " + year;
    const currentTime = "Hour " + hour + " : " + min.substr(-2) + " : " + sec.substr(-2);
    console.log(currentDay);
    console.log(currentTime);
  };

  handleOnWheel = (event) => {
    if (event.deltaY<0) {
      console.log("Scroll up");
    } else {
      console.log("Scroll down");
    }
    const timestamp = Date.now()/1000;
    const date = new Date(timestamp * 1000);
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];
    const day = date.getDate();
    const month = months[date.getMonth()];
    const year = date.getFullYear();
    const hour = date.getHours();
    const min = "0" + date.getMinutes();
    const sec = "0" + date.getSeconds();
    const currentDay = "Date " + day + " : " + month + " : " + year;
    const currentTime = "Hour " + hour + " : " + min.substr(-2) + " : " + sec.substr(-2);
    console.log(currentDay);
    console.log(currentTime);
  };

  handleDelete = (event) => {
    event.preventDefault();
    this.setState({delete:event.target.name, countDelete:this.state.countDelete+1});
    console.log("Delete");
  };

  handleClickSave = (event) => {
    event.preventDefault();
    this.setState({button:event.target.name, countClick:this.state.countClick+1});
    console.log("Click Save");
  };

  handleClickSubmit = (event) => {
    event.preventDefault();
    this.setState({button:event.target.name, countClick:this.state.countClick+1});
    console.log("Click Submit");
  };

  handleMouseMovement = (event) => {
    const newIntervalId = setInterval(() => {
      this.setState(prevState => {
        return {
          ...prevState,
          count: prevState.count + 1
        };
      });
    }, 1000);

    this.setState(prevState => {
      return {
        ...prevState,
        intervalId: newIntervalId
      };
    });

    //console.log("screen x: " + event.screenX);
    //console.log("screen y: " + event.screenY);

    let totalX = Math.abs(event.movementX);
    let totalY = Math.abs(event.movementY);
    let moveX = event.movementX;
    let moveY = event.movementY;
    //console.log("SpeedX : " + totalX + "px/s");
    //console.log("SpeedY : " + totalY + "px/s");
    //console.log("MovementX : " + moveX + "px");
    //console.log("MovementY : " + moveY + "px");
    moveX = moveY = totalX = totalY = 0;

    const timestamp = Date.now()/1000;
    const date = new Date(timestamp * 1000);
    const months = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];
    const day = date.getDate();
    const month = months[date.getMonth()];
    const year = date.getFullYear();
    const hour = date.getHours();
    const min = "0" + date.getMinutes();
    const sec = "0" + date.getSeconds();
    const currentDay = "Date " + day + " : " + month + " : " + year;
    const currentTime = "Hour " + hour + " : " + min.substr(-2) + " : " + sec.substr(-2);
    //console.log(currentDay);
    //console.log(currentTime);
  };

  handleOnIdle = (event) => {
    console.log("User is idle ");
    console.log("total idle time ", this.idleTimer.getTotalIdleTime()/1000 + "s");
    console.log("total active time ", this.idleTimer.getTotalActiveTime()/1000 + "s");

    const time = this.idleTimer.getTotalIdleTime()/1000;
    const char = this.onKeyDown.characterCount;
    const wpm = char/5/time;
    console.log(wpm);
    //this.wordcount = this.handleKeyDown.bind(this);
    //console.log(this.wordcount)
    //console.log(this.onKeyDown.characterCount)
    //console.log("Char per sec " + this.characterCount / (this.idleTimer.getTotalActiveTime()/1000))
  };

  render() {
       return (
      <>
        <div className="row container-fluid"
        onWheel={this.handleOnWheel}
        onMouseMove={this.handleMouseMovement}
        >
          <IdleTimer
            ref={ref => {this.idleTimer = ref;}}
            timeout={3000}
            onIdle={this.handleOnIdle}
          />

          <div className="col">
          <label htmlFor="tags" className="mr-1">
              <b className="heading">Questions:</b>
            </label>
            <div className="rounded square-lg questions">
              <p>print the song twinkle twinkle little star by only using 2 variables</p>
              <img src={imgQ1} alt={imgQ1} max-width="50%" />
            </div>
          </div>

          <div className="col">
            <label htmlFor="tags" className="mr-1">
              <b className="heading">Language:</b>
            </label>
            <select
              value={this.state.language_id}
              onChange={this.language}
              id="tags"
              className="form-control form-inline mb-2 language"
            >
              <option value="54">C++</option>
              <option value="50">C</option>
              <option value="62">Java</option>
              <option value="71">Python</option>
            </select>

            <label htmlFor="solution ">
            <label htmlFor="tags" className="mr-1">
              <b className="heading">Code Here:</b>
            </label>
            </label>
            <textarea
              required
              name="solution"
              id="source"
              onChange={this.input}
              className="rounded source"
              value={this.state.input}
              onKeyDown={this.handleKeyDown}
              onPaste={(e)=>{
                e.preventDefault();
                return false;
              }} onCopy={(e)=>{
                e.preventDefault();
                return false;
              }} onCut={(e)=>{
                e.preventDefault();
                return false;
              }}
            ></textarea>
            <button
              type="submit"
              id="run"
              className="btn btn-info"
              onClick={this.run}
            >Run
            </button>
            &nbsp;
            &nbsp;
            <button
              type="submit"
              id="save"
              className="btn btn-warning ml-2 mr-3"
              onClick={this.handleClickSave}
            >Save
            </button>

            <button
              type="submit"
              id="submit"
              className="btn btn-success ml-2 mr-3"
              onClick={this.handleClickSave}
            >Submit
            </button>
            <br />
            <br />
            <div className="text-center">
              <div className="questionNo">
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q1"}>1</button>&nbsp;&nbsp;
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q2"}>2</button>&nbsp;&nbsp;
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q3"}>3</button>&nbsp;&nbsp;
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q4"}>4</button>&nbsp;&nbsp;
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q5"}>5</button>&nbsp;&nbsp;
                <button type="button" className="btn btn-outline-info" onClick={event => window.location.href="/Q6"}>6</button>&nbsp;&nbsp;
              </div>
            </div>
          </div>

          <div className="col">
              <label htmlFor="tags" className="mr-1">
              <b className="heading">User Input:</b>
              </label>
              <br />
              <textarea className="rounded inputTxt" id="input" onChange={this.userInput}></textarea>
              <br />
              <label htmlFor="solution ">
            <label htmlFor="tags" className="mr-1">
              <b className="heading">Output:</b>
            </label>
            </label>
              <br />
              <textarea className="rounded" id="output"></textarea>
          </div>

          <div>

          </div>
        </div>

      </>
    );
  }
}

export default Q1;
