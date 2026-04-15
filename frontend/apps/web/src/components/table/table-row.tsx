interface TableRowProps {
	children: React.ReactNode;
	onClick?: () => void;
}

export default function TableRow({ children, onClick }: TableRowProps) {
	return (
		<tr
			className="cursor-pointer border-b border-gray-100 transition-colors last:border-b-0 hover:bg-gray-50"
			onClick={onClick}
		>
			{children}
		</tr>
	);
}
