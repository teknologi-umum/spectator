import React from "react";
import { connect } from "react-redux";

const Navbar = (props) => {
    return (
        <nav className="nav nav-wrapper grey darken-3">
            <div className="container brand-logo" style={{ paddingLeft: "20px" }}>
                Programming Quiz
            </div>
        </nav>
    );
};

const mapStateToProps = (state) => {
    return {
        // auth: state.firebase.auth,
        // profile: state.firebase.profile,
    };
};

export default connect(mapStateToProps)(Navbar);