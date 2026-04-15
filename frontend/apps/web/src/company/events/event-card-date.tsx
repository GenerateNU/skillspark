import { CalendarBlank } from "@phosphor-icons/react";
import { ICON_SIZE_SM } from "./constants";
import { formatEventDate } from "./formatters";

export default function EventCardDate({ startTime }: { startTime: string }) {
	return (
		<div className="flex items-center gap-1.5 text-xs text-gray-500">
			<CalendarBlank size={ICON_SIZE_SM} className="shrink-0" />
			{formatEventDate(startTime)}
		</div>
	);
}
