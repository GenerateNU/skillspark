/// <reference types="vite/client" />

declare module "@tailwindcss/vite";

interface ImportMetaEnv {
  readonly VITE_ADMIN_NAME: string;
  readonly VITE_ADMIN_EMAIL: string;
  readonly VITE_ADMIN_USERNAME: string;
  readonly VITE_ADMIN_LANG_PREFERENCE: string;
  readonly VITE_ADMIN_ROLE: string;
}
