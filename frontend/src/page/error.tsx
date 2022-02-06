import React from "react";
import Alert from "react-bootstrap/Alert";

interface ErrorMessage {
    msg?: string;
}

function ErrorPage(props: ErrorMessage) {
    return (
        <main style={{ padding: "1rem" }}>
            <Alert variant="danger">
                <Alert.Heading>Oops! Something must be wrong</Alert.Heading>
                <hr />
                <p> There is an error while loading the page </p>
                <p> Please go back and try again later </p>
                {props.msg && <p>{`Message from server: ${props.msg}`}</p>}
            </Alert>
        </main>
    );
}

export default ErrorPage;
