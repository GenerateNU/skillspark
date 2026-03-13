import type { Organization } from "@skillspark/api-client";
import { useState } from "react";
import { Routes, Route } from "react-router-dom";
import { Sidebar } from "./components/sidebar";
import { OrganizationsPage } from "./pages/organizationsPage";
import { ProfilePage } from "./pages/profilePage";

export default function App() {
  const [organization, setOrganization] = useState<Organization | null>(null);

  return (
    <div className="flex h-screen w-screen overflow-hidden bg-gray-50 font-sans">
      <Sidebar />
      <Routes>
        <Route
          path="/"
          element={
            <OrganizationsPage
              organization={organization}
              onOrganizationChange={setOrganization}
            />
          }
        />
        <Route path="/profile" element={<ProfilePage />} />
      </Routes>
    </div>
  );
}