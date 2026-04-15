import { CaretRight } from "@phosphor-icons/react";
import { Skeleton } from "@/components/ui/skeleton";

export function Breadcrumbs({
	page,
	companyName,
	isLoading,
}: {
	page: string;
	companyName?: string;
	isLoading?: boolean;
}) {
	if (isLoading) {
		return (
			<nav className="flex items-center gap-1.5 mb-4">
				<Skeleton className="h-4 w-24" />
				{page !== "Home" && (
					<>
						<CaretRight size={12} className="text-gray-300" />
						<Skeleton className="h-4 w-16" />
					</>
				)}
			</nav>
		);
	}

	const segments = [companyName ?? "Company"];
	if (page !== "Home") {
		segments.push(page);
	}

	return (
		<nav className="flex items-center gap-1.5 text-sm text-gray-400 mb-4">
			{segments.map((segment, i) => (
				<span key={i} className="flex items-center gap-1.5">
					{i > 0 && <CaretRight size={12} />}
					<span
						className={
							i === segments.length - 1
								? "text-gray-700 font-medium"
								: ""
						}
					>
						{segment}
					</span>
				</span>
			))}
		</nav>
	);
}
