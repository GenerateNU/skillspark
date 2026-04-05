import type { ManagerSignUpInputBody } from "@skillspark/api-client";
import ManagerFormRow from "./admin_managerFormRow";
import type { ManagerFormInput } from "./admin_createModal";

interface ManagerStepProps {
  managerInputs: ManagerFormInput[];
  updMgr: (index: number, k: keyof ManagerSignUpInputBody, v: string) => void;
}

export default function ManagerStep({
  managerInputs,
  updMgr,
}: ManagerStepProps) {
  return (
    <div className="flex flex-col gap-4">
      <ManagerFormRow
        mgr={managerInputs[0]}
        index={0}
        onChange={updMgr}
        onRemove={function () {}}
        canRemove={false}
      />
    </div>
  );
}
