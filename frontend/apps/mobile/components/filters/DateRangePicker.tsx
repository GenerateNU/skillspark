import { useState } from "react";
import { Pressable, Text, View } from "react-native";
import { IconSymbol } from "@/components/ui/icon-symbol";
import { useTranslation } from "react-i18next";

type Props = {
  startDate?: Date;
  endDate?: Date;
  onChange: (start?: Date, end?: Date) => void;
};

function sameDay(a: Date, b: Date) {
  return (
    a.getFullYear() === b.getFullYear() &&
    a.getMonth() === b.getMonth() &&
    a.getDate() === b.getDate()
  );
}

function startOfDay(d: Date): Date {
  const r = new Date(d);
  r.setHours(0, 0, 0, 0);
  return r;
}

function endOfDay(d: Date): Date {
  const r = new Date(d);
  r.setHours(23, 59, 59, 999);
  return r;
}

export function DateRangePicker({ startDate, endDate, onChange }: Props) {
  const { t } = useTranslation();
  const months: string[] = t("calendar.months", {
    returnObjects: true,
  }) as string[];
  const dayNames: string[] = t("calendar.days", {
    returnObjects: true,
  }) as string[];

  const today = startOfDay(new Date());
  const [viewYear, setViewYear] = useState(today.getFullYear());
  const [viewMonth, setViewMonth] = useState(today.getMonth());

  function prevMonth() {
    if (viewMonth === 0) {
      setViewYear((y) => y - 1);
      setViewMonth(11);
    } else setViewMonth((m) => m - 1);
  }

  function nextMonth() {
    if (viewMonth === 11) {
      setViewYear((y) => y + 1);
      setViewMonth(0);
    } else setViewMonth((m) => m + 1);
  }

  function handlePress(day: Date) {
    const d = startOfDay(day);
    if (d < today) return;
    if (startDate && endDate) {
      onChange(d, undefined);
      return;
    }
    if (!startDate) {
      onChange(d, undefined);
      return;
    }
    if (sameDay(d, startDate)) {
      onChange(undefined, undefined);
      return;
    }
    if (d < startDate) {
      onChange(d, undefined);
      return;
    }
    onChange(startDate, endOfDay(day));
  }

  const daysInMonth = new Date(viewYear, viewMonth + 1, 0).getDate();
  const firstDOW = new Date(viewYear, viewMonth, 1).getDay();

  const cells: (Date | null)[] = [
    ...Array<null>(firstDOW).fill(null),
    ...Array.from(
      { length: daysInMonth },
      (_, i) => new Date(viewYear, viewMonth, i + 1),
    ),
  ];
  while (cells.length % 7 !== 0) cells.push(null);

  const rows = Array.from({ length: cells.length / 7 }, (_, i) =>
    cells.slice(i * 7, i * 7 + 7),
  );

  const hasRange = !!startDate && !!endDate;

  const fmt = (d: Date) =>
    d.toLocaleDateString("en-GB", { day: "numeric", month: "short" });

  const selectionLabel =
    startDate && endDate
      ? `${fmt(startDate)} – ${fmt(endDate)}`
      : startDate
        ? t("calendar.from", { date: fmt(startDate) })
        : null;

  return (
    <View className="bg-[#F0F0F0] rounded-[18px] border border-[rgba(124,58,237,0.22)] px-4 pt-[14px] pb-3">
      {/* Label + selection summary */}
      <View className="flex-row justify-between items-center mb-3">
        <Text className="font-nunito-bold text-[15px] text-[#111]">Date</Text>
        <Text className="font-nunito text-[12px] text-[#6B7280]">
          {selectionLabel ?? t("calendar.any")}
        </Text>
      </View>

      {/* Month navigation */}
      <View className="flex-row justify-between items-center mb-2">
        <Pressable onPress={prevMonth} hitSlop={12}>
          <IconSymbol name="chevron.left" size={16} color="#111" />
        </Pressable>
        <Text className="font-nunito-semibold text-[13px] text-[#111]">
          {months[viewMonth]} {viewYear}
        </Text>
        <Pressable onPress={nextMonth} hitSlop={12}>
          <IconSymbol name="chevron.right" size={16} color="#111" />
        </Pressable>
      </View>

      {/* Day-of-week header */}
      <View className="flex-row mb-0.5">
        {dayNames.map((d) => (
          <Text
            key={d}
            className="flex-1 text-center font-nunito-semibold text-[11px] text-[#6B7280] py-1"
          >
            {d}
          </Text>
        ))}
      </View>

      {/* Day grid */}
      {rows.map((row, ri) => (
        <View key={ri} className="flex-row">
          {row.map((day, ci) => {
            if (!day) return <View key={ci} className="flex-1 h-9" />;

            const isPast = day < today;
            const isStart = !!startDate && sameDay(day, startDate);
            const isEnd = !!endDate && sameDay(day, endDate);
            const isInRange =
              hasRange &&
              day.getTime() > startDate!.getTime() &&
              day.getTime() < endDate!.getTime();
            const isToday = sameDay(day, today);

            return (
              <Pressable
                key={ci}
                onPress={() => handlePress(day)}
                disabled={isPast}
                className="flex-1 items-center"
              >
                {/* Row-height container — carries the range tint */}
                <View
                  className="w-full h-9 items-center justify-center"
                  style={{
                    backgroundColor: isInRange
                      ? "rgba(26,26,26,0.1)"
                      : "transparent",
                  }}
                >
                  {/* Right-half tint on start cell */}
                  {isStart && hasRange && (
                    <View
                      className="absolute right-0 top-0 bottom-0 w-1/2"
                      style={{ backgroundColor: "rgba(26,26,26,0.1)" }}
                      pointerEvents="none"
                    />
                  )}
                  {/* Left-half tint on end cell */}
                  {isEnd && hasRange && (
                    <View
                      className="absolute left-0 top-0 bottom-0 w-1/2"
                      style={{ backgroundColor: "rgba(26,26,26,0.1)" }}
                      pointerEvents="none"
                    />
                  )}

                  {/* Circle */}
                  <View
                    className="w-8 h-8 rounded-full items-center justify-center"
                    style={{
                      backgroundColor:
                        isStart || isEnd ? "#111" : "transparent",
                      borderWidth: isToday && !isStart && !isEnd ? 1.5 : 0,
                      borderColor: "#111",
                    }}
                  >
                    <Text
                      className="font-nunito-semibold text-[12px]"
                      style={{
                        color:
                          isStart || isEnd
                            ? "#fff"
                            : isPast
                              ? "#9CA3AF"
                              : "#111",
                      }}
                    >
                      {day.getDate()}
                    </Text>
                  </View>
                </View>
              </Pressable>
            );
          })}
        </View>
      ))}

      {/* Clear */}
      {(startDate || endDate) && (
        <Pressable
          onPress={() => onChange(undefined, undefined)}
          hitSlop={8}
          className="mt-2 self-center"
        >
          <Text className="font-nunito-semibold text-[12px] text-[#2563EB]">
            {t("calendar.clearDates")}
          </Text>
        </Pressable>
      )}
    </View>
  );
}
