import React from "react";

import Container from "react-bootstrap/Container";

import "./layout.css";

function Footer() {
    return (
        <Container className="blog-footer">
            {/* Footer */}
            <p>TinyBlog - Paul</p>
            <p>
                <a href="#">Back to top</a>
            </p>
        </Container>
    );
}

export default Footer;
