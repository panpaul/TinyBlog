import React from "react";

import { LinkContainer } from "react-router-bootstrap";

import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Button from "react-bootstrap/Button";
import Container from "react-bootstrap/Container";

import { BsSearch } from "react-icons/bs";
import { useAuth } from "../service/auth";

import "./layout.css";

function Header() {
    const auth = useAuth();

    return (
        <Container className="blog-header py-3">
            {/* Header */}
            <Row className="justify-content-end flex-nowrap align-items-center">
                <Col xs={{ span: 4 }} className="text-center">
                    <LinkContainer to="/">
                        <a className="blog-header-logo blog-head-font text-dark">
                            Paul &apos;s Blog
                        </a>
                    </LinkContainer>
                </Col>
                <Col
                    xs={{ span: 4 }}
                    className="d-flex justify-content-end align-items-center"
                >
                    <LinkContainer to="/search">
                        <a className="link-secondary">
                            <BsSearch className="mx-3" />
                        </a>
                    </LinkContainer>
                    <LinkContainer to={auth.isLogin ? "/logout" : "/login"}>
                        <Button variant="outline-secondary" size="sm">
                            {auth.isLogin ? "Logout" : "Login"}
                        </Button>
                    </LinkContainer>
                </Col>
            </Row>
        </Container>
    );
}

export default Header;
