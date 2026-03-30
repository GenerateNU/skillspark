import { useState } from "react";

export default function ValidatedInput({
  value,
  onChange,
  validate,
  placeholder,
  type = "text",
}: {
  value: string;
  onChange: (v: string) => void;
  validate: (v: string) => string | null;
  placeholder?: string;
  type?: string;
}) {
  const [error, setError] = useState<string | null>(null);
  const [touched, setTouched] = useState<boolean>(false);

  function handleBlur(): void {
    setTouched(true);
    setError(validate(value));
  }

  function handleChange(e: React.ChangeEvent<HTMLInputElement>): void {
    onChange(e.target.value);
    if (touched) {
      setError(validate(e.target.value));
    }
  }

  return (
    <div className="flex flex-col gap-1">
      <input
        type={type}
        value={value}
        onChange={handleChange}
        onBlur={handleBlur}
        placeholder={placeholder}
        className={`w-full border rounded-md px-3 py-2 text-sm bg-white outline-none transition focus:ring-2 focus:ring-blue-500 focus:border-blue-500 placeholder:text-gray-400 ${error ? "border-red-400 bg-red-50" : "border-gray-300"}`}
      />
      {error && <p className="text-xs text-red-600">{error}</p>}
    </div>
  );
}