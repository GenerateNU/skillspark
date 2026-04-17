import { useInfiniteQuery } from "@tanstack/react-query";
import {
  getReviewByEventId,
  getEventReviewsForOrganization,
  type GetReviewByEventIdSortBy,
  type GetEventReviewsForOrganizationSortBy,
} from "@skillspark/api-client";

const PAGE_SIZE = 10;

export function useInfiniteReviewsByEventId(
  eventId: string | undefined,
  sortBy?: GetReviewByEventIdSortBy,
) {
  return useInfiniteQuery({
    queryKey: ["reviews", "event", "infinite", eventId, sortBy],
    queryFn: ({ pageParam }) =>
      getReviewByEventId(eventId!, {
        page: pageParam,
        page_size: PAGE_SIZE,
        sort_by: sortBy,
      }),
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

export function useInfiniteEventReviewsForOrganization(
  orgId: string | undefined,
  sortBy?: GetEventReviewsForOrganizationSortBy,
) {
  return useInfiniteQuery({
    queryKey: ["reviews", "org", "infinite", orgId, sortBy],
    queryFn: ({ pageParam }) =>
      getEventReviewsForOrganization(orgId!, {
        page: pageParam,
        page_size: PAGE_SIZE,
        sort_by: sortBy,
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
