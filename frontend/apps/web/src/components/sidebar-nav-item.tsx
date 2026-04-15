import { Link } from "react-router-dom";
import type { Icon as PhosphorIcon } from "@phosphor-icons/react";

interface SidebarNavItemProps {
	label: string;
	icon: PhosphorIcon;
	to: string;
	isActive: boolean;
}

export default function SidebarNavItem({ label, icon: Icon, to, isActive }: SidebarNavItemProps) {
	return (
		<Link
			to={to}
			className={`flex items-center gap-2.5 rounded-md px-3 py-2 text-sm font-medium transition-colors ${
				isActive
					? "bg-blue-50 text-blue-700"
					: "text-gray-600 hover:bg-gray-100 hover:text-gray-900"
			}`}
		>
			<Icon size={18} weight={isActive ? "fill" : "regular"} />
			{label}
		</Link>
	);
}
