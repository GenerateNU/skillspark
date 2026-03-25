import { Link } from "react-router-dom";


export function Sidebar() {

  return (
    <aside className="w-56 shrink-0 bg-white border-r border-gray-200 flex flex-col h-full">
      <Link to="/admin" className="px-5 py-4 border-b border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors">
        <div className="w-7 h-7 rounded-md bg-blue-600 flex items-center justify-center shrink-0">
          <svg className="w-4 h-4 text-white" fill="currentColor" viewBox="0 0 20 20">
            <path d="M11.3 1.046A1 1 0 0112 2v5h4a1 1 0 01.82 1.573l-7 10A1 1 0 018 18v-5H4a1 1 0 01-.82-1.573l7-10a1 1 0 011.12-.38z" />
          </svg>
        </div>
        <div>
          <p className="text-sm font-bold text-gray-900 leading-none">SkillSpark</p>
          <p className="text-xs text-gray-400 mt-0.5">Admin Console</p>
        </div>
      </Link>

      <div className="flex-1" />

      <Link
        to="/admin/profile"
        className="w-full px-4 py-3 border-t border-gray-200 flex items-center gap-2.5 hover:bg-gray-50 transition-colors"
      >
        <div className="w-7 h-7 rounded-full bg-blue-600 flex items-center justify-center text-xs font-bold text-white shrink-0">A</div>
        <div className="min-w-0 text-left">
          <p className="text-xs font-semibold text-gray-800 truncate">Admin User</p>
          <p className="text-xs text-gray-400 truncate">admin@skillspark.co</p>
        </div>
      </Link>
    </aside>
  );
}