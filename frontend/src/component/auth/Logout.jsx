import React, {useEffect} from "react";

const Logout = () => {
    useEffect(() => {
        localStorage.removeItem("token-user")
        setTimeout(() => {
            window.location = "/login"
        }, 3000)
    })
    return (
        <div>
            <div className="d-flex justify-content-center align-items-center">
                <div
                    className="spinner-border text-primary"
                    style={{height: "3rem", width: "3rem"}}
                    role="status"
                >
                    <span className="sr-only">Loading...</span>
                </div>
            </div>
        </div>
    )
}

export default Logout;