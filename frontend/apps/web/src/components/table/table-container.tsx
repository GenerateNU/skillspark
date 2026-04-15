export default function TableContainer({ children }: { children: React.ReactNode }) {
	return (
		<div className="overflow-hidden rounded-[4px] border border-gray-200">
			<table className="w-full text-left text-sm">{children}</table>
		</div>
	);
}
