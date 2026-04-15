import type { EventOccurrence } from "@skillspark/api-client";
import { TableContainer, TableHeader } from "@/components/table";
import EventTableRow from "./event-table-row";

const COLUMNS = ["Status", "Event", "Date", "Location", "Enrolled", "Price"];

interface EventTableProps {
	events: EventOccurrence[];
	onEventClick?: (occurrence: EventOccurrence) => void;
}

export default function EventTable({ events, onEventClick }: EventTableProps) {
	return (
		<TableContainer>
			<TableHeader columns={COLUMNS} />
			<tbody>
				{events.map((occ) => (
					<EventTableRow
						key={occ.id}
						occurrence={occ}
						onClick={() => onEventClick?.(occ)}
					/>
				))}
			</tbody>
		</TableContainer>
	);
}
