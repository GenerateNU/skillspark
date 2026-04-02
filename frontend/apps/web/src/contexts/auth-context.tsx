import type {
  GuardianLoginOutputBody,
  GuardianSignUpOutputBody,
  loginGuardianResponse,
  signupGuardianResponse,
} from "@skillspark/api-client";
import {
  useLoginGuardian,
  useSignupGuardian,
} from "@skillspark/api-client";
import { useNavigate } from "react-router-dom";
import { createContext, useEffect, useState } from "react";

interface AuthContextType {
  guardianId: string | null;
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
    profile_picture_s3_key: string | undefined,
    onError: (msg: string) => void,
  ) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: React.ReactNode }) {
  const [guardianId, setGuardianId] = useState<string | null>(null);
  const [jwt, setJWT] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { mutate: loginFunc } = useLoginGuardian();
  const { mutate: signupFunc } = useSignupGuardian();
  const navigate = useNavigate();
  
  useEffect(() => {
    const checkAlreadyAuth = () => {
      const storedGuardianId = localStorage.getItem("guardian_id");
      const storedJwt = localStorage.getItem("jwt");

      if (storedJwt && storedGuardianId) {
        setJWT(storedJwt);
        setGuardianId(storedGuardianId);
      }

      setIsLoading(false);
    };
    checkAlreadyAuth();
  }, []);

  const login = (
    email: string,
    password: string,
    onError: (msg: string) => void,
  ) => {
    loginFunc(
      { data: { email, password } },
      {
        onSuccess: (resp: loginGuardianResponse) => {
          const success = resp.data as GuardianLoginOutputBody;
          localStorage.setItem("token", success.token);
          setJWT(success.token);
          localStorage.setItem("guardian_id", success.guardian_id);
          setGuardianId(success.guardian_id);
          navigate("/");
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string }};
          onError(fail.data?.message ?? "An unexpected error occurred");
        },
      },
    );
  };

  const signup = (
    name: string,
    email: string,
    username: string,
    password: string,
    language_preference: string,
    profile_picture_s3_key: string | undefined,
    onError: (msg: string) => void,
  ) => {
    signupFunc(
      {
        data: {
          name,
          email,
          username,
          password,
          language_preference,
          profile_picture_s3_key,
        },
      },
      {
        onSuccess: (resp: signupGuardianResponse) => {
          const success = resp.data as GuardianSignUpOutputBody;
          console.log(JSON.stringify(resp));
          localStorage.setItem("token", success.token);
          setJWT(success.token);
          localStorage.setItem("guardian_id", success.guardian_id);
          setGuardianId(success.guardian_id);
          navigate("/")
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string }};
          onError(
            fail.data?.message ?? "An unexpected error occurred",
          );
        },
      },
    );
  };

  const logout = () => {
    localStorage.removeItem("jwt");
    localStorage.removeItem("userId");
    setJWT(null);
    setGuardianId(null);
    navigate("/login");
  };

  return (
    <AuthContext.Provider
      value={{
        guardianId,
        jwt,
        isAuthenticated: !!(jwt && guardianId),
        isLoading,
        login,
        signup,
        logout,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}