type ButtonVariant = "primary" | "danger" | "ghost";

interface BtnProps extends React.ButtonHTMLAttributes<HTMLButtonElement> { variant?: ButtonVariant; size?: "sm" | "md"; icon?: React.ReactNode; }

export function Btn({ children, variant = "primary", size = "md", icon, className = "", ...rest }: BtnProps) {
  const variants: Record<ButtonVariant, string> = {
    primary: "bg-blue-600 hover:bg-blue-700 active:bg-blue-800 text-white border border-blue-600",
    danger:  "bg-white hover:bg-red-50 text-red-600 border border-gray-300 hover:border-red-300",
    ghost:   "bg-white hover:bg-gray-50 text-gray-700 border border-gray-300",
  };
  const sizes: Record<string, string> = { sm: "px-3 py-1.5 text-xs gap-1.5", md: "px-3.5 py-2 text-sm gap-2" };
  return (
    <button className={`inline-flex items-center justify-center font-medium rounded-md transition-colors focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-1 disabled:opacity-50 disabled:cursor-not-allowed ${variants[variant]} ${sizes[size]} ${className}`} {...rest}>
      {icon && <span className="shrink-0">{icon}</span>}
      {children}
    </button>
  );
}


interface FieldProps { label: string; error?: string; required?: boolean; children: React.ReactNode; }
export function Field({ label, error, required, children }: FieldProps) {
  return (
    <div className="flex flex-col gap-1">
      <label className="text-sm font-medium text-gray-700">{label}{required && <span className="text-red-500 ml-0.5">*</span>}</label>
      {children}
      {error && <p className="text-xs text-red-600 mt-0.5">{error}</p>}
    </div>
  );
}

interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> { error?: string; }
export function Input({ error, ...rest }: InputProps) {
  return <input className={`w-full border rounded-md px-3 py-2 text-sm bg-white outline-none transition focus:ring-2 focus:ring-blue-500 focus:border-blue-500 placeholder:text-gray-400 ${error ? "border-red-400 bg-red-50 focus:ring-red-400" : "border-gray-300"}`} {...rest} />;
}

interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> { error?: string; children: React.ReactNode; }
export function Select({ error, children, ...rest }: SelectProps) {
  return (
    <select className={`w-full border rounded-md px-3 py-2 text-sm bg-white outline-none transition focus:ring-2 focus:ring-blue-500 ${error ? "border-red-400" : "border-gray-300"}`} {...rest}>
      {children}
    </select>
  );
}

export function Divider({ label }: { label?: string }) {
  if (!label) return <div className="border-t border-gray-200 my-5" />;
  return (
    <div className="relative my-5">
      <div className="absolute inset-0 flex items-center"><div className="w-full border-t border-gray-200" /></div>
      <div className="relative flex justify-start"><span className="bg-white pr-3 text-xs font-semibold text-gray-500 uppercase tracking-wide">{label}</span></div>
    </div>
  );
}