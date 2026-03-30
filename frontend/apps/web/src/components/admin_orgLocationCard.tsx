import { useState } from "react";
import type { Organization, Location, CreateLocationInputBody } from "@skillspark/api-client";
import { postLocation, updateOrganization } from "@skillspark/api-client";

interface OrgLocationCardProps {
  org: Organization;
  orgLocation: Location | undefined;
  onOrgUpdate: (org: Organization) => void;
  onLocationUpdate: (location: Location) => void;
}

export default function OrgLocationCard({ org, orgLocation, onOrgUpdate, onLocationUpdate }: OrgLocationCardProps) {
  const [changingLocation, setChangingLocation] = useState<boolean>(false);
  const [locAddressLine1, setLocAddressLine1] = useState<string>("");
  const [locAddressLine2, setLocAddressLine2] = useState<string>("");
  const [locSubdistrict, setLocSubdistrict] = useState<string>("");
  const [locDistrict, setLocDistrict] = useState<string>("");
  const [locProvince, setLocProvince] = useState<string>("");
  const [locPostalCode, setLocPostalCode] = useState<string>("");
  const [locCountry, setLocCountry] = useState<string>("");
  const [savingLocation, setSavingLocation] = useState<boolean>(false);

  function startChangingLocation(): void {
    setLocAddressLine1("");
    setLocAddressLine2("");
    setLocSubdistrict("");
    setLocDistrict("");
    setLocProvince("");
    setLocPostalCode("");
    setLocCountry("");
    setChangingLocation(true);
  }

  const addressFields = orgLocation ? [
              { label: "Address", value: orgLocation.address_line1, mono: false },
              { label: "Address line 2", value: orgLocation.address_line2 || "—", mono: false },
              { label: "Subdistrict", value: orgLocation.subdistrict, mono: false },
              { label: "District", value: orgLocation.district, mono: false },
              { label: "Province", value: orgLocation.province, mono: false },
              { label: "Postal code", value: orgLocation.postal_code, mono: true },
              { label: "Country", value: orgLocation.country, mono: false },
              { label: "Coordinates", value: `${orgLocation.latitude}, ${orgLocation.longitude}`, mono: true },
            ] : [];

  function isLocationFormValid(): boolean {
    return (
      locAddressLine1.trim().length >= 5 &&
      locSubdistrict.trim().length >= 2 &&
      locDistrict.trim().length >= 2 &&
      locProvince.trim().length >= 2 &&
      locPostalCode.trim().length >= 3 &&
      locCountry.trim().length >= 2
    );
  }

  async function handleSaveLocation(): Promise<void> {
    if (!isLocationFormValid()) return;
    try {
      setSavingLocation(true);
      const locationInput: CreateLocationInputBody = {
        address_line1: locAddressLine1,
        ...(locAddressLine2.trim() ? { address_line2: locAddressLine2 } : {}),
        country: locCountry,
        district: locDistrict,
        subdistrict: locSubdistrict,
        province: locProvince,
        postal_code: locPostalCode,
      };
      const locationRes = await postLocation(locationInput);
      if (locationRes.status !== 200 && locationRes.status !== 201) throw new Error("Failed to create location");
      const newLocationId: string = (locationRes.data as { id: string }).id;

      const updateRes = await updateOrganization(org.id, {
        name: org.name,
        location_id: newLocationId,
        active: org.active,
      });
      if (updateRes.status !== 200) throw new Error("Failed to update organization");
      onOrgUpdate(updateRes.data as Organization);
      onLocationUpdate(locationRes.data as Location);
      setChangingLocation(false);
    } catch (e) {
      console.error(e);
    } finally {
      setSavingLocation(false);
    }
  }

  return (
    <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
      <div className="px-5 py-4 flex items-center justify-between">
        <h3 className="text-base font-semibold text-gray-700 uppercase tracking-wide">Location</h3>
        {!changingLocation ? (
          <button onClick={startChangingLocation} className="text-sm font-medium text-blue-600 hover:text-blue-800 transition-colors cursor-pointer">
            Change
          </button>
        ) : (
          <div className="flex items-center gap-2">
            <button onClick={function () { setChangingLocation(false); }} disabled={savingLocation}
              className="text-sm font-medium text-gray-500 hover:text-gray-700 transition-colors disabled:opacity-50 cursor-pointer">
              Cancel
            </button>
            <button onClick={handleSaveLocation} disabled={savingLocation || !isLocationFormValid()}
              className="px-3.5 py-1.5 text-sm font-medium rounded-md bg-blue-600 hover:bg-blue-700 text-white transition-colors disabled:opacity-50 cursor-pointer">
              {savingLocation ? "Saving…" : "Save location"}
            </button>
          </div>
        )}
      </div>

      {!changingLocation ? (
        !orgLocation ? (
          <p className="px-5 py-4 text-base text-gray-400">No location assigned.</p>
        ) : (
          <>
            {addressFields.map(function (row) {
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
        )
      ) : (
        <div className="px-5 py-4 flex flex-col gap-4">
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium text-gray-700">Address line 1 <span className="text-red-500">*</span></label>
            <input value={locAddressLine1} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocAddressLine1(e.target.value); }}
              placeholder="123 Sukhumvit Rd"
              className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium text-gray-700">Address line 2</label>
            <input value={locAddressLine2} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocAddressLine2(e.target.value); }}
              placeholder="Floor 4, Suite 401"
              className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
          </div>
          <div className="grid grid-cols-2 gap-3">
            <div className="flex flex-col gap-1">
              <label className="text-sm font-medium text-gray-700">Subdistrict <span className="text-red-500">*</span></label>
              <input value={locSubdistrict} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocSubdistrict(e.target.value); }}
                placeholder="Khlong Toei"
                className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
            </div>
            <div className="flex flex-col gap-1">
              <label className="text-sm font-medium text-gray-700">District <span className="text-red-500">*</span></label>
              <input value={locDistrict} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocDistrict(e.target.value); }}
                placeholder="Khlong Toei"
                className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
            </div>
          </div>
          <div className="grid grid-cols-2 gap-3">
            <div className="flex flex-col gap-1">
              <label className="text-sm font-medium text-gray-700">Province <span className="text-red-500">*</span></label>
              <input value={locProvince} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocProvince(e.target.value); }}
                placeholder="Bangkok"
                className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
            </div>
            <div className="flex flex-col gap-1">
              <label className="text-sm font-medium text-gray-700">Postal code <span className="text-red-500">*</span></label>
              <input value={locPostalCode} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocPostalCode(e.target.value); }}
                placeholder="10110"
                className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
            </div>
          </div>
          <div className="flex flex-col gap-1">
            <label className="text-sm font-medium text-gray-700">Country <span className="text-red-500">*</span></label>
            <input value={locCountry} onChange={function (e: React.ChangeEvent<HTMLInputElement>) { setLocCountry(e.target.value); }}
              placeholder="Thailand"
              className="w-full border border-gray-300 rounded-md px-3 py-2 text-base bg-white outline-none focus:ring-2 focus:ring-blue-500 focus:border-blue-500" />
          </div>
        </div>
      )}
    </div>
  );
}