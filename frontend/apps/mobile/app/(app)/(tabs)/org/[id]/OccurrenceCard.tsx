import {
  Animated,
  Image,
  Pressable,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { useRef, useState } from "react";
import { useRouter } from "expo-router";
import type { EventOccurrence } from "@skillspark/api-client";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { ReservationModal } from "@/components/ReservationModal";
import { RatingSmiley } from "@/components/RatingSmiley";
import {
  formatAgeRange,
  formatPrice,
  formatTime,
} from "../../../../../utils/format";
import { useTranslation } from "react-i18next";

const BUTTON_ROW_HEIGHT = 52;

export function OccurrenceCard({
  occurrence,
  avgRating,
}: {
  occurrence: EventOccurrence;
  avgRating: number | null;
}) {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const [expanded, setExpanded] = useState(false);
  const [reservationVisible, setReservationVisible] = useState(false);
  const progress = useRef(new Animated.Value(0)).current;

  const toggle = () => {
    const toValue = expanded ? 0 : 1;
    setExpanded((prev) => !prev);
    Animated.timing(progress, {
      toValue,
      duration: 250,
      useNativeDriver: false,
    }).start();
  };

  const buttonsAnimStyle = {
    height: progress.interpolate({
      inputRange: [0, 1],
      outputRange: [0, BUTTON_ROW_HEIGHT],
    }),
    opacity: progress,
    overflow: "hidden" as const,
  };

  const chevronAnimStyle = {
    transform: [
      {
        rotate: progress.interpolate({
          inputRange: [0, 1],
          outputRange: ["0deg", "180deg"],
        }),
      },
    ],
  };

  const ageRange = formatAgeRange(
    occurrence.event.age_range_min,
    occurrence.event.age_range_max
  );
  const locationText = [
    occurrence.location?.district,
    occurrence.location?.province,
  ]
    .filter(Boolean)
    .join(", ");

  return (
    <>
      <Pressable
        onPress={toggle}
        className="mx-4 mb-3 rounded-2xl bg-white"
        style={{
          shadowColor: "#000",
          shadowOpacity: 0.1,
          shadowRadius: 8,
          shadowOffset: { width: 0, height: 10 },
          elevation: 3,
        }}
      >
        <View className="rounded-2xl overflow-hidden">
          <View className="flex-row p-3 gap-3">
            {/* Thumbnail */}
            <View
              className="w-[80px] h-[80px] rounded-xl overflow-hidden flex-shrink-0"
              style={{ backgroundColor: AppColors.imagePlaceholder }}
            >
              {occurrence.event.presigned_url ? (
                <Image
                  source={{ uri: occurrence.event.presigned_url }}
                  style={{ width: "100%", height: "100%" }}
                />
              ) : null}
            </View>

            {/* Info */}
            <View className="flex-1 justify-center gap-[3px]">
              <Text
                numberOfLines={1}
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: FontSizes.base,
                  color: AppColors.primaryText,
                }}
              >
                {occurrence.event.title}
              </Text>

              <View className="flex-row items-center gap-1">
                <View className="pr-2">
                  <RatingSmiley rating={avgRating} width={16} height={16} />
                </View>
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {avgRating ? avgRating : "~"} {translate("occurrence.smiles")}
                </Text>
              </View>

              {!!ageRange && (
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {ageRange}
                </Text>
              )}

              {!!locationText && (
                <Text
                  numberOfLines={1}
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                  }}
                >
                  {locationText}
                </Text>
              )}
            </View>

            {/* Time & Price */}
            <View className="items-end justify-center gap-1 flex-shrink-0">
              <Text
                style={{
                  fontFamily: FontFamilies.semiBold,
                  fontSize: FontSizes.sm,
                  color: AppColors.primaryText,
                }}
              >
                {formatTime(occurrence.start_time)}
              </Text>
              <Text
                style={{
                  fontFamily: FontFamilies.semiBold,
                  fontSize: FontSizes.base,
                  color: AppColors.primaryText,
                }}
              >
                {formatPrice(occurrence.price, occurrence.currency)}
              </Text>
            </View>
          </View>

          {/* Action buttons (animated expand/collapse) */}
          <Animated.View style={buttonsAnimStyle}>
            <View className="flex-row gap-3 px-3 pb-3">
              <TouchableOpacity
                onPress={() => router.push(`/event/${occurrence.event.id}`)}
                activeOpacity={0.7}
                className="flex-1 rounded-full py-2.5 items-center"
                style={{ backgroundColor: "#99C0EE" }}
              >
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.base,
                  }}
                >
                  {translate("occurrence.learnMore")}
                </Text>
              </TouchableOpacity>
              <TouchableOpacity
                onPress={() => setReservationVisible(true)}
                activeOpacity={0.7}
                className="flex-1 rounded-full py-2.5 items-center"
                style={{ backgroundColor: AppColors.checkboxSelected }}
              >
                <Text
                  style={{
                    fontFamily: FontFamilies.semiBold,
                    fontSize: FontSizes.base,
                    color: "#fff",
                  }}
                >
                  {translate("occurrence.reserve")}
                </Text>
              </TouchableOpacity>
            </View>
          </Animated.View>

          {/* Expand/collapse chevron (rotates 180° when expanded) */}
          <View className="items-center pb-2 -mt-1">
            <Animated.View style={chevronAnimStyle}>
              <IconSymbol name="chevron.down" size={16} color="#000000" />
            </Animated.View>
          </View>
        </View>
      </Pressable>

      <ReservationModal
        visible={reservationVisible}
        onClose={() => setReservationVisible(false)}
        occurrence={occurrence}
      />
    </>
  );
}
