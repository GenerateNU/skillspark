import { useState } from "react";
import { Btn, Field, Input, Select } from "./common";
import { Drawer } from "./drawer";
import { IconCheck, IconPlus } from "./icons";
import { blankBiz, type BizErrors, blankMgr, type ManagerErrors, uid } from "./types";
import { validateBiz, validateMgr } from "./validation";
import { ManagerFormRow } from "./managerFormRow";
import type { CreateOrganizationBody, Manager, Organization } from "@skillspark/api-client";

interface CreateDrawerProps {
  onClose: () => void;
  onCreate: (org: Organization, managers: Manager[]) => void;
  existingEmails: Set<string>;
}

export function CreateDrawer({ onClose, onCreate, existingEmails }: CreateDrawerProps) {
  const [step, setStep] = useState<0 | 1>(0);
  const [biz, setBiz] = useState<CreateOrganizationBody>(blankBiz());
  const [bizErrors, setBizErrors] = useState<BizErrors>({});
  const [managers, setManagers] = useState<Manager[]>([blankMgr()]);
  const [mgrErrors, setMgrErrors] = useState<Record<string, ManagerErrors>>({});

  function updBiz(k: keyof CreateOrganizationBody, v: string): void {
    setBiz(function (prev: CreateOrganizationBody) {
      const next = Object.assign({}, prev);
      (next as unknown as Record<string, string>)[k] = v;
      return next;
    });
    setBizErrors(function (prev: BizErrors) {
      const next = Object.assign({}, prev);
      delete next[k];
      return next;
    });
  }

  function updMgr(id: string, k: keyof Manager, v: string): void {
    setManagers(function (prev: Manager[]) {
      return prev.map(function (m: Manager) {
        if (m.id !== id) return m;
        const next = Object.assign({}, m);
        (next as unknown as Record<string, string>)[k] = v;
        return next;
      });
    });
  }

  function goNext(): void {
    const e: BizErrors = validateBiz(biz);
    if (Object.keys(e).length) { setBizErrors(e); return; }
    setStep(1);
  }

  function submit(): void {
    const allErrors: Record<string, ManagerErrors> = {};
    let ok: boolean = true;
    managers.forEach(function (m: Manager) {
      const e: ManagerErrors = validateMgr(m);
      if (m.email && existingEmails.has(m.email.trim().toLowerCase())) e.email = "Email already in use";
      if (Object.keys(e).length) { allErrors[m.id] = e; ok = false; }
    });
    setMgrErrors(allErrors);
    if (!ok) return;
    const orgId: string = uid();
    const newOrg: Organization = {
      id: orgId,
      name: typeof biz.name === "string" ? biz.name : "",
      active: biz.active ?? true,
      location_id: typeof biz.location_id === "string" ? biz.location_id : undefined,
      pfp_s3_key: undefined,
      presigned_url: "",
      created_at: new Date().toISOString(),
      updated_at: new Date().toISOString(),
    };
    const newManagers: Manager[] = managers.map(function (m: Manager) {
      return Object.assign({}, m, { organization_id: orgId });
    });
    onCreate(newOrg, newManagers);
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
            <Btn onClick={goNext}>Continue</Btn>
          </>
        ) : (
          <>
            <Btn variant="ghost" onClick={function () { setStep(0); }}>Back</Btn>
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
                <span className={`text-sm ${active ? "font-semibold text-gray-900" : done ? "text-gray-500" : "text-gray-400"}`}>
                  {label}
                </span>
              </div>
              {i < stepLabels.length - 1 && <div className="w-8 h-px bg-gray-300 mx-3" />}
            </div>
          );
        })}
      </div>

      {step === 0 && (
        <div className="flex flex-col gap-4">
          <Field label="Organization name" error={bizErrors.name} required>
            <Input
              value={typeof biz.name === "string" ? biz.name : ""}
              error={bizErrors.name}
              placeholder="Acme Kids Academy"
              onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updBiz("name", e.target.value); }}
            />
          </Field>
          <Field label="Location ID" error={bizErrors.location_id}>
            <Input
              value={typeof biz.location_id === "string" ? biz.location_id : ""}
              error={bizErrors.location_id}
              placeholder="uuid"
              onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updBiz("location_id", e.target.value); }}
            />
          </Field>
          <Field label="Active">
            <Select
              value={biz.active ? "true" : "false"}
              onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { updBiz("active", e.target.value); }}
            >
              <option value="true">Yes</option>
              <option value="false">No</option>
            </Select>
          </Field>
        </div>
      )}

      {step === 1 && (
        <div className="flex flex-col gap-4">
          {managers.map(function (m: Manager, i: number) {
            return (
              <ManagerFormRow
                key={m.id} mgr={m} index={i}
                onChange={updMgr}
                onRemove={function (id: string) {
                  setManagers(function (prev: Manager[]) {
                    return prev.filter(function (x: Manager) { return x.id !== id; });
                  });
                }}
                canRemove={managers.length > 1}
                errors={mgrErrors[m.id] || {}}
              />
            );
          })}
          <button
            onClick={function () {
              setManagers(function (prev: Manager[]) { return [...prev, blankMgr()]; });
            }}
            className="flex items-center gap-2 text-sm text-blue-600 hover:text-blue-800 font-medium py-2"
          >
            <IconPlus />
            Add another manager
          </button>
        </div>
      )}
    </Drawer>
  );
}