import { Colors, AppColors } from "@/constants/theme";
import { useColorScheme } from "@/hooks/use-color-scheme";
import { PaymentMethod } from "@skillspark/api-client";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { View, TouchableOpacity } from "react-native";
import { ThemedText } from "./themed-text";

export default function PaymentMethodRow({
  method,
  onDelete,
}: {
  method: PaymentMethod;
  onDelete: (id: string) => void;
}) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  const [pressed, setPressed] = useState(false);
  const handleDelete = () => {
    onDelete(method.id);
  };

  return (
    <View className="flex flex-row items-center justify-between mb-8">
      <View>
        <ThemedText className="text-base font-nunito mb-[6px]">
          {method.card.brand.toUpperCase()} · {translate("payment.creditCard")}
        </ThemedText>
        <ThemedText className="text-base font-nunito mb-2 tracking-widest">
          **** **** **** {method.card.last4}
        </ThemedText>
        <ThemedText
          className="text-sm font-nunito"
          style={{ color: theme.icon }}
        >
          {translate("payment.expires")} {method.card.exp_month}/
          {method.card.exp_year}
        </ThemedText>
      </View>
      <TouchableOpacity
        className="px-4 py-2 rounded-lg border-[1.5px] items-center justify-center"
        style={{
          borderColor: pressed ? AppColors.danger : theme.text,
          backgroundColor: pressed ? AppColors.danger : "transparent",
        }}
        onPress={() => {
          handleDelete();
        }}
        onPressIn={() => setPressed(true)}
        onPressOut={() => setPressed(false)}
        activeOpacity={1}
      >
        <ThemedText
          className="text-[13px] font-nunito text-center"
          style={{ color: pressed ? AppColors.white : theme.text }}
        >
          {translate("payment.delete")}
        </ThemedText>
      </TouchableOpacity>
    </View>
  );
}
