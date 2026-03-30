import Field from "../common/field";
import Select from "../common/select";
import ValidatedInput from "../common/validatedInput";

interface OrgLocationStepProps {
  orgName: string;
  setOrgName: (v: string) => void;
  orgActive: boolean;
  setOrgActive: (v: boolean) => void;
  addressLine1: string;
  setAddressLine1: (v: string) => void;
  addressLine2: string;
  setAddressLine2: (v: string) => void;
  subdistrict: string;
  setSubdistrict: (v: string) => void;
  district: string;
  setDistrict: (v: string) => void;
  province: string;
  setProvince: (v: string) => void;
  postalCode: string;
  setPostalCode: (v: string) => void;
  country: string;
  setCountry: (v: string) => void;
}

export default function OrgLocationStep({
  orgName, setOrgName,
  orgActive, setOrgActive,
  addressLine1, setAddressLine1,
  addressLine2, setAddressLine2,
  subdistrict, setSubdistrict,
  district, setDistrict,
  province, setProvince,
  postalCode, setPostalCode,
  country, setCountry,
}: OrgLocationStepProps) {
  return (
    <div className="flex flex-col gap-4">
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
  );
}