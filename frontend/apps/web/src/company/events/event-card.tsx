import type { EventOccurrence } from "@skillspark/api-client";
import { cn } from "@/lib/utils";
import { formatPrice } from "./formatters";
import EventCardImage from "./event-card-image";
import EventCardDate from "./event-card-date";
import EventCardLocation from "./event-card-location";
import EnrollmentBar from "./enrollment-bar";

interface EventCardProps {
	occurrence: EventOccurrence;
	className?: string;
	onClick?: () => void;
}

export default function EventCard({ occurrence, className, onClick }: EventCardProps) {
	const { event, location } = occurrence;

	return (
		<div
			className={cn(
				"group flex h-full cursor-pointer flex-col overflow-hidden rounded-[4px] border border-gray-200 bg-white transition-shadow hover:shadow-md",
				className,
			)}
			onClick={onClick}
		>
			<EventCardImage
				presignedUrl={event.presigned_url}
				alt={event.title}
				categories={event.category}
				status={occurrence.status}
			/>
			<div className="flex flex-1 flex-col gap-2 p-4">
				<EventCardDate startTime={occurrence.start_time} />
				<h3 className="line-clamp-2 text-sm font-bold leading-snug text-gray-900">
					{event.title}
				</h3>
				<EventCardLocation
					addressLine1={location.address_line1}
					district={location.district}
				/>
				<EnrollmentBar
					enrolled={occurrence.curr_enrolled}
					max={occurrence.max_attendees}
				/>
				<p className="text-sm font-semibold text-gray-900">
					{formatPrice(occurrence.price, occurrence.currency)}
				</p>
			</div>
		</div>
	);
}
