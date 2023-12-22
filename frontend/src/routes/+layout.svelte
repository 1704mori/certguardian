<script lang="ts">
  import { theme, type Theme } from "$lib/stores/theme";
  import { onMount } from "svelte";
  import "../app.css";
  import ThemeSwitch from "$lib/components/ThemeSwitch.svelte";
  import { Toaster } from "svelte-sonner";

  onMount(() => {
    const _theme = localStorage.getItem("theme");
    if (theme) {
      $theme = _theme as Theme;
    }
  });
</script>

<Toaster theme={$theme} />
<ThemeSwitch />
<slot />
