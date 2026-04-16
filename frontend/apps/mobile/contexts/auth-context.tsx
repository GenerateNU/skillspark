import { useGuardian } from "@/hooks/use-guardian";
import i18n from "@/i18n";
import {
  createStripeCustomer,
  GuardianLoginOutputBody,
  GuardianSignUpOutputBody,
  loginGuardianResponse,
  setCurrentLanguage,
  signupGuardianResponse,
  updateGuardianResponse,
  useLoginGuardian,
  useSignupGuardian,
  useUpdateGuardian,
  usernameExists as checkUsernameExists,
  usernameExistsResponseError,
  UsernameExistsOutputBody,
} from "@skillspark/api-client";
import { useQueryClient } from "@tanstack/react-query";
import { router } from "expo-router";
import * as Linking from "expo-linking";
import * as SecureStore from "expo-secure-store";
import { createContext, ReactNode, useEffect, useState } from "react";

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
    onSuccess: () => void,
    onError: (msg: string) => void,
    id: string,
    email: string,
    language_preference: string,
    name: string,
    username: string,
    profile_picture_s3_key?: string | undefined,
    expo_push_token?: string | undefined,
    push_notifications?: boolean,
    email_notifications?: boolean,
  ) => void;
  usernameExists: (
    username: string,
    onError: (msg: string) => void,
  ) => Promise<boolean>;
}

export const AuthContext = createContext<AuthContextType | undefined>(
  undefined,
);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [guardianId, setGuardianId] = useState<string | null>(null);
  const [jwt, setJWT] = useState<string | null>(null);
  const [langPref, setLangPref] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(false);
  const queryClient = useQueryClient();
  const { mutate: loginFunc } = useLoginGuardian();
  const { mutate: signupFunc } = useSignupGuardian();
  const { mutate: updateFunc } = useUpdateGuardian();

  const { guardian } = useGuardian(guardianId);

  useEffect(() => {
    const checkAlreadyAuth = async () => {
      const storedJWT = await SecureStore.getItemAsync("token");
      const storedGuardianId = await SecureStore.getItemAsync("guardian_id");
      const storedLangPref = await SecureStore.getItemAsync(
        "language_preference",
      );
      if (storedLangPref) {
        await i18n.changeLanguage(storedLangPref);
        setCurrentLanguage(storedLangPref);
        queryClient.invalidateQueries({ refetchType: "all" });
      }
      if (storedJWT && storedGuardianId) {
        setJWT(storedJWT);
        setGuardianId(storedGuardianId);
        setLangPref(storedLangPref);
      }
      setIsLoading(false);
    };
    checkAlreadyAuth();
  }, [queryClient]);

  useEffect(() => {
    const getUpdatedLangPref = async () => {
      if (!guardian) return;
      const pref = guardian.language_preference ?? "en";
      await i18n.changeLanguage(pref);
      setCurrentLanguage(pref);
      await SecureStore.setItemAsync("language_preference", pref);
      setLangPref(pref);
      queryClient.invalidateQueries({ refetchType: "all" });
    };
    getUpdatedLangPref();
  }, [guardian, queryClient]);

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
          const initialUrl = await Linking.getInitialURL().catch(() => null);
          if (!initialUrl) {
            router.replace("/(app)/(tabs)");
          }
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string } };
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
          await createStripeCustomer(success.guardian_id);
          const initialUrl = await Linking.getInitialURL().catch(() => null);
          if (!initialUrl) {
            router.replace("/(app)/(tabs)");
          }
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string } };
          onError(fail.data?.message ?? "An unexpected error occurred");
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
    onSuccess: () => void,
    onError: (msg: string) => void,
    id: string,
    email: string,
    language_preference: string,
    name: string,
    username: string,
    profile_picture_s3_key?: string | undefined,
    expo_push_token?: string | undefined,
    push_notifications?: boolean,
    email_notifications?: boolean,
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
          push_notifications,
          email_notifications,
        },
      },
      {
        onSuccess: async (_resp: updateGuardianResponse) => {
          // refetch all getGuardian queries to show changes
          queryClient.invalidateQueries({
            queryKey: [`/api/v1/guardians/${id}`],
          });
          onSuccess();
        },
        onError: (err) => {
          const fail = err as unknown as { data?: { message?: string } };
          onError(fail.data?.message ?? "An unexpected error occurred");
        },
      },
    );
  };

  const usernameExists = async (
    username: string,
    onError: (msg: string) => void,
  ) => {
    try {
      const resp = await checkUsernameExists(username);
      const data = resp.data as UsernameExistsOutputBody;
      if (data.exists) {
        onError("Username is taken.");
        return false;
      }
      return true;
    } catch (err) {
      const typedErr = err as usernameExistsResponseError;
      onError(typedErr.data?.detail ?? "An unexpected error occurred.");
      return false;
    }
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
        update,
        usernameExists,
      }}
    >
      {children}
    </AuthContext.Provider>
  );
}
