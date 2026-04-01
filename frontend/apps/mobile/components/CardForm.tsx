import { Colors, AppColors } from "@/constants/theme";
import { useGuardian } from "@/hooks/use-guardian";
import {
  createGuardianSetupIntent,
  type CreateSetupIntentOutputBody,
} from "@skillspark/api-client";
import { useStripe, CardField } from "@stripe/stripe-react-native";
import { useState } from "react";
import { useTranslation } from "react-i18next";
import { useColorScheme, View, TouchableOpacity } from "react-native";
import { ThemedText } from "./themed-text";
import { ErrorMessage } from "./ErrorMessage";

export default function CardForm({
  onSave,
  onCancel,
}: {
  onSave: () => void;
  onCancel: () => void;
}) {
  const [cardComplete, setCardComplete] = useState(false);
  const [saving, setSaving] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const { confirmSetupIntent } = useStripe();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();
  const { guardian } = useGuardian();

  async function handleSave(): Promise<void> {
    try {
      setError(null);
      if (!guardian) throw new Error("No user is authenticated");
      setSaving(true);

      const res = await createGuardianSetupIntent(guardian.id);
      if (res.status !== 200 && res.status !== 201)
        throw new Error("Failed to create setup intent");

      const clientSecret = (res.data as CreateSetupIntentOutputBody)
        .client_secret;
      const { error: stripeError } = await confirmSetupIntent(clientSecret, {
        paymentMethodType: "Card",
      });
      if (stripeError)
        throw new Error("Failed to confirm payment method, please try again.");

      onSave();
    } catch (e) {
      setError(e instanceof Error ? e.message : "An unexpected error occurred");
    } finally {
      setSaving(false);
    }
  }

  return (
    <View className="mb-8">
      <CardField
        postalCodeEnabled={false}
        style={{ height: 50, width: "100%", marginBottom: 16 }}
        onCardChange={(details) => setCardComplete(details.complete)}
        cardStyle={{
          backgroundColor: theme.background,
          textColor: theme.text,
          borderColor: theme.text,
          borderWidth: 1,
          borderRadius: 8,
        }}
      />
      <View className="flex-row gap-3 mt-2">
        <TouchableOpacity
          className="flex-1 py-[14px] rounded-lg border-[1.5px] items-center justify-center"
          style={{ borderColor: theme.text }}
          onPress={onCancel}
          disabled={saving}
          activeOpacity={0.8}
        >
          <ThemedText
            className="text-[15px] font-nunito"
            style={{ color: theme.text }}
          >
            {translate("common.cancel")}
          </ThemedText>
        </TouchableOpacity>
        <TouchableOpacity
          className="flex-1 py-[14px] rounded-lg items-center justify-center"
          style={{
            backgroundColor: AppColors.primaryBlue,
            opacity: cardComplete && !saving ? 1 : 0.5,
          }}
          onPress={handleSave}
          disabled={!cardComplete || saving}
          activeOpacity={0.8}
        >
          <ThemedText className="text-white text-[15px] font-nunito-semibold">
            {saving ? translate("common.saving") : translate("payment.save")}
          </ThemedText>
        </TouchableOpacity>
      </View>
      {error && <ErrorMessage message={error} />}
    </View>
  );
}
