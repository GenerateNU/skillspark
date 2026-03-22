import { Routes, Route } from "react-router-dom";
import { Sidebar } from "../components/admin_sidebar";
import { OrganizationsPage } from "./organizations";
import HomePage from "./home";

export default function Admin() {
  return (
    <div className="flex h-screen w-screen overflow-hidden bg-gray-50 font-sans">
      <Sidebar />
      <Routes>
        <Route path="" element={<HomePage />} />
        <Route path="/organization" element={<OrganizationsPage />} />
      </Routes>
    </div>
  );
}