import type { ManagerSignUpInputBody } from "@skillspark/api-client";
import type { ManagerErrors } from "./validation";
import { Field, Input, Select } from "./common";
import { IconX } from "./icons";
import OtpCard from "./admin_otpCard";

interface ManagerFormRowProps {
  mgr: ManagerSignUpInputBody;
  index: number;
  onChange: (index: number, key: keyof ManagerSignUpInputBody, value: string) => void;
  onRemove: (index: number) => void;
  canRemove: boolean;
  errors?: ManagerErrors;
}
export default function ManagerFormRow({ mgr, index, onChange, onRemove, canRemove, errors = {} }: ManagerFormRowProps) {
  return (
    <div className="rounded-md border border-gray-200 bg-gray-50/60 p-4">
      <div className="flex items-center justify-between mb-3">
        <span className="text-xs font-semibold text-gray-500 uppercase tracking-wide">Manager {index + 1}</span>
        {canRemove && (
          <button onClick={function () { onRemove(index); }} className="text-xs text-red-500 hover:text-red-700 font-medium flex items-center gap-1">
            <IconX /> Remove
          </button>
        )}
      </div>
      <div className="grid grid-cols-2 gap-3 mb-3">
        <Field label="Full name" error={errors.name} required>
          <Input value={mgr.name} error={errors.name} placeholder="Jane Doe"
            onChange={function (e: React.ChangeEvent<HTMLInputElement>) { onChange(index, "name", e.target.value); }} />
        </Field>
        <Field label="Email address" error={errors.email} required>
          <Input type="email" value={mgr.email} error={errors.email} placeholder="jane@acme.com"
            onChange={function (e: React.ChangeEvent<HTMLInputElement>) { onChange(index, "email", e.target.value); }} />
        </Field>
      </div>
      <div className="grid grid-cols-2 gap-3 mb-3">
        <Field label="Username" error={errors.username} required>
          <Input value={mgr.username} error={errors.username} placeholder="janedoe"
            onChange={function (e: React.ChangeEvent<HTMLInputElement>) { onChange(index, "username", e.target.value); }} />
        </Field>
        <Field label="Role" error={errors.role} required>
          <Input value={mgr.role} error={errors.role} placeholder="manager"
            onChange={function (e: React.ChangeEvent<HTMLInputElement>) { onChange(index, "role", e.target.value); }} />
        </Field>
      </div>
      <Field label="Language preference" error={errors.language_preference}>
        <Select value={mgr.language_preference}
          onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { onChange(index, "language_preference", e.target.value); }}>
          <option value="en">English</option>
          <option value="th">Thai</option>
        </Select>
      </Field>
      <OtpCard password={mgr.password} name={mgr.name} />
    </div>
  );
}