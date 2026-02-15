export interface LocationPin {
    id: string;
    title: string;
    description: string;
    latitude: number;
    longitude: number;
    rating: number;     
    members: number;     
    image?: string;     
  }
  
  export const MOCK_LOCATIONS: LocationPin[] = [
    {
      id: '1',
      title: 'Soccer Club',
      description: 'Intramural soccer club with great community...',
      latitude: 13.7467,
      longitude: 100.5350,
      rating: 4.5,
      members: 55,
    },
    {
      id: '2',
      title: 'Junior Robotics',
      description: 'Learn to build and code robots.',
      latitude: 13.7563,
      longitude: 100.5018,
      rating: 5.0,
      members: 12,
    },
    {
      id: '3',
      title: 'Astronomy Club',
      description: 'Stargazing events every Friday.',
      latitude: 13.8200,
      longitude: 100.5600,
      rating: 4.8,
      members: 30,
    },
    {
      id: '4',
      title: 'Piano Lessons',
      description: 'Beginner friendly music classes.',
      latitude: 13.7300,
      longitude: 100.5240,
      rating: 4.2,
      members: 8,
    },
  ];