export function formatEventDate(startTime: string): string {
	const date = new Date(startTime);
	return (
		date.toLocaleDateString("en-US", {
			month: "short",
			day: "numeric",
			year: "numeric",
		}) +
		" — " +
		date.toLocaleTimeString("en-US", {
			hour: "numeric",
			minute: "2-digit",
			hour12: true,
		})
	);
}

export function formatPrice(price: number, currency: string): string {
	if (price === 0) return "Free";
	const amount = price / 100;
	const code = currency.toUpperCase();
	return `${code} ${amount.toLocaleString()}`;
}
