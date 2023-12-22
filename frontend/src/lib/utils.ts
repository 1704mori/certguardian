import { twMerge } from 'tailwind-merge';
import clsx from 'clsx';

export function cn(...classes: any[]) {
  return clsx(twMerge(...classes));
}