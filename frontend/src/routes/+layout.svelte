<script lang="ts">
  import { theme, type Theme } from "$lib/stores/theme";
  import { onMount } from "svelte";
  import "../app.css";
  import ThemeSwitch from "$lib/components/ThemeSwitch.svelte";

  onMount(() => {
    const _theme = localStorage.getItem("theme");
    if (theme) {
      $theme = _theme as Theme;
    }
  });
</script>

<ThemeSwitch />
<slot />