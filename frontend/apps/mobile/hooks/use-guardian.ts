import {
	useGetGuardianById,
	useGetChildrenByGuardianId,
	useGetEmergencyContactsByGuardianId,
} from "@skillspark/api-client";

export function useGuardian(guardianId: string | null) {
	const { data: guardianResponse, isLoading: guardianLoading } =
		useGetGuardianById(guardianId || "");
	const { data: childrenResponse, isLoading: childrenLoading } =
		useGetChildrenByGuardianId(guardianId || "");
	const { data: emergencycontactResponse, isLoading: emergencycontactLoading } =
		useGetEmergencyContactsByGuardianId(guardianId || "");

	const guardian =
		guardianResponse?.status === 200 ? guardianResponse.data : null;
	const children =
		childrenResponse?.status === 200 ? childrenResponse.data : [];
	const emergencyContacts =
		emergencycontactResponse?.status === 200
			? emergencycontactResponse.data
			: [];

	return {
		guardian,
		children,
		emergencyContacts,
		guardianId: guardianId,
		isLoading: guardianLoading || childrenLoading || emergencycontactLoading,
	};
}
