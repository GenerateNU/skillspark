import Carousel from "react-native-reanimated-carousel";
import { View } from "react-native";
import { type Event } from "@skillspark/api-client";
import { DiscoverBanner } from "./DiscoverBanner";

const CARD_PADDING = 22;

export default function CarouselCard({
  events,
  width,
  height,
}: {
  events: Event[];
  width: number;
  height: number;
}) {
  const cardHeight = Math.min(height * 0.3, 240);

  return (
    <Carousel
      width={width}
      height={cardHeight}
      data={events}
      loop
      autoPlay
      autoPlayInterval={3500}
      mode="parallax"
      modeConfig={{
        parallaxScrollingScale: 0.9,
        parallaxScrollingOffset: 100,
        parallaxAdjacentItemScale: 0.78,
      }}
      renderItem={({ item }) => (
        <View style={{ flex: 1, paddingHorizontal: CARD_PADDING }}>
          <View style={{ flex: 1, borderRadius: 24, overflow: "hidden" }}>
            <DiscoverBanner event={item} />
          </View>
        </View>
      )}
    />
  );
}
