import {
  GuardianLoginOutputBody,
  GuardianSignUpOutputBody,
  loginGuardianResponse,
  signupGuardianResponse,
  useLoginGuardian,
  useSignupGuardian,
  useGetGuardianById,
  Guardian,
  useUpdateGuardian,
  updateGuardianResponse,
} from "@skillspark/api-client";
import * as SecureStore from "expo-secure-store";
import { router } from "expo-router";
import { createContext, useState, useEffect, ReactNode } from "react";

interface AuthContextType {
  guardianId: string | null;
  jwt: string | null;
  langPref: string | null;
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
  update: (
    id: string,
    email: string,
    language_preference: string,
    name: string,
    username: string,
    onError: (msg: string) => void,
    profile_picture_s3_key?: string | undefined,
    expo_push_token?: string | undefined,
  ) => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [guardianId, setGuardianId] = useState<string | null>(null);
  const [jwt, setJWT] = useState<string | null>(null);
  const [langPref, setLangPref] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { mutate: loginFunc } = useLoginGuardian();
  const { mutate: signupFunc } = useSignupGuardian();
  const { mutate: updateFunc } = useUpdateGuardian();

  const { data: guardianData } = useGetGuardianById(guardianId!, {
    query: {
      enabled: !!guardianId,
    }
  });
  let guardian = (guardianData as unknown as { data: Guardian })?.data;

  useEffect(() => {
    const checkAlreadyAuth = async () => {
      const storedJWT = await SecureStore.getItemAsync("token");
      const storedGuardianId = await SecureStore.getItemAsync("guardian_id");
      const storedLangPref = await SecureStore.getItemAsync("language_preference");
      if (storedJWT && storedGuardianId) {
        setJWT(storedJWT);
        setGuardianId(storedGuardianId);
        setLangPref(storedLangPref);
      }
      setIsLoading(false);
    };
    checkAlreadyAuth();
  }, []);

  useEffect(() => {
    const getUpdatedLangPref = async () => {
      if (!guardian) return;
      const pref = guardian?.language_preference ?? "en"; // default to english
      await SecureStore.setItemAsync("language_preference", pref);
      setLangPref(pref);
    };
    getUpdatedLangPref();
  }, [guardian]);

  const login = (
    email: string,
    password: string,
    onError: (msg: string) => void,
  ) => {
    loginFunc(
      { data: { email, password } },
      {
        onSuccess: async (resp: loginGuardianResponse) => {
          const success = resp.data as GuardianLoginOutputBody;
          await SecureStore.setItemAsync("token", success.token);
          setJWT(success.token);
          await SecureStore.setItemAsync("guardian_id", success.guardian_id);
          setGuardianId(success.guardian_id);
          router.replace("/(app)/(tabs)");
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
        onSuccess: async (resp: signupGuardianResponse) => {
          const success = resp.data as GuardianSignUpOutputBody;
          await SecureStore.setItemAsync("token", success.token);
          setJWT(success.token);
          await SecureStore.setItemAsync("guardian_id", success.guardian_id);
          setGuardianId(success.guardian_id);
          router.replace("/(app)/(tabs)");
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

  const logout = async () => {
    router.replace("/(auth)/login");
    await SecureStore.deleteItemAsync("token");
    setJWT(null);
    await SecureStore.deleteItemAsync("guardian_id");
    setGuardianId(null);
    await SecureStore.deleteItemAsync("language_preference");
    setLangPref(null);
  };

  const update = (
    id: string,
    email: string,
    language_preference: string,
    name: string,
    username: string,
    onError: (msg: string) => void,
    profile_picture_s3_key?: string | undefined,
    expo_push_token?: string | undefined,
  ) => {
    updateFunc(
      { 
        id: id,
        data: {
            email,
            language_preference,
            name,
            profile_picture_s3_key,
            username,
            expo_push_token,
        },
      },
      {
        onSuccess: async (resp: updateGuardianResponse) => {
          guardian = resp.data as Guardian;
          console.log(JSON.stringify(resp));
          console.log(JSON.stringify(guardian));
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string }};
          onError(fail.data?.message ?? "An unexpected error occurred");
          console.log(JSON.stringify(err));
        },
      },
    );
  };

  return (
    <AuthContext.Provider
      value={{
        guardianId,
        jwt,
        langPref,
        isAuthenticated: !!(jwt && guardianId),
        isLoading,
        login,
        signup,
        logout,
        update
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
