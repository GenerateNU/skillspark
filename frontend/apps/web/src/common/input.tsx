interface InputProps extends React.InputHTMLAttributes<HTMLInputElement> {
  error?: string;
}

export default function Input({ error, ...rest }: InputProps) {
  return (
    <input
      className={`w-full border rounded-md px-3 py-2 text-sm bg-white outline-none transition focus:ring-2 focus:ring-primary focus:border-primary placeholder:text-gray-400 ${error ? "border-red-400 bg-red-50 focus:ring-red-400" : "border-gray-300"}`}
      {...rest}
    />
  );
}
