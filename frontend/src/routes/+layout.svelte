<script lang="ts">
  import { theme, type Theme } from "$lib/stores/theme";
  import { onMount } from "svelte";
  import "../app.css";

  onMount(() => {
    const _theme = localStorage.getItem("theme");
    if (theme) {
      $theme = _theme as Theme;
    }
  });
</script>

<slot />