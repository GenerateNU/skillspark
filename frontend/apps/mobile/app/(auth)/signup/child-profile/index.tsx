import { Button } from "@/components/Button";
import { AuthBackground } from "@/components/AuthBackground";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontSizes } from "@/constants/theme";
import { useRouter } from "expo-router";
import { useTranslation } from "react-i18next";
import { ScrollView, Text, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { useAuthContext } from "@/hooks/use-auth-context";
import { useGuardian } from "@/hooks/use-guardian";
import { ChildAvatar } from "@/components/ChildAvatar";
import { MONTHS } from "@/components/ChildProfileForm";

// 4. set up your child's profile
export default function ChildProfileScreen() {
  const router = useRouter();
  const { t: translate } = useTranslation();
  const insets = useSafeAreaInsets();
  const { guardianId } = useAuthContext();
  const { children } = useGuardian(guardianId);

  return (
    <View className="flex-1" style={{ paddingTop: insets.top }}>
      <AuthBackground />
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
      <View className="px-6 pt-2 pb-6 items-center">
        <ThemedText
          className="font-nunito-bold text-[#111] text-center"
          style={{
            fontSize: FontSizes.hero,
            lineHeight: FontSizes.hero + 8,
            letterSpacing: -0.5,
          }}
        >
          {translate("onboarding.setUpChild")}
        </ThemedText>
      </View>

      {/* Scrollable content */}
      <ScrollView
        className="flex-1"
        contentContainerStyle={{
          paddingHorizontal: 24,
          paddingBottom: 16,
          gap: 12,
        }}
        showsVerticalScrollIndicator={false}
      >
        {/* Added children */}
        {children.map((child: any) => (
          <TouchableOpacity
            className="shadow-sm flex-row items-center bg-white rounded-2xl py-3.5 px-4 gap-3"
            key={child.id}
            onPress={() =>
              router.push({
                pathname: "/(auth)/signup/child-profile/add-child",
                params: {
                  id: child.id,
                  name: child.name,
                  birth_month: child.birth_month,
                  birth_year: child.birth_year,
                  school_id: child.school_id,
                  interests: child.interests?.join(","),
                  avatar_face: child.avatar_face ?? "",
                  avatar_background: child.avatar_background ?? "",
                },
              })
            }
            activeOpacity={0.75}
          >
            <View className="mr-2">
              <ChildAvatar
                name={child.name ?? ""}
                avatarFace={child.avatar_face}
                avatarBackground={child.avatar_background}
                size={32}
              />
            </View>
            <View className="flex-1">
              <ThemedText
                className="text-sm font-nunito-medium"
                numberOfLines={1}
              >
                {child.name}
              </ThemedText>
              <ThemedText className="text-[10px] font-nunito text-[#6B7280]">
                {[
                  child.birth_month ? translate("months." + MONTHS[child.birth_month - 1]) : null,
                  child.birth_year,
                ]
                  .filter(Boolean)
                  .join(", ")}
              </ThemedText>
            </View>
            <IconSymbol
              name="chevron.right"
              size={18}
              color={AppColors.mutedText}
            />
          </TouchableOpacity>
        ))}

        {/* Add a new child profile button */}
        <TouchableOpacity
          className="shadow-sm bg-white rounded-2xl py-5 items-center"
          onPress={() => router.push("/(auth)/signup/child-profile/add-child")}
          activeOpacity={0.7}
        >
          <Text className="font-nunito-semibold text-[15px] text-[#111]">
            {translate("onboarding.addNewChild")}
          </Text>
          <Text className="font-nunito text-[22px] text-[#111] leading-7">
            +
          </Text>
        </TouchableOpacity>
      </ScrollView>

      <View
        className="items-center px-6 pt-4"
        style={{ paddingBottom: insets.bottom + 16 }}
      >
        <Button
          label={translate("onboarding.continue")}
          onPress={() => router.push("/(auth)/signup/emergency-contact")}
          disabled={false}
        />
      </View>
    </View>
  );
}
