<script lang="ts">
  import Button from "./ui/Button.svelte";
  import Dialog from "./ui/Dialog.svelte";
  import Input from "./ui/Input.svelte";

  let showModal = false;
</script>

<Dialog bind:showModal>
  <h2 slot="header">Add Domain</h2>

  <div class="flex flex-col w-full max-w-sm gap-1.5">
    <label class="font-medium text-xs uppercase" for="domain">Domain</label>
    <Input id="domain" placeholder="domain" />
  </div>
</Dialog>

<Button on:click={() => (showModal = true)}>Add Domain</Button>
