import { useEffect, useRef, useState } from "react";
import * as Location from "expo-location";

const DEFAULT_LAT = "13.7563";
const DEFAULT_LNG = "100.5018";

export function useGeolocation() {
  const [lat, setLat] = useState<string | undefined>(DEFAULT_LAT);
  const [lng, setLng] = useState<string | undefined>(DEFAULT_LNG);
  const mounted = useRef(true);

  useEffect(() => {
    mounted.current = true;
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") return;
      const loc = await Location.getCurrentPositionAsync({});
      if (!mounted.current) return;
      setLat(String(loc.coords.latitude));
      setLng(String(loc.coords.longitude));
    })();
    return () => {
      mounted.current = false;
    };
  }, []);

  return { lat, lng };
}
