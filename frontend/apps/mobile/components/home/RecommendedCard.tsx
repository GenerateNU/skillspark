import { Image, View, Text, Pressable } from "react-native";
import { useRouter } from "expo-router";
import { SvgXml } from "react-native-svg";
import { type Child, type EventOccurrence } from "@skillspark/api-client";
import { GLASSES_FACE_SVG, getSvgWithColor } from "@/constants/avatarFaces";
import { DEFAULT_AVATAR_COLOR } from "@/constants/avatarColors";

export function RecommendedCard({
  child,
  occurrences,
}: {
  child: Child;
  occurrences: EventOccurrence[];
}) {
  const router = useRouter();
  const firstName = child.name.split(" ")[0];
  const bgColor = child.avatar_background ?? DEFAULT_AVATAR_COLOR;
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
        <Text className="mb-2 font-nunito-bold text-xl text-[#111]">
          {firstName}
        </Text>
        <SvgXml
          xml={getSvgWithColor(GLASSES_FACE_SVG, bgColor)}
          width={38}
          height={38}
        />
        <Text className="mt-2 font-nunito text-[12px] text-[#6B7280]">
          {summaryText}
        </Text>
      </View>
      <View className="w-[160px] h-[130px]">
        <View
          className="absolute top-[30px] left-[60px] w-[95px] h-[95px] rounded-xl bg-[#DDD6FE] rotate-[35deg]"
        />
        <View
          className="absolute top-[15px] left-[40px] w-[100px] h-[100px] rounded-xl overflow-hidden rotate-[20deg]"
        >
          {img1?.event.presigned_url ? (
            <Image
              source={{ uri: img1.event.presigned_url }}
              className="w-full h-full"
              resizeMode="cover"
            />
          ) : (
            <View className="w-full h-full bg-[#D9D9D9]" />
          )}
        </View>
        {img0 && (
          <View
            className="absolute top-0 left-0 w-[120px] h-[120px] rounded-xl overflow-hidden"
          >
            {img0.event.presigned_url ? (
              <Image
                source={{ uri: img0.event.presigned_url }}
                className="w-full h-full"
                resizeMode="cover"
              />
            ) : (
              <View className="w-full h-full bg-[#D9D9D9]" />
            )}
          </View>
        )}
      </View>
    </Pressable>
  );
}
