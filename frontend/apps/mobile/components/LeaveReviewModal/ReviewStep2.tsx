import { AppColors } from "@/constants/theme";
import {
  ActivityIndicator,
  Switch,
  Text,
  TextInput,
  TouchableOpacity,
  View,
} from "react-native";

interface Props {
  description: string;
  setDescription: (value: string) => void;
  anonymous: boolean;
  setAnonymous: (value: boolean) => void;
  isPending: boolean;
  onBack: () => void;
  onSubmit: () => void;
  translate: (key: string) => string;
}

export function ReviewStep2({
  description,
  setDescription,
  anonymous,
  setAnonymous,
  isPending,
  onBack,
  onSubmit,
  translate,
}: Props) {
  return (
    <>
      <Text
        className="text-xl font-nunito-bold"
        style={{ color: AppColors.primaryText }}
      >
        {translate("review.tellUsMore")}{" "}
        <Text className="text-base" style={{ color: AppColors.subtleText }}>
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
        <Text className="text-base" style={{ color: AppColors.primaryText }}>
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
          onPress={onBack}
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
          onPress={onSubmit}
          disabled={isPending}
          className="flex-1 py-4 rounded-2xl items-center"
          style={{ backgroundColor: AppColors.primaryText }}
        >
          {isPending ? (
            <ActivityIndicator color="#fff" size="small" />
          ) : (
            <Text className="text-base font-nunito-bold" style={{ color: "#fff" }}>
              {translate("common.submit")}
            </Text>
          )}
        </TouchableOpacity>
      </View>
    </>
  );
}
