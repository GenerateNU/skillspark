import i18n from "@/i18n";
import type { EventOccurrence } from "@skillspark/api-client";

export function formatDuration(
  start: string,
  end: string,
  labels: { hr: string; min: string } = { hr: "hr", min: "min" },
) {
  const mins = Math.round(
    (new Date(end).getTime() - new Date(start).getTime()) / 60000,
  );
  return mins >= 60
    ? `${Math.round(mins / 60)} ${labels.hr}`
    : `${mins} ${labels.min}`;
}

export function isWithinNext7Days(dateStr: string): boolean {
  const date = new Date(dateStr);
  const now = new Date();
  const in7Days = new Date(now.getTime() + 7 * 24 * 60 * 60 * 1000);
  return date >= now && date <= in7Days;
}

export function formatEventDate(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString("en-US", {
    weekday: "long",
    month: "long",
    day: "numeric",
  });
}

export function formatEventTime(start: string, end: string): string {
  const fmt = (d: string) =>
    new Date(d).toLocaleTimeString("en-US", {
      hour: "2-digit",
      minute: "2-digit",
    });
  return `${fmt(start)} - ${fmt(end)}`;
}

export function formatModalTime(dateStr: string): string {
  return new Date(dateStr).toLocaleTimeString("en-US", {
    hour: "numeric",
    minute: "2-digit",
    hour12: true,
  });
}

export function formatSectionDate(dateStr: string): string {
  const date = new Date(dateStr);
  const today = new Date();
  if (date.toDateString() === today.toDateString()) return i18n.t("time.today");
  return date.toLocaleDateString("en-US", { weekday: "short", day: "numeric" });
}

export function formatSectionMonth(dateStr: string): string {
  return new Date(dateStr).toLocaleDateString("en-US", { month: "long" });
}

export function formatTime(dateStr: string): string {
  return new Date(dateStr).toLocaleTimeString("en-US", {
    hour: "numeric",
    minute: "2-digit",
    hour12: true,
  });
}

export function formatPrice(cents: number, currency: string): string {
  const amount = cents / 100;
  if (currency?.toUpperCase() === "THB")
    return `฿${amount % 1 === 0 ? amount.toFixed(0) : amount.toFixed(2)}`;
  return `$${amount % 1 === 0 ? amount.toFixed(0) : amount.toFixed(2)}`;
}

export function formatAgeRange(
  min: number | null | undefined,
  max: number | null | undefined,
): string {
  if (!min && !max) return "";
  if (!max) return i18n.t("occurrence.agesOpen", { min });
  if (min === max) return i18n.t("occurrence.ages", { min });
  return i18n.t("occurrence.agesRange", { min, max });
}

export function formatLocation(occurrence: EventOccurrence): string {
  const loc = occurrence.location;
  const parts = [loc.district, loc.province].filter(Boolean);
  return parts.join(", ") || "Location";
}

export function formatAddress(occurrence: EventOccurrence): string {
  const loc = occurrence.location;
  const parts = [loc?.address_line1, loc?.address_line2, loc?.district].filter(
    Boolean,
  );
  return parts.join(", ");
}

export function filterFutureOccurrences(
  occurrences: EventOccurrence[],
): EventOccurrence[] {
  const now = new Date();
  return occurrences.filter((o) => new Date(o.start_time) > now);
}

export function extractResponseData<T>(resp: unknown): T[] {
  const d = resp as { data: T[] } | undefined;
  return Array.isArray(d?.data) ? d!.data : [];
}

export function arrayToMap<T extends { id: string }>(
  arr: T[],
): Record<string, T> {
  const map: Record<string, T> = {};
  arr.forEach((item) => {
    map[item.id] = item;
  });
  return map;
}
