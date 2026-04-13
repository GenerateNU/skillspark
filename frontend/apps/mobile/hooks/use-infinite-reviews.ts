import { useInfiniteQuery } from "@tanstack/react-query";
import {
  getReviewByEventId,
  getReviewByOrganizationId,
} from "@skillspark/api-client";

const PAGE_SIZE = 10;

export function useInfiniteReviewsByEventId(eventId: string | undefined) {
  return useInfiniteQuery({
    queryKey: ["reviews", "event", "infinite", eventId],
    queryFn: ({ pageParam }) =>
      getReviewByEventId(eventId!, { page: pageParam, page_size: PAGE_SIZE }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      const items = lastPage.data;
      if (Array.isArray(items) && items.length === PAGE_SIZE) {
        return allPages.length + 1;
      }
      return undefined;
    },
    enabled: !!eventId,
  });
}

export function useInfiniteReviewsByOrganizationId(orgId: string | undefined) {
  return useInfiniteQuery({
    queryKey: ["reviews", "org", "infinite", orgId],
    queryFn: ({ pageParam }) =>
      getReviewByOrganizationId(orgId!, {
        page: pageParam,
        page_size: PAGE_SIZE,
      }),
    initialPageParam: 1,
    getNextPageParam: (lastPage, allPages) => {
      const items = lastPage.data;
      if (Array.isArray(items) && items.length === PAGE_SIZE) {
        return allPages.length + 1;
      }
      return undefined;
    },
    enabled: !!orgId,
  });
}
