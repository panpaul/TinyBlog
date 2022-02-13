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

type UserReq = {
    nickname?: string;
    password?: string;
    username?: string;
};

async function sendReq<D, T>(api: string, token: string, data?: D): Promise<T> {
    try {
        const response = await axios.post<ResponseWrap<T>>(api, data, {
            headers: { token: token },
        });
        if (response.data.code !== 0) return Promise.reject(response.data.msg);
        return response.data.body;
    } catch (err) {
        return Promise.reject(err);
    }
}

async function getArticleList(data: ArticleReq): Promise<ArticleListResp[]> {
    return sendReq<ArticleReq, ArticleListResp[]>("/article/list", "", data);
}

async function getArticlePages(data: ArticleReq): Promise<number> {
    return sendReq<ArticleReq, number>("/article/page", "", data);
}

async function getArticle(data: ArticleReq): Promise<ArticleResp> {
    return sendReq<ArticleReq, ArticleResp>("/article/content", "", data);
}

async function userLogin(data: UserReq): Promise<string> {
    return sendReq<UserReq, string>("/user/login", "", data);
}

async function userLogout(token: string): Promise<void> {
    return sendReq<void, void>("/user/logout", token);
}

export { getArticleList, getArticlePages, getArticle };
export { userLogin, userLogout };
export type { ArticleListResp, ArticleResp };
