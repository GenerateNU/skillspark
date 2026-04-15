import type { EventOccurrence } from "@skillspark/api-client";
import EventCard from "./event-card";
import EventCardSkeleton from "./event-card-skeleton";

interface EventCardGridProps {
	events: EventOccurrence[];
	isLoading?: boolean;
	onEventClick?: (occurrence: EventOccurrence) => void;
}

export default function EventCardGrid({
	events,
	isLoading,
	onEventClick,
}: EventCardGridProps) {
	if (isLoading) {
		return (
			<div className="flex flex-wrap gap-4">
				{Array.from({ length: 8 }).map((_, i) => (
					<div
						key={i}
						className="w-full sm:w-[calc(50%-0.5rem)] lg:w-[calc(33.333%-0.75rem)] xl:w-[calc(25%-0.75rem)]"
					>
						<EventCardSkeleton />
					</div>
				))}
			</div>
		);
	}

	return (
		<div className="flex flex-wrap gap-4">
			{events.map((occ) => (
				<div
					key={occ.id}
					className="w-full sm:w-[calc(50%-0.5rem)] lg:w-[calc(33.333%-0.75rem)] xl:w-[calc(25%-0.75rem)]"
				>
					<EventCard
						occurrence={occ}
						onClick={() => onEventClick?.(occ)}
					/>
				</div>
			))}
		</div>
	);
}
