import React, { useEffect, useState } from "react";

import { Navigate, useLocation } from "react-router-dom";
import { userLogin, userLogout } from "./api";

interface AuthContextType {
    token: string;
    isLogin: boolean;
    login: (
        username: string,
        password: string,
        onSuccess: VoidFunction,
        onFailed: (error: string) => void
    ) => void;
    logout: (
        onSuccess: VoidFunction,
        onFailed: (error: string) => void
    ) => void;
}

const AuthContext = React.createContext<AuthContextType>({
    token: "",
    isLogin: false,
    login: (
        username: string,
        password: string,
        onSuccess: VoidFunction,
        onFailed: (error: string) => void
    ) => {
        onFailed("not implemented");
    },
    logout: (onSuccess: VoidFunction, onFailed: (error: string) => void) => {
        onFailed("not implemented");
    },
});

function AuthProvider({ children }: { children: React.ReactNode }) {
    const [token, setToken] = useState<string>("");
    const [isLogin, setIsLogin] = useState<boolean>(false);

    useEffect(() => {
        const token = localStorage.getItem("token");
        if (token) {
            setToken(token);
            setIsLogin(true);
        }
    }, []);

    const login = (
        username: string,
        password: string,
        onSuccess: VoidFunction,
        onFailed: (error: string) => void
    ) => {
        userLogin({ username: username, password: password })
            .then((token) => {
                setToken(token);
                localStorage.setItem("token", token);
                setIsLogin(true);
                onSuccess();
            })
            .catch((err) => {
                console.error("[user] userLogin", err);
                setIsLogin(false);
                onFailed(err);
            });
    };

    const logout = (
        onSuccess: VoidFunction,
        onFailed: (error: string) => void
    ) => {
        localStorage.removeItem("token");
        setToken("");
        setIsLogin(false);

        if (!isLogin) {
            onFailed("not logged in");
            return;
        }

        userLogout(token)
            .then(() => {
                onSuccess();
            })
            .catch((err) => {
                console.error("[user] userLogout", err);
                onFailed(err);
            });
    };

    const value = { token, isLogin, login, logout };

    return (
        <AuthContext.Provider value={value}>{children}</AuthContext.Provider>
    );
}

function useAuth() {
    return React.useContext(AuthContext);
}

function RequireAuth({ children }: { children: JSX.Element }) {
    const auth = useAuth();
    const location = useLocation();

    if (!auth.token) {
        return <Navigate to="/login" state={{ from: location }} replace />;
    }

    return children;
}

export { AuthProvider, RequireAuth, useAuth };
