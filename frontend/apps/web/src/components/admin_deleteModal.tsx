interface DeleteModalProps {
  title: string;
  description: React.ReactNode;
  deleting: boolean;
  onConfirm: () => void;
  onClose: () => void;
}

export default function DeleteModal({
  title,
  description,
  deleting,
  onConfirm,
  onClose,
}: DeleteModalProps) {
  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/45">
      <div className="bg-white rounded-xl shadow-2xl w-full max-w-sm">
        <div className="px-6 py-5 border-b border-gray-200">
          <h2 className="text-lg font-semibold text-gray-900">{title}</h2>
        </div>
        <div className="px-6 py-5">
          <p className="text-base text-gray-600">{description}</p>
        </div>
        <div className="px-6 py-4 border-t border-gray-200 bg-gray-50 rounded-b-xl flex items-center justify-end gap-3">
          <button
            onClick={onClose}
            disabled={deleting}
            className="px-3.5 py-2 text-sm font-medium rounded-md bg-white border border-gray-300 text-gray-700 hover:bg-gray-50 transition-colors disabled:opacity-50 cursor-pointer"
          >
            Cancel
          </button>
          <button
            onClick={onConfirm}
            disabled={deleting}
            className="px-3.5 py-2 text-sm font-medium rounded-md bg-red-600 hover:bg-red-700 text-white transition-colors disabled:opacity-50 cursor-pointer"
          >
            {deleting ? "Deleting…" : "Delete"}
          </button>
        </div>
      </div>
    </div>
  );
}
