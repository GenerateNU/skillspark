import { GuardianLoginOutputBody, GuardianSignUpOutputBody, useLoginGuardian, useSignupGuardian } from '@skillspark/api-client';
import * as SecureStore from 'expo-secure-store';
import { router } from 'expo-router';
import { createContext, useState, useEffect, ReactNode } from 'react';

interface AuthContextType {
  guardianId: string | null;
  jwt: string | null;
  isAuthenticated: boolean;
  isLoading: boolean;
  login: (
    email: string, 
    password: string, 
    onError: (msg: string) => void
  ) => void;
  signup: (
    name: string, 
    email: string,
    username: string,
    password: string,
    language_preference: string,
    profile_picture_s3_key: string | undefined,
    onError: (msg: string) => void
  ) => void;
  logout: () => void;
}

export const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const [guardianId, setGuardianId] = useState<string | null>(null);
  const [jwt, setJWT] = useState<string | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const { mutate: loginFunc } = useLoginGuardian();
  const { mutate: signupFunc } = useSignupGuardian();

  useEffect(() => {
    const checkAlreadyAuth = async () => {
      const storedJWT = await SecureStore.getItemAsync('token');
      const storedGuardianId = await SecureStore.getItemAsync('guardian_id');
      if (storedJWT && storedGuardianId) {
        setJWT(storedJWT);
        setGuardianId(storedGuardianId);
      }
      setIsLoading(false);
    };
    checkAlreadyAuth();
  }, []);

  const login = (email: string, password: string, onError: (msg: string) => void) => {
    loginFunc(
      { data: { email, password } },
      {
        onSuccess: async (resp) => {
          const success = resp as unknown as GuardianLoginOutputBody;
          await SecureStore.setItemAsync('token', success.token);
          setJWT(success.token);
          await SecureStore.setItemAsync('guardian_id', success.guardian_id);
          setGuardianId(success.guardian_id);
          router.replace('/(app)/(tabs)');
        },
        onError: (err) => {
          const fail = err as Error;
          onError(fail.message ?? "An unexpected error occurred");
        }
      }
    );
  };

  const signup = (
    name: string, 
    email: string, 
    username: string, 
    password: string, 
    language_preference: string, 
    profile_picture_s3_key: string | undefined, 
    onError: (msg: string) => void) => {
    signupFunc(
      { data: { name, email, username, password, language_preference, profile_picture_s3_key } },
      {
        onSuccess: async (resp) => {
          const success = resp as unknown as GuardianSignUpOutputBody;
          await SecureStore.setItemAsync('token', success.token);
          setJWT(success.token);
          await SecureStore.setItemAsync('guardian_id', success.guardian_id);
          setGuardianId(success.guardian_id);
          router.replace('/(app)/(tabs)');
        },
        onError: (err) => {
          const fail = err as Error;
          onError(fail.message ?? "An unexpected error occurred");
        }
      }
    );
  }

  const logout = async () => {
    await SecureStore.deleteItemAsync('token');
    setJWT(null);
    await SecureStore.deleteItemAsync('guardian_id');
    setGuardianId(null);
    router.replace('/(auth)/login');
  }

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

