import { useEffect } from "react";
import { IconX } from "./icons";

interface DrawerProps {
  title: string;
  subtitle?: string;
  onClose: () => void;
  children: React.ReactNode;
  footer?: React.ReactNode;
  width?: string;
}
export function Drawer({
  title,
  subtitle,
  onClose,
  children,
  footer,
  width = "max-w-xl",
}: DrawerProps) {
  useEffect(
    function () {
      function handler(e: KeyboardEvent) {
        if (e.key === "Escape") onClose();
      }
      window.addEventListener("keydown", handler);
      return function () {
        window.removeEventListener("keydown", handler);
      };
    },
    [onClose],
  );
  return (
    <div className="fixed inset-0 z-50 flex">
      <div
        className="flex-1 bg-black/30"
        style={{ animation: "fadeIn 0.2s ease-out" }}
        onClick={onClose}
      />
      <div
        className={`relative flex flex-col w-full ${width} bg-white shadow-2xl h-full`}
        style={{ animation: "slideIn 0.2s ease-out" }}
      >
        <div className="flex items-start justify-between px-6 py-5 border-b border-gray-200 shrink-0">
          <div>
            <h2 className="text-base font-semibold text-gray-900">{title}</h2>
            {subtitle && (
              <p className="text-sm text-gray-500 mt-0.5">{subtitle}</p>
            )}
          </div>
          <button
            onClick={onClose}
            className="ml-4 p-1.5 rounded-md text-gray-400 hover:text-gray-600 hover:bg-gray-100"
          >
            <IconX />
          </button>
        </div>
        <div className="flex-1 overflow-y-auto px-6 py-6">{children}</div>
        {footer && (
          <div className="shrink-0 px-6 py-4 border-t border-gray-200 bg-gray-50 flex items-center justify-end gap-3">
            {footer}
          </div>
        )}
      </div>
      <style>{`
        @keyframes slideIn { from { transform: translateX(100%); } to { transform: translateX(0); } }
        @keyframes fadeIn  { from { opacity: 0; } to { opacity: 1; } }
      `}</style>
    </div>
  );
}
