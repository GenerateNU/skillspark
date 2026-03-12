import type { Organization } from "@skillspark/api-client";
import { Link } from "react-router-dom";

interface SidebarProps {
  organizations: Organization[];
  activeId: string | null;
  onSelect: (id: string) => void;
}

export function Sidebar({ organizations, activeId, onSelect }: SidebarProps) {
  return (
    <aside className="w-1/8 shrink-0 bg-white border-r border-gray-200 flex flex-col h-full">
      <Link
        to="/"
        className="px-5 py-4 border-b border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors"
      >
        <div className="w-7 h-7 rounded-md bg-blue-600 flex items-center justify-center shrink-0">
          <svg className="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" />
          </svg>
        </div>
        <div>
          <p className="text-base font-bold text-gray-900 leading-none">SkillSpark</p>
          <p className="text-base text-gray-400 mt-0.5">Admin Console</p>
        </div>
      </Link>

      <nav className="flex-1 px-3 py-4 flex flex-col gap-0.5 overflow-y-auto">
        <p className="text-base font-semibold text-gray-400 uppercase tracking-wider px-2 mb-2">Organizations</p>
        {organizations.length === 0 && (
          <p className="text-xs text-gray-400 px-2 py-1">No organizations yet</p>
        )}
        {organizations.map(function (org: Organization) {
          const active: boolean = org.id === activeId;
          const initial: string = org.name ? org.name.charAt(0).toUpperCase() : "?";
          return (
            <button
              key={org.id}
              onClick={function () { onSelect(org.id); }}
              className={`w-full flex items-center gap-2.5 px-2.5 py-2 rounded-md text-sm font-medium transition-colors text-left ${active ? "bg-blue-50 text-blue-700" : "text-gray-600 hover:bg-gray-100 hover:text-gray-900"}`}
            >
              <div className={`w-6 h-6 rounded-md flex items-center justify-center text-xs font-bold shrink-0 ${active ? "bg-blue-600 text-white" : "bg-gray-200 text-gray-600"}`}>
                {initial}
              </div>
              <span className="truncate">{org.name}</span>
              {!org.active && (
                <span className="ml-auto text-xs text-gray-400 font-normal shrink-0">Inactive</span>
              )}
            </button>
          );
        })}
      </nav>

      <Link
  to="/profile"
  className="w-full px-4 py-3 border-t border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors"
>
  <div className="w-7 h-7 rounded-full bg-blue-600 flex items-center justify-center text-xs font-bold text-white shrink-0">A</div>
  <div className="min-w-0 text-left">
    <p className="text-base font-semibold text-gray-800 truncate">Admin User</p>
    <p className="text-base text-gray-400 truncate">admin@skillspark.co</p>
  </div>
</Link>
    </aside>
  );
}