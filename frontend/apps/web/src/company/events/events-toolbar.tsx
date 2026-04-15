import { MagnifyingGlass, SquaresFour, List } from "@phosphor-icons/react";
import { Button } from "@/components/ui/button";
import { ICON_SIZE_MD } from "./constants";

export type StatusFilter = "all" | "scheduled" | "cancelled" | "past";
export type ViewMode = "grid" | "table";

interface EventsToolbarProps {
	statusFilter: StatusFilter;
	onStatusFilterChange: (filter: StatusFilter) => void;
	searchQuery: string;
	onSearchQueryChange: (query: string) => void;
	viewMode: ViewMode;
	onViewModeChange: (mode: ViewMode) => void;
	counts: { all: number; scheduled: number; cancelled: number; past: number };
}

const FILTERS: { key: StatusFilter; label: string }[] = [
	{ key: "all", label: "All" },
	{ key: "scheduled", label: "Scheduled" },
	{ key: "past", label: "Past" },
	{ key: "cancelled", label: "Cancelled" },
];

export default function EventsToolbar({
	statusFilter,
	onStatusFilterChange,
	searchQuery,
	onSearchQueryChange,
	viewMode,
	onViewModeChange,
	counts,
}: EventsToolbarProps) {
	return (
		<div className="flex flex-col gap-3 sm:flex-row sm:items-center sm:justify-between">
			<div className="flex gap-1">
				{FILTERS.map(({ key, label }) => (
					<Button
						key={key}
						variant={statusFilter === key ? "secondary" : "ghost"}
						size="sm"
						className="cursor-pointer"
						onClick={() => onStatusFilterChange(key)}
					>
						{label}
						<span className="ml-1 text-xs text-gray-400">
							{counts[key]}
						</span>
					</Button>
				))}
			</div>
			<div className="flex items-center gap-2">
				<div className="relative">
					<MagnifyingGlass
						size={ICON_SIZE_MD}
						className="absolute top-1/2 left-2.5 -translate-y-1/2 text-gray-400"
					/>
					<input
						type="text"
						placeholder="Search events..."
						value={searchQuery}
						onChange={(e) => onSearchQueryChange(e.target.value)}
						className="h-8 w-56 rounded-[4px] border border-gray-200 py-1.5 pl-9 pr-4 text-sm text-gray-900 placeholder:text-gray-400 focus:border-blue-500 focus:outline-none focus:ring-1 focus:ring-blue-500"
					/>
				</div>
				<div className="flex gap-0.5">
					<Button
						variant={viewMode === "grid" ? "secondary" : "ghost"}
						size="icon-sm"
						onClick={() => onViewModeChange("grid")}
					>
						<SquaresFour size={ICON_SIZE_MD} />
					</Button>
					<Button
						variant={viewMode === "table" ? "secondary" : "ghost"}
						size="icon-sm"
						onClick={() => onViewModeChange("table")}
					>
						<List size={ICON_SIZE_MD} />
					</Button>
				</div>
			</div>
		</div>
	);
}
