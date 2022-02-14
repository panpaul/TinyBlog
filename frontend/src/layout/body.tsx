import React from "react";

import { Route, Routes } from "react-router-dom";
import Container from "react-bootstrap/Container";

import ArticleList from "../page/list";
import ContentPage from "../page/content";
import ErrorPage from "../page/error";
import LoginPage from "../page/login";
import LogoutPage from "../page/logout";

import { RequireAuth } from "../service/auth";
import "./layout.css";

function Body() {
    const Search = () => <span>Search</span>;
    return (
        <Container className="py-3">
            {/* Body */}
            <Routes>
                <Route path="/search" element={<Search />} />
                <Route path="/login" element={<LoginPage />} />
                <Route path="/logout" element={<LogoutPage />} />
                <Route path="/article/:articleId" element={<ContentPage />} />
                <Route
                    path="/edit/:articleId"
                    element={
                        <RequireAuth>
                            <ContentPage />
                        </RequireAuth>
                    }
                />
                <Route path="/" element={<ArticleList />} />
                <Route path="*" element={<ErrorPage />} />
            </Routes>
        </Container>
    );
}

export default Body;
