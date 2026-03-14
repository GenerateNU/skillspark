import type { CreateOrganizationBody, Manager } from "@skillspark/api-client";
import { type BizErrors, isValidEmail, type ManagerErrors } from "./types";

export function validateBiz(b: CreateOrganizationBody): BizErrors {
  const e: BizErrors = {};
  const name: string = typeof b.name === "string" ? b.name : "";
  const locationId: string = typeof b.location_id === "string" ? b.location_id : "";
  if (!name.trim()) e.name = "Required";
  if (b.location_id && !locationId.trim()) e.location_id = "Invalid";
  return e;
}

export function validateMgr(m: Manager): ManagerErrors {
  const e: ManagerErrors = {};
  if (!m.name.trim())     e.name = "Required";
  if (!m.email.trim())    e.email = "Required";
  else if (!isValidEmail(m.email)) e.email = "Invalid email address";
  if (!m.username.trim()) e.username = "Required";
  if (!m.role.trim())     e.role = "Required";
  return e;
}