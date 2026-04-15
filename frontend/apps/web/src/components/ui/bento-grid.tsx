import { cn } from "@/lib/utils";

/**
 * Responsive grid that adapts columns by breakpoint.
 * Default: 1 col → md:2 → lg:3. Override with `cols` prop.
 *
 * Usage:
 *   <BentoGrid>              — 1 / 2 / 3 columns
 *   <BentoGrid cols={4}>     — 1 / 2 / 4 columns
 *   <BentoGrid cols={2}>     — 1 / 2 columns
 */
export function BentoGrid({
	cols = 3,
	className,
	children,
}: {
	cols?: 1 | 2 | 3 | 4;
	className?: string;
	children: React.ReactNode;
}) {
	const colStyles: Record<number, string> = {
		1: "grid-cols-1",
		2: "grid-cols-1 md:grid-cols-2",
		3: "grid-cols-1 md:grid-cols-2 lg:grid-cols-3",
		4: "grid-cols-1 md:grid-cols-2 lg:grid-cols-4",
	};

	return (
		<div className={cn("grid gap-3", colStyles[cols], className)}>
			{children}
		</div>
	);
}

/**
 * Card cell within a BentoGrid. Consistent styling with optional span.
 *
 * Usage:
 *   <BentoCard>...</BentoCard>
 *   <BentoCard span={2}>...</BentoCard>         — spans 2 columns on lg
 *   <BentoCard span="full">...</BentoCard>      — spans entire row
 */
export function BentoCard({
	span,
	className,
	children,
}: {
	span?: 1 | 2 | 3 | 4 | "full";
	className?: string;
	children: React.ReactNode;
}) {
	const spanStyles: Record<string | number, string> = {
		1: "",
		2: "md:col-span-2",
		3: "lg:col-span-3",
		4: "lg:col-span-4",
		full: "col-span-full",
	};

	return (
		<div
			className={cn(
				"rounded-[4px] border border-gray-200 bg-white p-5",
				span ? spanStyles[span] : "",
				className,
			)}
		>
			{children}
		</div>
	);
}
