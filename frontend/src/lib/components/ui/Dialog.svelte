<script lang="ts">
  import { cn } from "$lib/utils";

  export let showModal = false;

  let dialog: HTMLDialogElement;

  $: if (dialog && showModal) dialog.showModal();
  $: if (dialog && !showModal) dialog.close();

  let className: string | undefined = undefined;
  export { className as class };
</script>

<!-- svelte-ignore a11y-click-events-have-key-events a11y-no-noninteractive-element-interactions -->
<dialog
  bind:this={dialog}
  on:close={() => (showModal = false)}
  on:click|self={() => dialog.close()}
  class={cn(
    "max-w-2xl rounded-md border border-neutral-600 dark:border-neutral-800 p-0 bg-white dark:bg-neutral-950 dark:text-neutral-100",
    className
  )}
>
  <!-- svelte-ignore a11y-no-static-element-interactions -->
  <main on:submit|preventDefault class="flex flex-col gap-4 p-4 w-full">
    <div class="font-semibold">
      <slot name="header" />
    </div>
    <slot />
    <!-- svelte-ignore a11y-autofocus -->
    <div class="flex items-center gap-2 self-end">
      <button class="hover:underline" on:click={() => dialog.close()}>
        Close
      </button>
      <slot name="footer" />
    </div>
  </main>
</dialog>

<style>
  dialog::backdrop {
    background-color: rgba(0, 0, 0, 0.3);
    /* backdrop-filter: blur(5px); */
  }
  dialog[open] {
    animation: zoom 0.3s cubic-bezier(0.34, 1.56, 0.64, 1);
  }
  @keyframes zoom {
    from {
      transform: scale(0.95);
    }
    to {
      transform: scale(1);
    }
  }
  dialog[open]::backdrop {
    animation: fade 0.2s ease-out;
  }
  @keyframes fade {
    from {
      opacity: 0;
    }
    to {
      opacity: 1;
    }
  }
  button {
    display: block;
  }
</style>
