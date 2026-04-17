import { useInfiniteQuery } from "@tanstack/react-query";
import { getSavedByGuardianId } from "@skillspark/api-client";

const PAGE_SIZE = 10;

export const infiniteSavedQueryKey = (guardianId: string | undefined) =>
  ["saved", "infinite", guardianId] as const;

export function useInfiniteSavedByGuardianId(guardianId: string | undefined) {
  return useInfiniteQuery({
    queryKey: infiniteSavedQueryKey(guardianId),
    queryFn: ({ pageParam, signal }) =>
      getSavedByGuardianId(
        guardianId!,
        { page: pageParam, page_size: PAGE_SIZE },
        { signal },
      ),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      const items = lastPage.data;
      if (Array.isArray(items) && items.length === PAGE_SIZE) {
        return allPages.length + 1;
      }
      return undefined;
    },
    enabled: !!guardianId,
  });
}
