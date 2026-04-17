import { View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { type Child, type EventOccurrence } from "@skillspark/api-client";
import { ChildAvatar } from "@/components/ChildAvatar";
import { AppColors } from "@/constants/theme";
import { EventImage } from "@/components/EventImage";

export function RecommendedCard({
  child,
  occurrences,
}: {
  child: Child;
  occurrences: EventOccurrence[];
}) {
  const router = useRouter();
  const firstName = child.name.split(" ")[0];
  const titles = occurrences.map((o) => o.event.title);
  const summaryText =
    titles.length > 2
      ? `${titles[0]},\n${titles[1]},\nAnd More!`
      : titles.join(",\n");
  const img0 = occurrences[0];
  const img1 = occurrences[1];

  return (
    <Pressable
      onPress={() =>
        router.push({
          pathname: "/child/[id]",
          params: { id: child.id, name: firstName },
        })
      }
      className="mx-5 mb-3 rounded-2xl p-4 flex-row bg-white shadow-sm"
    >
      <View className="flex-1 pr-3">
        <Text
          className="mb-2 font-nunito-bold text-xl"
          style={{ color: AppColors.primaryText }}
        >
          {firstName}
        </Text>
        <ChildAvatar
          name={child.name}
          avatarFace={child.avatar_face}
          avatarBackground={child.avatar_background}
          size={38}
        />
        <Text
          className="mt-2 font-nunito text-[12px]"
          style={{ color: AppColors.mutedText }}
        >
          {summaryText}
        </Text>
      </View>
      <View className="w-[160px] h-[130px]">
        <View
          className="absolute top-[30px] left-[60px] w-[95px] h-[95px] rounded-xl rotate-[35deg]"
          style={{ backgroundColor: AppColors.violetPastel }}
        />
        <View className="absolute top-[15px] left-[40px] w-[100px] h-[100px] rounded-xl overflow-hidden rotate-[20deg]">
          <EventImage
            uri={img1?.event.presigned_url}
            style={{ width: "100%", height: "100%" }}
          />
        </View>
        {img0 && (
          <View className="absolute top-0 left-0 w-[120px] h-[120px] rounded-xl overflow-hidden">
            <EventImage
              uri={img0.event.presigned_url}
              style={{ width: "100%", height: "100%" }}
            />
          </View>
        )}
      </View>
    </Pressable>
  );
}
