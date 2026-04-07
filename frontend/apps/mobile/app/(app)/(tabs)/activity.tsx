import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors } from "@/constants/theme";
import { useMemo, useState } from "react";
import { Image, Pressable, ScrollView, TouchableOpacity, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import {
  useGetAllEventOccurrences,
  useGetRegistrationsByGuardianId,
  useGetChildrenByGuardianId,
  type EventOccurrence,
  type Registration,
  type Child,
} from "@skillspark/api-client";

type toggleValue = "upcoming" | "past" | undefined

interface ToggleProps {
  value: toggleValue
  onChange: (newValue: toggleValue) => void
}

function Toggle({ value, onChange }: ToggleProps) {
  return (
    <ThemedView className="w-11/12 border border-gray-200 rounded-md flex flex-row justify-between p-2 mt-2">
      <Pressable
        className={`${value === "upcoming" ? "bg-gray-200" : ""} w-6/12 flex items-center rounded-md py-1.5`}
        onPress={() => onChange("upcoming")}
      >
        <ThemedText>Upcoming</ThemedText>
      </Pressable>

      <Pressable
        className={`${value === "past" ? "bg-gray-200" : ""} w-6/12 flex items-center rounded-md py-1.5`}
        onPress={() => onChange("past")}
      >
        <ThemedText>Past</ThemedText>
      </Pressable>
    </ThemedView>
  );
}

interface RegistrationCardData {
  event_occurrence_id: string
  image_uri: string
  start_time: Date
  end_time: Date
  title: string
  children: Set<Child>
  location: string
  price: number
  onClickRemove: () => void
}

interface RegistrationCardProps {
  data: RegistrationCardData
}

const formatTime = (d: Date) =>
  d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" })
const formatDate = (d: Date) =>
  d.toLocaleDateString([], { weekday: "short", month: "short", day: "numeric" })

function UpcomingRegistrationCard({ data }: RegistrationCardProps) {
  const childNames = Array.from(data.children).map((c) => c.name).join(", ")

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4"
      style={{
        borderWidth: 1,
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
        className="w-full h-44 px-4 pt-4 rounded-md"
      />
      <View className="flex flex-row justify-between"> 
        <View className="px-4 pb-4 gap-1">
          <ThemedText type="subtitle">
            {data.title}
          </ThemedText>
          <View className="flex flex-row gap-2 items-center "> 
            <IconSymbol name="clock" color="black" size={20}/>
            <ThemedText className="text-sm" >
              {formatTime(data.start_time)} – {formatTime(data.end_time)}
            </ThemedText>
          </View>
          <View className="flex flex-row gap-2 items-center "> 
            <IconSymbol name="location" color="black" size={20}/>
            <ThemedText className="text-sm" >
              {data.location}
            </ThemedText>
          </View>
          {childNames ? (
            <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
              {childNames}
            </ThemedText>
          ) : null}
          <TouchableOpacity
            onPress={data.onClickRemove}
            className="mt-2 items-center rounded-lg py-2"
            style={{ backgroundColor: AppColors.danger }}
            activeOpacity={0.7}
          >
            <ThemedText className="text-sm font-medium text-white">
              Remove Registration
            </ThemedText>
          </TouchableOpacity>
        </View>
        <View>
          TODO: make this date box
        </View>
      </View>
    </View>
  );
}

function PastRegistrationCard({ data }: RegistrationCardProps) {
  const childNames = Array.from(data.children).map((c) => c.name).join(", ")
  const priceDisplay = `฿${(data.price / 100).toLocaleString()}`

  return (
    <View
      className="w-11/12 rounded-xl overflow-hidden mb-4"
      style={{
        borderWidth: 1,
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
        className="w-full h-44 px-4 pt-4 rounded-md"
      />
      <View className="px-4 pb-4">
        <ThemedText type="defaultSemiBold" className="text-xl">
          {data.title}
        </ThemedText>
        <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
          {formatDate(data.start_time)} · {formatTime(data.start_time)} – {formatTime(data.end_time)}
        </ThemedText>
        <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
          {data.location}
        </ThemedText>
        {childNames ? (
          <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
            {childNames}
          </ThemedText>
        ) : null}
        <ThemedText className="text-sm" style={{ color: AppColors.mutedText }}>
          {priceDisplay}
        </ThemedText>
        <TouchableOpacity
          onPress={data.onClickRemove}
          className="mt-2 items-center rounded-lg py-2"
          style={{ backgroundColor: AppColors.danger }}
          activeOpacity={0.7}
        >
          <ThemedText className="text-sm font-medium text-white">
            Remove Registration
          </ThemedText>
        </TouchableOpacity>
      </View>
    </View>
  );
}

export default function ActivityScreen() {
  const insets = useSafeAreaInsets();
  const [selection, setSelection] = useState<"upcoming" | "past">("upcoming");

  // TODO CHANGE THIS
  // const { guardianId } = useAuthContext();
  const guardianId = "11111111-1111-1111-1111-111111111111";

  const getOnRemove = (registrationId: string) => {
    return function () {
      console.log(registrationId)
    }
  }

  const { data: registrationsResp } = useGetRegistrationsByGuardianId(
    guardianId!,
    {
      query: { enabled: !!guardianId },
    },
  );
  const registrations: Registration[] = useMemo(() => {
    const d = registrationsResp as unknown as
      | { data: { registrations: Registration[] } }
      | undefined;
    return d?.data?.registrations ?? [];
  }, [registrationsResp]);

  const { data: occurrencesResp } = useGetAllEventOccurrences();
  const allOccurrences: EventOccurrence[] = useMemo(() => {
    const d = occurrencesResp as unknown as
      | { data: EventOccurrence[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [occurrencesResp]);

  const eventOccurrencesMap: Record<string, EventOccurrence> = useMemo(() => {
    const map: Record<string, EventOccurrence> = {}
    allOccurrences.forEach((o) => {
      map[o.id] = o
    })
    return map
  }, [allOccurrences])

  const { data: childrenResp } = useGetChildrenByGuardianId(
    guardianId!,
    {
      query: { enabled: !!guardianId },
    },
  )

  const children: Child[] = useMemo(() => {
    const d = childrenResp as unknown as
      | { data: Child[] }
      | undefined;
    return Array.isArray(d?.data) ? d.data : [];
  }, [childrenResp]);

  const childMap = useMemo(() => {
    const map: Record<string, Child> = {}
    children.forEach((c) => {
      map[c.id] = c
    })
    return map
  }, [children])

  const { upcomingRegistrations, pastRegistrations } = useMemo(() => {
    const now = new Date()
    const grouped: Record<string, RegistrationCardData> = {}
    registrations.filter((r) => r.status === "registered").forEach((r) => {
      const occurrence = eventOccurrencesMap[r.event_occurrence_id]
      if (!occurrence) return
      const child = childMap[r.child_id]
      if (grouped[r.event_occurrence_id]) {
        if (child) grouped[r.event_occurrence_id].children.add(child)
      } else {
        const startDate = new Date(occurrence.start_time)
        const endDate = new Date(occurrence.end_time)
        grouped[r.event_occurrence_id] = {
          event_occurrence_id: r.event_occurrence_id,
          image_uri: "https://picsum.photos/800/300",
          start_time: startDate,
          end_time: endDate,
          title: occurrence.event.title,
          children: new Set(child ? [child] : []),
          location: occurrence.location.address_line1,
          price: occurrence.price,
          onClickRemove: getOnRemove(r.id),
        }
      }
    })
    const all = Object.values(grouped)
    return {
      upcomingRegistrations: all.filter((r) => r.start_time >= now).sort((a, b) => a.start_time.getTime() - b.start_time.getTime()),
      pastRegistrations: all.filter((r) => r.start_time < now).sort((a, b) => b.start_time.getTime() - a.start_time.getTime()),
    }
  }, [registrations, eventOccurrencesMap, childMap])

  const toggleSelection = (newValue: toggleValue) => {
    setSelection(newValue!);
  };

  const displayed = selection === "upcoming" ? upcomingRegistrations : pastRegistrations;

  return (
    <ThemedView className="w-full flex-1" style={{ paddingTop: insets.top }}>
      <ThemedView className="w-full flex items-center border-b border-black/[0.5]">
        <ThemedText className="py-3">Activity</ThemedText>
      </ThemedView>

      <ThemedView className="w-full flex items-center">
        <Toggle value={selection} onChange={toggleSelection} />
      </ThemedView>

      <ScrollView
        contentContainerStyle={{ alignItems: "center", paddingTop: 16, paddingBottom: 32 }}
        showsVerticalScrollIndicator={false}
      >
        {displayed.length === 0 ? (
          <ThemedText className="mt-8" style={{ color: AppColors.mutedText }}>
            No {selection} registrations.
          </ThemedText>
        ) : (
          displayed.map((reg) =>
            selection === "upcoming"
              ? <UpcomingRegistrationCard key={reg.event_occurrence_id} data={reg} />
              : <PastRegistrationCard key={reg.event_occurrence_id} data={reg} />
          )
        )}
      </ScrollView>
    </ThemedView>
  );
}
