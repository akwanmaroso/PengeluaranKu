import React from "react";
import Footer from "./layouts/Footer";
import Header from "./layouts/Header";

import {BrowserRouter as Router, Redirect, Route} from "react-router-dom";
import Login from "./auth/Login";
import Transactions from "./Transactions/Transactions";
import {isAuth} from "../utils/isAuth";
import Logout from "./auth/Logout";
import Signup from "./auth/Signup";

const Home = () => {
    return (
        <div>
            <Router>
                <Header/>
                <Route exact path="/">
                    {isAuth() ? <Redirect to="/"/> : <Redirect to="/login"/>}
                    <Transactions/>
                </Route>
                <Route path="/login" component={Login}/>
                <Route path="/signup" component={Signup}/>
                <Route path="/logout" component={Logout}/>
                <Footer/>
            </Router>
        </div>
    );
};

export default Home;
