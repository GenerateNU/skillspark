import type { Organization, Manager } from "@skillspark/api-client";
import { useState, useCallback } from "react";
import { Drawer } from "../components/admin_drawer";
import { IconSearch, IconPlus, IconBuilding, IconChevronRight } from "../components/icons";
import Badge from "../components/badge";
import { Btn, Divider } from "../components/common";
import { CreateModal } from "../components/admin_createDrawer";

export function OrganizationsPage() {
  const [organizations, setOrganizations] = useState<Organization[]>([]);
  const [managers, setManagers] = useState<Manager[]>([]);
  const [showCreate, setShowCreate] = useState<boolean>(false);
  const [selected, setSelected] = useState<Organization | null>(null);
  const [search, setSearch] = useState<string>("");

  const fmtDate = (iso: string): string =>
  new Date(iso).toLocaleDateString("en-US", { month: "short", day: "numeric", year: "numeric" });

  const handleCreate = useCallback(function (org: Organization, newManagers: Manager[]): void {

    setOrganizations(function (prev: Organization[]) { return [org, ...prev]; });
    setManagers([...managers, ...newManagers]);
    setSelected(org);
  }, []);

  const handleRemoveManager = useCallback(function (orgId: string, email: string): void {
    setManagers(function (prev: Manager[]) {
      return prev.filter(function (m: Manager) {
        return !(m.organization_id === orgId && m.email === email);
      });
    });
  }, []);

  const filtered: Organization[] = organizations.filter(function (o: Organization) {
    const q: string = search.toLowerCase();
    return !q || o.name.toLowerCase().includes(q) || (o.location_id ?? "").toLowerCase().includes(q);
  });

  const selectedManagers: Manager[] = selected
    ? managers.filter(function (m: Manager) { return m.organization_id === selected.id; })
    : [];

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Toolbar */}
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center gap-4 shrink-0">
        <div>
          <h1 className="text-base font-semibold text-gray-900">Organizations</h1>
          <p className="text-xs text-gray-500">{organizations.length} {organizations.length === 1 ? "organization" : "organizations"} registered</p>
        </div>
        <div className="ml-auto flex items-center gap-3">
          <div className="relative">
            <span className="absolute left-3 top-1/2 -translate-y-1/2 pointer-events-none text-gray-400"><IconSearch /></span>
            <input
              className="border border-gray-300 rounded-md pl-9 pr-3 py-2 text-sm bg-white outline-none focus:ring-2 focus:ring-blue-500 w-56 placeholder:text-gray-400"
              placeholder="Search organizations…" value={search}
              onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setSearch(e.target.value); }}
            />
          </div>
          <Btn onClick={function () { setShowCreate(true); }} icon={<IconPlus />}>Add organization</Btn>
        </div>
      </div>

      {/* Table */}
      <div className="flex-1 overflow-auto bg-white">
        {filtered.length === 0 ? (
          <div className="flex flex-col items-center justify-center h-full text-center py-20">
            <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mb-3"><IconBuilding /></div>
            <p className="text-sm font-semibold text-gray-700">{organizations.length === 0 ? "No organizations yet" : "No results found"}</p>
            <p className="text-sm text-gray-400 mt-1">{organizations.length === 0 ? "Click \"Add organization\" to register the first one." : "Try adjusting your search."}</p>
          </div>
        ) : (
          <div>
            <div className="grid grid-cols-12 gap-4 px-6 py-3 border-b border-gray-200 bg-gray-50 sticky top-0">
              <div className="col-span-4 text-xs font-semibold text-gray-500 uppercase tracking-wide">Organization</div>
              <div className="col-span-2 text-xs font-semibold text-gray-500 uppercase tracking-wide">Status</div>
              <div className="col-span-2 text-xs font-semibold text-gray-500 uppercase tracking-wide">Stripe</div>
              <div className="col-span-2 text-xs font-semibold text-gray-500 uppercase tracking-wide">Managers</div>
              <div className="col-span-1 text-xs font-semibold text-gray-500 uppercase tracking-wide">Created</div>
              <div className="col-span-1" />
            </div>
            {filtered.map(function (o: Organization, i: number) {
              const borderClass: string = i < filtered.length - 1 ? "border-b border-gray-100" : "";
              const orgMgrCount: number = managers.filter(function (m: Manager) { return m.organization_id === o.id; }).length;
              return (
                <div key={o.id}
                  className={`grid grid-cols-12 gap-4 px-6 py-3.5 items-center hover:bg-gray-50 cursor-pointer transition-colors ${borderClass}`}
                  onClick={function () { setSelected(o); }}>
                  <div className="col-span-4 flex items-center gap-3 min-w-0">
                    <div className="w-8 h-8 rounded-md bg-blue-600 flex items-center justify-center text-white text-xs font-bold shrink-0">
                      {o.name.charAt(0).toUpperCase()}
                    </div>
                    <div className="min-w-0">
                      <p className="text-sm font-medium text-gray-900 truncate">{o.name}</p>
                      {o.location_id && <p className="text-xs text-gray-400 truncate font-mono">{o.location_id}</p>}
                    </div>
                  </div>
                  <div className="col-span-2">
                    <Badge color={o.active ? "green" : "gray"}>{o.active ? "Active" : "Inactive"}</Badge>
                  </div>
                  <div className="col-span-2">
                    <Badge color={o.stripe_account_activated ? "green" : "yellow"}>
                      {o.stripe_account_activated ? "Connected" : "Pending"}
                    </Badge>
                  </div>
                  <div className="col-span-2 text-sm text-gray-700">{orgMgrCount} manager{orgMgrCount !== 1 ? "s" : ""}</div>
                  <div className="col-span-1 text-xs text-gray-400">{fmtDate(o.created_at)}</div>
                  <div className="col-span-1 flex justify-end text-gray-300"><IconChevronRight /></div>
                </div>
              );
            })}
          </div>
        )}
      </div>

      {/* Detail drawer */}
      {selected && (
        <Drawer title={selected.name} subtitle={selected.active ? "Active" : "Inactive"} onClose={function () { setSelected(null); }} width="max-w-lg">
          <div className="grid grid-cols-2 gap-x-6 gap-y-3 text-sm mb-2">
            <div>
              <span className="text-xs font-medium text-gray-500 uppercase tracking-wide block mb-0.5">ID</span>
              <span className="text-gray-800 font-mono text-xs break-all">{selected.id}</span>
            </div>
            {selected.location_id && (
              <div>
                <span className="text-xs font-medium text-gray-500 uppercase tracking-wide block mb-0.5">Location ID</span>
                <span className="text-gray-800 font-mono text-xs">{selected.location_id}</span>
              </div>
            )}
            <div>
              <span className="text-xs font-medium text-gray-500 uppercase tracking-wide block mb-0.5">Stripe</span>
              <Badge color={selected.stripe_account_activated ? "green" : "yellow"}>
                {selected.stripe_account_activated ? "Connected" : "Not connected"}
              </Badge>
            </div>
            <div>
              <span className="text-xs font-medium text-gray-500 uppercase tracking-wide block mb-0.5">Created</span>
              <span className="text-gray-800">{fmtDate(selected.created_at)}</span>
            </div>
          </div>

          <Divider label="Managers" />

          {selectedManagers.length === 0 ? (
            <p className="text-sm text-gray-400">No managers assigned.</p>
          ) : (
            <div className="rounded-md border border-gray-200 overflow-hidden">
              {selectedManagers.map(function (m: Manager, i: number) {
                const isLast: boolean = selectedManagers.length === 1;
                const borderClass: string = i < selectedManagers.length - 1 ? "border-b border-gray-100" : "";
                return (
                  <div key={i} className={`px-4 py-3 flex items-center justify-between gap-3 ${borderClass}`}>
                    <div className="flex items-center gap-3 min-w-0">
                      <div className="w-8 h-8 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center text-xs font-bold shrink-0">
                        {m.name ? m.name.charAt(0).toUpperCase() : "?"}
                      </div>
                      <div className="min-w-0">
                        <p className="text-sm font-medium text-gray-900 truncate">{m.name}</p>
                        <p className="text-xs text-gray-500 truncate">{m.email}</p>
                        <p className="text-xs text-gray-400 truncate">{m.username} · {m.role}</p>
                      </div>
                    </div>
                    {isLast ? (
                      <Badge color="blue">Only manager</Badge>
                    ) : (
                      <button onClick={function () { handleRemoveManager(selected.id, m.email); }}
                        className="text-xs text-gray-400 hover:text-red-600 font-medium transition-colors">
                        Remove
                      </button>
                    )}
                  </div>
                );
              })}
            </div>
          )}
        </Drawer>
      )}

      {showCreate && (
        <CreateModal
          onClose={function () { setShowCreate(false); }}
          onCreate={handleCreate}
        />
      )}
    </div>
  );
}