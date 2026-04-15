import CompanyLayout from "./company-layout";
import { useCompany } from "./company-context";
import { BentoGrid, BentoCard } from "@/components/ui/bento-grid";

export default function CompanyHome() {
	const { companyId } = useCompany();

	return (
		<CompanyLayout page="Home">
			<BentoGrid>
				<BentoCard>
					<p className="text-xs font-medium text-gray-400 uppercase tracking-wide">
						Total Events
					</p>
					<p className="mt-2 text-2xl font-bold text-gray-900">--</p>
				</BentoCard>
				<BentoCard>
					<p className="text-xs font-medium text-gray-400 uppercase tracking-wide">
						Customers
					</p>
					<p className="mt-2 text-2xl font-bold text-gray-900">--</p>
				</BentoCard>
				<BentoCard>
					<p className="text-xs font-medium text-gray-400 uppercase tracking-wide">
						Revenue
					</p>
					<p className="mt-2 text-2xl font-bold text-gray-900">--</p>
				</BentoCard>
			</BentoGrid>
		</CompanyLayout>
	);
}
