import { useState } from "react";
import type { Organization, UpdateOrganizationBody } from "@skillspark/api-client";
import { updateOrganization } from "@skillspark/api-client";

interface OrgDetailsCardProps {
  org: Organization;
  onOrgUpdate: (org: Organization) => void;
  fmtDate: (iso: string) => string;
}

export default function OrgDetailsCard({ org, onOrgUpdate, fmtDate }: OrgDetailsCardProps) {
  const [editing, setEditing] = useState<boolean>(false);
  const [editName, setEditName] = useState<string>("");
  const [editActive, setEditActive] = useState<boolean>(true);
  const [saving, setSaving] = useState<boolean>(false);

  function startEditing(): void {
    setEditName(org.name);
    setEditActive(org.active);
    setEditing(true);
  }

  function cancelEditing(): void {
    setEditing(false);
  }

  async function handleSave(): Promise<void> {
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
        onOrgUpdate(res.data as Organization);
        setEditing(false);
      }
    } catch (e) {
      console.error(e);
    } finally {
      setSaving(false);
    }
  }

  return (
    <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
      <div className="px-5 py-4 flex items-center justify-between">
        <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">Details</h3>
        {!editing ? (
          <button onClick={startEditing} className="text-sm font-medium text-blue-600 hover:text-blue-800 transition-colors cursor-pointer">
            Edit
          </button>
        ) : (
          <div className="flex items-center gap-2">
            <button onClick={cancelEditing} disabled={saving}
              className="text-sm font-medium text-gray-500 hover:text-gray-700 transition-colors disabled:opacity-50 cursor-pointer">
              Cancel
            </button>
            <button onClick={handleSave} disabled={saving}
              className="px-3.5 py-1.5 text-sm font-medium rounded-md bg-blue-600 hover:bg-blue-700 text-white transition-colors disabled:opacity-50 cursor-pointer">
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
              className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 cursor-pointer"
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
  );
}