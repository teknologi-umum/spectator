import React, { Component } from 'react'
import { connect } from 'react-redux'
import { recordPersonalInfo } from '../../store/actions/personalInfoActions'
import "./PersonalInfo.css";

export class PersonalInfo extends Component {
    state = {
        stdNo: '',
        programmingExp: 0,
        programmingExercise: 0,
        programmingLanguage: '',
    }
    handleChange = (e) => {
        this.setState({
            [e.target.id]: e.target.value
        })
    }
    handleSubmit = (e) => {
        e.preventDefault();
        console.log('Participant Basic Info:')
        console.log(this.state);
        this.props.onRecordPersonalInfo(this.state);
        this.props.history.push('/instructions');

        if (this.handleValidation()) {
            alert("Form submitted");
          } else {
            alert("Field can't be blank.");
          }
    }

    handleChange(field, e) {
        let fields = this.state.fields;
        fields[field] = e.target.value;
        this.setState({ fields });
      }

    constructor(props) {
        super(props);
    
        this.state = {
          fields: {},
          errors: {},
        };
      }
    
      handleValidation() {
        let fields = this.state.fields;
        let errors = {};
        let formIsValid = true;

        if (!fields["stdNo"]) {
            formIsValid = false;
            errors["stdNo"] = "Cannot be empty";
          }

        if (!fields["programmingExp"]) {
            formIsValid = false;
            errors["programmingExp"] = "Cannot be empty";
          }

        if (!fields["programmingExercise"]) {
            formIsValid = false;
            errors["programmingExercise"] = "Cannot be empty";
          }

        if (!fields["programmingLanguage"]) {
            formIsValid = false;
            errors["programmingLanguage"] = "Cannot be empty";
          }
        
          this.setState({ errors: errors });
          return formIsValid;
      }

    render() {

        return (
            <div className="container">
                <form onSubmit={this.handleSubmit}>
                    <h5 className="grey-text text-darken-3">Personal Basic Info</h5>
                    <div>
                        <label style={{ fontSize: '16px', color: 'black' }}>Student Number</label>
                        <input
                            type="text"
                            id="stdNo"
                            value={this.state.fields["stdNo"]}
                            onChange={(value) => this.setState({
                                ...this.state,
                                stdNo: value.target.value,
                            })}
                        />
                    </div>

                    <div style={{ marginTop: '15px' }}>
                        <label style={{ fontSize: '16px', color: 'black' }}>How many years have you been doing programming</label>
                        <input
                            type="number"
                            id="programmingExp"
                            value={this.state.fields["programmingExp"]}
                            onChange={(value) => this.setState({
                                ...this.state,
                                programmingExp: value.target.value,
                            })}
                        />
                    </div>

                    <div style={{ marginTop: '15px' }}>
                        <label style={{ fontSize: '16px', color: 'black' }}>How many hours in a week do you practice programming</label>
                        <input
                            type="number"
                            id="programmingExercise"
                            value={this.state.fields["programmingExercise"]}
                            onChange={(value) => this.setState({
                                ...this.state,
                                programmingExp: value.target.value,
                            })}
                        />
                    </div>

                    <div style={{ marginTop: '15px' }}>
                        <label style={{ fontSize: '16px', color: 'black' }}>What programming languages are you familiar with (example : java, python, c)</label>
                        <input
                            type="text"
                            id="programmingLanguage"
                            value={this.state.fields["programmingLanguage"]}
                            onChange={(value) => this.setState({
                                ...this.state,
                                programmingExp: value.target.value,
                            })}
                        />
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
        personalInfo: state.personalInfo,
    }
}

const mapDispatchToProps = (dispatch) => {
    return {
        onRecordPersonalInfo: (info) => dispatch(recordPersonalInfo(info))
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(PersonalInfo);