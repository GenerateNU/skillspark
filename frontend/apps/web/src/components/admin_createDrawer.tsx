import { type Organization, type ManagerSignUpInputBody, type CreateOrganizationBody, createOrganization, type Manager, signupManager, type signupManagerResponse, getManagerByOrgId } from "@skillspark/api-client";
import { useState } from "react";
import { Drawer } from "./admin_drawer";
import { Btn, Field, Select } from "./common";
import { IconCheck, IconPlus } from "./icons";
import { blankMgr, validateMgr } from "./validation";
import ManagerFormRow from "./admin_managerFormRow";
import ValidatedInput from "./validatedInput";

interface CreateDrawerProps {
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

export function CreateDrawer({ onClose, onCreate }: CreateDrawerProps) {
  const [step, setStep] = useState<0 | 1>(0);
  const [orgName, setOrgName] = useState<string>("");
  const [locationId, setLocationId] = useState<string | undefined>();
  const [orgActive, setOrgActive] = useState<boolean>(true);

  const [managerInputs, setManagerInputs] = useState<ManagerFormInput[]>([blankMgr()]);

  function updMgr(index: number, k: keyof ManagerSignUpInputBody, v: string): void {
    setManagerInputs(function (prev: ManagerFormInput[]) {
      return prev.map(function (m: ManagerFormInput, i: number) {
        if (i !== index) return m;
        return Object.assign({}, m, { [k]: v });
      });
    });
  }

  const stepLabels: string[] = ["Organization details", "Managers"];

  return (
    <Drawer
      title={step === 0 ? "Register new organization" : "Add managers"}
      subtitle={step === 0 ? "Step 1 of 2 — Enter organization information" : "Step 2 of 2 — At least one manager required"}
      onClose={onClose}
      footer={
        step === 0 ? (
          <>
            <Btn variant="ghost" onClick={onClose}>Cancel</Btn>
            <Btn onClick={async () => {
              try {
                if (orgName.length === 0) {
                  throw new Error("Organization must have a name")
                }
                setStep(1);
              } catch (e) {
                console.error(e)
              }
            }}
            >Continue</Btn>
          </>
        ) : (
          <>
            <Btn variant="ghost" onClick={function () { setStep(0) }}>Back</Btn>
            <Btn onClick={async () => {
              try {
                const input: CreateOrganizationBody = {
                  name: orgName,
                  location_id: locationId,
                  active: orgActive,
                }
                const createdOrg = await createOrganization(input);
                if (createdOrg.status !== 201 && createdOrg.status !== 200) {
                  throw new Error("There was an error creating an organization")
                }
                const org: Organization = createdOrg.data as Organization;
                const completeManagerInputs: ManagerSignUpInputBody[] = managerInputs.map((man: ManagerFormInput) => {
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
                for (const input of completeManagerInputs) {
                  const newManResponse: signupManagerResponse = await signupManager(input)
                  if (newManResponse.status !== 200 && newManResponse.status !== 201) {
                    throw new Error("There was an error signing up a manager")
                  }
                }
                const completeManagers = await getManagerByOrgId(org!.id);
                if (completeManagers.status !== 200 && completeManagers.status !== 201) {
                  throw new Error("There was an error signing up a manager")
                }
                onCreate(org!, completeManagers.data as Manager[])
                onClose();
              } catch (e) {

              }
            }} icon={<IconCheck />}>Create organization</Btn>
          </>
        )
      }
    >
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
          <Field label="Organization name" required>
            <ValidatedInput
              value={orgName}
              onChange={function (v: string) { setOrgName(v); }}
              validate={function (v: string) { return v.trim() ? null : "Required"; }}
              placeholder="Acme Kids Academy"
            />
          </Field>
          <Field label="Location ID">
            <ValidatedInput
              value={locationId ?? ""}
              onChange={function (v: string) { setLocationId(v || undefined); }}
              validate={function (_v: string) { return null; }}
              placeholder="uuid"
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
        </div>
      )}

      {step === 1 && (
        <div className="flex flex-col gap-4">
          {managerInputs.map(function (m: ManagerFormInput, i: number) {
            return (
              <ManagerFormRow
                key={i} mgr={m} index={i}
                onChange={updMgr}
                onRemove={function (idx: number) {
                  setManagerInputs(function (prev: ManagerFormInput[]) {
                    return prev.filter(function (_: ManagerFormInput, j: number) { return j !== idx; });
                  });
                }}
                canRemove={managerInputs.length > 1}
              />
            );
          })}
          <button
            onClick={function () { setManagerInputs(function (prev: ManagerFormInput[]) { return [...prev, blankMgr()]; }); }}
            className="flex items-center gap-2 text-sm text-blue-600 hover:text-blue-800 font-medium py-2"
          >
            <IconPlus /> Add another manager
          </button>
        </div>
      )}
    </Drawer>
  );
}