import { Image } from "expo-image";
import { StyleSheet, ActivityIndicator, FlatList, View } from "react-native";
import ParallaxScrollView from "@/components/parallax-scroll-view";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useGetAllEventOccurrences } from "@skillspark/api-client";
import type { EventOccurrence } from "@skillspark/api-client";

function EventOccurrencesList() {
  const { data: eventOccurrences, isLoading, error } = useGetAllEventOccurrences();

  if (isLoading) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ActivityIndicator size="large" />
        <ThemedText>Loading events...</ThemedText>
      </ThemedView>
    );
  }

  if (error) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ThemedText type="defaultSemiBold" style={styles.errorText}>
          Error loading events
        </ThemedText>
        <ThemedText>{error.message || 'An error occurred'}</ThemedText>
      </ThemedView>
    );
  }

  // Check if eventOccurrences exists and is an array before filtering
  if (!eventOccurrences || !Array.isArray(eventOccurrences)) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ThemedText>No events available</ThemedText>
      </ThemedView>
    );
  }

  // Filter and sort upcoming events
  const upcomingEvents = eventOccurrences
    .filter(occurrence => {
      return new Date(occurrence.start_time) >= new Date();
    })
    .sort((a, b) => {
      return new Date(a.start_time).getTime() - new Date(b.start_time).getTime();
    });

  if (upcomingEvents.length === 0) {
    return (
      <ThemedView style={styles.centerContainer}>
        <ThemedText>No upcoming events</ThemedText>
      </ThemedView>
    );
  }

  const renderEventOccurrence = ({ item }: { item: EventOccurrence }) => (
    <ThemedView style={styles.eventCard}>
      <ThemedText type="subtitle">
        {item.event.title}
      </ThemedText>
      
      {item.event.description && (
        <ThemedText style={styles.eventDescription}>
          {item.event.description}
        </ThemedText>
      )}
      
      {item.location && (
        <ThemedText style={styles.eventDetail}>
          📍 {item.location.address_line1} {item.location.address_line2 ? `, ${item.location.address_line2}` : ''} {item.location.province}, {item.location.subdistrict} {item.location.postal_code}
        </ThemedText>
      )}
      
      <ThemedText style={styles.eventDetail}>
        🕒 {new Date(item.start_time).toLocaleDateString('en-US', {
          weekday: 'short',
          month: 'short',
          day: 'numeric',
          hour: 'numeric',
          minute: '2-digit',
        })}
      </ThemedText>
      
      <ThemedText style={styles.eventDetail}>
        ⏱️ Ends: {new Date(item.end_time).toLocaleTimeString('en-US', {
          hour: 'numeric',
          minute: '2-digit',
        })}
      </ThemedText>
      
      <ThemedText style={styles.eventDetail}>
        👥 {item.curr_enrolled} / {item.max_attendees} enrolled
      </ThemedText>
      
      {item.language && (
        <ThemedText style={styles.eventDetail}>
          🌐 Language: {item.language}
        </ThemedText>
      )}
    </ThemedView>
  );

  return (
    <FlatList
      data={upcomingEvents}
      renderItem={renderEventOccurrence}
      keyExtractor={(item) => item.id}
      contentContainerStyle={styles.listContainer}
      ItemSeparatorComponent={() => <View style={styles.separator} />}
    />
  );
}

export default function HomeScreen() {
  return (
    <ParallaxScrollView
      headerBackgroundColor={{ light: "#BBBBBB", dark: "#333333" }}
      headerImage={
        <Image
          source={require("@/assets/images/generate.png")}
          style={styles.generateLogo}
        />
      }
    >
      <ThemedView style={styles.stepContainer}>
        <ThemedText type="title">Welcome to SkillSpark</ThemedText>
        <ThemedText>Upcoming Events</ThemedText>
      </ThemedView>

      <EventOccurrencesList />
    </ParallaxScrollView>
  );
}

const styles = StyleSheet.create({
  titleContainer: {
    flexDirection: "row",
    alignItems: "center",
    gap: 8,
  },
  stepContainer: {
    gap: 8,
    marginBottom: 16,
  },
  generateLogo: {
    height: 300,
    width: 300,
    position: "absolute",
    bottom: -60,
    left: -35,
  },
  centerContainer: {
    padding: 16,
    alignItems: "center",
    gap: 8,
  },
  eventCard: {
    padding: 16,
    borderRadius: 8,
    gap: 8,
  },
  eventDescription: {
    opacity: 0.8,
    marginTop: 4,
  },
  eventDetail: {
    fontSize: 14,
    opacity: 0.7,
    marginTop: 2,
  },
  listContainer: {
    paddingBottom: 16,
  },
  separator: {
    height: 12,
  },
  errorText: {
    color: '#ff4444',
  },
});