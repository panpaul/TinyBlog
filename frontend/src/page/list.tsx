import React, { useEffect, useState } from "react";

import { useSearchParams } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap";
import Button from "react-bootstrap/Button";

import {
    ArticleListResp,
    getArticleList,
    getArticlePages,
} from "../service/api";
import ErrorPage from "./error";

import "./blog.css";

function ArticleList() {
    const [searchParams, setSearchParams] = useSearchParams();
    const [errorMsg, setErrorMsg] = useState<string>();
    const [descList, setDescList] = useState<ArticleListResp[]>([]);
    const [totalPages, setTotalPages] = useState(1);

    const author = searchParams.get("author") || "";
    const tag = searchParams.get("tag") || "";
    const currentPage = +(searchParams.get("currentPage") || "0");

    useEffect(() => {
        // get list uuid
        getArticleList({ author: author, tag: tag, page: currentPage })
            .then((res) => setDescList(res || []))
            .catch((err) => {
                console.error("[list] getArticleList", err);
                setErrorMsg((msg) =>
                    msg == undefined ? err : `${msg} | ${err}`
                );
            });
        // get pages
        getArticlePages({ author: author, tag: tag })
            .then((res) => setTotalPages(res))
            .catch((err) => {
                console.error("[list] getArticlePages", err);
                setErrorMsg((msg) =>
                    msg == undefined ? err : `${msg} | ${err}`
                );
            });
    }, [author, tag, currentPage]);

    const articleList = descList.map((item) => {
        return (
            <LinkContainer to={`/article/${item.uuid}`} key={item.uuid}>
                <article className="blog-post">
                    <h3 className="blog-post-title">{item.title}</h3>
                    <p className="blog-post-meta">
                        {new Date(item.created_at).toDateString()}
                    </p>
                    <p>{item.description}</p>
                </article>
            </LinkContainer>
        );
    });

    const params = { author: author, tag: tag };

    const navPrevEn = currentPage < totalPages - 1;
    const navPrevClass = navPrevEn ? "outline-primary" : "outline-secondary";
    const navPrevParams = Object.assign({}, params, {
        currentPage: String(currentPage + 1),
    });

    const navNextEn = currentPage > 0;
    const navNextClass = navNextEn ? "outline-primary" : "outline-secondary";
    const navNextParams = Object.assign({}, params, {
        currentPage: String(currentPage - 1),
    });

    const page = (
        <>
            {articleList}
            <nav className="blog-pagination">
                <Button
                    variant={navPrevClass}
                    disabled={!navPrevEn}
                    onClick={() => setSearchParams(navPrevParams)}
                >
                    Older
                </Button>{" "}
                <Button
                    variant={navNextClass}
                    disabled={!navNextEn}
                    onClick={() => setSearchParams(navNextParams)}
                >
                    Newer
                </Button>
            </nav>
        </>
    );

    return (
        <>
            {errorMsg || descList.length === 0 ? (
                <ErrorPage msg={errorMsg} />
            ) : (
                page
            )}
        </>
    );
}

export default ArticleList;
