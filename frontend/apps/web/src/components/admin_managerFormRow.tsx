import { isValidEmail, type ManagerErrors } from "./validation";
import { Field, Select } from "./common";
import { IconX } from "./icons";
import OtpCard from "./admin_otpCard";
import ValidatedInput from "./validatedInput";
import type { ManagerFormInput } from "./admin_createDrawer";

interface ManagerFormRowProps {
  mgr: ManagerFormInput;
  index: number;
  onChange: (index: number, key: keyof ManagerFormInput, value: string) => void;
  onRemove: (index: number) => void;
  canRemove: boolean;
  errors?: ManagerErrors;
}

export default function ManagerFormRow({ mgr, index, onChange, onRemove, canRemove }: ManagerFormRowProps) {
  function upd(key: keyof ManagerFormInput): (v: string) => void {
    return function (v: string) { onChange(index, key, v); };
  }

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
        <Field label="Full name" required>
          <ValidatedInput
            value={mgr.name}
            onChange={upd("name")}
            validate={function (v: string) { return v.trim() ? null : "Required"; }}
            placeholder="Jane Doe"
          />
        </Field>
        <Field label="Email address" required>
          <ValidatedInput
            type="email"
            value={mgr.email}
            onChange={upd("email")}
            validate={function (v: string) {
              if (!v.trim()) return "Required";
              if (!isValidEmail(v)) return "Invalid email address";
              return null;
            }}
            placeholder="ebk@skillspark.com"
          />
        </Field>
      </div>

      <div className="grid grid-cols-2 gap-3 mb-3">
        <Field label="Username" required>
          <ValidatedInput
            value={mgr.username}
            onChange={upd("username")}
            validate={function (v: string) { return v.trim() ? null : "Required"; }}
            placeholder="janedoe"
          />
        </Field>
        <Field label="Role" required>
          <ValidatedInput
            value={mgr.role}
            onChange={upd("role")}
            validate={function (v: string) { return v.trim() ? null : "Required"; }}
            placeholder="Manager"
          />
        </Field>
      </div>

      <Field label="Language preference">
        <Select
          value={mgr.language_preference}
          onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { onChange(index, "language_preference", e.target.value); }}
        >
          <option value="en">English</option>
          <option value="th">Thai</option>
        </Select>
      </Field>

      <OtpCard password={mgr.password} name={mgr.name} />
    </div>
  );
}