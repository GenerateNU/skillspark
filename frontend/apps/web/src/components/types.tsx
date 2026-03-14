import type { CreateOrganizationBody, Manager, ManagerSignUpInputBody } from "@skillspark/api-client";

// CreateManagerInput maps to ManagerSignUpInputBody
export type CreateManagerInput = ManagerSignUpInputBody;

export type ManagerErrors = Partial<Record<keyof Omit<Manager, "$schema" | "id" | "auth_id" | "user_id" | "organization_id" | "profile_picture_s3_key" | "created_at" | "updated_at">, string>>;
export type BizErrors = Partial<Record<keyof CreateOrganizationBody, string>>;
export type ButtonVariant = "primary" | "danger" | "ghost";
export type BadgeColor = "blue" | "green" | "yellow" | "gray";

export const uid = (): string => Math.random().toString(36).slice(2, 10);
export const genOtp = (): string => Math.random().toString(36).slice(2, 10).toUpperCase();
export const isValidEmail = (v: string): boolean => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v.trim());
export const isValidPhone = (v: string): boolean => v === "" || /^\+?[\d\s\-().]{7,20}$/.test(v.trim());

export const blankMgr = (): Manager => ({
  id: uid(),
  user_id: "",
  organization_id: "",
  role: "",
  name: "",
  email: "",
  username: "",
  profile_picture_s3_key: "",
  language_preference: "en",
  auth_id: "",
  created_at: "",
  updated_at: "",
});

export const blankBiz = (): CreateOrganizationBody => ({
  name: "",
  active: true,
  location_id: "",
});

export const fmtDate = (iso: string): string =>
  new Date(iso).toLocaleDateString("en-US", { month: "short", day: "numeric", year: "numeric" });