import ReviewsScreen from "@/components/ReviewsScreen";
import {
  useGetReviewAggregate,
  useGetReviewByEventId,
} from "@skillspark/api-client";
import { useLocalSearchParams } from "expo-router";

const useGetAggregate = (id: string) =>
  useGetReviewAggregate(id, { query: { enabled: !!id } });

const useGetReviews = (id: string) =>
  useGetReviewByEventId(id, undefined, { query: { enabled: !!id } });

export default function EventReviewsPage() {
  const { canReview } = useLocalSearchParams<{ canReview?: string }>();

  return (
    <ReviewsScreen
      useGetAggregate={useGetAggregate}
      useGetReviews={useGetReviews}
      canReview={canReview === "true"}
    />
  );
}