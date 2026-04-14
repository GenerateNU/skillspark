import type { Child } from "@skillspark/api-client";

export interface ChildRegistration {
  child: Child;
  registrationId: string;
}

export interface RegistrationCardData {
  event_id: string;
  event_occurrence_id: string;
  image_uri: string;
  start_time: Date;
  end_time: Date;
  title: string;
  childRegistrations: ChildRegistration[];
  childColorMap: Record<string, string>;
  location: string;
  price: number;
  hasReviewed: boolean;
  onClickRemove: () => void;
  onClickReview: () => void;
}

export interface RegistrationCardProps {
  data: RegistrationCardData;
}

export const formatTime = (d: Date) =>
  d.toLocaleTimeString([], { hour: "2-digit", minute: "2-digit" });

export const formatDate = (d: Date) =>
  d.toLocaleDateString([], { weekday: "short", month: "short", day: "numeric" });
