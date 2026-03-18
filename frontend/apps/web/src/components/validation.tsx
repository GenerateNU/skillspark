import type { ManagerSignUpInputBody, CreateOrganizationBody } from "@skillspark/api-client";

export type ManagerErrors = Partial<Record<keyof Omit<ManagerSignUpInputBody, "auth_id" | "organization_id" | "profile_picture_s3_key">, string>>;
export type OrgErrors = Partial<Record<keyof CreateOrganizationBody, string>>;

// ── Helpers ───────────────────────────────────────────────────────────────────
const genOtp = (): string => Math.random().toString(36).slice(2, 10).toUpperCase();
const isValidEmail = (v: string): boolean => /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v.trim());

export const blankMgr = (): ManagerSignUpInputBody => ({
  name: "",
  email: "",
  username: "",
  password: genOtp(),
  language_preference: "en",
  organization_id: "",
  role: "",
});

export const blankOrg = (): CreateOrganizationBody => ({
  name: "",
  active: true,
  location_id: "",
});

export function validateOrg(o: CreateOrganizationBody): OrgErrors {
  const e: OrgErrors = {};
  if (!(o.name as string).trim()) e.name = "Required";
  return e;
}

export function validateMgr(m: ManagerSignUpInputBody): ManagerErrors {
  const e: ManagerErrors = {};
  if (!m.name.trim())     e.name = "Required";
  if (!m.email.trim())    e.email = "Required";
  else if (!isValidEmail(m.email)) e.email = "Invalid email address";
  if (!m.username.trim()) e.username = "Required";
  if (!m.role.trim())     e.role = "Required";
  if (!m.language_preference.trim()) e.language_preference = "Required";
  return e;
}