<script lang="ts">
  import domain from "$lib/stores/domain";
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";
  import { toast } from "svelte-sonner";
  import Switch from "./ui/Switch.svelte";

  let showModal = false;
  let _domain = "";

  let enableCertsCron = false;

  let loading = false;

  async function handleSubmit() {
    loading = true;
    toast.promise(domain.addNewDomain(_domain), {
      success: (d) => d.message,
      error: (e: any) => e.message ?? "Something went wrong",
    });

    loading = false;
    showModal = false;
  }
</script>

<Dialog class="min-w-[32rem] max-w-[32rem]" bind:showModal>
  <h2 slot="header">Manage Crons</h2>

  <div class="flex items-center gap-2">
    <Switch bind:value={enableCertsCron} />
    <span class="font-medium">
      Set custom
    </span>
  </div>

  <div class="flex flex-col w-full gap-1.5">
    <label class="font-medium text-xs uppercase" for="domain">Domain</label>
    <Input id="domain" placeholder="example.com" bind:value={_domain} />
  </div>

  <Button slot="footer" {loading} on:click={handleSubmit}>Add</Button>
</Dialog>

<Button on:click={() => (showModal = true)}>Cron</Button>
