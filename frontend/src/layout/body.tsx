import React from "react";

import { Route, Routes } from "react-router-dom";

import Row from "react-bootstrap/Row";
import Col from "react-bootstrap/Col";
import Container from "react-bootstrap/Container";

import ContentPage from "../page/content";
import ArticleList from "../page/list";
import ErrorPage from "../page/error";
import LoginPage from "../page/login";
import LogoutPage from "../page/logout";

import { RequireAuth } from "../service/auth";
import "./layout.css";

function Body() {
    const Search = () => <span>Search</span>;
    const Admin = () => <span>Admin</span>;
    return (
        <Container className="py-3">
            {/* Body */}
            <Row>
                <Col md={8}>
                    <Routes>
                        <Route
                            path="/edit/:articleId"
                            element={
                                <RequireAuth>
                                    <Admin />
                                </RequireAuth>
                            }
                        />
                        <Route path="/search" element={<Search />} />
                        <Route path="/login" element={<LoginPage />} />
                        <Route path="/logout" element={<LogoutPage />} />
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
                            <h4 className="fst-italic blog-head-font">About</h4>
                            <p className="mb-0">
                                A place for taking down notes, recording life
                                and sharing experiences.
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
    );
}

export default Body;
