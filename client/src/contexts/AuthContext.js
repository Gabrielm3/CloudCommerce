import { useState, useEffect, createContext } from "react";

export const AuthContext = createContext();

export function AuthProvider(props) {
    const { children } = props;
    const [user, setUser] = useState(null);

    const data = {
        user,
    };

    return <AuthContext.Provider value={data}>{children}</AuthContext.Provider>
}