import React, { useState } from "react";
import axios from "axios";
import { notifError, notifSuccess } from "../../utils/notifications";

const initialState = {
  name: "",
  email: "",
  password: "",
  verifyPassword: "",
};
const Signup = () => {
  const [state, setState] = useState(initialState);

  const handleChangeEmail = (e) => {
    setState({
      ...state,
      email: e.target.value,
    });
  };
  const handleChangeName = (e) => {
    setState({
      ...state,
      name: e.target.value,
    });
  };
  const handleChangePassword = (e) => {
    setState({
      ...state,
      password: e.target.value,
    });
  };
  const handleVerifyPassword = (e) => {
    setState({
      ...state,
      verifyPassword: e.target.value,
    });
  };

  const handleSignup = (e) => {
    e.preventDefault();
    // const {name, password, email, verifyPassword} = state
    axios
      .post("http://localhost:9000/users", state)
      .then((res) => {
        notifSuccess("Signup successfully");
        setTimeout(() => {
          window.location = "/";
        }, 2000);
      })
      .catch((err) => {
        notifError(`${err.message}`);
      });
    console.log(state);
  };

  const validatePassword = () => {
    return state.password !== state.verifyPassword;
  };

  return (
    <div className="d-flex justify-content-center">
      <div className="card col-6">
        <div className="card-body">
          <h2 className="text-center">SignUp</h2>
          <br />
          <form>
            <div className="form-group">
              <label htmlFor="name">Name</label>
              <input
                type="text"
                id="name"
                className="form-control"
                name="name"
                onChange={handleChangeName}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="email">Email</label>
              <input
                type="text"
                id="email"
                className="form-control"
                name="email"
                onChange={handleChangeEmail}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="password">Password</label>
              <input
                type="password"
                id="password"
                className="form-control"
                name="password"
                onChange={handleChangePassword}
                required
              />
            </div>
            <div className="form-group">
              <label htmlFor="verifyPassword">Verify Password</label>
              <input
                type="password"
                id="verifyPassword"
                className="form-control"
                name="verifyPassword"
                onChange={handleVerifyPassword}
                required
              />
              {validatePassword() ? (
                <p className="text-danger">Password not match</p>
              ) : (
                ""
              )}
            </div>
            <div className="form-group">
              <button
                onClick={handleSignup}
                className="btn btn-success float-right"
              >
                SignUp
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  );
};

export default Signup;
