<script lang="ts">
  import type { CertList } from "$lib/api/certificates";
  import { certificates } from "$lib/stores/certificates";
  import Actions from "./Table/Actions.svelte";
  import Expired from "./Table/Expired.svelte";
  import Valid from "./Table/Valid.svelte";
  import Table from "./ui/Table.svelte";

  const processData = (certificates: CertList) => {
    return Object.values(certificates).flatMap(certsPath =>
      Object.values(certsPath).map(certInfo => ({
        ...certInfo,
        isExpired: certInfo.isExpired ? Expired : Valid,
        validFrom: new Date(certInfo.validFrom).toLocaleDateString(undefined, {
          year: "numeric",
          month: "long",
          day: "numeric",
        }),
        validTo: new Date(certInfo.validTo).toLocaleDateString(undefined, {
          year: "numeric",
          month: "long",
          day: "numeric",
        }),
        actions: {
          component: Actions,
          props: {
            data: certInfo
          }
        }
      }))
    );
  };

  $: processedData = processData($certificates);

</script>

{#if !processedData.length}
  <span
    class="flex justify-between items-center px-4 py-2 bg-neutral-50 dark:bg-neutral-950 rounded-md"
    >no certificates found. click on 'Manage Directories'.</span
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
    data={processedData}
  />
{/if}
