import { Routes, Route } from "react-router-dom";
import { Sidebar } from "../components/admin_sidebar";
import OrganizationDetailPage from "./organizations";
import HomePage from "./home";
import { ProfilePage } from "./profile";

export default function Admin() {
  return (
    <div className="flex h-screen w-screen overflow-hidden bg-gray-50 font-sans">
      <Sidebar />
      <Routes>
        <Route path="" element={<HomePage />} />
        <Route path="/organization" element={<OrganizationDetailPage />} />
        <Route path="/profile" element={<ProfilePage />} />
      </Routes>
    </div>
  );
}
