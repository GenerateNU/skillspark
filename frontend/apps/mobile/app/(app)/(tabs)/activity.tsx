import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { useState } from "react";
import { useSafeAreaInsets } from "react-native-safe-area-context";

type toggleValue = "upcoming" | "past" | undefined

interface ToggleProps {
  value: toggleValue
  onChange: (newValue: toggleValue) => void
}

function Toggle({ value, onChange } : ToggleProps) {
  
  return (
    <ThemedView className="w-11/12 border rounded-sm flex"> 
      <ThemedView className="bg-gray-400">
        <ThemedText> 
          Upcoming
        </ThemedText>
      </ThemedView>
      
      <ThemedView className="bg-gray-400">
        <ThemedText> 
          Past
        </ThemedText>
      </ThemedView>
    </ThemedView>
  )
}

export default function ActivityScreen() {
  const insets = useSafeAreaInsets();
  const [selection, setSelection] = useState<"upcoming" | "past">()
  
  const toggleSelection = (newValue : toggleValue) => {
    setSelection(newValue)
  }
  return (
    <ThemedView className="w-full" style={{ paddingTop: insets.top }}>
      <ThemedView className="w-full flex items-center border-b border-black/[0.5]"> 
        <ThemedText className="py-3">
          Activity
        </ThemedText>
      </ThemedView>
      <Toggle value={selection } onChange={setSelection}/>
    </ThemedView>
  );
}
