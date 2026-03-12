import type { Manager, Organization } from "@skillspark/api-client";
import { useCallback, useState } from "react";
import { Btn, Field, Input } from "../components/common";
import { CreateDrawer } from "../components/createDrawer";
import { IconPlus } from "../components/icons";
import { blankMgr, type ManagerErrors, fmtDate } from "../components/types";
import { validateMgr } from "../components/validation";

interface OrganizationsPageProps {
  organizations: Organization[];
  activeOrgId: string | null;
  onOrganizationsChange: (orgs: Organization[]) => void;
  onActiveOrgChange: (id: string) => void;
}

export function OrganizationsPage({ organizations, activeOrgId, onOrganizationsChange, onActiveOrgChange }: OrganizationsPageProps) {
  const [showCreate, setShowCreate] = useState<boolean>(false);
  const [addingMgr, setAddingMgr] = useState<boolean>(false);
  const [mgr, setMgr] = useState<Manager>(blankMgr());
  const [mgrErrors, setMgrErrors] = useState<ManagerErrors>({});

  const [managers, setManagers] = useState<Manager[]>([]);

  const handleRegister = useCallback(function (org: Organization): void {
    onOrganizationsChange([org, ...organizations]);
    onActiveOrgChange(org.id);
  }, [organizations, onOrganizationsChange, onActiveOrgChange]);

  const handleAddManager = useCallback(function (orgId: string, newMgr: Manager): void {
    setManagers(function (prev: Manager[]) {
      return [...prev, Object.assign({}, newMgr, { organization_id: orgId })];
    });
  }, []);

  const handleRemoveManager = useCallback(function (orgId: string, mgrId: string): void {
    setManagers(function (prev: Manager[]) {
      return prev.filter(function (m: Manager) { return !(m.id === mgrId && m.organization_id === orgId); });
    });
  }, []);

  const allEmails: Set<string> = new Set(
    managers.map(function (m: Manager) { return m.email.toLowerCase(); })
  );

  const activeOrg: Organization | null = organizations.find(function (o: Organization) { return o.id === activeOrgId; }) || null;

  const orgManagers: Manager[] = managers.filter(function (m: Manager) { return m.organization_id === activeOrgId; });

  function submitMgr(): void {
    if (!activeOrg) return;
    const e: ManagerErrors = validateMgr(mgr);
    const existingEmails: Set<string> = new Set(orgManagers.map(function (m: Manager) { return m.email.toLowerCase(); }));
    if (mgr.email && existingEmails.has(mgr.email.trim().toLowerCase())) e.email = "Email already in use";
    if (Object.keys(e).length) { setMgrErrors(e); return; }
    handleAddManager(activeOrg.id, mgr);
    setMgr(blankMgr());
    setAddingMgr(false);
    setMgrErrors({});
  }

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center gap-4 shrink-0">
        <div>
          <h1 className="text-base font-semibold text-gray-900">
            {activeOrg ? activeOrg.name : "Organizations"}
          </h1>
          <p className="text-xs text-gray-500">
            {organizations.length} {organizations.length === 1 ? "organization" : "organizations"} registered
          </p>
        </div>
        <div className="ml-auto">
          <Btn onClick={function () { setShowCreate(true); }} icon={<IconPlus />}>
            Add organization
          </Btn>
        </div>
      </div>

      <div className="flex-1 overflow-auto bg-gray-50 p-6">
        {!activeOrg ? (
          <div className="flex flex-col items-center justify-center h-full text-center">
            <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mb-3">
              <IconPlus />
            </div>
            <p className="text-sm font-semibold text-gray-700">No organization selected</p>
            <p className="text-sm text-gray-400 mt-1">
              {organizations.length === 0
                ? "Click \"Add organization\" to register the first one."
                : "Select an organization from the sidebar."}
            </p>
          </div>
        ) : (
          <div className="max-w-2xl flex flex-col gap-6">

            {/* Details card */}
            <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
              <div className="px-5 py-4 flex items-center justify-between">
                <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">Details</h2>
                <span className={`inline-flex items-center text-xs font-medium px-2 py-0.5 rounded ring-1 ${activeOrg.active ? "bg-green-50 text-green-700 ring-green-200" : "bg-gray-100 text-gray-500 ring-gray-200"}`}>
                  {activeOrg.active ? "Active" : "Inactive"}
                </span>
              </div>
              {[
                { label: "ID", value: activeOrg.id },
                { label: "Location ID", value: activeOrg.location_id || "—" },
                { label: "Profile Image Key", value: activeOrg.pfp_s3_key || "—" },
                { label: "Created", value: fmtDate(activeOrg.created_at) },
                { label: "Updated", value: fmtDate(activeOrg.updated_at) },
              ].map(function (row: { label: string; value: string }) {
                return (
                  <div key={row.label} className="px-5 py-3 grid grid-cols-3 gap-4">
                    <span className="text-xs font-medium text-gray-500">{row.label}</span>
                    <span className="col-span-2 text-sm text-gray-800 font-mono break-all">{row.value}</span>
                  </div>
                );
              })}
            </div>

            {/* Managers card */}
            <div className="bg-white rounded-lg border border-gray-200">
              <div className="px-5 py-4 border-b border-gray-100 flex items-center justify-between">
                <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">
                  Managers <span className="ml-1 text-xs font-medium text-gray-400 normal-case">({orgManagers.length})</span>
                </h2>
                {!addingMgr && (
                  <Btn size="sm" onClick={function () { setAddingMgr(true); }} icon={<IconPlus />}>
                    Add manager
                  </Btn>
                )}
              </div>

              {orgManagers.length === 0 && !addingMgr && (
                <p className="px-5 py-4 text-sm text-gray-400">No managers assigned.</p>
              )}
              {orgManagers.map(function (m: Manager, i: number) {
                const isLast: boolean = orgManagers.length === 1;
                const borderClass: string = i < orgManagers.length - 1 || addingMgr ? "border-b border-gray-100" : "";
                return (
                  <div key={m.id} className={`px-5 py-3 flex items-center justify-between gap-3 ${borderClass}`}>
                    <div className="flex items-center gap-3 min-w-0">
                      <div className="w-8 h-8 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center text-xs font-bold shrink-0">
                        {m.name ? m.name.charAt(0).toUpperCase() : "?"}
                      </div>
                      <div className="min-w-0">
                        <p className="text-sm font-medium text-gray-900 truncate">{m.name}</p>
                        <p className="text-xs text-gray-500 truncate">{m.email}</p>
                        <p className="text-xs text-gray-400 truncate">{m.username}</p>
                      </div>
                    </div>
                    {isLast ? (
                      <span className="text-xs font-medium px-2 py-0.5 rounded bg-blue-50 text-blue-700 ring-1 ring-blue-200">Only manager</span>
                    ) : (
                      <button
                        onClick={function () { handleRemoveManager(activeOrg.id, m.id); }}
                        className="text-xs text-gray-400 hover:text-red-600 font-medium transition-colors"
                      >
                        Remove
                      </button>
                    )}
                  </div>
                );
              })}

              {addingMgr && (
                <div className="px-5 py-4 flex flex-col gap-3">
                  <p className="text-sm font-semibold text-gray-800">New manager</p>
                  <div className="grid grid-cols-2 gap-3">
                    <Field label="Full name" error={mgrErrors.name} required>
                      <Input value={mgr.name} error={mgrErrors.name} placeholder="Jane Doe"
                        onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setMgr(function (p: Manager) { return Object.assign({}, p, { name: e.target.value }); }); }} />
                    </Field>
                    <Field label="Email" error={mgrErrors.email} required>
                      <Input type="email" value={mgr.email} error={mgrErrors.email} placeholder="jane@acme.com"
                        onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setMgr(function (p: Manager) { return Object.assign({}, p, { email: e.target.value }); }); }} />
                    </Field>
                  </div>
                  <div className="grid grid-cols-2 gap-3">
                    <Field label="Username" error={mgrErrors.username} required>
                      <Input value={mgr.username} error={mgrErrors.username} placeholder="janedoe"
                        onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setMgr(function (p: Manager) { return Object.assign({}, p, { username: e.target.value }); }); }} />
                    </Field>
                    <Field label="Role" error={mgrErrors.role} required>
                      <Input value={mgr.role} error={mgrErrors.role} placeholder="manager"
                        onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setMgr(function (p: Manager) { return Object.assign({}, p, { role: e.target.value }); }); }} />
                    </Field>
                  </div>
                  <Field label="Language preference">
                    <Input value={mgr.language_preference} placeholder="en"
                      onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setMgr(function (p: Manager) { return Object.assign({}, p, { language_preference: e.target.value }); }); }} />
                  </Field>
                  <div className="flex gap-2 mt-1">
                    <Btn onClick={submitMgr}>Add manager</Btn>
                    <Btn variant="ghost" onClick={function () { setAddingMgr(false); setMgr(blankMgr()); setMgrErrors({}); }}>Cancel</Btn>
                  </div>
                </div>
              )}
            </div>

          </div>
        )}
      </div>

      {showCreate && (
        <CreateDrawer
          onClose={function () { setShowCreate(false); }}
          onCreate={handleRegister}
          existingEmails={allEmails}
        />
      )}
    </div>
  );
}