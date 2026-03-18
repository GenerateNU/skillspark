export function formatDuration(start: string, end: string) {
  const mins = Math.round(
    (new Date(end).getTime() - new Date(start).getTime()) / 60000
  );
  return mins >= 60 ? `${Math.round(mins / 60)} hr` : `${mins} min`;
}
