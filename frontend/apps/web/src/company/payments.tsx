import CompanyLayout from "./company-layout";
import { BentoGrid, BentoCard } from "@/components/ui/bento-grid";

export default function CompanyPayments() {
	return (
		<CompanyLayout page="Payments">
			<BentoGrid cols={1}>
				<BentoCard>
					<p className="text-sm text-gray-500">No payments yet.</p>
				</BentoCard>
			</BentoGrid>
		</CompanyLayout>
	);
}
