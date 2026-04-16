import { Tabs } from "expo-router";
import React from "react";

import { FloatingTabBar } from "@/components/floating-tab-bar";
import { useTranslation } from "react-i18next";

export default function TabLayout() {
  const { t: translate } = useTranslation();

  return (
    <Tabs
      tabBar={(props) => <FloatingTabBar {...props} />}
      screenOptions={{
        headerShown: false,
        sceneStyle: { backgroundColor: "#fff" },
      }}
    >
      <Tabs.Screen name="index" options={{ title: translate("nav.home") }} />
      <Tabs.Screen name="map" options={{ title: translate("nav.map") }} />
      <Tabs.Screen
        name="activity"
        options={{ title: translate("nav.activity") }}
      />
      <Tabs.Screen
        name="profile"
        options={{ title: translate("nav.profile") }}
      />
      <Tabs.Screen name="org" options={{ href: null }} />
      <Tabs.Screen name="family" options={{ href: null }} />
      <Tabs.Screen name="settings" options={{ href: null }} />
      <Tabs.Screen name="language" options={{ href: null }} />
      <Tabs.Screen name="terms-and-conditions" options={{ href: null }} />
      <Tabs.Screen name="privacy-policy" options={{ href: null }} />
      <Tabs.Screen name="payment" options={{ href: null }} />
      <Tabs.Screen name="saved" options={{ href: null }} />
      <Tabs.Screen name="event/[id]" options={{ href: null }} />
      <Tabs.Screen name="child/[id]" options={{ href: null }} />
    </Tabs>
  );
}
