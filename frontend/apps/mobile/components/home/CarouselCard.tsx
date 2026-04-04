import Carousel from 'react-native-reanimated-carousel';
import { AppColors } from "@/constants/theme";
import { type EventOccurrence } from '@skillspark/api-client';
import { DiscoverBanner } from './DiscoverBanner';

export default function CarouselCard({ events, width }: { events: EventOccurrence[], width: number }) {
    // make height dynamic or something idk - make it match DiscoverBanner
  return (
    <Carousel
      width={width}
      height={188}
      data={events}
      loop
      renderItem={({ item }) => (
        <DiscoverBanner event={item} />
      )}
    />
  );
}