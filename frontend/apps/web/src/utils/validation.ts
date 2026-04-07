import type {
  ManagerSignUpInputBody,
  CreateOrganizationBody,
} from "@skillspark/api-client";
import type { ManagerFormInput } from "../components/admin_createModal";

export type ManagerErrors = Partial<
  Record<
    keyof Omit<
      ManagerSignUpInputBody,
      "auth_id" | "organization_id" | "profile_picture_s3_key"
    >,
    string
  >
>;
export type OrgErrors = Partial<Record<keyof CreateOrganizationBody, string>>;

const generateOnetimePassword = (): string => {
  const upper: string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ";
  const lower: string = "abcdefghijklmnopqrstuvwxyz";
  const numbers: string = "0123456789";
  const special: string = "!@#$%^&*";
  const all: string = upper + lower + numbers + special;

  const randomChar = (charset: string): string => {
    const array = new Uint32Array(1);
    crypto.getRandomValues(array);
    return charset[array[0] % charset.length];
  };

  // ensures that there is at least one uppercase, lowercase, number, and special character in the password
  const required: string[] = [
    randomChar(upper),
    randomChar(lower),
    randomChar(numbers),
    randomChar(special),
  ];

  const remaining: string[] = [];
  const array = new Uint32Array(8);
  crypto.getRandomValues(array);
  array.forEach(function (val: number) {
    remaining.push(all[val % all.length]);
  });

  const combined: string[] = [...required, ...remaining];
  const shuffled = new Uint32Array(combined.length);
  crypto.getRandomValues(shuffled);
  combined.sort(function (a, b) {
    return shuffled[combined.indexOf(a)] - shuffled[combined.indexOf(b)];
  });

  return combined.join("");
};

export const isValidEmail = (v: string): boolean =>
  /^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(v.trim());

export const isValidUUID = (v: string): boolean =>
  /^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$/i.test(
    v.trim(),
  );

export const blankMgr = (): ManagerSignUpInputBody => ({
  name: "",
  email: "",
  username: "",
  password: generateOnetimePassword(),
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

export function validateMgr(m: ManagerFormInput): ManagerErrors {
  const e: ManagerErrors = {};
  if (!m.name.trim()) e.name = "Required";
  if (!m.email.trim()) e.email = "Required";
  else if (!isValidEmail(m.email)) e.email = "Invalid email address";
  if (!m.username.trim()) e.username = "Required";
  if (!m.role.trim()) e.role = "Required";
  if (!m.language_preference.trim()) e.language_preference = "Required";
  return e;
}

export const isValidPassword = (v: string): boolean => {
  return (
    v.length >= 8 &&
    /[A-Z]/.test(v) &&
    /[a-z]/.test(v) &&
    /[0-9]/.test(v) &&
    /[!@#$%^&*]/.test(v)
  );
};
