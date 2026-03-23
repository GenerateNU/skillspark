import { TouchableOpacity } from "react-native";
import { useQueryClient } from "@tanstack/react-query";
import {
  useCreateSaved,
  useDeleteSaved,
  useGetSavedByGuardianId,
  getGetSavedByGuardianIdQueryKey,
} from "@skillspark/api-client";
import type { getSavedByGuardianIdResponse, Saved, Event } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AppColors } from "@/constants/theme";
import { useGuardian } from "@/hooks/use-guardian";

export function BookmarkButton({ eventId, event }: { eventId: string; event: Event }) {
  const { guardianId } = useGuardian();
  const queryClient = useQueryClient();
  const { data: savedResponse } = useGetSavedByGuardianId(guardianId);
  const savedQueryKey = getGetSavedByGuardianIdQueryKey(guardianId);

  const savedItems = savedResponse?.status === 200 ? savedResponse.data : [];
  const savedEntry = savedItems.find((s) => s.event?.id === eventId);
  const isBookmarked = !!savedEntry;

  const optimisticOptions = (updater: (items: Saved[]) => Saved[]) => ({
    mutation: {
      onMutate: async () => {
        await queryClient.cancelQueries({ queryKey: savedQueryKey });
        const previous = queryClient.getQueryData<getSavedByGuardianIdResponse>(savedQueryKey);
        queryClient.setQueryData<getSavedByGuardianIdResponse>(savedQueryKey, (old) => {
          if (!old || old.status !== 200) return old;
          return { ...old, data: updater(old.data) };
        });
        return { previous };
      },
      onError: (_err: unknown, _vars: unknown, context: { previous?: getSavedByGuardianIdResponse } | undefined) => {
        if (context?.previous) {
          queryClient.setQueryData(savedQueryKey, context.previous);
        }
      },
      onSettled: () => queryClient.invalidateQueries({ queryKey: savedQueryKey }),
    },
  });

  const createSaved = useCreateSaved(
    optimisticOptions((items) => [
      ...items,
      { id: "optimistic", guardian_id: guardianId, event, created_at: "", updated_at: "" },
    ])
  );
  const deleteSaved = useDeleteSaved(
    optimisticOptions((items) => items.filter((s) => s.event?.id !== eventId))
  );

  const handlePress = () => {
    if (isBookmarked && savedEntry) {
      deleteSaved.mutate({ id: savedEntry.id });
    } else {
      createSaved.mutate({ data: { event_id: eventId, guardian_id: guardianId } });
    }
  };

  return (
    <TouchableOpacity onPress={handlePress} activeOpacity={0.7}>
      <MaterialIcons
        name={isBookmarked ? "bookmark" : "bookmark-border"}
        size={40}
        color={AppColors.primaryText}
      />
    </TouchableOpacity>
  );
}
