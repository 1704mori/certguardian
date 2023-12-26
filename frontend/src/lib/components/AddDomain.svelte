<script lang="ts">
  import domain from "$lib/stores/domain";
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";
  import { toast } from "svelte-sonner";

  let showModal = false;
  let _domain = "";

  let loading = false;

  const domainRegex = /^((?!-)[A-Za-z0-9-]{1,63}(?<!-)\.)+[A-Za-z]{2,}$/;

  async function handleSubmit() {
    if (!_domain.trim()) {
      toast.error("Domain cannot be empty");
      return;
    }

    if (!domainRegex.test(_domain.trim())) {
      toast.error("Invalid domain name");
      return;
    }

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
  <h2 slot="header">Add Domain</h2>
  <form on:submit|preventDefault={handleSubmit}>
    <div class="flex flex-col w-full gap-1.5">
      <label class="font-medium text-xs uppercase" for="domain">Domain</label>
      <Input id="domain" placeholder="example.com" bind:value={_domain} />
    </div>

    <button type="submit" hidden>Add</button>
  </form>

  <Button slot="footer" {loading} on:click={handleSubmit}>Add</Button>
</Dialog>

<Button on:click={() => (showModal = true)}>Add Domain</Button>
