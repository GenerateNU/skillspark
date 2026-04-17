import {
  Modal,
  Share,
  Linking,
  View,
  Text,
  Pressable,
  Platform,
} from "react-native";
import * as Clipboard from "expo-clipboard";
import { CircleIcon } from "@/components/CircleIcon";
import { Image } from "expo-image";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import FontAwesome5 from "@expo/vector-icons/FontAwesome5";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";
import { Gesture, GestureDetector } from "react-native-gesture-handler";
import Animated, {
  useSharedValue,
  useAnimatedStyle,
  withSpring,
  withTiming,
  runOnJS,
} from "react-native-reanimated";
import { useEffect } from "react";
import { useTranslation } from "react-i18next";

const CIRCLE_BG = "#EFEFEF";

const MESSENGER_STORE_URL =
  Platform.OS === "android"
    ? "https://play.google.com/store/apps/details?id=com.facebook.orca"
    : "https://apps.apple.com/us/app/messenger/id454638411";

async function openMessenger(url: string) {
  const scheme = `fb-messenger://share?link=${encodeURIComponent(url)}`;
  try {
    const supported = await Linking.canOpenURL(scheme);
    if (supported) {
      await Linking.openURL(scheme);
    } else {
      await Linking.openURL(MESSENGER_STORE_URL);
    }
  } catch {
    await Linking.openURL(MESSENGER_STORE_URL);
  }
}

function makeAppIcons(text: string, url: string) {
  return [
    {
      name: "Messages",
      bg: "#34C759",
      onPress: () => {
        const smsUrl =
          Platform.OS === "ios"
            ? `sms:&body=${encodeURIComponent(text)}`
            : `sms:?body=${encodeURIComponent(text)}`;
        Linking.openURL(smsUrl).catch(() => {});
      },
      render: () => <IconSymbol name="message.fill" size={22} color="#fff" />,
    },
    {
      name: "Messenger",
      bg: "#0084FF",
      onPress: () => openMessenger(url),
      render: () => (
        <FontAwesome5 name="facebook-messenger" size={20} color="#fff" />
      ),
    },
    {
      name: "Line",
      bg: "#00B900",
      onPress: () =>
        Linking.openURL(
          `https://line.me/R/msg/text/?${encodeURIComponent(text)}`,
        ),
      render: () => <FontAwesome5 name="line" size={20} color="#fff" />,
    },
    {
      name: "WhatsApp",
      bg: "#25D366",
      onPress: () =>
        Linking.openURL(`https://wa.me/?text=${encodeURIComponent(text)}`),
      render: () => <FontAwesome5 name="whatsapp" size={20} color="#fff" />,
    },
  ];
}

interface ShareModalProps {
  visible: boolean;
  onClose: () => void;
  name: string;
  imageUrl?: string;
  shareUrl: string;
  message?: string;
}

export function ShareModal({
  visible,
  onClose,
  name,
  imageUrl,
  shareUrl,
  message,
}: ShareModalProps) {
  const insets = useSafeAreaInsets();
  const { t } = useTranslation();
  const shareMessage = message ?? t("share.defaultMessage", { name });
  const fullText = `${shareMessage}\n${shareUrl}`;
  const appIcons = makeAppIcons(fullText, shareUrl);
  const translateY = useSharedValue(0);

  useEffect(() => {
    if (visible) {
      translateY.value = 0;
    }
  }, [visible]);

  const panGesture = Gesture.Pan()
    .onUpdate((e) => {
      translateY.value = Math.max(0, e.translationY);
    })
    .onEnd((e) => {
      if (e.translationY > 100 || e.velocityY > 500) {
        translateY.value = withTiming(600, { duration: 250 }, () => {
          runOnJS(onClose)();
        });
      } else {
        translateY.value = withSpring(0, { damping: 20, stiffness: 200 });
      }
    });

  const animatedStyle = useAnimatedStyle(() => ({
    transform: [{ translateY: translateY.value }],
  }));

  const handleShare = async () => {
    await Share.share({ message: fullText, url: shareUrl });
  };

  const handleCopyUrl = async () => {
    await Clipboard.setStringAsync(shareUrl);
  };

  return (
    <Modal
      visible={visible}
      transparent
      animationType="slide"
      onRequestClose={onClose}
    >
      <View style={{ flex: 1, justifyContent: "flex-end" }}>
        {/* Backdrop */}
        <Pressable style={{ flex: 1 }} onPress={onClose} />

        {/* Sheet */}
        <GestureDetector gesture={panGesture}>
          <Animated.View
            style={[
              {
                height: "50%",
                backgroundColor: "#fff",
                borderTopLeftRadius: 20,
                borderTopRightRadius: 20,
                paddingBottom: insets.bottom,
                overflow: "hidden",
              },
              animatedStyle,
            ]}
          >
            {/* Grabber */}
            <View className="items-center pt-3 pb-2">
              <View
                className="w-10 h-1 rounded-full"
                style={{ backgroundColor: "#D1D1D6" }}
              />
            </View>

            {/* Thumbnail */}
            <View className="items-center mt-2 mb-3">
              <View
                className="w-[80px] h-[80px] rounded-2xl overflow-hidden"
                style={{ backgroundColor: AppColors.imagePlaceholder }}
              >
                {imageUrl ? (
                  <Image
                    source={{ uri: imageUrl }}
                    style={{ width: "100%", height: "100%" }}
                    contentFit="cover"
                  />
                ) : null}
              </View>
            </View>

            {/* Text */}
            <View className="px-6 items-center mb-5">
              <Text
                className="mb-3 text-center"
                style={{
                  fontFamily: FontFamilies.bold,
                  fontSize: 20,
                  color: AppColors.primaryText,
                }}
              >
                {name}
              </Text>
              <Text
                className="text-center mb-1"
                style={{
                  fontFamily: FontFamilies.regular,
                  fontSize: FontSizes.sm,
                  color: AppColors.secondaryText,
                  lineHeight: 20,
                }}
              >
                {shareMessage}
              </Text>
            </View>

            {/* App icons */}
            <View className="flex-row justify-around px-4 mb-4">
              {appIcons.map((app) => (
                <CircleIcon
                  key={app.name}
                  bg={app.bg}
                  label={app.name}
                  onPress={app.onPress}
                >
                  {app.render()}
                </CircleIcon>
              ))}
            </View>

            {/* Copy link + Share */}
            <View className="flex-row justify-center gap-8 px-4">
              <CircleIcon
                bg={CIRCLE_BG}
                label={t("share.copyLink")}
                onPress={handleCopyUrl}
              >
                <MaterialIcons
                  name="link"
                  size={22}
                  color={AppColors.primaryText}
                />
              </CircleIcon>
              <CircleIcon
                bg={CIRCLE_BG}
                label={t("share.share")}
                onPress={handleShare}
              >
                <IconSymbol
                  name="square.and.arrow.up"
                  size={22}
                  color={AppColors.primaryText}
                />
              </CircleIcon>
            </View>
          </Animated.View>
        </GestureDetector>
      </View>
    </Modal>
  );
}
