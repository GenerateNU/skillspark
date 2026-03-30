export default function Divider({ label }: { label?: string }) {
  if (!label) return <div className="border-t border-gray-200 my-5" />;
  return (
    <div className="relative my-5">
      <div className="absolute inset-0 flex items-center"><div className="w-full border-t border-gray-200" /></div>
      <div className="relative flex justify-start"><span className="bg-white pr-3 text-xs font-semibold text-gray-500 uppercase tracking-wide">{label}</span></div>
    </div>
  );
}