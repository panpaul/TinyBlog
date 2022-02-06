import React from "react";

import { Routes, Route } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap";

import Button from "react-bootstrap/Button";
import Col from "react-bootstrap/Col";
import Container from "react-bootstrap/Container";
import Row from "react-bootstrap/Row";
import { BsSearch } from "react-icons/bs";

import "./Layout.css";

import ErrorPage from "./page/error";
import ContentPage from "./page/content";
import ArticleList from "./page/list";

function Layout() {
    const Search = () => <span>Search</span>;
    const Admin = () => <span>Admin</span>;

    return (
        <>
            <Container className="blog-header py-3">
                {/* Header */}
                <Row className="justify-content-end flex-nowrap align-items-center">
                    <Col xs={{ span: 4 }} className="text-center">
                        <LinkContainer to="/">
                            <a className="blog-header-logo blog-head-font text-dark">
                                {/*Paul & apos;s Blog*/}
                                Title
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
                        <LinkContainer to="/admin">
                            <Button variant="outline-secondary" size="sm">
                                Admin
                            </Button>
                        </LinkContainer>
                    </Col>
                </Row>
            </Container>

            <Container className="py-3">
                {/* Body */}
                <Row>
                    <Col md={8}>
                        <Routes>
                            <Route path="/admin" element={<Admin />} />
                            <Route path="/search" element={<Search />} />
                            <Route
                                path="/article/:articleId"
                                element={<ContentPage />}
                            />
                            <Route path="/" element={<ArticleList />} />
                            <Route path="*" element={<ErrorPage />} />
                        </Routes>
                    </Col>
                    <Col md={4}>
                        <div className="position-sticky">
                            <div className="p-4 mb-3 bg-light rounded">
                                <h4 className="fst-italic blog-head-font">
                                    About
                                </h4>
                                <p className="mb-0">
                                    A place for taking down notes, recording
                                    life and sharing experiences.
                                </p>
                            </div>
                            <div className="p-4">
                                <h4 className="fst-italic blog-head-font">
                                    Elsewhere
                                </h4>
                                <ol className="list-unstyled">
                                    <li>
                                        <a href="https://github.com/panpaul">
                                            GitHub
                                        </a>
                                    </li>
                                </ol>
                            </div>
                        </div>
                    </Col>
                </Row>
            </Container>

            <Container className="blog-footer">
                {/* Footer */}
                <p>TinyBlog - Paul</p>
                <p>
                    <a href="#">Back to top</a>
                </p>
            </Container>
        </>
    );
}

export default Layout;
