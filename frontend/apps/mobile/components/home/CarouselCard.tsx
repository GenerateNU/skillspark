import Carousel from "react-native-reanimated-carousel";
import { View } from "react-native";
import { useState } from "react";
import { type Event } from "@skillspark/api-client";
import { AppColors } from "@/constants/theme";
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
  const [currentIndex, setCurrentIndex] = useState(0);

  return (
    <View>
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
        onSnapToItem={setCurrentIndex}
        renderItem={({ item }) => (
          <View style={{ flex: 1, paddingHorizontal: CARD_PADDING }}>
            <View style={{ flex: 1, borderRadius: 24, overflow: "hidden" }}>
              <DiscoverBanner event={item} />
            </View>
          </View>
        )}
      />
      <View
        style={{
          flexDirection: "row",
          justifyContent: "center",
          alignItems: "center",
          gap: 6,
          marginTop: 10,
        }}
      >
        {events.map((_, i) => (
          <View
            key={i}
            style={{
              width: i === currentIndex ? 16 : 6,
              height: 6,
              borderRadius: 3,
              backgroundColor:
                i === currentIndex ? AppColors.primaryText : AppColors.surfaceGray,
            }}
          />
        ))}
      </View>
    </View>
  );
}
