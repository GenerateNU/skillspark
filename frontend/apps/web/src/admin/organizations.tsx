import { useLocation, useNavigate } from "react-router-dom";
import type { Organization, Manager, Location, UpdateOrganizationBody } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { getManagerByOrgId, deleteOrganization, updateOrganization, getLocationById, deleteManager } from "@skillspark/api-client";

const fmtDate = (iso: string): string =>
  new Date(iso).toLocaleDateString("en-US", { month: "short", day: "numeric", year: "numeric" });

export default function OrganizationDetailPage() {
  const [org, setOrg] = useState<Organization>();
  const location = useLocation();
  const navigate = useNavigate();

  const [manager, setManager] = useState<Manager>();
  const [orgLocation, setOrgLocation] = useState<Location>();
  const [loadingMgr, setLoadingMgr] = useState<boolean>(true);
  const [showDeleteModal, setShowDeleteModal] = useState<boolean>(false);
  const [deleting, setDeleting] = useState<boolean>(false);

  const [editing, setEditing] = useState<boolean>(false);
  const [editName, setEditName] = useState<string>("");
  const [editActive, setEditActive] = useState<boolean>(true);
  const [saving, setSaving] = useState<boolean>(false);

  useEffect(function () {
    const orgFromState = location.state?.org as Organization;
    if (!orgFromState) {
      navigate("/admin", { replace: true });
      return;
    }
    setOrg(orgFromState);
    setEditName(orgFromState.name);
    setEditActive(orgFromState.active);

    async function fetchManager(): Promise<void> {
      try {
        const res = await getManagerByOrgId(orgFromState.id);
        if (res.status === 200) {
          setManager(res.data as Manager);
        }
      } catch (e) {
        console.error(e);
      } finally {
        setLoadingMgr(false);
      }
    }

    fetchManager();

    async function fetchLocation() {
      if (!orgFromState.location_id) return;
      try {
        const res = await getLocationById(orgFromState.location_id as string);
        if (res.status === 200 || res.status === 201) {
          setOrgLocation(res.data as Location);
        }
      } catch (e) {
        console.error(e);
      }
    }

    fetchLocation();
  }, []);

  function startEditing(): void {
    if (!org) return;
    setEditName(org.name);
    setEditActive(org.active);
    setEditing(true);
  }

  function cancelEditing(): void {
    setEditing(false);
  }

  async function handleSave(): Promise<void> {
    if (!org) return;
    if (!editName.trim()) return;
    try {
      setSaving(true);
      const input: UpdateOrganizationBody = {
        name: editName,
        location_id: org.location_id,
        active: editActive,
      };
      const res = await updateOrganization(org.id, input);
      if (res.status === 200) {
        setOrg(res.data as Organization);
        setEditing(false);
      }
    } catch (e) {
      console.error(e);
    } finally {
      setSaving(false);
    }
  }

  async function handleDelete(): Promise<void> {
    if (!org) return;
    try {
      setDeleting(true);
      if (manager != undefined) {
        const managerRes = await deleteManager(manager!.id);
        if (managerRes.status !== 200) {
          throw new Error("Failed to delete manager");
        }
      }
      const res = await deleteOrganization(org.id);
      if (res.status === 200) {
        navigate("/admin", { replace: true });
      }
    } catch (e) {
      console.error(e);
    } finally {
      setDeleting(false);
      setShowDeleteModal(false);
    }
  }

  if (!org) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <div className="text-center">
          <p className="text-base font-semibold text-gray-700">No organization selected</p>
          <button onClick={function () { navigate("/admin"); }} className="mt-3 text-sm text-blue-600 hover:underline">
            Back to organizations
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center gap-3 shrink-0">
        <button onClick={function () { navigate("/admin"); }} className="text-sm text-gray-400 hover:text-gray-600">
          Organizations
        </button>
        <span className="text-gray-300">›</span>
        <h1 className="text-base font-semibold text-gray-900">{org.name}</h1>
        <div className="ml-auto flex items-center gap-3">
          <span className={`inline-flex items-center text-sm font-medium px-2.5 py-1 rounded ring-1 ${org.active ? "bg-green-50 text-green-700 ring-green-200" : "bg-gray-100 text-gray-500 ring-gray-200"}`}>
            {org.active ? "Active" : "Inactive"}
          </span>
          <button
            onClick={function () { setShowDeleteModal(true); }}
            className="px-3.5 py-2 text-sm font-medium rounded-md bg-white border border-red-300 text-red-600 hover:bg-red-50 transition-colors"
          >
            Delete
          </button>
        </div>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-auto bg-gray-50 p-6 flex justify-center">
        <div className="w-full max-w-lg flex flex-col gap-6">

          {/* Title */}
          <h2 className="text-2xl font-bold text-gray-900 text-center">{org.name}</h2>

          {/* Details card */}
          <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
            <div className="px-5 py-4 flex items-center justify-between">
              <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">Details</h3>
              {!editing ? (
                <button onClick={startEditing} className="text-sm font-medium text-blue-600 hover:text-blue-800 transition-colors">
                  Edit
                </button>
              ) : (
                <div className="flex items-center gap-2">
                  <button onClick={cancelEditing} disabled={saving}
                    className="text-sm font-medium text-gray-500 hover:text-gray-700 transition-colors disabled:opacity-50">
                    Cancel
                  </button>
                  <button onClick={handleSave} disabled={saving}
                    className="px-3.5 py-1.5 text-sm font-medium rounded-md bg-blue-600 hover:bg-blue-700 text-white transition-colors disabled:opacity-50">
                    {saving ? "Saving…" : "Save changes"}
                  </button>
                </div>
              )}
            </div>

            {!editing ? (
              <>
                {[
                  { label: "ID", value: org.id, mono: true },
                  { label: "Created", value: fmtDate(org.created_at), mono: false },
                  { label: "Updated", value: fmtDate(org.updated_at), mono: false },
                ].map(function (row: { label: string; value: string; mono: boolean }) {
                  return (
                    <div key={row.label} className="px-5 py-3.5 grid grid-cols-3 gap-4">
                      <span className="text-sm font-medium text-gray-500">{row.label}</span>
                      <span className={`col-span-2 text-base text-gray-800 break-all ${row.mono ? "font-mono" : ""}`}>
                        {row.value}
                      </span>
                    </div>
                  );
                })}
                <div className="px-5 py-3.5 grid grid-cols-3 gap-4">
                  <span className="text-sm font-medium text-gray-500">Stripe</span>
                  <span className="col-span-2">
                    <span className={`inline-flex items-center text-sm font-medium px-2.5 py-1 rounded ring-1 ${org.stripe_account_activated ? "bg-green-50 text-green-700 ring-green-200" : "bg-yellow-50 text-yellow-700 ring-yellow-200"}`}>
                      {org.stripe_account_activated ? "Connected" : "Not connected"}
                    </span>
                  </span>
                </div>
              </>
            ) : (
              <div className="px-5 py-4 flex flex-col gap-4">
                <div className="flex flex-col gap-1">
                  <label className="text-sm font-medium text-gray-700">Organization name <span className="text-red-500">*</span></label>
                  <input
                    value={editName}
                    onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setEditName(e.target.value); }}
                    placeholder="Acme Kids Academy"
                    className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500"
                  />
                </div>
                <div className="flex flex-col gap-1">
                  <label className="text-sm font-medium text-gray-700">Active</label>
                  <select
                    value={editActive ? "true" : "false"}
                    onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { setEditActive(e.target.value === "true"); }}
                    className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500"
                  >
                    <option value="true">Yes</option>
                    <option value="false">No</option>
                  </select>
                </div>
                <div className="grid grid-cols-3 gap-4 py-1">
                  <span className="text-sm font-medium text-gray-500 self-center">Created</span>
                  <span className="col-span-2 text-base text-gray-400">{fmtDate(org.created_at)}</span>
                </div>
                <div className="grid grid-cols-3 gap-4 py-1">
                  <span className="text-sm font-medium text-gray-500 self-center">Stripe</span>
                  <span className="col-span-2">
                    <span className={`inline-flex items-center text-sm font-medium px-2.5 py-1 rounded ring-1 ${org.stripe_account_activated ? "bg-green-50 text-green-700 ring-green-200" : "bg-yellow-50 text-yellow-700 ring-yellow-200"}`}>
                      {org.stripe_account_activated ? "Connected" : "Not connected"}
                    </span>
                  </span>
                </div>
              </div>
            )}
          </div>

          {/* Location card */}
          <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
            <div className="px-5 py-4 border-b border-gray-100">
              <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">Location</h3>
            </div>
            {!orgLocation ? (
              <p className="px-5 py-4 text-base text-gray-400">No location assigned.</p>
            ) : (
              <>
                {[
                  { label: "Address", value: orgLocation.address_line1, mono: false },
                  { label: "Address line 2", value: orgLocation.address_line2 || "—", mono: false },
                  { label: "Subdistrict", value: orgLocation.subdistrict, mono: false },
                  { label: "District", value: orgLocation.district, mono: false },
                  { label: "Province", value: orgLocation.province, mono: false },
                  { label: "Postal code", value: orgLocation.postal_code, mono: true },
                  { label: "Country", value: orgLocation.country, mono: false },
                  { label: "Coordinates", value: `${orgLocation.latitude}, ${orgLocation.longitude}`, mono: true },
                ].map(function (row) {
                  return (
                    <div key={row.label} className="px-5 py-3.5 grid grid-cols-3 gap-4">
                      <span className="text-sm font-medium text-gray-500">{row.label}</span>
                      <span className={`col-span-2 text-base text-gray-800 break-all ${row.mono ? "font-mono" : ""}`}>
                        {row.value}
                      </span>
                    </div>
                  );
                })}
              </>
            )}
          </div>

          {/* Manager card */}
          <div className="bg-white rounded-lg border border-gray-200">
            <div className="px-5 py-4 border-b border-gray-100">
              <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">Manager</h3>
            </div>
            {loadingMgr ? (
              <p className="px-5 py-4 text-base text-gray-400">Loading manager…</p>
            ) : !manager ? (
              <p className="px-5 py-4 text-base text-gray-400">No manager assigned.</p>
            ) : (
              <div className="px-5 py-4 flex items-center gap-4">
                <div className="w-12 h-12 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center text-base font-bold shrink-0">
                  {manager.name ? manager.name.charAt(0).toUpperCase() : "?"}
                </div>
                <div className="min-w-0">
                  <p className="text-base font-semibold text-gray-900">{manager.name}</p>
                  <p className="text-sm text-gray-500">{manager.email}</p>
                  <p className="text-sm text-gray-400">{manager.username} · {manager.role}</p>
                </div>
              </div>
            )}
          </div>

        </div>
      </div>

      {/* Delete confirmation modal */}
      {showDeleteModal && (
        <div className="fixed inset-0 z-50 flex items-center justify-center p-4" style={{ background: "rgba(0,0,0,0.45)" }}>
          <div className="bg-white rounded-xl shadow-2xl w-full max-w-sm">
            <div className="px-6 py-5 border-b border-gray-200">
              <h2 className="text-lg font-semibold text-gray-900">Delete organization</h2>
            </div>
            <div className="px-6 py-5">
              <p className="text-base text-gray-600">
                Are you sure you want to delete <span className="font-semibold text-gray-900">{org.name}</span>? This action cannot be undone.
              </p>
            </div>
            <div className="px-6 py-4 border-t border-gray-200 bg-gray-50 rounded-b-xl flex items-center justify-end gap-3">
              <button onClick={function () { setShowDeleteModal(false); }} disabled={deleting}
                className="px-3.5 py-2 text-sm font-medium rounded-md bg-white border border-gray-300 text-gray-700 hover:bg-gray-50 transition-colors disabled:opacity-50">
                Cancel
              </button>
              <button onClick={handleDelete} disabled={deleting}
                className="px-3.5 py-2 text-sm font-medium rounded-md bg-red-600 hover:bg-red-700 text-white transition-colors disabled:opacity-50">
                {deleting ? "Deleting…" : "Delete"}
              </button>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}