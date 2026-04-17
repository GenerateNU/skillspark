import { Button } from "@/components/Button";
import { ErrorMessage } from "@/components/ErrorMessage";
import { ThemedText } from "@/components/themed-text";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useRouter } from "expo-router";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity, View } from "react-native";
import * as ImagePicker from "expo-image-picker";
import { AuthBackground } from "@/components/AuthBackground";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useGuardian } from "@/hooks/use-guardian";
import { FontSizes } from "@/constants/theme";
import { IconSymbol } from "@/components/ui/icon-symbol";

export default function PhotoScreen() {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const insets = useSafeAreaInsets();
  const [errorText, setErrorText] = useState("");
  const [image, setImage] = useState<string | undefined>(undefined);
  const { update, guardianId } = useAuthContext();
  const { guardian } = useGuardian(guardianId);

  const pickImage = async () => {
    const result = await ImagePicker.launchImageLibraryAsync({
      mediaTypes: ["images"],
      allowsEditing: true,
      aspect: [1, 1],
    });
    if (!result.canceled) {
      setImage(result.assets[0].uri);
    }
  };

  const onSubmit = () => {
    if (!guardian) {
      setErrorText("ERROR: Could not fetch guardian ID");
    } else {
      update(
        () => router.push("/(auth)/signup/child-profile"),
        setErrorText,
        guardian.id,
        guardian.email,
        guardian.language_preference,
        guardian.name,
        guardian.username,
        image,
      );
    }
  };

  return (
    <View className="absolute inset-0">
      <AuthBackground />
      <View className="flex-1" style={{ paddingTop: insets.top }}>
        {/* Back button */}
        <TouchableOpacity
          onPress={() => router.back()}
          className="flex-row items-center px-5 py-3 gap-1"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={18} color="#11181C" />
          <ThemedText className="text-base font-nunito">
            {translate("onboarding.back")}
          </ThemedText>
        </TouchableOpacity>

        {/* Title */}
        <View className="px-6 pt-4 pb-2 items-center">
          <ThemedText
            className="font-nunito-bold text-[#111] text-center"
            style={{
              fontSize: FontSizes.hero,
              lineHeight: FontSizes.hero + 8,
              letterSpacing: -0.5,
            }}
          >
            {translate("onboarding.addPhoto")}
          </ThemedText>
        </View>

        {image ? (
          /* — Photo selected state — */
          <View className="flex-1 items-center justify-center px-6">
            <TouchableOpacity onPress={pickImage} activeOpacity={0.85}>
              <View
                className="rounded-full overflow-hidden"
                style={{ width: 180, height: 180 }}
              >
                <Image
                  source={{ uri: image }}
                  style={{ width: "100%", height: "100%" }}
                  resizeMode="cover"
                />
              </View>
            </TouchableOpacity>
            <ThemedText className="font-nunito text-[#6B7280] text-center mt-4 text-[15px]">
              {translate("onboarding.cropAndSubmit")}
            </ThemedText>
          </View>
        ) : (
          /* — Empty state — */
          <View className="flex-1 items-center justify-center px-6 gap-5">
            {/* Dashed circle placeholder */}
            <TouchableOpacity onPress={pickImage} activeOpacity={0.7}>
              <View
                className="items-center justify-center rounded-full"
                style={{
                  width: 180,
                  height: 180,
                  borderWidth: 2,
                  borderColor: "#9CA3AF",
                  borderStyle: "dashed",
                }}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={52}
                  color="#9CA3AF"
                />
              </View>
            </TouchableOpacity>

            {/* Choose Photo button */}
            <Button
              label={translate("onboarding.choosePhoto")}
              onPress={pickImage}
              disabled={false}
              bgColor="#FFFFFF"
              textColor="#1B1B1B"
              width="60%"
            />

            {/* Caption */}
            <ThemedText className="font-nunito text-[#6B7280] text-center text-[14px]">
              {translate("onboarding.personalize")}
            </ThemedText>
          </View>
        )}

        {/* Bottom buttons */}
        <View
          className="items-center px-6 pt-2 gap-3"
          style={{ paddingBottom: insets.bottom + 16 }}
        >
          {image ? (
            <View className="flex-row gap-3 w-full">
              <Button
                label={translate("onboarding.cancel")}
                onPress={() => setImage(undefined)}
                disabled={false}
                bgColor="#FFFFFF"
                textColor="#1B1B1B"
                width="48%"
              />
              <Button
                label={translate("onboarding.save")}
                onPress={onSubmit}
                disabled={false}
                bgColor="#1B1B1B"
                textColor="#FFFFFF"
                width="48%"
              />
            </View>
          ) : (
            <Button
              label={translate("onboarding.skip")}
              onPress={onSubmit}
              disabled={false}
              bgColor="#FFFFFF"
              textColor="#1B1B1B"
            />
          )}
          <ErrorMessage message={errorText} />
        </View>
      </View>
    </View>
  );
}
