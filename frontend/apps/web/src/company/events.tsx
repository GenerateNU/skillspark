import { useState, useMemo } from "react";
import type { EventOccurrence } from "@skillspark/api-client";
import CompanyLayout from "./company-layout";
import { MOCK_EVENTS } from "./events/mock-events";
import EventsToolbar, {
	type StatusFilter,
	type ViewMode,
} from "./events/events-toolbar";
import EventCardGrid from "./events/event-card-grid";
import EventTable from "./events/event-table";
import CreateEventButton from "./events/create-event-button";
import { useDebounce } from "@/hooks/use-debounce";

export default function CompanyEvents() {
	const [viewMode, setViewMode] = useState<ViewMode>("grid");
	const [statusFilter, setStatusFilter] = useState<StatusFilter>("all");
	const [searchQuery, setSearchQuery] = useState("");
	const debouncedSearchQuery = useDebounce(searchQuery, 200);

	// TODO: swap with useGetEventOccurrencesByOrganizationId(companyId) when ready
	const events: EventOccurrence[] = MOCK_EVENTS;
	const isLoading = false;

	const counts = useMemo(() => {
		const now = new Date();
		return {
			all: events.length,
			scheduled: events.filter(
				(e) => e.status === "scheduled" && new Date(e.start_time) >= now,
			).length,
			past: events.filter(
				(e) => e.status === "scheduled" && new Date(e.start_time) < now,
			).length,
			cancelled: events.filter((e) => e.status === "cancelled").length,
		};
	}, [events]);

	const filtered = useMemo(() => {
		const now = new Date();
		let result = events;

		if (statusFilter === "scheduled") {
			result = result.filter(
				(e) => e.status === "scheduled" && new Date(e.start_time) >= now,
			);
		} else if (statusFilter === "past") {
			result = result.filter(
				(e) => e.status === "scheduled" && new Date(e.start_time) < now,
			);
		} else if (statusFilter === "cancelled") {
			result = result.filter((e) => e.status === "cancelled");
		}

		if (debouncedSearchQuery.trim()) {
			const q = debouncedSearchQuery.toLowerCase();
			result = result.filter(
				(e) =>
					e.event.title.toLowerCase().includes(q) ||
					e.location.address_line1.toLowerCase().includes(q) ||
					e.location.district.toLowerCase().includes(q),
			);
		}

		return result;
	}, [events, statusFilter, debouncedSearchQuery]);

	return (
		<CompanyLayout page="Events">
			<div className="flex flex-col gap-4">
				<div className="flex items-center justify-between">
					<h2 className="text-lg font-semibold text-gray-900">Events</h2>
					<CreateEventButton />
				</div>

				<EventsToolbar
					statusFilter={statusFilter}
					onStatusFilterChange={setStatusFilter}
					searchQuery={searchQuery}
					onSearchQueryChange={setSearchQuery}
					viewMode={viewMode}
					onViewModeChange={setViewMode}
					counts={counts}
				/>

				{isLoading ? (
					<EventCardGrid events={[]} isLoading skeletonCount={events.length} />
				) : filtered.length === 0 ? (
					<p className="py-12 text-center text-sm text-gray-500">
						No events found.
					</p>
				) : viewMode === "grid" ? (
					<EventCardGrid events={filtered} />
				) : (
					<EventTable events={filtered} />
				)}
			</div>
		</CompanyLayout>
	);
}
