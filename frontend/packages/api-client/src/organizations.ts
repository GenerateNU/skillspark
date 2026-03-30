import { type UseMutationOptions, type MutationFunction, QueryClient, type UseMutationResult, useMutation } from "@tanstack/react-query";
import { customInstance } from "./apiClient";
import type { createOrganizationResponseSuccess, createOrganizationResponseError, HTTPStatusCodes } from "./generated/organizations/organizations";
import type { CreateOrganizationBody, ErrorModel, Organization, UpdateOrganizationBody } from "./generated/skillSparkAPI.schemas";

type SecondParameter<T extends (...args: never) => unknown> = Parameters<T>[1];

export type createOrganizationResponse = (createOrganizationResponseSuccess | createOrganizationResponseError)

const getCreateOrganizationUrl = () => {
  return `/api/v1/organizations`
}

export const createOrganization = async (createOrganizationBody: CreateOrganizationBody, options?: RequestInit): Promise<createOrganizationResponse> => {
  const formData = new FormData();
  if (createOrganizationBody.active !== undefined) {
    formData.append(`active`, createOrganizationBody.active.toString())
  }
  if (createOrganizationBody.location_id !== undefined) {
    formData.append(`location_id`, createOrganizationBody.location_id);
  }
  formData.append(`name`, createOrganizationBody.name);
  if (createOrganizationBody.profile_image !== undefined) {
    formData.append(`profile_image`, createOrganizationBody.profile_image);
  }

  return customInstance<createOrganizationResponse>(getCreateOrganizationUrl(),
    {
      ...options,
      method: 'POST'
      ,
      body:
        formData,
    }
  );
}

export const getCreateOrganizationMutationOptions = <TError = ErrorModel,
  TContext = unknown>(options?: { mutation?: UseMutationOptions<Awaited<ReturnType<typeof createOrganization>>, TError, { data: CreateOrganizationBody }, TContext>, request?: SecondParameter<typeof customInstance> }
  ): UseMutationOptions<Awaited<ReturnType<typeof createOrganization>>, TError, { data: CreateOrganizationBody }, TContext> => {

  const mutationKey = ['createOrganization'];
  const { mutation: mutationOptions, request: requestOptions } = options ?
    options.mutation && 'mutationKey' in options.mutation && options.mutation.mutationKey ?
      options
      : { ...options, mutation: { ...options.mutation, mutationKey } }
    : { mutation: { mutationKey, }, request: undefined };

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof createOrganization>>, { data: CreateOrganizationBody }> = (props) => {
    const { data } = props ?? {};

    return createOrganization(data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type CreateOrganizationMutationResult = NonNullable<Awaited<ReturnType<typeof createOrganization>>>
export type CreateOrganizationMutationBody = CreateOrganizationBody
export type CreateOrganizationMutationError = ErrorModel

/**
* @summary Create a new organization
*/
export const useCreateOrganization = <TError = ErrorModel,
  TContext = unknown>(options?: { mutation?: UseMutationOptions<Awaited<ReturnType<typeof createOrganization>>, TError, { data: CreateOrganizationBody }, TContext>, request?: SecondParameter<typeof customInstance> }
    , queryClient?: QueryClient): UseMutationResult<
      Awaited<ReturnType<typeof createOrganization>>,
      TError,
      { data: CreateOrganizationBody },
      TContext
    > => {
  return useMutation(getCreateOrganizationMutationOptions(options), queryClient);
}

/**
* Updates an existing organization with the provided fields (partial update)
* @summary Update an organization
*/
export type updateOrganizationResponse200 = {
  data: Organization
  status: 200
}

export type updateOrganizationResponseDefault = {
  data: ErrorModel
  status: Exclude<HTTPStatusCodes, 200>
}

export type updateOrganizationResponseSuccess = (updateOrganizationResponse200) & {
  headers: Headers;
};
export type updateOrganizationResponseError = (updateOrganizationResponseDefault) & {
  headers: Headers;
};

export type updateOrganizationResponse = (updateOrganizationResponseSuccess | updateOrganizationResponseError)

export const getUpdateOrganizationUrl = (id: string,) => {
  return `/api/v1/organizations/${id}`
}

export const updateOrganization = async (id: string,
  updateOrganizationBody: UpdateOrganizationBody, options?: RequestInit): Promise<updateOrganizationResponse> => {
  const formData = new FormData();
  if (updateOrganizationBody.active !== undefined) {
    formData.append(`active`, updateOrganizationBody.active.toString())
  }
  if (updateOrganizationBody.location_id !== undefined) {
    formData.append(`location_id`, updateOrganizationBody.location_id instanceof Blob ? updateOrganizationBody.location_id : new Blob([updateOrganizationBody.location_id], { type: 'text/plain' }));
  }
  formData.append(`name`, updateOrganizationBody.name instanceof Blob ? updateOrganizationBody.name : new Blob([updateOrganizationBody.name], { type: 'text/plain' }));
  if (updateOrganizationBody.profile_image !== undefined) {
    formData.append(`profile_image`, updateOrganizationBody.profile_image);
  }

  return customInstance<updateOrganizationResponse>(getUpdateOrganizationUrl(id),
    {
      ...options,
      method: 'PATCH'
      ,
      body:
        formData,
    }
  );
}

export const getUpdateOrganizationMutationOptions = <TError = ErrorModel,
  TContext = unknown>(options?: { mutation?: UseMutationOptions<Awaited<ReturnType<typeof updateOrganization>>, TError, { id: string; data: UpdateOrganizationBody }, TContext>, request?: SecondParameter<typeof customInstance> }
  ): UseMutationOptions<Awaited<ReturnType<typeof updateOrganization>>, TError, { id: string; data: UpdateOrganizationBody }, TContext> => {

  const mutationKey = ['updateOrganization'];
  const { mutation: mutationOptions, request: requestOptions } = options ?
    options.mutation && 'mutationKey' in options.mutation && options.mutation.mutationKey ?
      options
      : { ...options, mutation: { ...options.mutation, mutationKey } }
    : { mutation: { mutationKey, }, request: undefined };

  const mutationFn: MutationFunction<Awaited<ReturnType<typeof updateOrganization>>, { id: string; data: UpdateOrganizationBody }> = (props) => {
    const { id, data } = props ?? {};

    return updateOrganization(id, data, requestOptions)
  }

  return { mutationFn, ...mutationOptions }
}

export type UpdateOrganizationMutationResult = NonNullable<Awaited<ReturnType<typeof updateOrganization>>>
export type UpdateOrganizationMutationBody = UpdateOrganizationBody
export type UpdateOrganizationMutationError = ErrorModel

/**
* @summary Update an organization
*/
export const useUpdateOrganization = <TError = ErrorModel,
  TContext = unknown>(options?: { mutation?: UseMutationOptions<Awaited<ReturnType<typeof updateOrganization>>, TError, { id: string; data: UpdateOrganizationBody }, TContext>, request?: SecondParameter<typeof customInstance> }
    , queryClient?: QueryClient): UseMutationResult<
      Awaited<ReturnType<typeof updateOrganization>>,
      TError,
      { id: string; data: UpdateOrganizationBody },
      TContext
    > => {
  return useMutation(getUpdateOrganizationMutationOptions(options), queryClient);
}