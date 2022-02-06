import axios from "axios";

axios.defaults.baseURL = "http://127.0.0.1:8080/api/v1";

type ResponseWrap<T> = {
    code: number;
    msg: string;
    body: T;
};

type ArticleReq = {
    author?: string;
    page?: number;
    tag?: string;
    uuid?: string;
};

type ArticleListResp = {
    uuid: string;
    created_at: string;
    title: string;
    description: string;
};

type ArticleResp = {
    uuid: string;
    CreatedAt: string;
    UpdatedAt: string;
    author: {
        nick_name: string;
        user_name: string;
    };
    title: string;
    content: string;
    description: string;
    tags: string[];
};

async function articleRequest<T>(
    address: string,
    data: ArticleReq
): Promise<T> {
    try {
        const response = await axios.post<ResponseWrap<T>>(address, data);
        if (response.data.code !== 0) return Promise.reject(response.data.msg);
        return response.data.body;
    } catch (err) {
        return Promise.reject(err);
    }
}

async function getArticleList(data: ArticleReq): Promise<ArticleListResp[]> {
    return articleRequest<ArticleListResp[]>("/article/list", data);
}

async function getArticlePages(data: ArticleReq): Promise<number> {
    return articleRequest<number>("/article/page", data);
}

async function getArticle(data: ArticleReq): Promise<ArticleResp> {
    return articleRequest<ArticleResp>("/article/content", data);
}

export { getArticleList, getArticlePages, getArticle };
export type { ArticleListResp, ArticleResp };
