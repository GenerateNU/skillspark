import { BentoCard } from "@/components/ui/bento-grid";

interface StatCardProps {
	label: string;
	value: string;
}

export default function StatCard({ label, value }: StatCardProps) {
	return (
		<BentoCard>
			<p className="text-xs font-medium uppercase tracking-wide text-gray-400">
				{label}
			</p>
			<p className="mt-2 text-2xl font-bold text-gray-900">{value}</p>
		</BentoCard>
	);
}
