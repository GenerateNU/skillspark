import { useQuery, useQueryClient } from "@tanstack/react-query";
import { type GetAllEventOccurrencesParams } from "@skillspark/api-client";

const FILTERS_KEY = ["filters"] as const;
const EMPTY: GetAllEventOccurrencesParams = {};

export function useFilters() {
  const queryClient = useQueryClient();

  const { data: filters = EMPTY } = useQuery<GetAllEventOccurrencesParams>({
    queryKey: FILTERS_KEY,
    queryFn: () => EMPTY,
    staleTime: Infinity,
    gcTime: Infinity,
  });

  const hasActiveFilters = Object.values(filters).some((v) => v !== undefined);

  function setFilters(newFilters: GetAllEventOccurrencesParams) {
    queryClient.setQueryData<GetAllEventOccurrencesParams>(FILTERS_KEY, newFilters);
  }

  function clearFilters() {
    queryClient.setQueryData<GetAllEventOccurrencesParams>(FILTERS_KEY, EMPTY);
  }

  return { filters, setFilters, clearFilters, hasActiveFilters };
}
