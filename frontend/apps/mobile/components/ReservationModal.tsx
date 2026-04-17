import { useRef, useState } from "react";
import {
  ActivityIndicator,
  Modal,
  Pressable,
  ScrollView,
  Text,
  TouchableOpacity,
  View,
} from "react-native";
import { TermsAndConditionsModal } from "./TermsAndConditionsModal";
import { ReservationSuccessModal } from "./ReservationSuccessModal";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import {
  createRegistrationResponseError,
  getGetRegistrationsByGuardianIdQueryKey,
  useCreateRegistration,
  useGetChildrenByGuardianId,
  type Child,
  type EventOccurrence,
} from "@skillspark/api-client";
import { useQueryClient } from "@tanstack/react-query";
import { useTranslation } from "react-i18next";
import { useAuthContext } from "@/hooks/use-auth-context";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { ChildAvatar } from "./ChildAvatar";
import { EventPreviewSection } from "./EventPreviewSection";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";

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
  const queryClient = useQueryClient();

  const isSuccessTransition = useRef(false);
  const [selectedChildId, setSelectedChildId] = useState<string | null>(null);
  const [reservationError, setReservationError] = useState<string | null>(null);
  const [errorDetails, setErrorDetails] = useState<string[]>([]);
  const [errorExpanded, setErrorExpanded] = useState(false);
  const [termsAccepted, setTermsAccepted] = useState(false);
  const [showTermsModal, setShowTermsModal] = useState(false);
  const [showSuccessModal, setShowSuccessModal] = useState(false);

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
    setErrorDetails([]);
    setErrorExpanded(false);
    createRegistration(
      {
        data: {
          child_id: selectedChildId,
          event_occurrence_id: occurrence.id,
          guardian_id: guardianId,
          status: "registered",
        },
      },
      {
        onSuccess: () => {
          queryClient.invalidateQueries({
            queryKey: getGetRegistrationsByGuardianIdQueryKey(guardianId!),
          });
          isSuccessTransition.current = true;
          setShowSuccessModal(true);
          onClose();
        },
        onError: (error: unknown) => {
          try {
            const typedError = error as createRegistrationResponseError;
            const errorMsgs = (typedError.data.errors || [])
              .map((e) => e.message)
              .filter((m): m is string => !!m);
            setErrorDetails(errorMsgs);
            setReservationError(translate("reservation.paymentFailed"));
          } catch {
            setReservationError(translate("reservation.paymentFailed"));
          }
        },
      }
    );
  }

  function resetState() {
    setSelectedChildId(null);
    setReservationError(null);
    setErrorDetails([]);
    setErrorExpanded(false);
    setTermsAccepted(false);
  }

  function handleClose() {
    resetState();
    if (!isSuccessTransition.current) {
      onClose();
    }
    isSuccessTransition.current = false;
  }

  function handleSuccessClose() {
    setShowSuccessModal(false);
    onClose();
  }

  return (
    <>
      <Modal
        visible={visible}
        transparent
        animationType="fade"
        onRequestClose={handleClose}
      >
        <TermsAndConditionsModal
          visible={showTermsModal}
          onClose={() => setShowTermsModal(false)}
        />
        <ReservationSuccessModal
          visible={showSuccessModal}
          occurrence={occurrence}
          onClose={handleSuccessClose}
        />
        <Pressable
          style={{
            flex: 1,
            backgroundColor: "rgba(0,0,0,0.5)",
            justifyContent: "center",
            alignItems: "center",
          }}
          onPress={handleClose}
        >
          <Pressable
            onPress={() => {}}
            style={{
              width: "90%",
              maxHeight: "85%",
              backgroundColor: "#fff",
              borderRadius: 20,
              overflow: "hidden",
            }}
          >
            {/* X button */}
            <TouchableOpacity
              onPress={handleClose}
              style={{
                position: "absolute",
                top: 16,
                right: 16,
                zIndex: 10,
                padding: 4,
              }}
              hitSlop={8}
            >
              <MaterialIcons
                name="close"
                size={24}
                color={AppColors.primaryText}
              />
            </TouchableOpacity>

            <ScrollView
              showsVerticalScrollIndicator={false}
              contentContainerStyle={{
                padding: 20,
                paddingTop: 24,
                paddingBottom: insets.bottom + 24,
              }}
            >
              <EventPreviewSection occurrence={occurrence} />

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
                      <Pressable
                        key={child.id}
                        onPress={() =>
                          setSelectedChildId(
                            selectedChildId === child.id ? null : child.id
                          )
                        }
                        style={
                          selectedChildId === child.id
                            ? {
                                borderRadius: 999,
                                borderWidth: 2,
                                borderColor: AppColors.primaryBlue,
                              }
                            : undefined
                        }
                      >
                        <ChildAvatar
                          name={child.name}
                          avatarFace={child.avatar_face}
                          avatarBackground={child.avatar_background}
                        />
                      </Pressable>
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
                <TouchableOpacity onPress={() => setShowTermsModal(true)}>
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
                </TouchableOpacity>
                <TouchableOpacity
                  onPress={() => setTermsAccepted((prev) => !prev)}
                  activeOpacity={0.7}
                  className="flex-row items-center gap-2 mt-3 self-center"
                >
                  <View
                    style={{
                      width: 20,
                      height: 20,
                      borderRadius: 4,
                      borderWidth: 2,
                      borderColor: termsAccepted
                        ? AppColors.checkboxSelected
                        : AppColors.borderLight,
                      backgroundColor: termsAccepted
                        ? AppColors.checkboxSelected
                        : "transparent",
                      alignItems: "center",
                      justifyContent: "center",
                    }}
                  >
                    {termsAccepted && (
                      <Text
                        style={{ color: "#fff", fontSize: 12, lineHeight: 14 }}
                      >
                        ✓
                      </Text>
                    )}
                  </View>
                  <Text
                    className="text-xs"
                    style={{
                      fontFamily: FontFamilies.regular,
                      color: AppColors.primaryText,
                    }}
                  >
                    {translate("reservation.termsAgree")}
                  </Text>
                </TouchableOpacity>
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
                  {errorDetails.length > 0 && (
                    <>
                      <TouchableOpacity
                        onPress={() => setErrorExpanded((prev) => !prev)}
                        className="mt-2 items-center"
                      >
                        <Text
                          className="text-xs"
                          style={{
                            fontFamily: FontFamilies.semiBold,
                            color: "#DC2626",
                            textDecorationLine: "underline",
                          }}
                        >
                          {errorExpanded ? "See less" : "See more"}
                        </Text>
                      </TouchableOpacity>
                      {errorExpanded && (
                        <View className="mt-2 gap-1">
                          {errorDetails.map((detail, i) => (
                            <Text
                              key={i}
                              className="text-xs"
                              style={{
                                fontFamily: FontFamilies.regular,
                                color: "#DC2626",
                              }}
                            >
                              {"\u2022"} {detail}
                            </Text>
                          ))}
                        </View>
                      )}
                    </>
                  )}
                </View>
              )}

              {/* Reserve button */}
              <TouchableOpacity
                onPress={handleReserve}
                activeOpacity={0.8}
                disabled={!selectedChildId || !termsAccepted || isPending}
                className="rounded-2xl py-4 items-center"
                style={{
                  backgroundColor:
                    selectedChildId && termsAccepted
                      ? "#000000"
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
            </ScrollView>
          </Pressable>
        </Pressable>
      </Modal>
    </>
  );
}
