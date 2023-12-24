<script lang="ts">
  import type { DomainList } from "$lib/api/domain";
  import domain from "$lib/stores/domain";
  import { toast } from "svelte-sonner";
  import Icon from "../ui/icons/Icon.svelte";
  import Trash from "../ui/icons/Trash.svelte";

  export let data: DomainList;

  async function handleDelete() {
    toast.promise(domain.deleteDomain(data.commonName), {
      success: (d) => d.message,
      error: (e: any) => e.message ?? "Something went wrong",
    });
  }
</script>

<div class="flex items-center justify-center gap-2">
  <button type="button" on:click={handleDelete}>
    <Icon size="20">
      <Trash />
    </Icon>
  </button>
</div>
