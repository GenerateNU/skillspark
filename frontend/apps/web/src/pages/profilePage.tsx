import { useState } from "react";
import { Btn, Field, Input, Select } from "../components/common";

export function ProfilePage() {
  const [editing, setEditing] = useState<boolean>(false);
  const [form, setForm] = useState({
    name: "Admin User",
    email: "admin@skillspark.co",
    username: "admin",
    language_preference: "en",
    role: "admin",
  });
  const [saved, setSaved] = useState(Object.assign({}, form));

  function updForm(k: keyof typeof form, v: string): void {
    setForm(function (prev: typeof form) { return Object.assign({}, prev, { [k]: v }); });
  }

  function save(): void {
    setSaved(Object.assign({}, form));
    setEditing(false);
  }

  function cancel(): void {
    setForm(Object.assign({}, saved));
    setEditing(false);
  }

  const languageLabels: Record<string, string> = { en: "English", th: "Thai" };

  return (
    <div className="flex-1 flex flex-col overflow-hidden">
      <div className="bg-white border-b border-gray-200 px-6 py-4 shrink-0">
        <h1 className="text-base font-semibold text-gray-900">Profile</h1>
      </div>

      <div className="flex-1 overflow-auto bg-gray-50 p-6 flex justify-center">
        <div className="w-full max-w-lg flex flex-col gap-6">

          <div className="bg-white rounded-lg border border-gray-200 px-5 py-6 flex flex-col items-center gap-3 text-center">
            <div className="w-20 h-20 rounded-full bg-blue-600 flex items-center justify-center text-3xl font-bold text-white">
              {saved.name.charAt(0).toUpperCase()}
            </div>
            <div>
              <p className="text-base font-semibold text-gray-900">{saved.name}</p>
              <p className="text-sm text-gray-500">{saved.email}</p>
              <span className="inline-flex items-center text-xs font-medium px-2 py-0.5 rounded ring-1 bg-blue-50 text-blue-700 ring-blue-200 mt-1">
                {saved.role}
              </span>
            </div>
          </div>

          <div className="bg-white rounded-lg border border-gray-200 divide-y divide-gray-100">
            <div className="px-5 py-4">
              <h2 className="text-sm font-semibold text-gray-700 uppercase tracking-wide">Account details</h2>
            </div>
            {!editing ? (
              <>
                {[
                  { label: "Name", value: saved.name },
                  { label: "Email", value: saved.email },
                  { label: "Username", value: saved.username },
                  { label: "Language", value: languageLabels[saved.language_preference] ?? saved.language_preference },
                  { label: "Role", value: saved.role },
                ].map(function (row: { label: string; value: string }) {
                  return (
                    <div key={row.label} className="px-5 py-3 grid grid-cols-3 gap-4">
                      <span className="text-xs font-medium text-gray-500">{row.label}</span>
                      <span className="col-span-2 text-sm text-gray-800">{row.value}</span>
                    </div>
                  );
                })}
              </>
            ) : (
              <div className="px-5 py-5 flex flex-col gap-4">
                <div className="grid grid-cols-2 gap-4">
                  <Field label="Name" required>
                    <Input value={form.name} placeholder="Jane Doe"
                      onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updForm("name", e.target.value); }} />
                  </Field>
                  <Field label="Username" required>
                    <Input value={form.username} placeholder="janedoe"
                      onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updForm("username", e.target.value); }} />
                  </Field>
                </div>
                <Field label="Email" required>
                  <Input type="email" value={form.email} placeholder="jane@acme.com"
                    onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updForm("email", e.target.value); }} />
                </Field>
                <div className="grid grid-cols-2 gap-4">
                  <Field label="Language preference">
                    <Select
                      value={form.language_preference}
                      onChange={function (e: React.ChangeEvent<HTMLSelectElement>) { updForm("language_preference", e.target.value); }}
                    >
                      <option value="en">English</option>
                      <option value="th">Thai</option>
                    </Select>
                  </Field>
                  <Field label="Role">
                    <Input value={form.role} placeholder="admin"
                      onChange={function (e: React.ChangeEvent<HTMLInputElement>) { updForm("role", e.target.value); }} />
                  </Field>
                </div>
              </div>
            )}
          </div>

          <div className="flex justify-center gap-3">
            {!editing ? (
              <Btn variant="ghost" onClick={function () { setEditing(true); }}>Edit profile</Btn>
            ) : (
              <>
                <Btn variant="ghost" onClick={cancel}>Cancel</Btn>
                <Btn onClick={save}>Save changes</Btn>
              </>
            )}
          </div>

        </div>
      </div>
    </div>
  );
}