import { EventCard } from "@/components/EventCard";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useGetSavedByGuardianId } from '@skillspark/api-client';
import { useRouter } from "expo-router";
import { ActivityIndicator, FlatList, ScrollView, useColorScheme, View } from "react-native";
import { useSafeAreaInsets } from "react-native-safe-area-context";

const GUARDIAN_ID = '88888888-8888-8888-8888-888888888888';

export default function SavedScreen() {

    const insets = useSafeAreaInsets();
    const colorScheme = useColorScheme();
    const router = useRouter();

    const { data: savedResponse, isLoading: isLoading } = useGetSavedByGuardianId(GUARDIAN_ID);
    const saved = savedResponse?.status === 200 ? savedResponse.data : null;

    if (isLoading) {
        return (
            <View style={{ flex: 1, alignItems: "center", justifyContent: "center", gap: 8 }}>
                <ActivityIndicator size="large" />
                <ThemedText>Loading saved events...</ThemedText>
            </View>
        );
    }

    return (
        <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
            <ScrollView showsVerticalScrollIndicator={false}
                contentContainerStyle={{ paddingTop: 10, paddingBottom: 20 }}
                bounces={false}>
                    <FlatList 
                        data={saved}
                        renderItem={({ item }) => <EventCard item={item} />}
keyExtractor={(item) => item.id}
                    ></FlatList>
            </ScrollView>
        </ThemedView>
    )

};