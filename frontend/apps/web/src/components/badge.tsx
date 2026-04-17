interface BadgeProps {
  children: React.ReactNode;
  color?: BadgeColor;
}

type BadgeColor = "blue" | "green" | "yellow" | "gray";

export default function Badge({ children, color = "gray" }: BadgeProps) {
  const styles: Record<BadgeColor, string> = {
    blue: "bg-primary/20 text-text ring-1 ring-primary/40",
    green: "bg-green-50 text-green-700 ring-1 ring-green-200",
    yellow: "bg-yellow-50 text-yellow-700 ring-1 ring-yellow-200",
    gray: "bg-gray-100 text-gray-600 ring-1 ring-gray-200",
  };
  return (
    <span
      className={`inline-flex items-center text-xs font-medium px-2 py-0.5 rounded ${styles[color]}`}
    >
      {children}
    </span>
  );
}
