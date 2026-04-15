import type { EventOccurrence } from "@skillspark/api-client";
import { TableRow, TableCell } from "@/components/table";
import StatusBadge from "./status-badge";
import { formatEventDate, formatPrice } from "./formatters";

interface EventTableRowProps {
	occurrence: EventOccurrence;
	onClick?: () => void;
}

export default function EventTableRow({ occurrence, onClick }: EventTableRowProps) {
	return (
		<TableRow onClick={onClick}>
			<TableCell>
				<StatusBadge status={occurrence.status} />
			</TableCell>
			<TableCell>
				<div className="flex items-center gap-3">
					<img
						src={occurrence.event.presigned_url || undefined}
						alt=""
						className="h-8 w-12 rounded-[4px] bg-gray-100 object-cover"
						onError={(e) => {
							(e.target as HTMLImageElement).style.display = "none";
						}}
					/>
					<span className="font-medium text-gray-900">
						{occurrence.event.title}
					</span>
				</div>
			</TableCell>
			<TableCell className="text-gray-600">
				{formatEventDate(occurrence.start_time)}
			</TableCell>
			<TableCell className="text-gray-600">
				{occurrence.location.address_line1}, {occurrence.location.district}
			</TableCell>
			<TableCell className="text-gray-600">
				{occurrence.curr_enrolled}/{occurrence.max_attendees}
			</TableCell>
			<TableCell className="font-medium text-gray-900">
				{formatPrice(occurrence.price, occurrence.currency)}
			</TableCell>
		</TableRow>
	);
}
