import { writable } from 'svelte/store';

export type Theme = 'dark' | 'light' | 'system';

function getInitialTheme(): Theme {
  if (typeof window !== 'undefined') {
    const storedTheme = localStorage.getItem('theme');

    if (storedTheme && ['dark', 'light', 'system'].includes(storedTheme)) {
      return storedTheme as Theme;
    }
  }
  return 'dark';
}


function createThemeStore() {
  const { subscribe, set } = writable<Theme>(getInitialTheme());

  return {
    subscribe,
    set: (value: Theme) => {
      localStorage.setItem('theme', value);
      window.document.documentElement.classList.remove('dark', 'light');
      window.document.documentElement.classList.add(value);
      set(value);
    },
  };
}

export const theme = createThemeStore();
