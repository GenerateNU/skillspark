import { MapPin } from "@phosphor-icons/react";
import { ICON_SIZE_SM } from "./constants";

interface EventCardLocationProps {
	addressLine1: string;
	district: string;
}

export default function EventCardLocation({ addressLine1, district }: EventCardLocationProps) {
	return (
		<div className="flex items-center gap-1.5 text-xs text-gray-500">
			<MapPin size={ICON_SIZE_SM} className="shrink-0" />
			<span className="truncate">
				{addressLine1}, {district}
			</span>
		</div>
	);
}
