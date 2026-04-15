import { Skeleton } from "@/components/ui/skeleton";

export function CompanyHeading({
	name,
	about,
	avatarUrl,
	isLoading,
}: {
	name?: string;
	about?: string;
	avatarUrl?: string;
	isLoading?: boolean;
}) {
	if (isLoading) {
		return (
			<div className="flex items-center gap-4 pt-4">
				<Skeleton className="w-12 h-12" />
				<div className="flex flex-col gap-2">
					<Skeleton className="h-5 w-40" />
					<Skeleton className="h-3.5 w-64" />
				</div>
			</div>
		);
	}

	return (
		<div className="flex items-center gap-4 pt-4">
			{avatarUrl ? (
				<img
					src={avatarUrl}
					alt={name}
					className="w-12 h-12 rounded-[4px] object-cover"
				/>
			) : (
				<div className="w-12 h-12 rounded-[4px] bg-gray-200 flex items-center justify-center text-lg font-bold text-gray-500">
					{name?.charAt(0) ?? "?"}
				</div>
			)}
			<div>
				<h1 className="text-xl font-bold text-gray-900">
					{name ?? "Company"}
				</h1>
				{about && (
					<p className="text-sm text-gray-500 mt-0.5">{about}</p>
				)}
			</div>
		</div>
	);
}
