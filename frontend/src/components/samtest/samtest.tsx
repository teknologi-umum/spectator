import React, { Component } from 'react'
import { connect } from 'react-redux'
//import { recordPretest } from '../../store/actions/pretestAction'
//import Select from 'react-select';
import "./samtest.css"
import { submitPreTest, submitPostTest } from '../../store/actions/questionActions'

//arousal img
import Arousal1 from "./assets/Arousal/Arousal-1.png"
import Arousal2 from "./assets/Arousal/Arousal-2.png"
import Arousal3 from "./assets/Arousal/Arousal-3.png"
import Arousal4 from "./assets/Arousal/Arousal-4.png"
import Arousal5 from "./assets/Arousal/Arousal-5.png"
import Arousal6 from "./assets/Arousal/Arousal-6.png"
import Arousal7 from "./assets/Arousal/Arousal-7.png"
import Arousal8 from "./assets/Arousal/Arousal-8.png"
import Arousal9 from "./assets/Arousal/Arousal-9.png"
//pleasure img
import Pleasure1 from "./assets/Pleasure/Pleasure-1.png"
import Pleasure2 from "./assets/Pleasure/Pleasure-2.png"
import Pleasure3 from "./assets/Pleasure/Pleasure-3.png"
import Pleasure4 from "./assets/Pleasure/Pleasure-4.png"
import Pleasure5 from "./assets/Pleasure/Pleasure-5.png"
import Pleasure6 from "./assets/Pleasure/Pleasure-6.png"
import Pleasure7 from "./assets/Pleasure/Pleasure-7.png"
import Pleasure8 from "./assets/Pleasure/Pleasure-8.png"
import Pleasure9 from "./assets/Pleasure/Pleasure-9.png"
//dominance img
import Dominance1 from "./assets/Dominance/Dominance-1.png"
import Dominance2 from "./assets/Dominance/Dominance-2.png"
import Dominance3 from "./assets/Dominance/Dominance-3.png"
import Dominance4 from "./assets/Dominance/Dominance-4.png"
import Dominance5 from "./assets/Dominance/Dominance-5.png"
import Dominance6 from "./assets/Dominance/Dominance-6.png"
import Dominance7 from "./assets/Dominance/Dominance-7.png"
import Dominance8 from "./assets/Dominance/Dominance-8.png"
import Dominance9 from "./assets/Dominance/Dominance-9.png"


export class samtest extends Component {

    state = {
        aroused: 0,
        dominant: 0,
        pleasure: 0
    }

    finishQuestions = (samTestScore) => {
        const {question, personalInfo} = this.props;
        this.props.history.push('/Q1');
    }

    handleSubmit = (e) => {
        e.preventDefault();
        const { aroused, dominant, pleasure } = this.state;
        let samTestScore = 0;

        if (this.props.nextQuestion === 1) {
            this.props.history.push('/Q1');
        } else if (this.props.nextQuestion === 2) {
            this.props.history.push('/Q2');
        } else if (this.props.nextQuestion === 3) {
            this.props.history.push('/Q3');
        } else if (this.props.nextQuestion === 4) {
            this.props.history.push('/Q4');
        } else if (this.props.nextQuestion === 5) {
            this.props.history.push('/Q5');
        } else if (this.props.nextQuestion === 6) {
            this.props.history.push('/Q6');
        } else {
            this.props.onSubmitPostTest(samTestScore);
            this.finishQuestions(samTestScore);
        }
    }

    render() {

        const arousalImg = (fieldName, selectedStateKey) => (
            <table>
                <tr style={{ border: 'none' }}>
                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={1}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 1,
                            })}
                            checked={selectedStateKey === 1}
                        />
                        <img src={Arousal1} alt="gambar arousal-1" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={2}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 2,
                            })}
                            checked={selectedStateKey === 2}
                        />
                        <img src={Arousal2} alt="gambar arousal-2" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={3}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 3,
                            })}
                            checked={selectedStateKey === 3}
                        />
                        <img src={Arousal3} alt="gambar arousal-3" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={4}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 4,
                            })}
                            checked={selectedStateKey === 4}
                        />
                        <img src={Arousal4} alt="gambar arousal-4" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={5}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 5,
                            })}
                            checked={selectedStateKey === 5}
                        />
                        <img src={Arousal5} alt="gambar arousal-5" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={6}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 6,
                            })}
                            checked={selectedStateKey === 6}
                        />
                        <img src={Arousal6} alt="gambar arousal-6" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={7}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 7,
                            })}
                            checked={selectedStateKey === 7}
                        />
                        <img src={Arousal7} alt="gambar arousal-7" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={8}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 8,
                            })}
                            checked={selectedStateKey === 8}
                        />
                        <img src={Arousal8} alt="gambar arousal-8" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={9}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 9,
                            })}
                            checked={selectedStateKey === 9}
                        />
                        <img src={Arousal9} alt="gambar arousal-9" />
                    </label>
                    </td>
                </tr>
            </table>
        )


        const pleasureImg = (fieldName, selectedStateKey) => (
            <table>
                <tr style={{ border: 'none' }}>
                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={1}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 1,
                            })}
                            checked={selectedStateKey === 1}
                        />
                        <img src={Pleasure1} alt="gambar Pleasure-1" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={2}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 2,
                            })}
                            checked={selectedStateKey === 2}
                        />
                        <img src={Pleasure2} alt="gambar Pleasure-2" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={3}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 3,
                            })}
                            checked={selectedStateKey === 3}
                        />
                        <img src={Pleasure3} alt="gambar Pleasure-3" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={4}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 4,
                            })}
                            checked={selectedStateKey === 4}
                        />
                        <img src={Pleasure4} alt="gambar Pleasure-4" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={5}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 5,
                            })}
                            checked={selectedStateKey === 5}
                        />
                        <img src={Pleasure5} alt="gambar Pleasure-5" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={6}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 6,
                            })}
                            checked={selectedStateKey === 6}
                        />
                        <img src={Pleasure6} alt="gambar Pleasure-6" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={7}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 7,
                            })}
                            checked={selectedStateKey === 7}
                        />
                        <img src={Pleasure7} alt="gambar Pleasure-7" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={8}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 8,
                            })}
                            checked={selectedStateKey === 8}
                        />
                        <img src={Pleasure8} alt="gambar Pleasure-8" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={9}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 9,
                            })}
                            checked={selectedStateKey === 9}
                        />
                        <img src={Pleasure9} alt="gambar Pleasure-9" />
                    </label>
                    </td>
                </tr>
            </table>
        )

        const dominantImg = (fieldName, selectedStateKey) => (
            <table>
                <tr style={{ border: 'none' }}>
                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={1}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 1,
                            })}
                            checked={selectedStateKey === 1}
                        />
                        <img src={Dominance1} alt="gambar Dominance-1" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={2}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 2,
                            })}
                            checked={selectedStateKey === 2}
                        />
                        <img src={Dominance2} alt="gambar Dominance-2" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={3}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 3,
                            })}
                            checked={selectedStateKey === 3}
                        />
                        <img src={Dominance3} alt="gambar Dominance-3" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={4}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 4,
                            })}
                            checked={selectedStateKey === 4}
                        />
                        <img src={Dominance4} alt="gambar Dominance-4" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={5}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 5,
                            })}
                            checked={selectedStateKey === 5}
                        />
                        <img src={Dominance5} alt="gambar Dominance-5" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={6}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 6,
                            })}
                            checked={selectedStateKey === 6}
                        />
                        <img src={Dominance6} alt="gambar Dominance-6" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={7}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 7,
                            })}
                            checked={selectedStateKey === 7}
                        />
                        <img src={Dominance7} alt="gambar Dominance-7" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={8}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 8,
                            })}
                            checked={selectedStateKey === 8}
                        />
                        <img src={Dominance8} alt="gambar Dominance-8" />
                    </label>
                    </td>

                    <td>
                    <label style={{ marginLeft: 15, fontSize: '16px', color: 'purple', marginRight: -70 }}>
                        <input
                            style={{ opacity: 'initial', marginTop: 7, pointerEvents: 'all' }}
                            type="radio"
                            value={9}
                            onChange={() => this.setState({
                                ...this.state,
                                [fieldName]: 9,
                            })}
                            checked={selectedStateKey === 9}
                        />
                        <img src={Dominance9} alt="gambar Dominance-9" />
                    </label>
                    </td>
                </tr>
            </table>
        )



        return (
            <div className="container-fluid white">
                <form onSubmit={this.handleSubmit}>
                    <h5 className="grey-text text-darken-3 header">Self Assessment Manikin Test (SAM Test)</h5>
                    <div>
                        <label style={{ fontSize: '20px', color: 'black' }}>How aroused are you now?</label>
                        <p>Arousal refer to how aroused are you generally in the meantime</p>
                        {arousalImg('aroused', this.state.aroused)}
                    </div>

                    <div style={{ marginTop: '15px' }}>
                        <label style={{ fontSize: '20px', color: 'black' }}>How pleased are you now?</label>
                        <p>Pleasure refer to how pleased are you generally in the meantime</p>
                        {pleasureImg('pleasure', this.state.pleasure)}
                    </div>

                    <div style={{ marginTop: '20px' }}>
                        <label style={{ fontSize: '16px', color: 'black' }}>How dominant are you now?</label>
                        <p>Dominance refer to how dominant are you generally in the meantime</p>
                        {dominantImg('dominant', this.state.dominant)}
                    </div>

                    <div className="input-field">
                        <button className="btn btn-primary lighten-1 z-depth-0">Submit</button>
                    </div>
                </form>
            </div>
        )
    }
}

const mapStateToProps = (state) => {
    return {
        nextQuestion: state.question.nextQuestion,
        question: state.question,
        personalInfo: state.personalInfo,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        onSubmitPostTest: (score) => dispatch(submitPostTest(score))
    }
}

export default connect (mapStateToProps, mapDispatchToProps)(samtest);