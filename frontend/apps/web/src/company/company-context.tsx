import { createContext, useContext } from "react";
import { useParams } from "react-router-dom";
import {
	useGetOrganization,
	type Organization,
} from "@skillspark/api-client";

interface CompanyContextValue {
	companyId: string;
	organization: Organization | undefined;
	isLoading: boolean;
}

const CompanyContext = createContext<CompanyContextValue | null>(null);

export function CompanyProvider({ children }: { children: React.ReactNode }) {
	const { id } = useParams<{ id: string }>();
	const { data, isLoading } = useGetOrganization(id!);

	const organization = data?.data as Organization | undefined;

	return (
		<CompanyContext.Provider
			value={{
				companyId: id!,
				organization,
				isLoading,
			}}
		>
			{children}
		</CompanyContext.Provider>
	);
}

export function useCompany() {
	const ctx = useContext(CompanyContext);
	if (!ctx) {
		throw new Error("useCompany must be used within a CompanyProvider");
	}
	return ctx;
}
