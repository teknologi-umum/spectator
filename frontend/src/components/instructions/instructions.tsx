import React, {Component} from "react";
import arousal from "./Arousal-copy.png";
import "./instruction.css";

export class instructions extends Component {
    render() {
        return (
        <div className="container white">
            <h4> General Instructions</h4>
                <p>
                This experiment contains of two part.
                The first part is SAM Test and the second part is coding test.
                SAM Test is a self assessment test to measure your emotion.
                This test contains three question that will assess your emotion and will take time less than 5 minutes.
                The second test is coding test that consist of six programming questions.
                In this part you have to answer the questions and finish it within 90 minutes. 
                </p>

            <h5>1) SAM Test</h5>
                 <p>
                In this part there will be three question that you are required to answer by clicking one out of nine images available.
                The images represent how do you feel about the question.
                The image that is located on the far left indicates that you are very disagree with the statement and the image on the far right indicates that you are very agree with the statement meanwhile whe 5th image indicates that you are neutral towards the statement.
                You are required to fill this test two times. The first is before you take the test and the second test is after you finish programming test.
                The first SAM Test will be asking your current emotion meanwhile the second SAM Test will be asking your emotion during programming test.
                </p>
                <img className="arousal" src={arousal} alt={arousal} />
                <label className="caption">SAM Test Example</label>

            <h5>2) Programming Test</h5>
                <p>
                In this part there will be six programming questions that you are required to answer within 90 minutes by using java programming language.
                You are not allowed to search the answer somewhere else or get help from other people.
                If you are not able to answer the questions, you are allowed to left the page empty and go to the next question.
                You are allowed to go back to previous questions when there is still time left.
                </p>

            <p>p.s : this test will not affect your mark</p>

            <button type="submit" className="btn btn-primary lighten-1 z-depth-0" onClick={event => window.location.href="/samtest"}>Submit</button>

        </div>

        );
    }
}

export default instructions;