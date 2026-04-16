import { useQuery, useQueryClient } from "@tanstack/react-query";

export type OrgScheduleFilters = {
  min_start_minutes?: number;
  max_start_minutes?: number;
  min_duration?: number;
  max_duration?: number;
  min_price?: number; // in cents
  max_price?: number; // in cents
  min_age?: number;
  max_age?: number;
  class_name?: string;
};

const EMPTY: OrgScheduleFilters = {};

export function useOrgScheduleFilters(orgId: string) {
  const queryClient = useQueryClient();
  const KEY = ["org-schedule-filters", orgId] as const;

  const { data: filters = EMPTY } = useQuery<OrgScheduleFilters>({
    queryKey: KEY,
    queryFn: () => EMPTY,
    staleTime: Infinity,
    gcTime: Infinity,
  });

  const activeCount = Object.values(filters).filter(
    (v) => v !== undefined,
  ).length;

  function setFilters(newFilters: OrgScheduleFilters) {
    queryClient.setQueryData<OrgScheduleFilters>(KEY, newFilters);
  }

  function clearFilters() {
    queryClient.setQueryData<OrgScheduleFilters>(KEY, EMPTY);
  }

  return { filters, setFilters, clearFilters, activeCount };
}
