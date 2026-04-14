import { Image } from "expo-image";
import { useCallback, useEffect, useRef, useState } from "react";
import {
  ActivityIndicator,
  Pressable,
  StyleSheet,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import Animated, {
  useAnimatedReaction,
  useAnimatedStyle,
  useSharedValue,
} from "react-native-reanimated";
import type { SharedValue } from "react-native-reanimated";
import {
  BottomSheetModal,
  BottomSheetScrollView,
  type BottomSheetBackdropProps,
} from "@gorhom/bottom-sheet";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import {
  useCreateRegistration,
  useGetChildrenByGuardianId,
  type Child,
  type EventOccurrence,
} from "@skillspark/api-client";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { ChildAvatar } from "./ChildAvatar";
import { EventPreviewSection } from "./EventPreviewSection";
import { StaticBackdrop } from "./StaticBackground";

interface ReservationModalProps {
  visible: boolean;
  onClose: () => void;
  occurrence: EventOccurrence;
}

export function ReservationModal({
  visible,
  onClose,
  occurrence,
}: ReservationModalProps) {
  const insets = useSafeAreaInsets();
  const { t: translate } = useTranslation();
  const { guardianId } = useAuthContext();

  const sheetRef = useRef<BottomSheetModal>(null);
  const [selectedChildId, setSelectedChildId] = useState<string | null>(null);
  const [step, setStep] = useState<"select" | "done">("select");
  const [reservationError, setReservationError] = useState<string | null>(null);

  useEffect(() => {
    if (visible) {
      sheetRef.current?.present();
    } else {
      sheetRef.current?.dismiss();
    }
  }, [visible]);

  const { data: childrenResp, isLoading: childrenLoading } =
    useGetChildrenByGuardianId(guardianId!, {
      query: { enabled: !!guardianId && visible },
    });

  const children: Child[] =
    childrenResp?.status === 200 ? (childrenResp.data as Child[]) : [];

  const { mutate: createRegistration, isPending } = useCreateRegistration();

  function handleReserve() {
    if (!selectedChildId || !guardianId) return;
    setReservationError(null);
    createRegistration(
      {
        data: {
          child_id: selectedChildId,
          event_occurrence_id: occurrence.id,
          guardian_id: guardianId,
          payment_method_id: "",
          status: "registered",
        },
      },
      {
        onSuccess: () => {
          setStep("done");
        },
        onError: (error: unknown) => {
          const message =
            error instanceof Error
              ? error.message
              : translate("reservation.paymentFailed");
          setReservationError(message);
        },
      }
    );
  }

  function handleClose() {
    setStep("select");
    setSelectedChildId(null);
    setReservationError(null);
    onClose();
  }

  const renderBackdrop = useCallback(
    (props: BottomSheetBackdropProps) => (
      <StaticBackdrop
        animatedIndex={props.animatedIndex}
        style={props.style}
        onPress={handleClose}
      />
    ),
    // handleClose intentionally omitted — state setters are stable and
    // onClose identity changing should not recreate the backdrop component.
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  return (
    <BottomSheetModal
      ref={sheetRef}
      snapPoints={["78%"]}
      enablePanDownToClose
      overDragResistanceFactor={0}
      onDismiss={handleClose}
      backdropComponent={renderBackdrop}
      handleIndicatorStyle={{ backgroundColor: AppColors.borderLight }}
      backgroundStyle={{ borderTopLeftRadius: 20, borderTopRightRadius: 20 }}
    >
      <BottomSheetScrollView
        className="px-5"
        showsVerticalScrollIndicator={false}
        contentContainerStyle={{ paddingBottom: insets.bottom + 24 }}
      >
        <EventPreviewSection
          occurrence={occurrence}
          titleOverride={
            step === "done" ? translate("reservation.completed") : undefined
          }
        />

        {step === "select" ? (
          <>
            {/* Child selection */}
            <View className="mb-6">
              <Text
                className="text-sm mb-3"
                style={{
                  fontFamily: FontFamilies.semiBold,
                  color: AppColors.secondaryText,
                }}
              >
                {translate("reservation.selectChildLabel")}
              </Text>
              {childrenLoading ? (
                <ActivityIndicator size="small" />
              ) : children.length === 0 ? (
                <Text
                  style={{
                    fontFamily: FontFamilies.regular,
                    fontSize: FontSizes.sm,
                    color: AppColors.mutedText,
                  }}
                >
                  {translate("reservation.noChildren")}
                </Text>
              ) : (
                <View className="flex-row gap-4 flex-wrap">
                  {children.map((child) => (
                    <ChildAvatar
                      key={child.id}
                      child={child}
                      selected={selectedChildId === child.id}
                      onPress={() =>
                        setSelectedChildId(
                          selectedChildId === child.id ? null : child.id
                        )
                      }
                    />
                  ))}
                </View>
              )}
            </View>

            {/* Terms */}
            <View className="mb-8">
              <Text
                className="text-xs text-center leading-5"
                style={{
                  fontFamily: FontFamilies.regular,
                  color: AppColors.mutedText,
                }}
              >
                {translate("reservation.termsNote")}
              </Text>
              <Text
                className="text-xs text-center mt-1"
                style={{
                  fontFamily: FontFamilies.semiBold,
                  color: AppColors.primaryText,
                  textDecorationLine: "underline",
                }}
              >
                {translate("reservation.terms")}
              </Text>
            </View>

            {/* Error message */}
            {!!reservationError && (
              <View
                className="rounded-xl px-4 py-3 mb-4"
                style={{ backgroundColor: "#FEE2E2" }}
              >
                <Text
                  className="text-sm text-center"
                  style={{
                    fontFamily: FontFamilies.regular,
                    color: "#DC2626",
                  }}
                >
                  {reservationError}
                </Text>
              </View>
            )}

            {/* Reserve button */}
            <TouchableOpacity
              onPress={handleReserve}
              activeOpacity={0.8}
              disabled={!selectedChildId || isPending}
              className="rounded-2xl py-4 items-center"
              style={{
                backgroundColor: selectedChildId
                  ? AppColors.checkboxSelected
                  : AppColors.borderLight,
              }}
            >
              {isPending ? (
                <ActivityIndicator size="small" color="#fff" />
              ) : (
                <Text
                  className="text-white text-base"
                  style={{ fontFamily: FontFamilies.bold }}
                >
                  {translate("reservation.payAndReserve")}
                </Text>
              )}
            </TouchableOpacity>
          </>
        ) : (
          <>
            {/* Completed state */}
            <Text
              className="text-base text-center mb-8"
              style={{
                fontFamily: FontFamilies.semiBold,
                color: AppColors.mutedText,
              }}
            >
              {translate("reservation.seeYouSoon")}
            </Text>
            <TouchableOpacity
              onPress={handleClose}
              activeOpacity={0.8}
              className="rounded-2xl py-4 items-center"
              style={{ backgroundColor: AppColors.checkboxSelected }}
            >
              <Text
                className="text-white text-base"
                style={{ fontFamily: FontFamilies.bold }}
              >
                {translate("common.close")}
              </Text>
            </TouchableOpacity>
          </>
        )}
      </BottomSheetScrollView>
    </BottomSheetModal>
  );
}
