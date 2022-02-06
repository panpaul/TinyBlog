import React, { useEffect, useState } from "react";

import { useParams } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap";

import Badge from "react-bootstrap/Badge";
import Table from "react-bootstrap/Table";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Prism } from "react-syntax-highlighter";
import remarkMath from "remark-math";
import rehypeKatex from "rehype-katex";
import rehypeRaw from "rehype-raw";
import emoji from "remark-emoji";

import { ArticleResp, getArticle } from "../service/api";
import ErrorPage from "./error";

import "katex/dist/katex.min.css";
import "./blog.css";

function ContentPage() {
    const uuid = useParams().articleId;
    const [errorMsg, setErrorMsg] = useState<string>();
    const [detail, setDetail] = useState<ArticleResp>({
        CreatedAt: Date(),
        UpdatedAt: Date(),
        author: { nick_name: "loading", user_name: "loading" },
        content: "# Loading",
        description: "loading",
        tags: ["loading"],
        title: "loading...",
        uuid: "00000000-0000-0000-0000-000000000000",
    });

    useEffect(() => {
        getArticle({ uuid: uuid })
            .then((res) => setDetail(res))
            .catch((err) => {
                console.error("[content]", err);
                setErrorMsg(err);
            });
    }, [uuid]);

    const tags = detail.tags.map((item) => (
        <LinkContainer
            to={`/?tag=${item}`}
            key={item}
            style={{ marginRight: "0.5rem" }}
        >
            <Badge bg="info">{item}</Badge>
        </LinkContainer>
    ));

    const page = (
        <article className="blog-post">
            <h2 className="blog-post-title">{detail.title}</h2>
            <p className="blog-post-meta">
                {new Date(detail.CreatedAt).toDateString()}
                {" | "}
                <LinkContainer to={`/?author=${detail.author.user_name}`}>
                    <a href="#">{detail.author.nick_name}</a>
                </LinkContainer>
            </p>
            <p>{detail.description}</p>
            <hr />
            <ReactMarkdown
                remarkPlugins={[remarkGfm, remarkMath, emoji]}
                rehypePlugins={[rehypeKatex, rehypeRaw]}
                components={{
                    // eslint-disable-next-line @typescript-eslint/no-unused-vars
                    code: ({ node, inline, className, children, ...props }) => {
                        const match = /language-(\w+)/.exec(className || "");
                        return !inline && match ? (
                            <Prism language={match[1]} PreTag="div" {...props}>
                                {String(children).replace(/\n$/, "")}
                            </Prism>
                        ) : (
                            <code className={className} {...props}>
                                {children}
                            </code>
                        );
                    },
                    // eslint-disable-next-line @typescript-eslint/no-unused-vars
                    table: ({ node, ...props }) => (
                        <Table striped bordered hover {...props} />
                    ),
                }}
            >
                {detail.content}
            </ReactMarkdown>
            <hr />
            <div>{tags}</div>
        </article>
    );

    return <>{errorMsg ? <ErrorPage msg={errorMsg} /> : page}</>;
}

export default ContentPage;
