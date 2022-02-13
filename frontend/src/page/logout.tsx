import React, { useEffect, useState } from "react";

import { useNavigate } from "react-router-dom";
import Alert from "react-bootstrap/Alert";

import { useAuth } from "../service/auth";
import ErrorPage from "./error";

function LogoutPage() {
    const navigate = useNavigate();
    const auth = useAuth();
    const [errMsg, setErrMsg] = useState("");

    useEffect(() => {
        auth.logout(
            () => navigate("/", { replace: true }),
            (err) => setErrMsg(err)
        );
    }, []);

    return (
        <>
            {errMsg ? (
                <ErrorPage msg={errMsg} />
            ) : (
                <Alert variant="info">Logging out ...</Alert>
            )}
        </>
    );
}

export default LogoutPage;
