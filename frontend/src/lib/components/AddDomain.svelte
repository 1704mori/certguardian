<script lang="ts">
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";
  import { domain as sDomain } from "$lib/stores/domain";

  let showModal = false;
  let domain = "";

  async function handleSubmit() {
    const { addDomain } = sDomain();
    await addDomain(domain);
  }
</script>

<Dialog bind:showModal on:submit={handleSubmit}>
  <h2 slot="header">Add Domain</h2>

  <div class="flex flex-col w-full max-w-sm gap-1.5">
    <label class="font-medium text-xs uppercase" for="domain">Domain</label>
    <Input id="domain" placeholder="example.com" bind:value={domain} />
  </div>

  <Button type="submit" slot="footer">Add</Button>
</Dialog>

<Button on:click={() => (showModal = true)}>Add Domain</Button>
