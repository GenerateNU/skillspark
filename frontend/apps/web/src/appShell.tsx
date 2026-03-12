import type { Organization } from "@skillspark/api-client";
import { useState } from "react";
import { useNavigate, Routes, Route } from "react-router-dom";
import { Sidebar } from "./components/sidebar";
import { OrganizationsPage } from "./pages/businessesPage";
import { ProfilePage } from "./pages/profliePage";

export function AppShell() {
  const navigate = useNavigate();
  const [organizations, setOrganizations] = useState<Organization[]>([]);
  const [activeOrgId, setActiveOrgId] = useState<string | null>(null);

  return (
    <div className="flex h-screen w-screen overflow-hidden bg-gray-50 font-sans">
      <Sidebar
        organizations={organizations}
        activeId={activeOrgId}
        onSelect={function (id: string) { setActiveOrgId(id); navigate("/"); }}
        onProfileClick={function () { navigate("/profile"); }}
      />
      <Routes>
        <Route
          path="/"
          element={
            <OrganizationsPage
              organizations={organizations}
              activeOrgId={activeOrgId}
              onOrganizationsChange={setOrganizations}
              onActiveOrgChange={setActiveOrgId}
            />
          }
        />
        <Route path="/profile" element={<ProfilePage />} />
      </Routes>
    </div>
  );
}