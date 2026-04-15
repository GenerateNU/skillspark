import type { EventOccurrence } from "@skillspark/api-client";
import StatusBadge from "./status-badge";
import { formatEventDate, formatPrice } from "./formatters";

interface EventTableProps {
	events: EventOccurrence[];
	onEventClick?: (occurrence: EventOccurrence) => void;
}

export default function EventTable({ events, onEventClick }: EventTableProps) {
	return (
		<div className="overflow-hidden rounded-[4px] border border-gray-200">
			<table className="w-full text-left text-sm">
				<thead>
					<tr className="border-b border-gray-200 bg-gray-50 text-xs font-medium text-gray-500">
						<th className="px-4 py-3">Status</th>
						<th className="px-4 py-3">Event</th>
						<th className="px-4 py-3">Date</th>
						<th className="px-4 py-3">Location</th>
						<th className="px-4 py-3">Enrolled</th>
						<th className="px-4 py-3">Price</th>
					</tr>
				</thead>
				<tbody>
					{events.map((occ) => (
						<tr
							key={occ.id}
							className="cursor-pointer border-b border-gray-100 transition-colors last:border-b-0 hover:bg-gray-50"
							onClick={() => onEventClick?.(occ)}
						>
							<td className="px-4 py-3">
								<StatusBadge status={occ.status} />
							</td>
							<td className="px-4 py-3">
								<div className="flex items-center gap-3">
									<img
										src={occ.event.presigned_url || undefined}
										alt=""
										className="h-8 w-12 rounded-[4px] bg-gray-100 object-cover"
										onError={(e) => {
											(e.target as HTMLImageElement).style.display = "none";
										}}
									/>
									<span className="font-medium text-gray-900">
										{occ.event.title}
									</span>
								</div>
							</td>
							<td className="px-4 py-3 text-gray-600">
								{formatEventDate(occ.start_time)}
							</td>
							<td className="px-4 py-3 text-gray-600">
								{occ.location.address_line1}, {occ.location.district}
							</td>
							<td className="px-4 py-3 text-gray-600">
								{occ.curr_enrolled}/{occ.max_attendees}
							</td>
							<td className="px-4 py-3 font-medium text-gray-900">
								{formatPrice(occ.price, occ.currency)}
							</td>
						</tr>
					))}
				</tbody>
			</table>
		</div>
	);
}
