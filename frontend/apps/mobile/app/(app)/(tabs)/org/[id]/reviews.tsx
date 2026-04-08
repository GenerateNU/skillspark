import ReviewsScreen from "@/components/ReviewsScreen";
import {
  useGetReviewAggregateOrganization,
  useGetReviewByOrganizationId,
} from "@skillspark/api-client";

const useGetAggregate = (id: string) =>
  useGetReviewAggregateOrganization(id, { query: { enabled: !!id } });

const useGetReviews = (id: string) =>
  useGetReviewByOrganizationId(id, undefined, { query: { enabled: !!id } });

export default function OrgReviewsPage() {
  return <ReviewsScreen useGetAggregate={useGetAggregate} useGetReviews={useGetReviews} />;
}