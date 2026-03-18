import { type Organization, type ManagerSignUpInputBody, type CreateOrganizationBody, createOrganization } from "@skillspark/api-client";
import { useState } from "react";
import { Drawer } from "./admin_drawer";
import Btn, { Field, Input, Select } from "./common";
import { IconCheck, IconPlus } from "./icons";
import { blankMgr } from "./validation";
import ManagerFormRow from "./admin_managerFormRow";

interface CreateDrawerProps {
  onClose: () => void;
  onCreate: (org: Organization, managers: ManagerSignUpInputBody[]) => void;
}

export function CreateDrawer({ onClose, onCreate }: CreateDrawerProps) {
  const [step, setStep] = useState<0 | 1>(0);
  const [orgName, setOrgName] = useState<string>("");
  const [locationId, setLocationId] = useState<string | undefined>();
  const [orgActive, setOrgActive] = useState<boolean>(true);

  const [org, setOrg] = useState<Organization | null>();
  const [managers, setManagers] = useState<ManagerSignUpInputBody[]>([blankMgr()]);

  function updMgr(index: number, k: keyof ManagerSignUpInputBody, v: string): void {
    setManagers(function (prev: ManagerSignUpInputBody[]) {
      return prev.map(function (m: ManagerSignUpInputBody, i: number) {
        if (i !== index) return m;
        return Object.assign({}, m, { [k]: v });
      });
    });
  }

  function submit(): void {
    console.log(onCreate)
    console.log(org)
    // const newManagers: ManagerSignUpInputBody[] = managers.map(function (m: ManagerSignUpInputBody) {
    //   return Object.assign({}, m, { organization_id: org!.id });
    // });
    // onCreate(newOrg, newManagers);
    onClose();
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
                        return;
                    }
                    const input : CreateOrganizationBody = {
                        name: orgName,
                        location_id: locationId,
                        active: orgActive
                    }
                    const createdOrg = await createOrganization(input);
                    if (createdOrg.status !== 201) {
                        throw new Error("There was an error creating an organization")
                    }
                    setOrg(createdOrg.data as Organization);
                    setStep(1);
                    console.log(org);
                } catch (e) {
                    console.error(e)
                }
                }}
            >Continue</Btn>
          </>
        ) : (
          <>
            <Btn variant="ghost" onClick={function () { setStep(0)}}>Back</Btn>
            <Btn onClick={submit} icon={<IconCheck />}>Create organization</Btn>
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
            <Input
              value={orgName ?? ""}
              placeholder="Acme Kids Academy"
              onChange={ (e) => {setOrgName(e.target.value)}}
            />
          </Field>
          <Field label="Location ID">
            <Input
              value={locationId ?? ""}
              placeholder="uuid"
              onChange={(e) => {setLocationId(e.target.value)}}
            />
          </Field>
          <Field label="Active">
            <Select
              value={orgActive ? 'true' : 'false'}
              onChange={(e) => {setOrgActive(e.target.value === 'true')}}
            >
              <option value="true">Yes</option>
              <option value="false">No</option>
            </Select>
          </Field>
        </div>
      )}

      {step === 1 && (
        <div className="flex flex-col gap-4">
          {managers.map(function (m: ManagerSignUpInputBody, i: number) {
            return (
              <ManagerFormRow
                key={i} mgr={m} index={i}
                onChange={updMgr}
                onRemove={function (idx: number) {
                  setManagers(function (prev: ManagerSignUpInputBody[]) {
                    return prev.filter(function (_: ManagerSignUpInputBody, j: number) { return j !== idx; });
                  });
                }}
                canRemove={managers.length > 1}
              />
            );
          })}
          <button
            onClick={function () { setManagers(function (prev: ManagerSignUpInputBody[]) { return [...prev, blankMgr()]; }); }}
            className="flex items-center gap-2 text-sm text-blue-600 hover:text-blue-800 font-medium py-2"
          >
            <IconPlus /> Add another manager
          </button>
        </div>
      )}
    </Drawer>
  );
}