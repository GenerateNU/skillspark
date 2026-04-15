import Carousel from "react-native-reanimated-carousel";
import { View } from "react-native";
import { AppColors } from "@/constants/theme";
import { type EventOccurrence } from "@skillspark/api-client";
import { DiscoverBanner } from "./DiscoverBanner";

export default function CarouselCard({
  events,
  width,
  height,
}: {
  events: EventOccurrence[];
  width: number;
  height: number;
}) {
  return (
    <Carousel
      width={width / 1.1}
      height={height / 4}
      data={events}
      loop
      autoPlay
      autoPlayInterval={3000}
      mode="vertical-stack"
      modeConfig={{
        snapDirection: "left",
        stackInterval: 0,
      }}
      renderItem={({ item }) => (
        <View className="rounded-3xl overflow-hidden flex-1">
          <DiscoverBanner event={item} />
        </View>
      )}
      style={{ alignSelf: "center" }}
    />
  );
}
