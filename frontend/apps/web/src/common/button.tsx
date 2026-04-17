type ButtonVariant = "primary" | "danger" | "ghost";

interface BtnProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: ButtonVariant;
  size?: "sm" | "md";
  icon?: React.ReactNode;
}

export default function Btn({
  children,
  variant = "primary",
  size = "md",
  icon,
  className = "",
  ...rest
}: BtnProps) {
  const variants: Record<ButtonVariant, string> = {
    primary:
      "bg-primary hover:bg-primary/80 active:bg-primary/70 text-text border border-primary",
    danger:
      "bg-white hover:bg-red-50 text-red-600 border border-gray-300 hover:border-red-300",
    ghost: "bg-white hover:bg-gray-50 text-text border border-gray-300",
  };
  const sizes: Record<string, string> = {
    sm: "px-3 py-1.5 text-xs gap-1.5",
    md: "px-3.5 py-2 text-sm gap-2",
  };
  return (
    <button
      className={`inline-flex items-center justify-center font-medium rounded-md transition-colors focus:outline-none focus:ring-2 focus:ring-primary focus:ring-offset-1 disabled:opacity-50 disabled:cursor-not-allowed ${variants[variant]} ${sizes[size]} ${className}`}
      {...rest}
    >
      {icon && <span className="shrink-0">{icon}</span>}
      {children}
    </button>
  );
}
