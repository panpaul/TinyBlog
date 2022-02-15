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
import { PluggableList } from "react-markdown/lib/react-markdown";
import remarkGfm from "remark-gfm";
import rehypeRaw from "rehype-raw";
import emoji from "remark-emoji";
import { PrismAsync } from "react-syntax-highlighter";
import { materialLight } from "react-syntax-highlighter/dist/esm/styles/prism";

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
        uuid: "00000000-0000-0000-0000-000000000000",
        author: { nick_name: "loading", user_name: "loading" },
        title: "loading...",
        description: "loading",
        content: "# Loading",
        enable_math: false,
        tags: ["loading"],
    });

    useEffect(() => {
        getArticle({ uuid: uuid })
            .then((res) => setDetail(res))
            .catch((err) => {
                console.error("[content]", err);
                setErrorMsg(err);
            });
    }, [uuid]);

    const EditorCol = (
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
            <h4>Enable Math Render: </h4>
            <div className="form-check form-switch">
                <input
                    className="form-check-input"
                    type="checkbox"
                    checked={detail.enable_math}
                    onChange={(e) =>
                        setDetail((d) => ({
                            ...d,
                            enable_math: e.target.checked,
                        }))
                    }
                />
            </div>
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

    const tags = detail.tags.map((item) => (
        <LinkContainer
            to={`/?tag=${item}`}
            key={item}
            style={{ marginRight: "0.5rem" }}
        >
            <Badge bg="info">{item}</Badge>
        </LinkContainer>
    ));

    const [remark, setRemark] = useState<PluggableList | undefined>([
        remarkGfm,
        emoji,
    ]);
    const [rehype, setRehype] = useState<PluggableList | undefined>([
        rehypeRaw,
    ]);

    useEffect(() => {
        if (detail.enable_math) {
            import("remark-math").then((r) => {
                setRemark([remarkGfm, emoji, r.default]);
            });
            import("rehype-katex").then((r) => {
                setRehype([rehypeRaw, r.default]);
            });
        } else {
            setRemark([remarkGfm, emoji]);
            setRehype([rehypeRaw]);
        }
    }, [detail.enable_math]);

    const ArticleCol = (
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
                remarkPlugins={remark}
                rehypePlugins={rehype}
                components={{
                    // eslint-disable-next-line @typescript-eslint/no-unused-vars
                    code: ({ node, inline, className, children, ...props }) => {
                        const match = /language-(\w+)/.exec(className || "");
                        return !inline && match ? (
                            <PrismAsync
                                language={match[1]}
                                wrapLines={true}
                                showLineNumbers={true}
                                wrapLongLines={true}
                                style={materialLight}
                                PreTag="div"
                                {...props}
                            >
                                {String(children).replace(/\n$/, "")}
                            </PrismAsync>
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
                <Col>{ArticleCol}</Col>
                {showEditor && <Col>{EditorCol}</Col>}
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
