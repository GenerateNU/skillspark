import { useQuery, useQueryClient } from "@tanstack/react-query";

export type OrgSortOption = "distance" | "rating" | "members";

interface OrgMapFilters {
  sort_by?: OrgSortOption;
}

const KEY = ["org-map-filters"] as const;
const EMPTY: OrgMapFilters = {};

export function useOrgMapFilters() {
  const queryClient = useQueryClient();

  const { data: filters = EMPTY } = useQuery<OrgMapFilters>({
    queryKey: KEY,
    queryFn: () => EMPTY,
    staleTime: Infinity,
    gcTime: Infinity,
  });

  function setFilters(next: OrgMapFilters) {
    queryClient.setQueryData<OrgMapFilters>(KEY, next);
  }

  function clearFilters() {
    queryClient.setQueryData<OrgMapFilters>(KEY, EMPTY);
  }

  return { filters, setFilters, clearFilters };
}
