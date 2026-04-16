export const RATING_OPTIONS = [
  {
    rating: 5,
    image: require("@/assets/images/ratings/great.png"),
    labelKey: "review.excellent",
  },
  {
    rating: 4,
    image: require("@/assets/images/ratings/good.png"),
    labelKey: "review.good",
  },
  {
    rating: 3,
    image: require("@/assets/images/ratings/okay.png"),
    labelKey: "review.okay",
  },
  {
    rating: 2,
    image: require("@/assets/images/ratings/bad.png"),
    labelKey: "review.bad",
  },
  {
    rating: 1,
    image: require("@/assets/images/ratings/terrible.png"),
    labelKey: "review.terrible",
  },
  {
    rating: null,
    image: require("@/assets/images/ratings/noreview.png"),
    labelKey: "review.noReviews",
  },
];
