import { RATING_OPTIONS } from "@/constants/ratings";
import { Image } from "expo-image";

type RatingSmileyProps = {
  /** The numeric rating (1–5). Pass null when no ratings exist — renders nothing. */
  rating: number | null;
  width: number;
  height: number;
};

/** Renders the smiley image that corresponds to a given rating (1–5). */
export function RatingSmiley({ rating, width, height }: RatingSmileyProps) {
  const rounded = rating && Math.min(5, Math.max(1, Math.round(rating)));
  const option = RATING_OPTIONS.find((o) => o.rating === rounded);
  if (!option) return null;
  return <Image source={option.image} style={{ width, height }} />;
}
