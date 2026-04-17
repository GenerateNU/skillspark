export { setCurrentLanguage } from "./apiClient";

// Re-export everything from generated
export {
  type HTTPStatusCodes,
  type HTTPStatusCode1xx,
  type HTTPStatusCode2xx,
  type HTTPStatusCode3xx,
  type HTTPStatusCode4xx,
  type HTTPStatusCode5xx,
} from "./generated/event-occurrences/event-occurrences";
export * from "./generated/event-occurrences/event-occurrences";
export * from "./generated/events/events";
export * from "./generated/skillSparkAPI.schemas";
export * from "./generated/child/child";
export * from "./generated/guardians/guardians";
export * from "./generated/managers/managers";
export * from "./generated/schools/schools";
export * from "./generated/saved/saved";
export * from "./generated/auth/auth";
export * from "./generated/registrations/registrations";
export * from "./generated/payments/payments";
export * from "./generated/emergency-contacts/emergency-contacts";
export * from "./generated/organizations/organizations";
export * from "./generated/locations/locations";
export * from "./generated/review/review";
export * from "./generated/user/user";
export * from "./generated/recommendations/recommendations";
export * from "./generated/search/search";
