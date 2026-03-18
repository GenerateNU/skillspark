import { useState } from "react";

export default function OtpCard({ password, name }: { password: string; name: string }) {
  const [copied, setCopied] = useState<boolean>(false);
  function copy(): void {
    if (navigator.clipboard) navigator.clipboard.writeText(password);
    setCopied(true);
    setTimeout(function () { setCopied(false); }, 2000);
  }
  return (
    <div className="mt-3 rounded-md border border-yellow-200 bg-yellow-50 p-3">
      <p className="text-xs font-semibold text-yellow-800 mb-1">One-Time Password — share once</p>
      <div className="flex items-center gap-3">
        <code className="font-mono text-sm font-bold text-yellow-900 tracking-widest">{password}</code>
        <button onClick={copy} className="text-xs text-yellow-700 hover:text-yellow-900 font-medium border border-yellow-300 px-2 py-0.5 rounded hover:bg-yellow-100 transition-colors">
          {copied ? "✓ Copied" : "Copy"}
        </button>
      </div>
      <p className="text-xs text-yellow-700 mt-1">{name || "This user"} must change it on first login.</p>
    </div>
  );
}
