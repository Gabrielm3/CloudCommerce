import { useState, useEffect, createContext } from "react";
import { userApi } from "@/api";

export const AuthContext = createContext();

export function AuthProvider(props) {
  const { children } = props;
  const [user, setUser] = useState(null);
  const [isAdmin, setIsAdmin] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    setLoading(false);
  }, []);

  const login = async () => {
    try {
      const response = await userApi.me();
      console.log(response);
      setUser(response);
      setIsAdmin(response.userStatus === 0);
      setLoading(false);
    } catch (error) {
      console.log(error);
    }
  };

  const data = {
    user,
    login,
  };

  if (loading) return null;

  return <AuthContext.Provider value={data}>{children}</AuthContext.Provider>;
}
