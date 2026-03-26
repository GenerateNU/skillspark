import { useLocation, useNavigate } from "react-router-dom";
import type { Organization, Manager, Location } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { getManagerByOrgId, deleteOrganization, getLocationById, deleteManager } from "@skillspark/api-client";
import DeleteModal from "../components/admin_deleteModal";
import OrgDetailsCard from "../components/admin_orgDetailsCard";
import OrgLocationCard from "../components/admin_orgLocationCard";
import OrgManagerCard from "../components/admin_orgManagerCard";

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

  useEffect(function () {
    const orgFromState = location.state?.org as Organization;
    if (!orgFromState) {
      navigate("/admin", { replace: true });
      return;
    }
    setOrg(orgFromState);

    async function fetchManager(): Promise<void> {
      try {
        const res = await getManagerByOrgId(orgFromState.id);
        if (res.status === 200) setManager(res.data as Manager);
      } catch (e) {
        console.error(e);
      } finally {
        setLoadingMgr(false);
      }
    }

    async function fetchLocation(): Promise<void> {
      if (!orgFromState.location_id) return;
      try {
        const res = await getLocationById(orgFromState.location_id as string);
        if (res.status === 200 || res.status === 201) setOrgLocation(res.data as Location);
      } catch (e) {
        console.error(e);
      }
    }

    fetchManager();
    fetchLocation();
  }, [location.state?.org, navigate]);

  async function handleDelete(): Promise<void> {
    if (!org) return;
    try {
      setDeleting(true);
      if (manager != undefined) {
        const managerRes = await deleteManager(manager!.id);
        if (managerRes.status !== 200) throw new Error("Failed to delete manager");
      }
      const res = await deleteOrganization(org.id);
      if (res.status === 200) navigate("/admin", { replace: true });
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
          <button onClick={function () { navigate("/admin"); }} className="mt-3 text-sm text-blue-600 hover:underline cursor-pointer">
            Back to organizations
          </button>
        </div>
      </div>
    );
  }

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 flex items-center gap-3 shrink-0">
        <button onClick={function () { navigate("/admin"); }} className="text-sm text-gray-400 hover:text-gray-600 cursor-pointer">
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
            className="px-3.5 py-2 text-sm font-medium rounded-md bg-white border border-red-300 text-red-600 hover:bg-red-50 transition-colors cursor-pointer"
          >
            Delete
          </button>
        </div>
      </div>

      <div className="flex-1 overflow-auto bg-gray-50 p-6 flex justify-center">
        <div className="w-full max-w-lg flex flex-col gap-6">
          <h2 className="text-2xl font-bold text-gray-900 text-center">{org.name}</h2>
          <OrgDetailsCard org={org} onOrgUpdate={setOrg} fmtDate={fmtDate} />
          <OrgLocationCard org={org} orgLocation={orgLocation} onOrgUpdate={setOrg} onLocationUpdate={setOrgLocation} />
          <OrgManagerCard manager={manager} loadingMgr={loadingMgr} />
        </div>
      </div>

      {showDeleteModal && (
        <DeleteModal
          title="Delete organization"
          description={<>Are you sure you want to delete <span className="font-semibold text-gray-900">{org.name}</span>? This action cannot be undone.</>}
          deleting={deleting}
          onConfirm={handleDelete}
          onClose={function () { setShowDeleteModal(false); }}
        />
      )}
    </div>
  );
}