<script lang="ts">
  import { domains } from "$lib/stores/domain";
  import Actions from "./Table/Actions.svelte";
  import Expired from "./Table/Expired.svelte";
  import Valid from "./Table/Valid.svelte";
  import Table from "./ui/Table.svelte";
</script>

{#if !$domains || !$domains.length}
  <span
    class="flex justify-between items-center px-4 py-2 bg-neutral-50 dark:bg-neutral-950 rounded-md"
    >no domains found</span
  >
{:else}
  <Table
    label={true}
    columns={[
      {
        key: "commonName",
        label: "Common Name",
      },
      {
        key: "validFrom",
        label: "From",
      },
      {
        key: "validTo",
        label: "To",
      },
      {
        key: "isExpired",
        label: "Expired",
      },
      {
        key: "issuer",
        label: "Issuer",
      },
      {
        key: "actions",
        label: "Actions",
      },
    ]}
    data={$domains.map((data) => ({
      ...data,
      isExpired: data.isExpired ? Expired : Valid,
      validFrom: new Date(data.validFrom).toLocaleDateString(undefined, {
        year: "numeric",
        month: "long",
        day: "numeric",
      }),
      validTo: new Date(data.validTo).toLocaleDateString(undefined, {
        year: "numeric",
        month: "long",
        day: "numeric",
      }),
      actions: {
        component: Actions,
        props: {
          data,
        },
      },
    }))}
  />
{/if}
