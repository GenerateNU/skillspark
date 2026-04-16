import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { Image, TouchableOpacity, View } from "react-native";
import {
  type RegistrationCardProps,
  formatTime,
} from "@/components/RegistrationCard.types";
import { ChildAvatar } from "@/components/ChildAvatar";

export function UpcomingRegistrationCard({ data }: RegistrationCardProps) {
  const { t } = useTranslation();
  const children = data.childRegistrations.map((cr) => cr.child);

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4 border"
      style={{
        borderColor: AppColors.borderLight,
        backgroundColor: AppColors.white,
        shadowColor: "#000",
        shadowOffset: { width: 0, height: 1 },
        shadowOpacity: 0.06,
        shadowRadius: 4,
        elevation: 2,
      }}
    >
      <Image
        source={{ uri: data.image_uri }}
        style={{ width: "100%", height: 176, borderRadius: 6 }}
        resizeMode="cover"
      />
      <View className="flex flex-row justify-between items-center p-1">
        <View className="px-4 pb-4 gap-1 flex flex-col justify-between">
          <ThemedText type="subtitle">{data.title}</ThemedText>
          <View className="flex flex-row gap-2 items-center ">
            <IconSymbol name="clock" color="black" size={18} />
            <ThemedText className="text-sm">
              {formatTime(data.start_time)} – {formatTime(data.end_time)}
            </ThemedText>
          </View>
          <View className="flex flex-row gap-2 items-center ">
            <IconSymbol name="location" color="black" size={18} />
            <ThemedText className="text-sm">{data.location}</ThemedText>
          </View>
        </View>
        <View className="flex flex-col justify-center items-center bg-[#E69BF040] w-20 h-20 mr-2 rounded-md">
          <ThemedText type="subtitle" className="font-bold leading-none mr-1">
            {data.start_time.getDate() < 10
              ? "0" + data.start_time.getDate().toString()
              : data.start_time.getDate().toString()}
          </ThemedText>
          <ThemedText
            type="subtitle"
            className=" font-semibold tracking-widest "
          >
            {data.start_time.toLocaleString("default", { month: "short" })}
          </ThemedText>
        </View>
      </View>
      <View className="flex flex-row justify-between items-center px-4 py-3">
        <View className="flex flex-row gap-2">
          {children.map((child) => (
            <ChildAvatar
              key={child.id}
              name={child.name}
              avatarFace={child.avatar_face}
              avatarBackground={child.avatar_background}
              size={32}
            />
          ))}
        </View>
        <TouchableOpacity
          onPress={data.onClickRemove}
          className="px-6 py-2 rounded-full bg-black"
          activeOpacity={0.7}
        >
          <ThemedText lightColor="white" className="text-sm font-medium">
            {t("activity.remove")}
          </ThemedText>
        </TouchableOpacity>
      </View>
    </View>
  );
}
