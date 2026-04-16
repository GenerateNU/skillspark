import React from "react";
import { View, TouchableOpacity, useColorScheme, Modal } from "react-native";
import { ThemedText } from "@/components/themed-text";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors, AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";

export default function DeletePaymentMethodModal({
  visible,
  deleting,
  onConfirm,
  onClose,
}: {
  visible: boolean;
  deleting: boolean;
  onConfirm: () => void;
  onClose: () => void;
}) {
  const { t: translate } = useTranslation();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];

  return (
    <Modal
      visible={visible}
      transparent
      animationType="fade"
      onRequestClose={onClose}
    >
      <View
        className="flex-1 items-center justify-center px-6"
        style={{ backgroundColor: "rgba(0,0,0,0.45)" }}
      >
        <View
          className="w-full rounded-2xl px-6 pt-6 pb-4"
          style={{ backgroundColor: theme.background }}
        >
          {/* Header */}
          <View className="flex-row items-center gap-2 mb-3">
            <IconSymbol name="creditcard" size={20} color={theme.text} />
            <ThemedText className="text-lg font-nunito-bold">
              {translate("payment.deleteModal.title")}
            </ThemedText>
          </View>

          {/* Description */}
          <ThemedText
            className="text-sm font-nunito mb-5"
            style={{ color: theme.icon }}
          >
            {translate("payment.deleteModal.description")}
          </ThemedText>

          {/* Actions */}
          <View className="flex-row justify-end gap-6 pt-2">
            <TouchableOpacity
              onPress={onClose}
              disabled={deleting}
              activeOpacity={0.6}
            >
              <ThemedText
                className="text-[15px] font-nunito"
                style={{ color: theme.icon }}
              >
                {translate("common.cancel")}
              </ThemedText>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={onConfirm}
              disabled={deleting}
              activeOpacity={0.6}
            >
              <ThemedText
                className="text-[15px] font-nunito-semibold"
                style={{
                  color: deleting ? AppColors.danger + "80" : AppColors.danger,
                }}
              >
                {deleting
                  ? translate("payment.deleteModal.deleting")
                  : translate("payment.deleteModal.confirm")}
              </ThemedText>
            </TouchableOpacity>
          </View>
        </View>
      </View>
    </Modal>
  );
}
