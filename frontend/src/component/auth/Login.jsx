import React, { useState } from "react";
import axios from "axios";
import { toast } from "react-toastify";
import { isAuth } from "../../utils/isAuth";
import { useHistory } from "react-router-dom";

const Login = () => {
  let history = useHistory();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const onChangeEmail = (e) => {
    const email = e.target.value;
    setEmail(email);
  };

  const onChangePassword = (e) => {
    const password = e.target.value;
    setPassword(password);
  };

  const handleLogin = (e) => {
    e.preventDefault();
    const data = {
      email,
      password,
    };
    axios
      .post("http://localhost:9000/login", data)
      .then((res) => {
        localStorage.setItem("token-user", res.data);
        toast.success("Login Success!", {
          position: "top-center",
          autoClose: 2000,
          hideProgressBar: true,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
        });
        window.location = "/";
      })
      .catch((err) => {
        toast.error(`Error: ${err.message}!`, {
          position: "top-center",
          autoClose: 2000,
          hideProgressBar: true,
          closeOnClick: true,
          pauseOnHover: true,
          draggable: true,
          progress: undefined,
        });
      });
  };

  if (isAuth()) {
    history.push("/");
  }

  return (
    <>
      <div className="d-flex justify-content-center">
        <div className="card col-6">
          <div className="card-body p-8">
            <h2 className="text-center">Login</h2>
            <br />
            <form>
              <div className="form-group">
                <label htmlFor="email">Email</label>
                <input
                  type="text"
                  id="email"
                  className="form-control"
                  name="email"
                  onChange={onChangeEmail}
                />
              </div>
              <div className="form-group">
                <label htmlFor="password">Password</label>
                <input
                  type="password"
                  id="password"
                  className="form-control"
                  name="password"
                  onChange={onChangePassword}
                />
              </div>
              <div className="form-group">
                <button
                  onClick={handleLogin}
                  className="btn btn-success float-right"
                >
                  Login
                </button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  );
};
export default Login;
