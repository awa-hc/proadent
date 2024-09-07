// src/utils/date-utils.ts
import { format } from 'date-fns';

export function localToUtc(localDate: string): string {
  const localDateTime = new Date(localDate);
  const utcDateTime = new Date(localDateTime.toUTCString()); // Convertir a UTC
  return utcDateTime.toISOString(); // Formato ISO 8601
}

export function utcToLocal(utcDate: string): string {
  const utcDateTime = new Date(utcDate);
  const localDateTime = new Date(
    utcDateTime.getTime() - utcDateTime.getTimezoneOffset() * 60000
  );
  return localDateTime.toISOString(); // Formato ISO 8601
}

export function formatUtcToLocal(utcDate: string, formatStr: string): string {
  const utcDateTime = new Date(utcDate);
  const localDateTime = new Date(
    utcDateTime.getTime() - utcDateTime.getTimezoneOffset() * 60000
  );

  return format(localDateTime, formatStr);
}
