import type { EventOccurrence } from "@skillspark/api-client";
import i18n from "@/i18n";

export function formatSectionDate(dateStr: string): string {
  const date = new Date(dateStr);
  const today = new Date();
  if (date.toDateString() === today.toDateString()) return i18n.t("time.today");
  return date.toLocaleDateString("en-US", { weekday: "short", day: "numeric" });
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

export function formatAgeRange(min: number, max: number): string {
  if (!min && !max) return "";
  if (min === max) return i18n.t("occurrence.ages", { min });
  return i18n.t("occurrence.agesRange", { min, max });
}

export function groupOccurrencesByDate(
  occurrences: EventOccurrence[]
): { label: string; items: EventOccurrence[] }[] {
  const sorted = [...occurrences].sort(
    (a, b) =>
      new Date(a.start_time).getTime() - new Date(b.start_time).getTime()
  );

  const groups: Map<string, EventOccurrence[]> = new Map();
  for (const occ of sorted) {
    const key = new Date(occ.start_time).toDateString();
    if (!groups.has(key)) groups.set(key, []);
    groups.get(key)!.push(occ);
  }

  return Array.from(groups.entries()).map(([, items]) => ({
    label: formatSectionDate(items[0].start_time),
    items,
  }));
}
