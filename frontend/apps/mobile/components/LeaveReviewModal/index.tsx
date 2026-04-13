import { useCreateReview } from "@skillspark/api-client";
import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  KeyboardAvoidingView,
  Modal,
  Platform,
  ScrollView,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { EventPreview } from "./EventPreview";
import { ReviewDoneScreen } from "./ReviewDoneScreen";
import { ReviewHeader } from "./ReviewHeader";
import { ReviewStep1 } from "./ReviewStep1";
import { ReviewStep2 } from "./ReviewStep2";

interface Props {
  visible: boolean;
  onClose: () => void;
  eventId: string;
  eventName: string;
  eventLocation: string;
  eventImageUrl?: string;
  registrationId: string;
  guardianId: string;
  initialRating?: number;
}

export function LeaveReviewModal({
  visible,
  onClose,
  eventId,
  eventName,
  eventLocation,
  eventImageUrl,
  registrationId,
  guardianId,
  initialRating,
}: Props) {
  const { t: translate } = useTranslation();
  const insets = useSafeAreaInsets();
  const queryClient = useQueryClient();

  const [step, setStep] = useState<1 | 2 | "done">(1);
  const [rating, setRating] = useState<number | null>(initialRating ?? null);
  const [selectedCategories, setSelectedCategories] = useState<string[]>([]);
  const [description, setDescription] = useState("");
  const [anonymous, setAnonymous] = useState(false);

  useEffect(() => {
    if (visible) {
      setStep(1);
      setRating(initialRating ?? null);
      setSelectedCategories([]);
      setDescription("");
      setAnonymous(false);
    }
  }, [visible, initialRating]);

  function toggleCategory(value: string) {
    setSelectedCategories((prev) =>
      prev.includes(value) ? prev.filter((c) => c !== value) : [...prev, value],
    );
  }

  const { mutate: createReview, isPending } = useCreateReview();

  function handleSubmit() {
    if (!rating) return;

    createReview(
      {
        data: {
          guardian_id: anonymous ? "" : guardianId,
          registration_id: registrationId,
          rating,
          categories: selectedCategories,
          description,
        },
      },
      {
        onSuccess: () => {
          queryClient.invalidateQueries({
            queryKey: [`/api/v1/review/event/${eventId}`],
          });
          queryClient.invalidateQueries({
            queryKey: [`/api/v1/review/event_aggregate/${eventId}`],
          });
          setStep("done");
        },
      },
    );
  }

  return (
    <Modal
      visible={visible}
      animationType="slide"
      presentationStyle="pageSheet"
      onRequestClose={onClose}
    >
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : undefined}
        style={{ flex: 1 }}
      >
        <View className="flex-1 bg-white" style={{ paddingTop: insets.top }}>
          {step === "done" ? (
            <ReviewDoneScreen
              rating={rating}
              onClose={onClose}
              translate={translate}
            />
          ) : (
            <>
              <ReviewHeader onClose={onClose} translate={translate} />
              <ScrollView
                className="flex-1 px-5 pt-5"
                showsVerticalScrollIndicator={false}
                contentContainerStyle={{ paddingBottom: 40 }}
                keyboardShouldPersistTaps="handled"
              >
                <EventPreview
                  eventName={eventName}
                  eventLocation={eventLocation}
                  eventImageUrl={eventImageUrl}
                />
                {step === 1 ? (
                  <ReviewStep1
                    rating={rating}
                    setRating={setRating}
                    selectedCategories={selectedCategories}
                    toggleCategory={toggleCategory}
                    translate={translate}
                    onNext={() => setStep(2)}
                  />
                ) : (
                  <ReviewStep2
                    description={description}
                    setDescription={setDescription}
                    anonymous={anonymous}
                    setAnonymous={setAnonymous}
                    isPending={isPending}
                    onBack={() => setStep(1)}
                    onSubmit={handleSubmit}
                    translate={translate}
                  />
                )}
              </ScrollView>
            </>
          )}
        </View>
      </KeyboardAvoidingView>
    </Modal>
  );
}
