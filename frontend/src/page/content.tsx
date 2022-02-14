import React, { useEffect, useState } from "react";

import { useParams } from "react-router-dom";
import { LinkContainer } from "react-router-bootstrap";

import Badge from "react-bootstrap/Badge";
import Button from "react-bootstrap/Button";
import Col from "react-bootstrap/Col";
import Row from "react-bootstrap/Row";
import Table from "react-bootstrap/Table";
import Toast from "react-bootstrap/Toast";

import ReactMarkdown from "react-markdown";
import remarkGfm from "remark-gfm";
import { Prism } from "react-syntax-highlighter";
import remarkMath from "remark-math";
import rehypeKatex from "rehype-katex";
import rehypeRaw from "rehype-raw";
import emoji from "remark-emoji";

import CodeMirror from "@uiw/react-codemirror";
import { markdown, markdownLanguage } from "@codemirror/lang-markdown";
import { languages } from "@codemirror/language-data";

import { ArticleResp, getArticle, modifyArticle } from "../service/api";
import { useAuth } from "../service/auth";
import ErrorPage from "./error";

import "katex/dist/katex.min.css";
import "./blog.css";

function ContentPage() {
    const uuid = useParams().articleId;
    const auth = useAuth();
    const [errorMsg, setErrorMsg] = useState("");
    const [showEditor, setShowEditor] = useState(false);
    const [editorMsg, setEditorMsg] = useState("");
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

    const editorCol = (
        <>
            <h4>Title: </h4>
            <input
                className="form-control"
                value={detail.title}
                onChange={(e) =>
                    setDetail((d) => ({ ...d, title: e.target.value }))
                }
            />
            <hr />
            <h4>Description: </h4>
            <textarea
                className="form-control"
                value={detail.description}
                onChange={(e) =>
                    setDetail((d) => ({ ...d, description: e.target.value }))
                }
            />
            <hr />
            <h4>Tags: </h4>
            <input
                className="form-control"
                value={detail.tags.join(",")}
                onChange={(e) =>
                    setDetail((d) => ({
                        ...d,
                        tags: e.target.value.split(","),
                    }))
                }
            />
            <hr />
            <h4>Content: </h4>
            <CodeMirror
                value={detail.content}
                extensions={[
                    markdown({
                        base: markdownLanguage,
                        codeLanguages: languages,
                    }),
                ]}
                onChange={(md) => setDetail((d) => ({ ...d, content: md }))}
            />
            <hr />
            <Button
                onClick={() => {
                    modifyArticle(auth.token, {
                        ...detail,
                        tags: detail.tags
                            .map((item) => item.trim())
                            .filter((item) => item !== ""),
                    })
                        .then(() =>
                            setEditorMsg(
                                "Edit Success. You may need to refresh the page."
                            )
                        )
                        .catch((err) => {
                            console.error("[modify]", err);
                            setEditorMsg(`Edit Failed with Error: ${err}`);
                        });
                }}
            >
                Submit
            </Button>
        </>
    );

    const articleCol = (
        <article className="blog-post">
            <h2 className="blog-post-title">{detail.title}</h2>
            {auth.isLogin && (
                <Button
                    variant="primary"
                    onClick={() => setShowEditor((before) => !before)}
                >
                    Edit
                </Button>
            )}
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

    const renderedPage = (
        <>
            <Row>
                <Col>{articleCol}</Col>
                {showEditor && <Col>{editorCol}</Col>}
            </Row>
            <div className="blog-editor-toast">
                <Toast
                    onClose={() => setEditorMsg("")}
                    show={editorMsg !== ""}
                    delay={3000}
                    autohide
                >
                    <Toast.Header>
                        <strong className="me-auto">Message</strong>
                    </Toast.Header>
                    <Toast.Body>{editorMsg}</Toast.Body>
                </Toast>
            </div>
        </>
    );

    return <>{errorMsg ? <ErrorPage msg={errorMsg} /> : renderedPage}</>;
}

export default ContentPage;
