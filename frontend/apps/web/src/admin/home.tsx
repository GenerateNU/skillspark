import { Link } from "react-router-dom";
import { IconBuilding } from "../components/icons";

export default function HomePage() {
  return (
    <div className="flex-1 flex flex-col p-8 bg-gray-50">
      <div className="max-w-2xl">
        <h1 className="text-2xl font-bold text-gray-900 mb-1">Welcome back</h1>
        <p className="text-gray-500 text-sm mb-8">Manage your platform from the admin console.</p>
        <div className="grid grid-cols-2 gap-4">
          <Link
            to="/organizations"
            className="group text-left bg-white border border-gray-200 rounded-lg p-5 hover:border-blue-300 hover:shadow-sm transition-all"
          >
            <div className="w-9 h-9 rounded-md bg-blue-50 text-blue-600 flex items-center justify-center mb-3">
              <IconBuilding />
            </div>
            <p className="text-sm font-semibold text-gray-900 group-hover:text-blue-700">Organizations</p>
            <p className="text-xs text-gray-500 mt-1">Register and manage organizations and their managers.</p>
          </Link>
        </div>
      </div>
    </div>
  );
}