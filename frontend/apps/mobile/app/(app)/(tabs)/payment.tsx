import React, { useEffect, useState } from "react";
import { View, TouchableOpacity, useColorScheme, Modal, ActivityIndicator } from "react-native";
import { useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors, AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import { CardField, confirmSetupIntent } from "@stripe/stripe-react-native";
import {
  createGuardianSetupIntent,
  type CreateSetupIntentOutputBody,
  getGuardianPaymentMethods,
  type GetPaymentMethodsByGuardianIDOutputBody,
  type PaymentMethod,
} from "@skillspark/api-client";
import { useGuardian } from "@/hooks/use-guardian";

function CardFormModal({ visible, onClose }: { visible: boolean; onClose: () => void }) {
  const [cardComplete, setCardComplete] = useState(false);
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();
  const { guardian } = useGuardian();


  async function handleSave(): Promise<void> {
    try {
      if (!guardian) {
        throw new Error("No user is authenticated");
      }
      const res = await createGuardianSetupIntent(guardian.id);
      if (res.status !== 200 && res.status !== 201) throw res.data;

      const clientSecret = (res.data as CreateSetupIntentOutputBody).client_secret;
      const { error } = await confirmSetupIntent(clientSecret, { paymentMethodType: "Card" });
      if (error) throw new Error(error.message);

      console.log("submitted");
      onClose();
    } catch (e) {
      console.error(e);
    }
  }

  return (
    <Modal visible={visible} animationType="slide" presentationStyle="pageSheet" onRequestClose={onClose}>
      <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
        <View className="flex-row items-center justify-between px-5 py-[14px]">
          <TouchableOpacity
            onPress={onClose}
            className="w-10 justify-center items-start"
            hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
          >
            <IconSymbol name="xmark" size={24} color={theme.text} />
          </TouchableOpacity>
          <ThemedText className="text-xl text-center font-nunito-bold">
            {translate("payment.updateBilling")}
          </ThemedText>
          <View className="w-10" />
        </View>

        <View className="px-5 pt-5 flex flex-col gap-6">
          <CardField
            postalCodeEnabled={false}
            style={{ height: 50, width: "100%" }}
            onCardChange={(details) => setCardComplete(details.complete)}
            cardStyle={{
              backgroundColor: theme.background,
              textColor: theme.text,
              borderColor: theme.text,
              borderWidth: 1,
              borderRadius: 8,
            }}
          />
          <TouchableOpacity
            className="py-[14px] rounded-lg items-center justify-center"
            style={{ backgroundColor: AppColors.primaryBlue, opacity: cardComplete ? 1 : 0.5 }}
            onPress={handleSave}
            disabled={!cardComplete}
            activeOpacity={0.8}
          >
            <ThemedText className="text-white text-[15px] font-nunito-semibold">
              {translate("payment.save")}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </ThemedView>
    </Modal>
  );
}

function PaymentMethodRow({ method }: { method: PaymentMethod }) {
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();

  return (
    <View>
      <ThemedText className="text-base font-nunito mb-[6px]">
        {method.card.brand.toUpperCase()} · {translate("payment.creditCard")}
      </ThemedText>
      <ThemedText className="text-base font-nunito mb-2 tracking-widest">
        **** **** **** {method.card.last4}
      </ThemedText>
      <ThemedText className="text-sm font-nunito mb-8" style={{ color: theme.icon }}>
        {translate("payment.expires")} {method.card.exp_month}/{method.card.exp_year}
      </ThemedText>
    </View>
  );
}

// ─── Payment Screen ───────────────────────────────────────────────────────────

export default function PaymentScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const colorScheme = useColorScheme();
  const theme = Colors[colorScheme ?? "light"];
  const { t: translate } = useTranslation();
  const [showCardModal, setShowCardModal] = useState<boolean>(false);
  const [paymentMethods, setPaymentMethods] = useState<PaymentMethod[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const { guardian } = useGuardian();

  const handleDelete = () => {};

  useEffect(() => {
    async function fetchPaymentMethods(): Promise<void> {
      try {
        if (!guardian) {
          throw new Error("No guardian is set.")
        }
        const res = await getGuardianPaymentMethods(guardian.id);
        if (res.status !== 200 && res.status !== 201) throw res.data;
        setPaymentMethods((res.data as GetPaymentMethodsByGuardianIDOutputBody).payment_methods);
      } catch (e) {
        console.error(e);
      } finally {
        setLoading(false);
      }
    }
    fetchPaymentMethods();
  }, []);

  const primaryMethod = paymentMethods[0];

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => router.navigate("/profile")}
          className="w-10 justify-center items-start"
          hitSlop={{ top: 10, bottom: 10, left: 10, right: 10 }}
        >
          <IconSymbol name="chevron.left" size={24} color={theme.text} />
        </TouchableOpacity>
        <ThemedText className="text-xl text-center font-nunito-bold">
          {translate("payment.title")}
        </ThemedText>
        <View className="w-10" />
      </View>

      <View className="px-5 pt-5">
        <ThemedText className="text-[22px] font-nunito-bold mb-5">
          {translate("payment.manageBilling")}
        </ThemedText>

        {loading ? (
          <ActivityIndicator size="small" color={AppColors.primaryBlue} style={{ marginBottom: 32 }} />
        ) : primaryMethod ? (
          <PaymentMethodRow method={primaryMethod} />
        ) : (
          <ThemedText className="text-base font-nunito mb-8" style={{ color: theme.icon }}>
            {translate("payment.noCard")}
          </ThemedText>
        )}

        <View className="flex-row gap-4">
          <TouchableOpacity
            className="flex-1 py-[14px] rounded-lg items-center justify-center"
            style={{ backgroundColor: AppColors.primaryBlue }}
            onPress={() => setShowCardModal(true)}
            activeOpacity={0.8}
          >
            <ThemedText className="text-white text-[15px] font-nunito-semibold">
              {translate("payment.updateBilling")}
            </ThemedText>
          </TouchableOpacity>

          <TouchableOpacity
            className="flex-1 py-[14px] rounded-lg border-[1.5px] items-center justify-center"
            style={{ borderColor: theme.text }}
            onPress={handleDelete}
            activeOpacity={0.8}
          >
            <ThemedText className="text-[15px] font-nunito" style={{ color: theme.text }}>
              {translate("payment.delete")}
            </ThemedText>
          </TouchableOpacity>
        </View>
      </View>

      <CardFormModal visible={showCardModal} onClose={() => setShowCardModal(false)} />
    </ThemedView>
  );
}