import { createContext, useState, useEffect, type ReactNode } from "react";
import { login as loginRequest, signup as signupRequest } from "../services/auth";
import type { LoginRequest, SignupRequest } from "../services/auth";

interface AuthContextType {
  token: string | null;
  isAuthenticated: boolean;
  login: (data: LoginRequest) => Promise<void>;
  signup: (data: SignupRequest) => Promise<void>;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

const TOKEN_KEY = "niavo_token";

export function AuthProvider({ children }: { children: ReactNode }) {
  const [token, setToken] = useState<string | null>(() =>
    localStorage.getItem(TOKEN_KEY)
  );

  useEffect(() => {
    if (token) {
      localStorage.setItem(TOKEN_KEY, token);
    } else {
      localStorage.removeItem(TOKEN_KEY);
    }
  }, [token]);

  async function login(data: LoginRequest) {
    const response = await loginRequest(data);
    setToken(response.token);
  }

  async function signup(data: SignupRequest) {
    await signupRequest(data);
  }

  function logout() {
    setToken(null);
  }

  const value: AuthContextType = {
    token,
    isAuthenticated: token !== null,
    login,
    signup,
    logout,
  };

  return <AuthContext.Provider value={value}>{children}</AuthContext.Provider>;
}