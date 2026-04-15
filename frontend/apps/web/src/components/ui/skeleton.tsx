import { cn } from "@/lib/utils";

export function Skeleton({ className }: { className?: string }) {
	return (
		<div
			className={cn(
				"animate-pulse rounded-[4px] bg-gray-200",
				className,
			)}
		/>
	);
}
