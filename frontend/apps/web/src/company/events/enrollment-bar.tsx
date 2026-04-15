interface EnrollmentBarProps {
	enrolled: number;
	max: number;
}

export default function EnrollmentBar({ enrolled, max }: EnrollmentBarProps) {
	const pct = max > 0 ? (enrolled / max) * 100 : 0;

	return (
		<div className="mt-auto pt-2">
			<div className="h-1.5 w-full overflow-hidden rounded-full bg-gray-100">
				<div
					className="h-full rounded-full bg-blue-300 transition-all"
					style={{ width: `${Math.min(pct, 100)}%` }}
				/>
			</div>
			<p className="mt-1 text-xs text-gray-500">
				{Math.round(pct)}% enrolled &middot; {enrolled}/{max}
			</p>
		</div>
	);
}
