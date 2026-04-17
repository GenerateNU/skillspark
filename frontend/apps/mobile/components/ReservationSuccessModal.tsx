import { Image } from "expo-image";
import { Modal, Text, TouchableOpacity, View } from "react-native";
import { AppColors, FontFamilies } from "@/constants/theme";
import { formatModalTime } from "@/utils/format";
import type { EventOccurrence } from "@skillspark/api-client";
import { useTranslation } from "react-i18next";

interface ReservationSuccessModalProps {
  visible: boolean;
  occurrence: EventOccurrence;
  onClose: () => void;
}

export function ReservationSuccessModal({
  visible,
  occurrence,
  onClose,
}: ReservationSuccessModalProps) {
  const { t: translate } = useTranslation();
  const timeLabel = translate("occurrence.classTime", {
    time: formatModalTime(occurrence.start_time),
  });

  return (
    <Modal
      visible={visible}
      transparent
      animationType="fade"
      onRequestClose={onClose}
    >
      <View
        className="flex-1 items-center justify-center px-8"
        style={{ backgroundColor: "rgba(0,0,0,0.55)" }}
      >
        <View
          className="w-full rounded-3xl px-6 pt-6 pb-6 items-center"
          style={{ backgroundColor: "#fff" }}
        >
          {/* Event image */}
          <View
            className="w-[200px] h-[200px] rounded-2xl overflow-hidden mb-5"
            style={{ backgroundColor: AppColors.imagePlaceholder }}
          >
            {occurrence.event.presigned_url ? (
              <Image
                source={{ uri: occurrence.event.presigned_url }}
                style={{ width: "100%", height: "100%" }}
                contentFit="cover"
              />
            ) : null}
          </View>

          {/* Completed title */}
          <Text
            className="text-2xl mb-3 text-center"
            style={{
              fontFamily: FontFamilies.bold,
              color: AppColors.primaryText,
            }}
          >
            {translate("reservation.completed")}
          </Text>

          {/* Description */}
          {!!occurrence.event.description && (
            <Text
              className="text-sm text-center leading-5 mb-2"
              style={{
                fontFamily: FontFamilies.regular,
                color: AppColors.secondaryText,
              }}
            >
              {occurrence.event.description}
            </Text>
          )}

          {/* Time */}
          <Text
            className="text-sm text-center mb-4"
            style={{
              fontFamily: FontFamilies.regular,
              color: AppColors.mutedText,
            }}
          >
            {timeLabel}
          </Text>

          {/* See you soon */}
          <Text
            className="text-base text-center mb-6"
            style={{
              fontFamily: FontFamilies.semiBold,
              color: AppColors.mutedText,
            }}
          >
            {translate("reservation.seeYouSoon")}
          </Text>

          {/* Close button */}
          <TouchableOpacity
            onPress={onClose}
            activeOpacity={0.8}
            className="w-full rounded-2xl py-4 items-center"
            style={{ backgroundColor: AppColors.primaryText }}
          >
            <Text
              className="text-white text-base"
              style={{ fontFamily: FontFamilies.bold }}
            >
              {translate("common.close")}
            </Text>
          </TouchableOpacity>
        </View>
      </View>
    </Modal>
  );
}
