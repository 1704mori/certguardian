<script lang="ts">
  import domain from "$lib/stores/domain";
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";
  import { toast } from "svelte-sonner";

  let showModal = false;
  let _domain = "";

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

<Dialog bind:showModal on:submit={handleSubmit}>
  <h2 slot="header">Add Domain</h2>

  <div class="flex flex-col w-full max-w-sm gap-1.5">
    <label class="font-medium text-xs uppercase" for="domain">Domain</label>
    <Input id="domain" placeholder="example.com" bind:value={_domain} />
  </div>

  <Button type="submit" slot="footer" {loading}>Add</Button>
</Dialog>

<Button on:click={() => (showModal = true)}>Add Domain</Button>
