import { Link } from "react-router-dom";
import { IconBuilding, IconPlus } from "../components/icons";
import { listOrganizations } from "@skillspark/api-client";
import type { Organization } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { CreateModal } from "../components/admin_createModal";

export default function HomePage() {
  const [loading, setLoading] = useState<boolean>(true);
  const [isError, setIsError] = useState<boolean>(false);
  const [organizations, setOrganizations] = useState<Organization[]>([]);
  const [showCreate, setShowCreate] = useState<boolean>(false);

  const fmtDate = (iso: string): string =>
    new Date(iso).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });

  useEffect(function () {
    async function getOrgs(): Promise<void> {
      try {
        const orgResponse = await listOrganizations();
        if (orgResponse.status !== 200) {
          setIsError(true);
          setLoading(false);
          return;
        }
        setOrganizations(orgResponse.data);
        setLoading(false);
      } catch {
        setIsError(true);
        setLoading(false);
      }
    }
    getOrgs();
  }, []);

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center justify-between shrink-0">
        <div>
          <h1 className="text-base font-semibold text-gray-900">
            Organizations
          </h1>
          <p className="text-xs text-gray-500">
            {organizations.length}{" "}
            {organizations.length === 1 ? "organization" : "organizations"}{" "}
            registered
          </p>
        </div>
        <button
          onClick={function () {
            setShowCreate(true);
          }}
          className="inline-flex items-center gap-2 px-3.5 py-2 text-sm font-medium rounded-md bg-blue-600 hover:bg-blue-700 text-white transition-colors cursor-pointer"
        >
          <IconPlus /> Add organization
        </button>
      </div>

      <div className="flex-1 overflow-auto bg-gray-50 p-6">
        {loading && (
          <div className="flex items-center justify-center h-full">
            <p className="text-sm text-gray-400">Loading organizations…</p>
          </div>
        )}
        {isError && (
          <div className="flex items-center justify-center h-full">
            <p className="text-sm text-red-500">
              Failed to load organizations.
            </p>
          </div>
        )}
        {!loading && !isError && organizations.length === 0 && (
          <div className="flex flex-col items-center justify-center h-full text-center">
            <div className="w-12 h-12 rounded-full bg-gray-100 flex items-center justify-center mb-3">
              <IconBuilding />
            </div>
            <p className="text-sm font-semibold text-gray-700">
              No organizations yet
            </p>
            <p className="text-sm text-gray-400 mt-1">
              Click "Add organization" to register the first one.
            </p>
          </div>
        )}
        {!loading && organizations.length > 0 && (
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
            {organizations.map(function (org: Organization) {
              return (
                <Link
                  key={org.id}
                  to={"/admin/organization/"}
                  state={{ org }}
                  className="group bg-white rounded-lg border border-gray-200 p-5 hover:border-blue-300 hover:shadow-sm transition-all"
                >
                  <div className="flex items-center gap-3 mb-3">
                    <div className="w-9 h-9 rounded-md bg-blue-600 flex items-center justify-center text-white text-sm font-bold shrink-0">
                      {org.name.charAt(0).toUpperCase()}
                    </div>
                    <div className="min-w-0">
                      <p className="text-sm font-semibold text-gray-900 group-hover:text-blue-700 truncate">
                        {org.name}
                      </p>
                      <p className="text-xs text-gray-400">
                        {org.active ? "Active" : "Inactive"}
                      </p>
                    </div>
                  </div>
                  <div className="flex items-center justify-between mt-2">
                    <span
                      className={`inline-flex items-center text-xs font-medium px-2 py-0.5 rounded ring-1 ${org.stripe_account_activated ? "bg-green-50 text-green-700 ring-green-200" : "bg-yellow-50 text-yellow-700 ring-yellow-200"}`}
                    >
                      {org.stripe_account_activated
                        ? "Stripe connected"
                        : "Stripe pending"}
                    </span>
                    <span className="text-xs text-gray-400">
                      {fmtDate(org.created_at)}
                    </span>
                  </div>
                </Link>
              );
            })}
          </div>
        )}
      </div>

      {showCreate && (
        <CreateModal
          onClose={function () {
            setShowCreate(false);
          }}
          onCreate={function (org: Organization) {
            setOrganizations(function (prev: Organization[]) {
              return [org, ...prev];
            });
            setShowCreate(false);
          }}
        />
      )}
    </div>
  );
}
