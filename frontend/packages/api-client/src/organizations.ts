import { customInstance } from "./apiClient";
import type { CreateOrganizationBody, ErrorModel, Organization, UpdateOrganizationBody } from "./generated/skillSparkAPI.schemas";
import type { createOrganizationResponseSuccess, createOrganizationResponseError, HTTPStatusCodes } from "./generated/organizations/organizations";

// ─── Create Organization ──────────────────────────────────────────────────────

export type createOrganizationResponse = createOrganizationResponseSuccess | createOrganizationResponseError;

export const createOrganization = async (
  body: CreateOrganizationBody,
  options?: RequestInit
): Promise<createOrganizationResponse> => {
  const formData = new FormData();
  if (body.active !== undefined) formData.append("active", body.active.toString());
  if (body.location_id !== undefined) formData.append("location_id", body.location_id);
  formData.append("name", body.name);
  if (body.profile_image !== undefined) formData.append("profile_image", body.profile_image);

  return customInstance<createOrganizationResponse>("/api/v1/organizations", {
    ...options,
    method: "POST",
    body: formData,
  });
};

// ─── Update Organization ──────────────────────────────────────────────────────

export type updateOrganizationResponse200 = { data: Organization; status: 200 };
export type updateOrganizationResponseDefault = { data: ErrorModel; status: Exclude<HTTPStatusCodes, 200> };
export type updateOrganizationResponseSuccess = updateOrganizationResponse200 & { headers: Headers };
export type updateOrganizationResponseError = updateOrganizationResponseDefault & { headers: Headers };
export type updateOrganizationResponse = updateOrganizationResponseSuccess | updateOrganizationResponseError;

export const updateOrganization = async (
  id: string,
  body: UpdateOrganizationBody,
  options?: RequestInit
): Promise<updateOrganizationResponse> => {
  const formData = new FormData();
  if (body.active !== undefined) formData.append("active", body.active.toString());
  if (body.location_id !== undefined) formData.append("location_id", body.location_id);
  formData.append("name", body.name);
  if (body.profile_image !== undefined) formData.append("profile_image", body.profile_image);

  return customInstance<updateOrganizationResponse>(`/api/v1/organizations/${id}`, {
    ...options,
    method: "PATCH",
    body: formData,
  });
};