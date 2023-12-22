<script lang="ts">
  import Button from "./ui/Button.svelte";
import Dialog from "./ui/Dialog.svelte";

  let showModal = false;
</script>


<Dialog bind:showModal>
	<h2 slot="header">
		Add Site
	</h2>

  todo
</Dialog>

<Button
  on:click={() => (showModal = true)}
>
  Add Site
</Button>