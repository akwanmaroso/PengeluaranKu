import React from "react";
import {NavLink} from "react-router-dom";
import {isAuth} from "../../utils/isAuth";

const Header = () => {
    return (
        <nav className="navbar navbar-expand-lg navbar-light bg-light mb-3">
            <div className="container">
                <NavLink className="navbar-brand" to="/">
                    PengeluaranKu
                </NavLink>
                <button
                    className="navbar-toggler"
                    type="button"
                    data-toggle="collapse"
                    data-target="#navbarSupportedContent"
                    aria-controls="navbarSupportedContent"
                    aria-expanded="false"
                    aria-label="Toggle navigation"
                >
                    <span className="navbar-toggler-icon"/>
                </button>

                <div className="collapse navbar-collapse" id="navbarSupportedContent">
                    <ul className="navbar-nav mr-auto">
                        <li className="nav-item">
                            <NavLink className="nav-link" to="/" activeClassName="active  ">
                                Home
                            </NavLink>
                        </li>
                    </ul>
                    <ul className="navbar-nav ml-auto">
                        {
                            isAuth() ? (
                                <li className="nav-item">
                                    <NavLink
                                        className="nav-link"
                                        to="/logout"
                                        activeClassName="active"
                                    >
                                        Logout
                                    </NavLink>
                                </li>
                            ) : (
                                <>
                                    <li className="nav-item">
                                        <NavLink
                                            className="nav-link"
                                            to="/login"
                                            activeClassName="active"
                                        >
                                            Login
                                        </NavLink>
                                    </li>
                                    <li className="nav-item">
                                        <NavLink
                                            className="nav-link"
                                            to="/signup"
                                            activeClassName="active"
                                        >
                                            SignUp
                                        </NavLink>
                                    </li>
                                </>
                            )
                        }
                    </ul>
                </div>
            </div>
        </nav>
    );
};

export default Header;
