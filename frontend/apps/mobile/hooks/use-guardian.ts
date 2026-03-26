import { useGetGuardianById, useGetChildrenByGuardianId } from '@skillspark/api-client';

// TODO: Replace with authenticated user's guardian ID
const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export function useGuardian() {
  const { data: guardianResponse, isLoading: guardianLoading } = useGetGuardianById(GUARDIAN_ID);
  const { data: childrenResponse, isLoading: childrenLoading } = useGetChildrenByGuardianId(GUARDIAN_ID);

  const guardian = guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children = childrenResponse?.status === 200 ? childrenResponse.data : [];
  return {
    guardian,
    children,
    guardianId: GUARDIAN_ID,
    isLoading: guardianLoading || childrenLoading,
  };
}