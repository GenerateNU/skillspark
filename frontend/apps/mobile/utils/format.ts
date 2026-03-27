export function formatDuration(
  start: string,
  end: string,
  labels: { hr: string; min: string } = { hr: 'hr', min: 'min' }
) {
  const mins = Math.round(
    (new Date(end).getTime() - new Date(start).getTime()) / 60000
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
    new Date(d).toLocaleTimeString("en-US", { hour: "2-digit", minute: "2-digit" });
  return `${fmt(start)} - ${fmt(end)}`;
}
