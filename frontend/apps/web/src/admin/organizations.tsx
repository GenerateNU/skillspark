import { useLocation, useNavigate } from "react-router-dom";
import type { Organization, Manager } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { getManagerByOrgId } from "@skillspark/api-client";
import Badge from "../components/badge";
import { IconChevronRight } from "../components/icons";

const fmtDate = (iso: string): string =>
  new Date(iso).toLocaleDateString("en-US", { month: "short", day: "numeric", year: "numeric" });

export default function OrganizationDetailPage() {
  const location = useLocation();
  const navigate = useNavigate();
  const org: Organization | undefined = location.state?.org;

  const [managers, setManagers] = useState<Manager[]>([]);
  const [loadingMgrs, setLoadingMgrs] = useState<boolean>(true);

  useEffect(function () {
    if (!org) {
      navigate("/admin", { replace: true });
    }
  }, []);

  useEffect(function () {
    if (!org) return;
    async function fetchManagers(): Promise<void> {
      try {
        const res = await getManagerByOrgId(org!.id);
        if (res.status === 200) {
          setManagers(res.data as Manager[]);
        }
      } catch (e) {
        console.error(e);
      } finally {
        setLoadingMgrs(false);
      }
    }
    fetchManagers();
  }, [org?.id]);

  if (!org) {
    return (
      <div className="flex-1 flex items-center justify-center">
        <div className="text-center">
          <p className="text-sm font-semibold text-gray-700">No organization selected</p>
          <button
            onClick={function () { navigate("/admin"); }}
            className="mt-3 text-sm text-blue-600 hover:underline"
          >
            Back to organizations
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      {/* Toolbar */}
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center gap-3 shrink-0">
        <button
          onClick={function () { navigate("/admin"); }}
          className="text-sm text-gray-400 hover:text-gray-600 flex items-center gap-1"
        >
          Organizations
        </button>
        <span className="text-gray-300"><IconChevronRight /></span>
        <h1 className="text-base font-semibold text-gray-900">{org.name}</h1>
        <div className="ml-auto">
          <Badge color={org.active ? "green" : "gray"}>{org.active ? "Active" : "Inactive"}</Badge>
        </div>
      </div>

      {/* Content */}
      <div className="flex-1 overflow-auto bg-gray-50 p-6">
        <div className="max-w-2xl flex flex-col gap-6">

          {/* Details card */}
          <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
            <div className="px-5 py-4">
              <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">Details</h2>
            </div>
            {[
              { label: "ID", value: org.id, mono: true },
              { label: "Location ID", value: org.location_id || "—", mono: true },
              { label: "Created", value: fmtDate(org.created_at), mono: false },
              { label: "Updated", value: fmtDate(org.updated_at), mono: false },
            ].map(function (row: { label: string; value: string; mono: boolean }) {
              return (
                <div key={row.label} className="px-5 py-3 grid grid-cols-3 gap-4">
                  <span className="text-xs font-medium text-gray-500">{row.label}</span>
                  <span className={`col-span-2 text-sm text-gray-800 break-all ${row.mono ? "font-mono" : ""}`}>
                    {row.value}
                  </span>
                </div>
              );
            })}
            <div className="px-5 py-3 grid grid-cols-3 gap-4">
              <span className="text-xs font-medium text-gray-500">Stripe</span>
              <span className="col-span-2">
                <Badge color={org.stripe_account_activated ? "green" : "yellow"}>
                  {org.stripe_account_activated ? "Connected" : "Not connected"}
                </Badge>
              </span>
            </div>
          </div>

          {/* Managers card */}
          <div className="bg-white rounded-lg border border-gray-200">
            <div className="px-5 py-4 border-b border-gray-100">
              <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">
                Managers <span className="ml-1 text-xs font-medium text-gray-400 normal-case">({managers.length})</span>
              </h2>
            </div>
            {loadingMgrs ? (
              <p className="px-5 py-4 text-sm text-gray-400">Loading managers…</p>
            ) : managers.length === 0 ? (
              <p className="px-5 py-4 text-sm text-gray-400">No managers assigned.</p>
            ) : (
              managers.map(function (m: Manager, i: number) {
                const borderClass: string = i < managers.length - 1 ? "border-b border-gray-100" : "";
                return (
                  <div key={m.id} className={`px-5 py-3 flex items-center gap-3 ${borderClass}`}>
                    <div className="w-8 h-8 rounded-full bg-blue-100 text-blue-700 flex items-center justify-center text-xs font-bold shrink-0">
                      {m.name ? m.name.charAt(0).toUpperCase() : "?"}
                    </div>
                    <div className="min-w-0">
                      <p className="text-sm font-medium text-gray-900 truncate">{m.name}</p>
                      <p className="text-xs text-gray-500 truncate">{m.email}</p>
                      <p className="text-xs text-gray-400 truncate">{m.username} · {m.role}</p>
                    </div>
                  </div>
                );
              })
            )}
          </div>

        </div>
      </div>
    </div>
  );
}