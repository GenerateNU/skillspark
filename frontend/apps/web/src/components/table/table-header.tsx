interface TableHeaderProps {
	columns: string[];
}

export default function TableHeader({ columns }: TableHeaderProps) {
	return (
		<thead>
			<tr className="border-b border-gray-200 bg-gray-50 text-xs font-medium text-gray-500">
				{columns.map((col) => (
					<th key={col} className="px-4 py-3">
						{col}
					</th>
				))}
			</tr>
		</thead>
	);
}
