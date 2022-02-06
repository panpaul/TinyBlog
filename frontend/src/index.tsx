import React from "react";
import ReactDOM from "react-dom";

import { BrowserRouter } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";

import Layout from "./Layout";

ReactDOM.render(
    <React.StrictMode>
        <BrowserRouter>
            <Layout />
        </BrowserRouter>
    </React.StrictMode>,
    document.getElementById("root")
);
