import { cn } from "@/lib/utils";

interface TableCellProps {
	children: React.ReactNode;
	className?: string;
}

export default function TableCell({ children, className }: TableCellProps) {
	return <td className={cn("px-4 py-3", className)}>{children}</td>;
}
