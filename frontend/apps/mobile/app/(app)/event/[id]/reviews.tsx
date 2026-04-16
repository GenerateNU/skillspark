import ReviewsScreen from "@/components/ReviewsScreen";
import {
  useGetReviewAggregate,
} from "@skillspark/api-client";
import { useLocalSearchParams } from "expo-router";

const useGetAggregate = (id: string) =>
  useGetReviewAggregate(id, { query: { enabled: !!id } });

export default function EventReviewsPage() {
  const { canReview, occurrenceId, eventName, eventLocation, eventImageUrl } =
    useLocalSearchParams<{
      canReview?: string;
      occurrenceId?: string;
      eventName?: string;
      eventLocation?: string;
      eventImageUrl?: string;
    }>();

  return (
    <ReviewsScreen
      useGetAggregate={useGetAggregate}
      canReview={canReview === "true"}
      occurrenceId={occurrenceId}
      eventName={eventName}
      eventLocation={eventLocation}
      eventImageUrl={eventImageUrl}
    />
  );
}
