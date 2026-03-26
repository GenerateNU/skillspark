import { type Organization, type ManagerSignUpInputBody, type CreateOrganizationBody, type CreateLocationInputBody, createOrganization, type Manager, signupManager, type signupManagerResponse, getManagerByOrgId, postLocation } from "@skillspark/api-client";
import { useState, useEffect } from "react";
import { Btn, Field, Select } from "./common";
import { IconCheck, IconX } from "./icons";
import { blankMgr, validateMgr } from "./validation";
import ManagerFormRow from "./admin_managerFormRow";
import ValidatedInput from "./validatedInput";

interface CreateModalProps {
  onClose: () => void;
  onCreate: (org: Organization, managers: Manager[]) => void;
}

export interface ManagerFormInput {
  name: string;
  email: string;
  username: string;
  role: string;
  language_preference: string;
  password: string;
}

export function CreateModal({ onClose, onCreate }: CreateModalProps) {
  const [step, setStep] = useState<0 | 1>(0);

  // org fields
  const [orgName, setOrgName] = useState<string>("");
  const [orgActive, setOrgActive] = useState<boolean>(true);

  // location fields
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

  const stepLabels: string[] = ["Organization details", "Managers"];

  return (
    <div className="fixed inset-0 z-50 flex items-center justify-center p-4" style={{ background: "rgba(0,0,0,0.45)" }}>
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
          <button onClick={onClose} className="ml-4 p-1.5 rounded-md text-gray-400 hover:text-gray-600 hover:bg-gray-100 transition-colors">
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
            <div className="flex flex-col gap-4">
              {/* Organization */}
              <Field label="Organization name" required>
                <ValidatedInput
                  value={orgName}
                  onChange={function (v: string) { setOrgName(v); }}
                  validate={function (v: string) { return v.trim() ? null : "Required"; }}
                  placeholder="Acme Kids Academy"
                />
              </Field>
              <Field label="Active">
                <Select
                  value={orgActive ? "true" : "false"}
                  onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { setOrgActive(e.target.value === "true"); }}
                >
                  <option value="true">Yes</option>
                  <option value="false">No</option>
                </Select>
              </Field>

              {/* Divider */}
              <div className="relative my-1">
                <div className="absolute inset-0 flex items-center"><div className="w-full border-t border-gray-200" /></div>
                <div className="relative flex justify-start"><span className="bg-white pr-3 text-xs font-semibold text-gray-500 uppercase tracking-wide">Location</span></div>
              </div>

              <Field label="Address line 1" required>
                <ValidatedInput
                  value={addressLine1}
                  onChange={function (v: string) { setAddressLine1(v); }}
                  validate={function (v: string) {
                    if (!v.trim()) return "Required";
                    if (v.trim().length < 5) return "Must be at least 5 characters";
                    return null;
                  }}
                  placeholder="123 Sukhumvit Rd"
                />
              </Field>
              <Field label="Address line 2">
                <ValidatedInput
                  value={addressLine2}
                  onChange={function (v: string) { setAddressLine2(v); }}
                  validate={function (v: string) {
                    if (v && v.trim().length < 5) return "Must be at least 5 characters";
                    return null;
                  }}
                  placeholder="Floor 4, Suite 401"
                />
              </Field>
              <div className="grid grid-cols-2 gap-3">
                <Field label="Subdistrict" required>
                  <ValidatedInput
                    value={subdistrict}
                    onChange={function (v: string) { setSubdistrict(v); }}
                    validate={function (v: string) { return v.trim().length >= 2 ? null : "Required"; }}
                    placeholder="Khlong Toei"
                  />
                </Field>
                <Field label="District" required>
                  <ValidatedInput
                    value={district}
                    onChange={function (v: string) { setDistrict(v); }}
                    validate={function (v: string) { return v.trim().length >= 2 ? null : "Required"; }}
                    placeholder="Khlong Toei"
                  />
                </Field>
              </div>
              <div className="grid grid-cols-2 gap-3">
                <Field label="Province" required>
                  <ValidatedInput
                    value={province}
                    onChange={function (v: string) { setProvince(v); }}
                    validate={function (v: string) { return v.trim().length >= 2 ? null : "Required"; }}
                    placeholder="Bangkok"
                  />
                </Field>
                <Field label="Postal code" required>
                  <ValidatedInput
                    value={postalCode}
                    onChange={function (v: string) { setPostalCode(v); }}
                    validate={function (v: string) { return v.trim().length >= 3 ? null : "Required"; }}
                    placeholder="10110"
                  />
                </Field>
              </div>
              <Field label="Country" required>
                <ValidatedInput
                  value={country}
                  onChange={function (v: string) { setCountry(v); }}
                  validate={function (v: string) { return v.trim().length >= 2 ? null : "Required"; }}
                  placeholder="Thailand"
                />
              </Field>
            </div>
          )}

          {step === 1 && (
            <div className="flex flex-col gap-4">
              <ManagerFormRow
                mgr={managerInputs[0]}
                index={0}
                onChange={updMgr}
                onRemove={function () { }}
                canRemove={false}
              />
            </div>
          )}
        </div>

        {/* Footer */}
        <div className="shrink-0 px-6 py-4 border-t border-gray-200 bg-gray-50 flex items-center justify-end gap-3 rounded-b-xl">
          {step === 0 ? (
            <>
              <Btn variant="ghost" onClick={onClose}>Cancel</Btn>
              <Btn
                onClick={function () {
                  if (!isStep0Valid()) return;
                  setStep(1);
                }}
                disabled={!isStep0Valid()}
              >
                Continue
              </Btn>
            </>
          ) : (
            <>
              <Btn variant="ghost" onClick={function () { setStep(0); }}>Back</Btn>
              <Btn onClick={async function () {
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
                  if (locationRes.status !== 200 && locationRes.status !== 201) {
                    throw new Error("Failed to create location");
                  }
                  const locationId: string = (locationRes.data as { id: string }).id;
                  console.log(locationId)

                  const orgInput: CreateOrganizationBody = {
                    name: orgName,
                    location_id: locationId,
                    active: orgActive,
                    profile_image: new Blob([], { type: "image/png" }),
                  };
                  const createdOrg = await createOrganization(orgInput);
                  if (createdOrg.status !== 200 && createdOrg.status !== 201) {
                    throw new Error("Failed to create organization");
                  }
                  const org: Organization = createdOrg.data as Organization;

                  const completeManagerInputs: ManagerSignUpInputBody[] = managerInputs.map(function (man: ManagerFormInput) {
                    const errors = validateMgr(man);
                    if (Object.keys(errors).length > 0) {
                      throw new Error(`Manager ${man.name || "unknown"} has incomplete or invalid inputs`);
                    }
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
                    if (res.status !== 200 && res.status !== 201) {
                      throw new Error(`Failed to sign up manager ${mgr.name}`);
                    }
                  }

                  const managersRes = await getManagerByOrgId(org.id);
                  if (managersRes.status !== 200 && managersRes.status !== 201) {
                    throw new Error("Failed to fetch managers");
                  }
                  onCreate(org, managersRes.data as Manager[]);
                  onClose();
                } catch (e) {
                  console.error(e);
                }
              }} icon={<IconCheck />}>Create organization</Btn>
            </>
          )}
        </div>

      </div>
    </div>
  );
}