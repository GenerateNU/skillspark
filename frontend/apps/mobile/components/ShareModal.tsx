import {
  Modal,
  Share,
  Linking,
  View,
  Text,
  Pressable,
} from "react-native";
import { Image } from "expo-image";
import { useSafeAreaInsets } from "react-native-safe-area-context";
import FontAwesome5 from "@expo/vector-icons/FontAwesome5";
import MaterialIcons from "@expo/vector-icons/MaterialIcons";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { AppColors, FontFamilies, FontSizes } from "@/constants/theme";

const CIRCLE_BG = "#EFEFEF";

const APP_STORE_URLS: Record<string, string> = {
  messenger: "https://apps.apple.com/app/messenger/id454638411",
  line:      "https://apps.apple.com/app/line/id443904275",
  whatsapp:  "https://apps.apple.com/app/whatsapp-messenger/id310633997",
};

async function openAppOrStore(appUrl: string, storeKey: string) {
  try {
    const supported = await Linking.canOpenURL(appUrl);
    if (supported) {
      await Linking.openURL(appUrl);
    } else {
      await Linking.openURL(APP_STORE_URLS[storeKey]);
    }
  } catch {
    await Linking.openURL(APP_STORE_URLS[storeKey]);
  }
}

function makeAppIcons(text: string, url: string) {
  return [
    {
      name: "Messages",
      bg: "#34C759",
      onPress: () => Linking.openURL(`sms:?body=${encodeURIComponent(text)}`),
      render: () => <IconSymbol name="message.fill" size={22} color="#fff" />,
    },
    {
      name: "Messenger",
      bg: "#0084FF",
      onPress: () => openAppOrStore(
        `fb-messenger://share?link=${encodeURIComponent(url)}`,
        "messenger"
      ),
      render: () => <FontAwesome5 name="facebook-messenger" size={20} color="#fff" />,
    },
    {
      name: "Line",
      bg: "#00B900",
      onPress: () => openAppOrStore(
        `line://msg/text/${encodeURIComponent(text)}`,
        "line"
      ),
      render: () => <FontAwesome5 name="line" size={20} color="#fff" />,
    },
    {
      name: "WhatsApp",
      bg: "#25D366",
      onPress: () => openAppOrStore(
        `whatsapp://send?text=${encodeURIComponent(text)}`,
        "whatsapp"
      ),
      render: () => <FontAwesome5 name="whatsapp" size={20} color="#fff" />,
    },
  ];
}

function CircleIcon({
  bg,
  children,
  label,
  onPress,
}: {
  bg: string;
  children: React.ReactNode;
  label: string;
  onPress: () => void;
}) {
  return (
    <Pressable onPress={onPress} className="items-center gap-1.5" style={{ width: 64 }}>
      <View
        className="w-[54px] h-[54px] rounded-full items-center justify-center"
        style={{ backgroundColor: bg }}
      >
        {children}
      </View>
      <Text
        numberOfLines={1}
        style={{
          fontFamily: FontFamilies.regular,
          fontSize: 11,
          color: AppColors.secondaryText,
        }}
      >
        {label}
      </Text>
    </Pressable>
  );
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
  const shareMessage = message ?? `Check out ${name} on SkillSpark!`;
  const fullText = `${shareMessage}\n${shareUrl}`;
  const appIcons = makeAppIcons(fullText, shareUrl);

  const handleShare = async () => {
    await Share.share({ message: fullText, url: shareUrl });
  };

  const handleCopyUrl = async () => {
    await Share.share({ message: shareUrl, url: shareUrl });
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
        <View
          style={{
            height: "50%",
            backgroundColor: "#fff",
            borderTopLeftRadius: 20,
            borderTopRightRadius: 20,
            paddingBottom: insets.bottom,
            overflow: "hidden",
          }}
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
              className="mb-0.5 text-center"
              style={{
                fontFamily: FontFamilies.bold,
                fontSize: 20,
                color: AppColors.primaryText,
              }}
            >
              {name}
            </Text>
            <Text
              className="mb-3 text-center"
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: FontSizes.sm,
                color: AppColors.subtleText,
              }}
            >
              4.5 Stars (140)
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
            <Text
              style={{
                fontFamily: FontFamilies.regular,
                fontSize: 12,
                color: AppColors.subtleText,
              }}
            >
              Limitations apply, see attached
            </Text>
          </View>

          {/* App icons */}
          <View className="flex-row justify-around px-4 mb-4">
            {appIcons.map((app) => (
              <CircleIcon key={app.name} bg={app.bg} label={app.name} onPress={app.onPress}>
                {app.render()}
              </CircleIcon>
            ))}
          </View>

          {/* Copy link + Share */}
          <View className="flex-row justify-center gap-8 px-4">
            <CircleIcon bg={CIRCLE_BG} label="Copy link" onPress={handleCopyUrl}>
              <MaterialIcons name="link" size={22} color={AppColors.primaryText} />
            </CircleIcon>
            <CircleIcon bg={CIRCLE_BG} label="Share" onPress={handleShare}>
              <MaterialIcons name="ios-share" size={22} color={AppColors.primaryText} />
            </CircleIcon>
          </View>
        </View>
      </View>
    </Modal>
  );
}
