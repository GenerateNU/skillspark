import ReviewsScreen from "@/components/ReviewsScreen";
import { useInfiniteReviewsByOrganizationId } from "@/hooks/use-infinite-reviews";
import { useGetReviewAggregateOrganization } from "@skillspark/api-client";

const useGetAggregate = (id: string) =>
  useGetReviewAggregateOrganization(id, { query: { enabled: !!id } });

const useGetReviews = (id: string) => useInfiniteReviewsByOrganizationId(id);

export default function OrgReviewsPage() {
  return (
    <ReviewsScreen
      useGetAggregate={useGetAggregate}
      useGetReviews={useGetReviews}
    />
  );
}
