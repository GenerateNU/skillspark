import CompanyLayout from "./company-layout";
import { BentoGrid } from "@/components/ui/bento-grid";
import StatCard from "@/components/ui/stat-card";

export default function CompanyHome() {
	return (
		<CompanyLayout page="Home" showHeading>
			<BentoGrid>
				<StatCard label="Total Events" value="--" />
				<StatCard label="Customers" value="--" />
				<StatCard label="Revenue" value="--" />
			</BentoGrid>
		</CompanyLayout>
	);
}
