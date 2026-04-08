import { RATING_OPTIONS } from "@/constants/ratings";
import { AppColors } from "@/constants/theme";
import { useCreateReview } from "@skillspark/api-client";
import { Image } from "expo-image";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { useQueryClient } from "@tanstack/react-query";
import { useEffect, useState } from "react";
import { useTranslation } from "react-i18next";
import {
  ActivityIndicator,
  KeyboardAvoidingView,
  Modal,
  Platform,
  ScrollView,
  Switch,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

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

const CONTENT_OPTIONS = [
  { value: "fun", labelKey: "review.fun" },
  { value: "engaging", labelKey: "review.engaging" },
  { value: "interesting", labelKey: "review.interesting" },
  { value: "informative", labelKey: "review.informative" },
  { value: "interactive", labelKey: "review.interactive" },
  { value: "immersive", labelKey: "review.immersive" },
  { value: "educational", labelKey: "review.educational" },
  { value: "insightful", labelKey: "review.insightful" },
  { value: "well structured", labelKey: "review.wellStructured" },
];

const DIFFICULTY_OPTIONS = [
  { value: "beginner friendly", labelKey: "review.beginnerFriendly" },
  { value: "intermediate", labelKey: "review.intermediate" },
  { value: "too easy", labelKey: "review.tooEasy" },
  { value: "advanced", labelKey: "review.advanced" },
  { value: "challenging", labelKey: "review.challenging" },
];

const INSTRUCTOR_OPTIONS = [
  { value: "punctual", labelKey: "review.punctual" },
  { value: "clear instruction", labelKey: "review.clearInstruction" },
  { value: "welcoming", labelKey: "review.welcoming" },
  { value: "inclusive", labelKey: "review.inclusive" },
  { value: "collaborative", labelKey: "review.collaborative" },
];

const VALUE_OPTIONS = [
  { value: "affordable", labelKey: "review.affordable" },
  { value: "fair", labelKey: "review.fair" },
  { value: "expensive", labelKey: "review.expensive" },
];

function CategoryPill({
  label,
  selected,
  onPress,
}: {
  label: string;
  selected: boolean;
  onPress: () => void;
}) {
  return (
    <TouchableOpacity
      onPress={onPress}
      className="px-4 py-2 rounded-full border"
      style={{
        borderColor: selected ? AppColors.primaryText : AppColors.borderLight,
        backgroundColor: selected ? AppColors.primaryText : "transparent",
      }}
    >
      <Text
        className="text-sm"
        style={{ color: selected ? "#fff" : AppColors.secondaryText }}
      >
        {label}
      </Text>
    </TouchableOpacity>
  );
}

function CategorySection({
  title,
  options,
  selected,
  onToggle,
  translate,
}: {
  title: string;
  options: { value: string; labelKey: string }[];
  selected: string[];
  onToggle: (value: string) => void;
  translate: (key: string) => string;
}) {
  return (
    <>
      <Text
        className="text-base font-nunito-bold mb-3"
        style={{ color: AppColors.primaryText }}
      >
        {title}
      </Text>
      <View className="flex-row flex-wrap gap-2 mb-5">
        {options.map(({ value, labelKey }) => (
          <CategoryPill
            key={value}
            label={translate(labelKey)}
            selected={selected.includes(value)}
            onPress={() => onToggle(value)}
          />
        ))}
      </View>
    </>
  );
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
      prev.includes(value) ? prev.filter((c) => c !== value) : [...prev, value]
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

  const submittedEmoji = RATING_OPTIONS.find((r) => r.rating === rating);

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
            <View className="flex-1 items-center justify-center px-8">
              <Text
                className="text-3xl font-nunito-bold mb-6"
                style={{ color: AppColors.primaryText }}
              >
                {translate("review.thankYou")}
              </Text>
              {submittedEmoji && (
                <Image
                  source={submittedEmoji.image}
                  style={{ width: 80, height: 80, marginBottom: 24 }}
                />
              )}
              <Text
                className="text-base text-center mb-10"
                style={{ color: AppColors.secondaryText, lineHeight: 24 }}
              >
                {translate("review.thankYouMessage")}
              </Text>
              <TouchableOpacity
                onPress={onClose}
                className="w-full py-4 rounded-2xl items-center"
                style={{ backgroundColor: AppColors.primaryText }}
              >
                <Text className="text-base font-nunito-bold text-white">
                  {translate("review.close")}
                </Text>
              </TouchableOpacity>
            </View>
          ) : (
            <>
              {/* Header */}
              <View
                className="flex-row items-center justify-between px-5 py-4 border-b"
                style={{ borderColor: AppColors.borderLight }}
              >
                <TouchableOpacity
                  onPress={onClose}
                  hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
                >
                  <MaterialIcons
                    name="chevron-left"
                    size={28}
                    color={AppColors.primaryText}
                  />
                </TouchableOpacity>
                <Text
                  className="text-lg font-nunito-bold"
                  style={{ color: AppColors.primaryText }}
                >
                  {translate("review.leaveReview")}
                </Text>
                <View style={{ width: 28 }} />
              </View>

              <ScrollView
                className="flex-1 px-5 pt-5"
                showsVerticalScrollIndicator={false}
                contentContainerStyle={{ paddingBottom: 40 }}
                keyboardShouldPersistTaps="handled"
              >
                {/* Event header */}
                <View className="flex-row items-center gap-3 mb-6">
                  <View
                    className="w-12 h-12 rounded-xl overflow-hidden"
                    style={{ backgroundColor: AppColors.imagePlaceholder }}
                  >
                    {!!eventImageUrl && (
                      <Image
                        source={{ uri: eventImageUrl }}
                        style={{ width: 48, height: 48 }}
                        contentFit="cover"
                      />
                    )}
                  </View>
                  <View className="flex-1">
                    <Text
                      className="text-base font-nunito-bold"
                      style={{ color: AppColors.primaryText }}
                    >
                      {eventName}
                    </Text>
                    <View className="flex-row items-center gap-1 mt-0.5">
                      <MaterialIcons
                        name="location-on"
                        size={12}
                        color={AppColors.subtleText}
                      />
                      <Text
                        className="text-xs"
                        style={{ color: AppColors.subtleText }}
                        numberOfLines={1}
                      >
                        {eventLocation}
                      </Text>
                    </View>
                  </View>
                </View>

                {step === 1 ? (
                  <>
                    <Text
                      className="text-xl font-nunito-bold mb-5"
                      style={{ color: AppColors.primaryText }}
                    >
                      {translate("review.howWasExperience")}
                    </Text>

                    {/* Rating smileys */}
                    <View className="flex-row justify-between mb-7">
                      {RATING_OPTIONS.map(({ rating: r, image, labelKey }) => (
                        <TouchableOpacity
                          key={r}
                          onPress={() => setRating(r)}
                          className="flex-1 items-center gap-2"
                          style={{
                            opacity: rating === null || rating === r ? 1 : 0.3,
                          }}
                        >
                          <Image source={image} style={{ width: 55, height: 55 }} />
                          <Text
                            className="text-sm text-center"
                            style={{ color: AppColors.secondaryText }}
                          >
                            {translate(labelKey)}
                          </Text>
                        </TouchableOpacity>
                      ))}
                    </View>

                    <CategorySection
                      title={translate("review.content")}
                      options={CONTENT_OPTIONS}
                      selected={selectedCategories}
                      onToggle={toggleCategory}
                      translate={translate}
                    />

                    <CategorySection
                      title={translate("review.difficulty")}
                      options={DIFFICULTY_OPTIONS}
                      selected={selectedCategories}
                      onToggle={toggleCategory}
                      translate={translate}
                    />

                    <CategorySection
                      title={translate("review.instructor")}
                      options={INSTRUCTOR_OPTIONS}
                      selected={selectedCategories}
                      onToggle={toggleCategory}
                      translate={translate}
                    />

                    <CategorySection
                      title={translate("review.value")}
                      options={VALUE_OPTIONS}
                      selected={selectedCategories}
                      onToggle={toggleCategory}
                      translate={translate}
                    />

                    <TouchableOpacity
                      onPress={() => setStep(2)}
                      disabled={!rating}
                      className="py-4 rounded-2xl items-center mt-3"
                      style={{
                        backgroundColor: rating
                          ? AppColors.primaryText
                          : AppColors.borderLight,
                      }}
                    >
                      <Text
                        className="text-base font-nunito-bold"
                        style={{ color: rating ? "#fff" : AppColors.subtleText }}
                      >
                        {translate("review.next")}
                      </Text>
                    </TouchableOpacity>
                  </>
                ) : (
                  <>
                    <Text
                      className="text-xl font-nunito-bold"
                      style={{ color: AppColors.primaryText }}
                    >
                      {translate("review.tellUsMore")}{" "}
                      <Text
                        className="text-base"
                        style={{ color: AppColors.subtleText }}
                      >
                        {translate("review.optional")}
                      </Text>
                    </Text>

                    <TextInput
                      className="mt-4 p-4 rounded-xl border text-sm"
                      style={{
                        borderColor: AppColors.borderLight,
                        color: AppColors.primaryText,
                        height: 120,
                        textAlignVertical: "top",
                        fontFamily: "NunitoSans_400Regular",
                      }}
                      placeholder={translate("review.typeDescription")}
                      placeholderTextColor={AppColors.placeholderText}
                      multiline
                      value={description}
                      onChangeText={setDescription}
                    />

                    <View className="flex-row items-center justify-between mt-5 mb-8">
                      <Text
                        className="text-base"
                        style={{ color: AppColors.primaryText }}
                      >
                        {translate("review.submitAnonymously")}
                      </Text>
                      <Switch
                        value={anonymous}
                        onValueChange={setAnonymous}
                        trackColor={{
                          false: AppColors.borderLight,
                          true: AppColors.primaryText,
                        }}
                        thumbColor="#fff"
                      />
                    </View>

                    <View className="flex-row gap-3">
                      <TouchableOpacity
                        onPress={() => setStep(1)}
                        className="flex-1 py-4 rounded-2xl items-center border"
                        style={{ borderColor: AppColors.borderLight }}
                      >
                        <Text
                          className="text-base font-nunito-bold"
                          style={{ color: AppColors.primaryText }}
                        >
                          {translate("review.back")}
                        </Text>
                      </TouchableOpacity>
                      <TouchableOpacity
                        onPress={handleSubmit}
                        disabled={isPending}
                        className="flex-1 py-4 rounded-2xl items-center"
                        style={{ backgroundColor: AppColors.primaryText }}
                      >
                        {isPending ? (
                          <ActivityIndicator color="#fff" size="small" />
                        ) : (
                          <Text
                            className="text-base font-nunito-bold"
                            style={{ color: "#fff" }}
                          >
                            {translate("common.submit")}
                          </Text>
                        )}
                      </TouchableOpacity>
                    </View>
                  </>
                )}
              </ScrollView>
            </>
          )}
        </View>
      </KeyboardAvoidingView>
    </Modal>
  );
}