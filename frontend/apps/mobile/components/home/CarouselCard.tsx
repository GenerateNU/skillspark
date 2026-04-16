import { FlatList, TouchableOpacity, View } from "react-native";
import { useRef, useState, useEffect, useCallback } from "react";
import { type EventOccurrence } from "@skillspark/api-client";
import { DiscoverBanner } from "./DiscoverBanner";
import { AppColors } from "@/constants/theme";

const PEEK_OFFSET = 40;
const ITEM_GAP = 12;

export default function CarouselCard({
  events,
  width,
  height,
}: {
  events: EventOccurrence[];
  width: number;
  height: number;
}) {
  const cardHeight = Math.min(height * 0.3, 240);
  const itemWidth = width - 2 * PEEK_OFFSET;
  const snapInterval = itemWidth + ITEM_GAP;

  const flatListRef = useRef<FlatList>(null);
  const [currentIndex, setCurrentIndex] = useState(0);
  const currentIndexRef = useRef(0);

  useEffect(() => {
    if (events.length <= 1) return;
    const interval = setInterval(() => {
      const next = (currentIndexRef.current + 1) % events.length;
      flatListRef.current?.scrollToOffset({
        offset: next * snapInterval,
        animated: true,
      });
      currentIndexRef.current = next;
      setCurrentIndex(next);
    }, 3500);
    return () => clearInterval(interval);
  }, [events.length, snapInterval]);

  const onMomentumScrollEnd = useCallback(
    (e: { nativeEvent: { contentOffset: { x: number } } }) => {
      const idx = Math.round(e.nativeEvent.contentOffset.x / snapInterval);
      currentIndexRef.current = idx;
      setCurrentIndex(idx);
    },
    [snapInterval],
  );

  return (
    <View>
      <FlatList
        ref={flatListRef}
        data={events}
        horizontal
        showsHorizontalScrollIndicator={false}
        snapToInterval={snapInterval}
        snapToAlignment="start"
        decelerationRate="fast"
        contentContainerStyle={{ paddingHorizontal: PEEK_OFFSET }}
        ItemSeparatorComponent={() => <View style={{ width: ITEM_GAP }} />}
        scrollEventThrottle={16}
        onMomentumScrollEnd={onMomentumScrollEnd}
        keyExtractor={(_, i) => String(i)}
        renderItem={({ item }) => (
          <View style={{ width: itemWidth, height: cardHeight }}>
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
        {events.map((_, index) => (
          <TouchableOpacity
            key={index}
            onPress={() => {
              flatListRef.current?.scrollToOffset({
                offset: index * snapInterval,
                animated: true,
              });
              currentIndexRef.current = index;
              setCurrentIndex(index);
            }}
            hitSlop={{ top: 8, bottom: 8, left: 4, right: 4 }}
            style={{
              width: currentIndex === index ? 16 : 6,
              height: 6,
              borderRadius: 3,
              backgroundColor:
                currentIndex === index
                  ? AppColors.primaryText
                  : AppColors.borderLight,
            }}
          />
        ))}
      </View>
    </View>
  );
}
