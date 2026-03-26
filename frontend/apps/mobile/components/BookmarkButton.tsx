import { TouchableOpacity } from "react-native";
import { useQueryClient } from "@tanstack/react-query";
import {
  useCreateSaved,
  useDeleteSaved,
  useGetSavedByGuardianId,
  getGetSavedByGuardianIdQueryKey,
} from "@skillspark/api-client";
import type { getSavedByGuardianIdResponse, Saved } from "@skillspark/api-client";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { AppColors } from "@/constants/theme";
import { useAuthContext } from "@/hooks/use-auth-context";
import i18n from "@/i18n";

interface BookmarkButtonProps {
  eventId: string;
  // Controlled mode: pass these to skip the internal fetch
  isBookmarked?: boolean;
  savedEntryId?: string;
}

export function BookmarkButton({ eventId, isBookmarked: isBookmarkedProp, savedEntryId }: BookmarkButtonProps) {
  const queryClient = useQueryClient();
  const { guardianId } = useAuthContext();

  const controlled = isBookmarkedProp !== undefined;

  const { data: savedResponse } = useGetSavedByGuardianId(guardianId!, undefined, {
    query: { enabled: !!guardianId && !controlled },
  });

  const savedQueryKey = getGetSavedByGuardianIdQueryKey(guardianId!);

  const savedItems = savedResponse?.status === 200 ? savedResponse.data : [];
  const savedEntry = controlled
    ? (savedEntryId ? { id: savedEntryId } : undefined)
    : savedItems.find((s) => s.event.id === eventId);
  const isBookmarked = controlled ? isBookmarkedProp : !!savedEntry;

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

  const optimisticSaved: Saved = {
    id: "optimistic",
    guardian_id: guardianId!,
    event: { id: eventId } as Saved["event"],
    created_at: "",
    updated_at: "",
  };

  const createSaved = useCreateSaved(
    optimisticOptions((items) => [...items, optimisticSaved])
  );

  const deleteSaved = useDeleteSaved(
    optimisticOptions((items) => items.filter((s) => s.event.id !== eventId))
  );

  if (!guardianId) return null;

  const isPending = createSaved.isPending || deleteSaved.isPending;

  const handlePress = () => {
    if (isPending) return;
    if (isBookmarked && savedEntry && savedEntry.id !== "optimistic") {
      deleteSaved.mutate({ id: savedEntry.id });
    } else if (!isBookmarked) {
      createSaved.mutate({ data: { event_id: eventId, guardian_id: guardianId } });
    }
  };

  return (
    <TouchableOpacity onPress={handlePress} activeOpacity={0.7} disabled={isPending}>
      <MaterialIcons
        name={isBookmarked ? "bookmark" : "bookmark-border"}
        size={40}
        color={isPending ? AppColors.placeholderText : AppColors.primaryText}
      />
    </TouchableOpacity>
  );
}
