import { Skeleton } from "@/components/ui/skeleton";

export default function EventCardSkeleton() {
	return (
		<div className="overflow-hidden rounded-[4px] border border-gray-200 bg-white">
			<Skeleton className="h-44 w-full rounded-none" />
			<div className="flex flex-col gap-2 p-4">
				<Skeleton className="h-3 w-36" />
				<Skeleton className="h-4 w-full" />
				<Skeleton className="h-4 w-3/4" />
				<Skeleton className="h-3 w-40" />
				<Skeleton className="mt-1 h-1.5 w-full" />
				<Skeleton className="h-3 w-24" />
				<Skeleton className="h-4 w-20" />
			</div>
		</div>
	);
}
