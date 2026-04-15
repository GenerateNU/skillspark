import { useRouter } from "expo-router";

type BackNavigationParams = {
  from?: string;
  category?: string;
};

export function useEventBackNavigation({ from, category }: BackNavigationParams) {
  const router = useRouter();

  return () => {
    switch (from) {
      case "event-categories":
        router.push({ pathname: "/(app)/(tabs)/event-categories", params: { category } });
        break;
      default:
        router.navigate("/");
    }
  };
}
