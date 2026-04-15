import Badge from "@/components/badge";

export default function CategoryBadges({ categories }: { categories: string[] }) {
	if (categories.length === 0) return null;

	return (
		<div className="absolute top-2 left-2 flex gap-1">
			{categories.slice(0, 2).map((cat) => (
				<Badge key={cat} color="blue">
					{cat.charAt(0).toUpperCase() + cat.slice(1)}
				</Badge>
			))}
		</div>
	);
}
