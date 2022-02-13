import React from "react";

import Header from "./layout/header";
import Footer from "./layout/footer";
import Body from "./layout/body";

import { AuthProvider } from "./service/auth";

function App() {
    return (
        <AuthProvider>
            <Header />
            <Body />
            <Footer />
        </AuthProvider>
    );
}

export default App;
