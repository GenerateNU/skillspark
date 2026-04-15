import type { EventOccurrenceStatus } from "@skillspark/api-client";
import Badge from "@/components/badge";

export default function StatusBadge({ status }: { status: EventOccurrenceStatus }) {
	if (status === "scheduled") {
		return (
			<Badge color="green">
				<span className="mr-1 inline-block h-1.5 w-1.5 rounded-full bg-green-500" />
				Active
			</Badge>
		);
	}

	return (
		<Badge color="gray">
			<span className="mr-1 inline-block h-1.5 w-1.5 rounded-full bg-gray-400" />
			Cancelled
		</Badge>
	);
}
