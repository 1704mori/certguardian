<script lang="ts">
  import { theme } from "$lib/stores/theme";

  const toggleTheme = () => {
    if ($theme === "light") {
      $theme = "dark";
    } else {
      $theme = "light";
    }
  };
</script>

<button
  type="button"
  class="fixed bottom-12 left-12 w-10 h-10 rounded-full flex justify-center items-center bg-neutral-50 dark:bg-neutral-950 border border-neutral-200 dark:border-neutral-800 shadow-md"
  on:click={toggleTheme}
>
  <svg
    xmlns="http://www.w3.org/2000/svg"
    width="24"
    height="24"
    viewBox="0 0 24 24"
    fill="none"
    stroke="currentColor"
    stroke-width="2"
    stroke-linecap="round"
    stroke-linejoin="round"
    class="lucide lucide-sun-moon"
  >
    <!-- moon -->
    <!-- <path d="M12 8a2.83 2.83 0 0 0 4 4 4 4 0 1 1-4-4" /> -->
    <!-- sun -->
    <!-- <circle cx="12" cy="12" r="4"/> -->

    {#if $theme === "light"}
      <path d="M12 8a2.83 2.83 0 0 0 4 4 4 4 0 1 1-4-4" />
    {:else}
      <circle cx="12" cy="12" r="4" />
    {/if}

    <path d="M12 2v2" />
    <path d="M12 20v2" /><path d="m4.9 4.9 1.4 1.4" />
    <path d="m17.7 17.7 1.4 1.4" />
    <path d="M2 12h2" />
    <path d="M20 12h2" />
    <path d="m6.3 17.7-1.4 1.4" />
    <path d="m19.1 4.9-1.4 1.4" />
  </svg>
</button>
