import { Plus } from "@phosphor-icons/react";
import { Button } from "@/components/ui/button";
import { ICON_SIZE_MD } from "./constants";

export default function CreateEventButton() {
	return (
		<Button variant="default" size="default" className="cursor-pointer gap-1.5 px-4 transition-all duration-200 hover:bg-white hover:text-blue-600 hover:ring-1 hover:ring-blue-600">
			<Plus size={ICON_SIZE_MD} weight="bold" />
			Create Event
		</Button>
	);
}
