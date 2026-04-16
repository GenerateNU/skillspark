import { useMemo } from "react";
import {
  useGetAllEventOccurrences,
  useGetRegistrationsByGuardianId,
  useGetChildrenByGuardianId,
  useGetReviewByGuardianId,
  type EventOccurrence,
  type Registration,
  type Child,
  type Review,
} from "@skillspark/api-client";
import { useAuthContext } from "@/hooks/use-auth-context";
import { extractResponseData, arrayToMap } from "@/utils/format";

export function useActivityData() {
  const { guardianId } = useAuthContext();
  const { data: registrationsResp } = useGetRegistrationsByGuardianId(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const registrations: Registration[] = useMemo(() => {
    const d = registrationsResp as unknown as
      | { data: { registrations: Registration[] } }
      | undefined;
    return d?.data?.registrations ?? [];
  }, [registrationsResp]);

  const { data: occurrencesResp } = useGetAllEventOccurrences({ limit: 100 });
  const eventOccurrencesMap: Record<string, EventOccurrence> = useMemo(
    () => arrayToMap(extractResponseData<EventOccurrence>(occurrencesResp)),
    [occurrencesResp],
  );

  const { data: childrenResp } = useGetChildrenByGuardianId(guardianId!, {
    query: { enabled: !!guardianId },
  });
  const children: Child[] = useMemo(
    () => extractResponseData<Child>(childrenResp),
    [childrenResp],
  );

  const childMap = useMemo(() => arrayToMap(children), [children]);

  const { data: guardianReviewsResp } = useGetReviewByGuardianId(guardianId!, undefined, {
    query: { enabled: !!guardianId },
  });
  const reviewedEventIds = useMemo(() => {
    const list =
      guardianReviewsResp?.status === 200 ? (guardianReviewsResp.data as Review[]) : [];
    return new Set(list.map((r) => r.event_id));
  }, [guardianReviewsResp]);

  return {
    guardianId,
    registrations,
    eventOccurrencesMap,
    children,
    childMap,
    reviewedEventIds,
  };
}
