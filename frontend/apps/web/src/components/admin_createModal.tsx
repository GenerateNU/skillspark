import { type Organization, type ManagerSignUpInputBody, type CreateOrganizationBody, type CreateLocationInputBody, createOrganization, type Manager, signupManager, type signupManagerResponse, getManagerByOrgId, postLocation } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { IconCheck, IconX } from "./icons";
import { validateMgr } from "../utils/validation";
import { blankMgr } from "../utils/validation";
import OrgLocationStep from "./admin_createModalStep0";
import ManagerStep from "./admin_createModalStep1";
import Btn from "../common/button";

export interface ManagerFormInput {
  name: string;
  email: string;
  username: string;
  role: string;
  language_preference: string;
  password: string;
}

interface CreateModalProps {
  onClose: () => void;
  onCreate: (org: Organization, managers: Manager[]) => void;
}

export function CreateModal({ onClose, onCreate }: CreateModalProps) {
  const [step, setStep] = useState<0 | 1>(0);

  const [orgName, setOrgName] = useState<string>("");
  const [orgActive, setOrgActive] = useState<boolean>(true);
  const [addressLine1, setAddressLine1] = useState<string>("");
  const [addressLine2, setAddressLine2] = useState<string>("");
  const [country, setCountry] = useState<string>("");
  const [district, setDistrict] = useState<string>("");
  const [subdistrict, setSubdistrict] = useState<string>("");
  const [province, setProvince] = useState<string>("");
  const [postalCode, setPostalCode] = useState<string>("");

  const [managerInputs, setManagerInputs] = useState<ManagerFormInput[]>([blankMgr()]);

  useEffect(function () {
    function handler(e: KeyboardEvent): void {
      if (e.key === "Escape") onClose();
    }
    window.addEventListener("keydown", handler);
    return function () { window.removeEventListener("keydown", handler); };
  }, [onClose]);

  function updMgr(index: number, k: keyof ManagerSignUpInputBody, v: string): void {
    setManagerInputs(function (prev: ManagerFormInput[]) {
      return prev.map(function (m: ManagerFormInput, i: number) {
        if (i !== index) return m;
        return Object.assign({}, m, { [k]: v });
      });
    });
  }

  function isStep0Valid(): boolean {
    return (
      orgName.trim().length > 0 &&
      addressLine1.trim().length >= 5 &&
      country.trim().length >= 2 &&
      district.trim().length >= 2 &&
      subdistrict.trim().length >= 2 &&
      province.trim().length >= 2 &&
      postalCode.trim().length >= 3
    );
  }

  async function handleCreate(): Promise<void> {
    try {
      const locationInput: CreateLocationInputBody = {
        address_line1: addressLine1,
        ...(addressLine2.trim() ? { address_line2: addressLine2 } : {}),
        country,
        district,
        subdistrict,
        province,
        postal_code: postalCode,
      };
      const locationRes = await postLocation(locationInput);
      if (locationRes.status !== 200 && locationRes.status !== 201) throw new Error("Failed to create location");
      const locationId: string = (locationRes.data as { id: string }).id;

      const orgInput: CreateOrganizationBody = {
        name: orgName,
        location_id: locationId,
        active: orgActive,
        profile_image: new Blob([], { type: "image/png" }),
      };
      const createdOrg = await createOrganization(orgInput);
      if (createdOrg.status !== 200 && createdOrg.status !== 201) throw new Error("Failed to create organization");
      const org: Organization = createdOrg.data as Organization;

      const completeManagerInputs: ManagerSignUpInputBody[] = managerInputs.map(function (man: ManagerFormInput) {
        const errors = validateMgr(man);
        if (Object.keys(errors).length > 0) throw new Error(`Manager ${man.name || "unknown"} has incomplete or invalid inputs`);
        return {
          name: man.name,
          email: man.email,
          username: man.username,
          password: man.password,
          role: man.role,
          language_preference: man.language_preference,
          organization_id: org.id,
        };
      });
      for (const mgr of completeManagerInputs) {
        const res: signupManagerResponse = await signupManager(mgr);
        if (res.status !== 200 && res.status !== 201) throw new Error(`Failed to sign up manager ${mgr.name}`);
      }

      const managersRes = await getManagerByOrgId(org.id);
      if (managersRes.status !== 200 && managersRes.status !== 201) throw new Error("Failed to fetch managers");
      onCreate(org, managersRes.data as Manager[]);
      onClose();
    } catch (e) {
      console.error(e);
    }
  }

  const stepLabels: string[] = ["Organization details", "Managers"];

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4 bg-black/45">
      <div className="bg-white rounded-xl shadow-2xl flex flex-col w-full max-w-lg max-h-[90vh]">

        {/* Header */}
        <div className="flex items-start justify-between px-6 py-5 border-b border-gray-200 shrink-0">
          <div>
            <h2 className="text-base font-semibold text-gray-900">
              {step === 0 ? "Register new organization" : "Add managers"}
            </h2>
            <p className="text-sm text-gray-500 mt-0.5">
              {step === 0 ? "Step 1 of 2 — Enter organization information" : "Step 2 of 2 — At least one manager required"}
            </p>
          </div>
          <button onClick={onClose} className="ml-4 p-1.5 rounded-md text-gray-400 hover:text-gray-600 hover:bg-gray-100 transition-colors cursor-pointer">
            <IconX />
          </button>
        </div>

        {/* Body */}
        <div className="flex-1 overflow-y-auto px-6 py-6">
          {/* Step indicator */}
          <div className="flex items-center gap-0 mb-6">
            {stepLabels.map(function (label: string, i: number) {
              const done: boolean = i < step;
              const active: boolean = i === step;
              return (
                <div key={i} className="flex items-center">
                  <div className="flex items-center gap-2">
                    <div className={`w-6 h-6 rounded-full flex items-center justify-center text-xs font-bold shrink-0 ${active || done ? "bg-blue-600 text-white" : "bg-gray-200 text-gray-500"}`}>
                      {done ? <IconCheck /> : i + 1}
                    </div>
                    <span className={`text-sm ${active ? "font-semibold text-gray-900" : done ? "text-gray-500" : "text-gray-400"}`}>{label}</span>
                  </div>
                  {i < stepLabels.length - 1 && <div className="w-8 h-px bg-gray-300 mx-3" />}
                </div>
              );
            })}
          </div>

          {step === 0 && (
            <OrgLocationStep
              orgName={orgName} setOrgName={setOrgName}
              orgActive={orgActive} setOrgActive={setOrgActive}
              addressLine1={addressLine1} setAddressLine1={setAddressLine1}
              addressLine2={addressLine2} setAddressLine2={setAddressLine2}
              subdistrict={subdistrict} setSubdistrict={setSubdistrict}
              district={district} setDistrict={setDistrict}
              province={province} setProvince={setProvince}
              postalCode={postalCode} setPostalCode={setPostalCode}
              country={country} setCountry={setCountry}
            />
          )}

          {step === 1 && (
            <ManagerStep
              managerInputs={managerInputs}
              updMgr={updMgr}
            />
          )}
        </div>

        {/* Footer */}
        <div className="shrink-0 px-6 py-4 border-t border-gray-200 bg-gray-50 flex items-center justify-end gap-3 rounded-b-xl">
          {step === 0 ? (
            <>
              <Btn variant="ghost" onClick={onClose}>Cancel</Btn>
              <Btn onClick={function () { if (!isStep0Valid()) return; setStep(1); }} disabled={!isStep0Valid()}>
                Continue
              </Btn>
            </>
          ) : (
            <>
              <Btn variant="ghost" onClick={function () { setStep(0); }}>Back</Btn>
              <Btn onClick={handleCreate} icon={<IconCheck />}>Create organization</Btn>
            </>
          )}
        </div>
      </div>
    </div>
  );
}