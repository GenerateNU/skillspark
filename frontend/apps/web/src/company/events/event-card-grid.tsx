import type { EventOccurrence } from "@skillspark/api-client";
import EventCard from "./event-card";
import EventCardSkeleton from "./event-card-skeleton";
import { GRID_COL_CLASS } from "./constants";

interface EventCardGridProps {
	events: EventOccurrence[];
	isLoading?: boolean;
	skeletonCount?: number;
	onEventClick?: (occurrence: EventOccurrence) => void;
}

export default function EventCardGrid({
	events,
	isLoading,
	skeletonCount = 4,
	onEventClick,
}: EventCardGridProps) {
	if (isLoading) {
		return (
			<div className="flex flex-wrap gap-4">
				{Array.from({ length: skeletonCount }).map((_, i) => (
					<div key={i} className={GRID_COL_CLASS}>
						<EventCardSkeleton />
					</div>
				))}
			</div>
		);
	}

	return (
		<div className="flex flex-wrap gap-4">
			{events.map((occ) => (
				<div key={occ.id} className={GRID_COL_CLASS}>
					<EventCard
						occurrence={occ}
						onClick={() => onEventClick?.(occ)}
					/>
				</div>
			))}
		</div>
	);
}
