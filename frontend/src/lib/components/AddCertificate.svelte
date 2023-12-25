<script lang="ts">
  import certificates from "$lib/stores/certificates";
  import { onMount } from "svelte";
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";
  import { toast } from "svelte-sonner";

  let _dirs: HTMLSpanElement[] = [];
  let dirs: string[] = [];

  let showModal = false;
  let dir = "";

  $: isValidDir = dir.trim() && !dirs.includes(dir.trim());
  $: console.log(_dirs);

  let loading = false;

  function addDir() {
    if (!dir.trim()) return;
    if (dirs.includes(dir)) return;

    dirs = [...dirs, dir.trim()];
    dir = "";
  }

  function removeDir(index: number) {
    toast.promise(certificates.deleteDir(dirs[index]), {
      success: (d) => d.message,
      error: (e: any) => e.message ?? "Something went wrong",
    });

    dirs = dirs.slice(0, index).concat(dirs.slice(index + 1));
    _dirs = _dirs.slice(0, index).concat(_dirs.slice(index + 1));
  }

  function getDimensions(index: number) {
    return _dirs?.[index].getBoundingClientRect();
  }

  async function handleSubmit() {
    loading = true;
    toast.promise(certificates.addNewDir(dirs), {
      success: (d) => d.message,
      error: (e: any) => e.message ?? "Something went wrong",
    });

    loading = false;
    showModal = false;
    dirs = [];
  }

  onMount(() => {
    dirs = Object.keys($certificates).map((path) => path);
  });
</script>

<Dialog
  class="min-w-[32rem] max-w-[32rem]"
  bind:showModal
  on:submit={handleSubmit}
>
  <h2 slot="header">Add Certificate</h2>

  <div class="flex flex-col">
    <div class="flex flex-col w-full gap-1.5">
      <label class="font-medium text-xs uppercase" for="dir">Directory</label>
      <Input
        id="dir"
        placeholder="e.g: /etc/letsencrypt/live/example.com"
        bind:value={dir}
        on:keydown={(e) => e.key == "Enter" && addDir()}
        on:blur={addDir}
      />
    </div>
    {#if isValidDir}
      <div
        class="flex items-center flex-wrap gap-2 bg-neutral-100 dark:bg-neutral-900 rounded-md p-1 w-full opacity-60"
      >
        <div
          class="bg-white dark:bg-black border border-neutral-800 dark:border-neutral-800 hover:border-neutral-500 dark:hover:border-neutral-600 rounded-md p-1 break-words max-w-[32rem] truncate"
        >
          {dir}
        </div>
      </div>
      <span class="text-xs text-neutral-700 dark:text-neutral-500 mt-1">
        Press 'Enter' or click outside the input to add a directory
      </span>
    {/if}
  </div>
  {#if dirs.length}
    <div
      class="flex items-center flex-1 flex-wrap overflow-hidden gap-2 bg-neutral-100 dark:bg-neutral-900 rounded-md p-1 w-full"
    >
      {#each dirs as dir, index}
        <button
          type="button"
          class="inline-flex items-center gap-2 bg-white dark:bg-black border border-neutral-800 dark:border-neutral-800 hover:border-neutral-500 dark:hover:border-neutral-600 rounded-md p-1 group max-h-8 min-h-8 cursor-pointer max-w-[32rem] truncate"
          on:click={() => removeDir(index)}
        >
          <span class="max-w-[32rem] truncate" bind:this={_dirs[index]}>
            {dir}
          </span>
          <button type="button"> X </button>
          <!-- {#if _dirs.length && _dirs[index]}
            <span
              class="hidden group-hover:flex justify-center items-center"
              style={`width: ${getDimensions(index)?.width}px; height: ${
                getDimensions(index)?.height
              }px`}
            >
              <Icon size="16">
                <Trash />
              </Icon>
            </span>
          {/if} -->
        </button>
      {/each}
    </div>
  {/if}

  <Button type="button" slot="footer" {loading} on:click={handleSubmit}
    >Add</Button
  >
</Dialog>

<Button on:click={() => (showModal = true)}>Manage Directories</Button>
