import CompanyLayout from "./company-layout";
import { BentoGrid, BentoCard } from "@/components/ui/bento-grid";

export default function CompanyCustomers() {
	return (
		<CompanyLayout page="Customers">
			<BentoGrid cols={1}>
				<BentoCard>
					<p className="text-sm text-gray-500">No customers yet.</p>
				</BentoCard>
			</BentoGrid>
		</CompanyLayout>
	);
}
