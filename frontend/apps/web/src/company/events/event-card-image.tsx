import { useState } from "react";
import { ImageSquare } from "@phosphor-icons/react";
import type { EventOccurrenceStatus } from "@skillspark/api-client";
import { ICON_SIZE_LG } from "./constants";
import CategoryBadges from "./category-badges";
import StatusBadge from "./status-badge";

interface EventCardImageProps {
	presignedUrl: string;
	alt: string;
	categories: string[];
	status: EventOccurrenceStatus;
}

export default function EventCardImage({
	presignedUrl,
	alt,
	categories,
	status,
}: EventCardImageProps) {
	const [imgError, setImgError] = useState(false);
	const hasImage = presignedUrl && !imgError;

	return (
		<div className="relative h-44 bg-gray-100">
			{hasImage ? (
				<img
					src={presignedUrl}
					alt={alt}
					className="h-full w-full object-cover"
					onError={() => setImgError(true)}
				/>
			) : (
				<div className="flex h-full w-full items-center justify-center text-gray-300">
					<ImageSquare size={ICON_SIZE_LG} />
				</div>
			)}
			<CategoryBadges categories={categories} />
			<div className="absolute top-2 right-2">
				<StatusBadge status={status} />
			</div>
		</div>
	);
}
