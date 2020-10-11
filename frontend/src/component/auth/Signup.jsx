import React from "react";

const Signup = () => {
    return (
        <div className="d-flex justify-content-center">
            <div className="row">
                <div className="card col-12">
                    <div className="card-body">
                        <h2 className="text-center">SignUp</h2>
                        <br/>
                        <form>
                            <div className="form-group">
                                <label htmlFor="name">Name</label>
                                <input
                                    type="text"
                                    id="name"
                                    className="form-control"
                                    name="name"
                                    // onChange={onChangeEmail}
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
                                    // onChange={onChangeEmail}
                                    required
                                />
                            </div>
                            <div className="form-group">
                                <label htmlFor="password">Password</label>
                                <input
                                    type="text"
                                    id="password"
                                    className="form-control"
                                    name="password"
                                    // onChange={onChangePassword}
                                    required
                                />
                            </div>
                            <div className="form-group">
                                <label htmlFor="Vpassword">Verify Password</label>
                                <input
                                    type="text"
                                    id="Vpassword"
                                    className="form-control"
                                    name="Vpassword"
                                    // onChange={onChangePassword}
                                    required
                                />
                            </div>
                            <div className="form-group">
                                <button
                                    // onClick={handleLogin}
                                    className="btn btn-success">
                                    SignUp
                                </button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    )
}

export default Signup;