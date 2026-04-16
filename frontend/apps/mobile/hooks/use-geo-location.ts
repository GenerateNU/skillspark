import { useEffect, useState } from "react";
import * as Location from "expo-location";

const DEFAULT_LAT = 13.7563;
const DEFAULT_LNG = 100.5018;

export function useGeoLocation() {
  const [lat, setLat] = useState(DEFAULT_LAT);
  const [lng, setLng] = useState(DEFAULT_LNG);

  useEffect(() => {
    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();
      if (status !== "granted") return;
      const loc = await Location.getCurrentPositionAsync({});
      setLat(loc.coords.latitude);
      setLng(loc.coords.longitude);
    })();
  }, []);

  return { lat, lng };
}
