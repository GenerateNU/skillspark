import { RATING_OPTIONS } from "@/constants/ratings";

export type RatingOption = typeof RATING_OPTIONS[number];

/** Returns the RATING_OPTIONS entry matching the rounded rating, falling back to the null-rating option. */
export function getRatingOption(rating: number | null | undefined): RatingOption {
  if (rating != null) {
    return (
      RATING_OPTIONS.find((r) => r.rating === Math.round(rating)) ??
      RATING_OPTIONS.find((r) => r.rating === null)!
    );
  }
  return RATING_OPTIONS.find((r) => r.rating === null)!;
}
