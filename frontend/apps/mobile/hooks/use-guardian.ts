import {
  useGetGuardianById,
  useGetChildrenByGuardianId,
} from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";

export function useGuardian() {
  const { guardianId } = useAuthContext();
  const { data: guardianResponse, isLoading: guardianLoading } =
    useGetGuardianById(guardianId || "");
  const { data: childrenResponse, isLoading: childrenLoading } =
    useGetChildrenByGuardianId(guardianId || "");

  const guardian =
    guardianResponse?.status === 200 ? guardianResponse.data : null;
  const children =
    childrenResponse?.status === 200 ? childrenResponse.data : [];
  return {
    guardian,
    children,
    guardianId: guardianId,
    isLoading: guardianLoading || childrenLoading,
  };
}
