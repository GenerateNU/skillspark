import {
	useGetGuardianById,
	useGetChildrenByGuardianId,
} from "@skillspark/api-client";

export function useGuardian(guardianId: string | null) {
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
