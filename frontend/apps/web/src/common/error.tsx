import { IconX } from "../components/icons";

interface ErrorModalProps {
  error: string;
  setError: (e: string | null) => void;
}

export default function ErrorModal({ error, setError }: ErrorModalProps) {
  return (
    <div className="mb-4 flex items-start gap-3 px-4 py-3 rounded-lg bg-red-50 border border-red-200">
      <span className="text-red-500 shrink-0 mt-0.5">⚠</span>
      <p className="text-sm text-red-700 flex-1">{error}</p>
      <button
        onClick={function () {
          setError(null);
        }}
        className="text-red-400 hover:text-red-600 cursor-pointer shrink-0"
      >
        <IconX />
      </button>
    </div>
  );
}
