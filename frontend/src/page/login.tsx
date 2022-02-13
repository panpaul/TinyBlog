import React, { useState } from "react";

import { useNavigate } from "react-router-dom";
import Alert from "react-bootstrap/Alert";

import { useAuth } from "../service/auth";

import "./login.css";

function LoginPage() {
    const navigate = useNavigate();
    const auth = useAuth();
    const [errMsg, setErrMsg] = useState("");

    function handleSubmit(event: React.FormEvent<HTMLFormElement>) {
        event.preventDefault();

        const formData = new FormData(event.currentTarget);
        const username = formData.get("username") as string;
        const password = formData.get("password") as string;

        auth.login(
            username,
            password,
            () => navigate("/", { replace: true }),
            (err) => setErrMsg(err)
        );
    }

    return (
        <div className="form-login">
            {errMsg && (
                <Alert
                    variant="danger"
                    onClose={() => setErrMsg("")}
                    dismissible
                >
                    {errMsg}
                </Alert>
            )}
            <form onSubmit={handleSubmit}>
                <h1 className="h3 mb-3 fw-normal">Please login</h1>

                <div className="form-floating">
                    <input
                        type="text"
                        name="username"
                        className="form-control"
                        placeholder="username"
                        required
                    />
                    <label htmlFor="floatingInput">Username</label>
                </div>

                <div className="form-floating">
                    <input
                        type="password"
                        name="password"
                        className="form-control"
                        placeholder="Password"
                        required
                    />
                    <label htmlFor="floatingPassword">Password</label>
                </div>

                <button className="w-100 btn btn-lg btn-primary" type="submit">
                    Sign in
                </button>
            </form>
        </div>
    );
}

export default LoginPage;
