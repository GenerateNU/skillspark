import { useState } from "react";
import {
	CalendarBlank,
	MapPin,
	ImageSquare,
} from "@phosphor-icons/react";
import type { EventOccurrence } from "@skillspark/api-client";
import Badge from "@/components/badge";
import { cn } from "@/lib/utils";
import { ICON_SIZE_SM, ICON_SIZE_LG } from "./constants";

export function formatEventDate(startTime: string): string {
	const date = new Date(startTime);
	return date.toLocaleDateString("en-US", {
		month: "short",
		day: "numeric",
		year: "numeric",
	}) +
		" — " +
		date.toLocaleTimeString("en-US", {
			hour: "numeric",
			minute: "2-digit",
			hour12: true,
		});
}

export function formatPrice(price: number, currency: string): string {
	if (price === 0) return "Free";
	const amount = price / 100;
	const code = currency.toUpperCase();
	return `${code} ${amount.toLocaleString()}`;
}

interface EventCardProps {
	occurrence: EventOccurrence;
	className?: string;
	onClick?: () => void;
}

export default function EventCard({
	occurrence,
	className,
	onClick,
}: EventCardProps) {
	const [imgError, setImgError] = useState(false);
	const { event, location, status } = occurrence;
	const enrollmentPct =
		occurrence.max_attendees > 0
			? (occurrence.curr_enrolled / occurrence.max_attendees) * 100
			: 0;

	const barColor = "bg-blue-300";

	const hasImage = event.presigned_url && !imgError;

	return (
		<div
			className={cn(
				"group flex h-full cursor-pointer flex-col overflow-hidden rounded-[4px] border border-gray-200 bg-white transition-shadow hover:shadow-md",
				className,
			)}
			onClick={onClick}
		>
			{/* Image area */}
			<div className="relative h-44 bg-gray-100">
				{hasImage ? (
					<img
						src={event.presigned_url}
						alt={event.title}
						className="h-full w-full object-cover"
						onError={() => setImgError(true)}
					/>
				) : (
					<div className="flex h-full w-full items-center justify-center text-gray-300">
						<ImageSquare size={ICON_SIZE_LG} />
					</div>
				)}

				{/* Category badges — top-left */}
				{event.category.length > 0 && (
					<div className="absolute top-2 left-2 flex gap-1">
						{event.category.slice(0, 2).map((cat) => (
							<Badge key={cat} color="blue">
								{cat.charAt(0).toUpperCase() + cat.slice(1)}
							</Badge>
						))}
					</div>
				)}

				{/* Status badge — top-right */}
				<div className="absolute top-2 right-2">
					{status === "scheduled" ? (
						<Badge color="green">
							<span className="mr-1 inline-block h-1.5 w-1.5 rounded-full bg-green-500" />
							Active
						</Badge>
					) : (
						<Badge color="gray">
							<span className="mr-1 inline-block h-1.5 w-1.5 rounded-full bg-gray-400" />
							Cancelled
						</Badge>
					)}
				</div>
			</div>

			{/* Content */}
			<div className="flex flex-1 flex-col gap-2 p-4">
				{/* Date */}
				<div className="flex items-center gap-1.5 text-xs text-gray-500">
					<CalendarBlank size={ICON_SIZE_SM} className="shrink-0" />
					{formatEventDate(occurrence.start_time)}
				</div>

				{/* Title */}
				<h3 className="line-clamp-2 text-sm font-bold leading-snug text-gray-900">
					{event.title}
				</h3>

				{/* Location */}
				<div className="flex items-center gap-1.5 text-xs text-gray-500">
					<MapPin size={ICON_SIZE_SM} className="shrink-0" />
					<span className="truncate">
						{location.address_line1}, {location.district}
					</span>
				</div>

				{/* Enrollment bar */}
				<div className="mt-auto pt-2">
					<div className="h-1.5 w-full overflow-hidden rounded-full bg-gray-100">
						<div
							className={cn("h-full rounded-full transition-all", barColor)}
							style={{ width: `${Math.min(enrollmentPct, 100)}%` }}
						/>
					</div>
					<p className="mt-1 text-xs text-gray-500">
						{Math.round(enrollmentPct)}% enrolled &middot;{" "}
						{occurrence.curr_enrolled}/{occurrence.max_attendees}
					</p>
				</div>

				{/* Price */}
				<p className="text-sm font-semibold text-gray-900">
					{formatPrice(occurrence.price, occurrence.currency)}
				</p>
			</div>
		</div>
	);
}
