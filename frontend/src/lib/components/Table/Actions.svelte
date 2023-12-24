<script lang="ts">
  import type { DomainList } from "$lib/api/domain";
  import domain from "$lib/stores/domain";
  import { Trash } from "lucide-svelte";
  import { toast } from "svelte-sonner";

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
    <Trash />
  </button>
</div>
