import { useRouter } from "expo-router";

type BackNavigationParams = {
  from?: string;
};

export function useEventBackNavigation({ from }: BackNavigationParams) {
  const router = useRouter();

  return () => {
    switch (from) {
      case "event-categories":
        router.back();
        break;
      default:
        router.navigate("/");
    }
  };
}
