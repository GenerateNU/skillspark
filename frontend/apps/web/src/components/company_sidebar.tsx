import { Link, useLocation } from "react-router-dom";
import {
	House,
	CalendarDots,
	CreditCard,
	Users,
} from "@phosphor-icons/react";
import { useCompany } from "@/company/company-context";

const navItems = [
	{ label: "Home", icon: House, path: "" },
	{ label: "Events", icon: CalendarDots, path: "/events" },
	{ label: "Payments", icon: CreditCard, path: "/payments" },
	{ label: "Customers", icon: Users, path: "/customers" },
] as const;

export function CompanySidebar() {
	const { companyId } = useCompany();
	const location = useLocation();
	const basePath = `/company/${companyId}`;

	return (
		<aside className="w-56 shrink-0 bg-white border-r border-gray-200 flex flex-col h-full">
			<Link
				to={basePath}
				className="px-5 py-4 border-b border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors"
			>
				<div className="w-7 h-7 rounded-md bg-blue-600 flex items-center justify-center shrink-0">
					<svg
						className="w-4 h-4 text-white"
						fill="currentColor"
						viewBox="0 0 20 20"
					>
						<path d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" />
					</svg>
				</div>
				<div>
					<p className="text-sm font-bold text-gray-900 leading-none">
						SkillSpark
					</p>
					<p className="text-xs text-gray-400 mt-0.5">Company</p>
				</div>
			</Link>

			<nav className="flex flex-col gap-0.5 px-2 py-3">
				{navItems.map(({ label, icon: Icon, path }) => {
					const fullPath = basePath + path;
					const isActive =
						path === ""
							? location.pathname === basePath ||
								location.pathname === basePath + "/"
							: location.pathname.startsWith(fullPath);

					return (
						<Link
							key={label}
							to={fullPath}
							className={`flex items-center gap-2.5 rounded-md px-3 py-2 text-sm font-medium transition-colors ${
								isActive
									? "bg-blue-50 text-blue-700"
									: "text-gray-600 hover:bg-gray-100 hover:text-gray-900"
							}`}
						>
							<Icon
								size={18}
								weight={isActive ? "fill" : "regular"}
							/>
							{label}
						</Link>
					);
				})}
			</nav>

			<div className="flex-1" />

			<Link
				to="/admin/profile"
				className="w-full px-4 py-3 border-t border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors"
			>
				<div className="w-7 h-7 rounded-full bg-blue-600 flex items-center justify-center text-xs font-bold text-white shrink-0">
					A
				</div>
				<div className="min-w-0 text-left">
					<p className="text-xs font-semibold text-gray-800 truncate">
						Admin User
					</p>
					<p className="text-xs text-gray-400 truncate">
						admin@skillspark.co
					</p>
				</div>
			</Link>
		</aside>
	);
}
