import { Breadcrumbs } from "./breadcrumbs";
import { CompanyHeading } from "./company-heading";
import { useCompany } from "./company-context";

export default function CompanyLayout({
	page,
	children,
	showHeading = false,
}: {
	page: string;
	children: React.ReactNode;
	showHeading?: boolean;
}) {
	const { organization, isLoading } = useCompany();

	return (
		<main className="flex-1 overflow-y-auto px-8 pt-6 pb-8">
			<Breadcrumbs
				page={page}
				companyName={organization?.name}
				isLoading={isLoading}
			/>
			{showHeading && (
				<CompanyHeading
					name={organization?.name}
					about={organization?.about}
					avatarUrl={organization?.presigned_url}
					isLoading={isLoading}
				/>
			)}
			<div className="mt-6">{children}</div>
		</main>
	);
}
