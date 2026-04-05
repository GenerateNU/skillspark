import type { Manager } from "@skillspark/api-client";

interface OrgManagerCardProps {
  manager: Manager | undefined;
  loadingMgr: boolean;
}

export default function OrgManagerCard({
  manager,
  loadingMgr,
}: OrgManagerCardProps) {
  return (
    <div className="bg-white rounded-lg border border-gray-200">
      <div className="px-5 py-4 border-b border-gray-100">
        <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">
          Manager
        </h3>
      </div>
      {loadingMgr ? (
        <p className="px-5 py-4 text-base text-gray-400">Loading manager…</p>
      ) : !manager ? (
        <p className="px-5 py-4 text-base text-gray-400">
          No manager assigned.
        </p>
      ) : (
        <div className="px-5 py-4 flex items-center gap-4">
          <div className="w-12 h-12 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center text-base font-bold shrink-0">
            {manager.name ? manager.name.charAt(0).toUpperCase() : "?"}
          </div>
          <div className="min-w-0">
            <p className="text-base font-semibold text-gray-900">
              {manager.name}
            </p>
            <p className="text-sm text-gray-500">{manager.email}</p>
            <p className="text-sm text-gray-400">
              {manager.username} · {manager.role}
            </p>
          </div>
        </div>
      )}
    </div>
  );
}
