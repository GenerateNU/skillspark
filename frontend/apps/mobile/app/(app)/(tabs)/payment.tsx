import React, { useCallback, useEffect, useState } from "react";
import {
  View,
  TouchableOpacity,
  useColorScheme,
  ActivityIndicator,
} from "react-native";
import { useFocusEffect, useRouter } from "expo-router";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import { ThemedText } from "@/components/themed-text";
import { ThemedView } from "@/components/themed-view";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { Colors, AppColors } from "@/constants/theme";
import { useTranslation } from "react-i18next";
import {
  detachGuardianPaymentMethod,
  getGuardianPaymentMethods,
  type GetPaymentMethodsByGuardianIDOutputBody,
  type PaymentMethod,
} from "@skillspark/api-client";
import { useGuardian } from "@/hooks/use-guardian";
import PaymentMethodRow from "@/components/PaymentMethodRow";
import CardForm from "@/components/CardForm";
import { ErrorMessage } from "@/components/ErrorMessage";
import DeletePaymentMethodModal from "@/components/DeletePaymentMethodModal";
import { useAuthContext } from "@/hooks/use-auth-context";

export default function PaymentScreen() {
  const router = useRouter();
  const insets = useSafeAreaInsets();
  const theme = Colors.light;
  const { t: translate } = useTranslation();
  const [editingCard, setEditingCard] = useState<boolean>(false);
  const [paymentMethods, setPaymentMethods] = useState<PaymentMethod[]>([]);
  const [loading, setLoading] = useState<boolean>(true);
  const [error, setError] = useState<string | null>(null);
  const [pendingDeleteId, setPendingDeleteId] = useState<string | null>(null);
  const [deleting, setDeleting] = useState<boolean>(false);
  const { guardianId } = useAuthContext();
  const { guardian } = useGuardian(guardianId);

  async function fetchPaymentMethods(): Promise<void> {
    if (!guardian) return;
    try {
      const res = await getGuardianPaymentMethods(guardian.id);
      if (res.status !== 200 && res.status !== 201) throw res.data;
      const methods = (res.data as GetPaymentMethodsByGuardianIDOutputBody)
        .payment_methods;
      setPaymentMethods(methods ? methods : []);
      setError(null);
    } catch (e) {
      setError("Failed to fetch payment methods");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    fetchPaymentMethods();
  }, [guardian]);

  useFocusEffect(
    useCallback(() => {
      return () => {
        setError(null);
      };
    }, [])
  );

  async function confirmDelete(): Promise<void> {
    if (!pendingDeleteId) return;
    try {
      if (!guardian) {
        throw new Error("No guardian is logged in.")
      }
      setDeleting(true);
      const res = await detachGuardianPaymentMethod({
        payment_method_id: pendingDeleteId,
        guardian_id: guardian!.id
      });
      if (res.status !== 200 && res.status !== 204) {
        throw res.data;
      }
      setPaymentMethods(prev => prev.filter(pm => pm.id !== pendingDeleteId));
    } catch (e) {
      setError("Failed to delete payment method");
    } finally {
      setDeleting(false);
      setPendingDeleteId(null);
    }
  }

  return (
    <ThemedView className="flex-1" style={{ paddingTop: insets.top }}>
      <View className="flex-row items-center justify-between px-5 py-[14px]">
        <TouchableOpacity
          onPress={() => {
            setEditingCard(false);
            router.navigate("/profile");
          }}
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
          <ActivityIndicator
            size="small"
            color={AppColors.primaryBlue}
            style={{ marginBottom: 32 }}
          />
        ) : editingCard ? (
          <CardForm
            onSave={async () => {
              setLoading(true);
              await new Promise(resolve => setTimeout(resolve, 2000));
              await fetchPaymentMethods();
              setEditingCard(false);
              setLoading(false);
            }}
            onCancel={() => setEditingCard(false)}
          />
        ) : paymentMethods.length > 0 ? (
          paymentMethods.map((method) => (
            <PaymentMethodRow
              method={method}
              onDelete={(id) => setPendingDeleteId(id)}
              key={method.id}
            />
          ))
        ) : (
          <ThemedText
            className="text-base font-nunito mb-8"
            style={{ color: theme.icon }}
          >
            {translate("payment.noCard")}
          </ThemedText>
        )}

        {!editingCard && (
          <View className="flex-row gap-4">
            <TouchableOpacity
              className="flex-1 py-[14px] rounded-lg items-center justify-center"
              style={{ backgroundColor: AppColors.primaryBlue }}
              onPress={() => setEditingCard(true)}
              activeOpacity={0.8}
            >
              <ThemedText className="text-white text-[15px] font-nunito-semibold">
                {translate("payment.updateBilling")}
              </ThemedText>
            </TouchableOpacity>
          </View>
        )}
      </View>

      {error && <ErrorMessage message={error} />}

      <DeletePaymentMethodModal
        visible={pendingDeleteId !== null}
        deleting={deleting}
        onConfirm={confirmDelete}
        onClose={() => setPendingDeleteId(null)}
      />
    </ThemedView>
  );
}
