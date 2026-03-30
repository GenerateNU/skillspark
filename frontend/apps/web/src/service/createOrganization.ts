import { type CreateLocationInputBody, postLocation, type CreateOrganizationBody, createOrganization, type Organization, type ManagerSignUpInputBody, type signupManagerResponse, signupManager, getManagerByOrgId } from "@skillspark/api-client";
import type { ManagerFormInput } from "../components/admin_createModal";
import { validateMgr } from "../utils/validation";

export type CreateOrganziationLocationAndManagerInput = {
    orgName: string;
    orgActive: boolean;
    addressLine1: string;
    addressLine2: string;
    country: string;
    district: string;
    subdistrict: string;
    province: string;
    postalCode: string;
    managerInputs: ManagerFormInput[];
}

export async function createOrganziationLocationAndManager (
    {orgName,
    orgActive,
    addressLine1,
    addressLine2,
    country,
    district,
    subdistrict,
    province,
    postalCode,
    managerInputs
} : CreateOrganziationLocationAndManagerInput) {
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
}