<script lang="ts">
  import { theme } from "$lib/stores/theme";
  import "../app.css";
</script>

<svelte:head>
  <meta name="color-scheme" content={$theme == 'system' ? 'light dark' :
  $theme}/> <link rel="stylesheet" href={`/theme/${$theme}.css`} />
</svelte:head>

<slot />