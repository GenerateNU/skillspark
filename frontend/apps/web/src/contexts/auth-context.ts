import { createContext } from "react";

export interface AuthContextType {
  managerId: string | null;
  jwt: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (
    email: string,
    password: string,
    onError: (msg: string) => void,
  ) => void;
  signup: (
    name: string,
    email: string,
    username: string,
    password: string,
    language_preference: string,
    organization_id: string,
    role: string,
    profile_picture_s3_key: string | undefined,
    onError: (msg: string) => void,
  ) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);
