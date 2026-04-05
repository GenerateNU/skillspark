interface SelectProps extends React.SelectHTMLAttributes<HTMLSelectElement> {
  error?: string;
  children: React.ReactNode;
}
export default function Select({ error, children, ...rest }: SelectProps) {
  return (
    <select
      className={`w-full border rounded-md px-3 py-2 text-sm bg-white outline-none transition focus:ring-2 focus:ring-blue-500 ${error ? "border-red-400" : "border-gray-300"}`}
      {...rest}
    >
      {children}
    </select>
  );
}
