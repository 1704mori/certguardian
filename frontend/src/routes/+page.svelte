<script lang="ts">
  import AddDomain from "$lib/components/AddDomain.svelte";
  import Alerts from "$lib/components/Alerts.svelte";
  import Certificates from "$lib/components/Certificates.svelte";
  import Cron from "$lib/components/Cron.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Table from "$lib/components/ui/Table.svelte";
  import Actions from "$lib/components/Table/Actions.svelte";
  import { domains } from "$lib/stores/domain";
  import Expired from "$lib/components/Table/Expired.svelte";
  import Valid from "$lib/components/Table/Valid.svelte";
  import Icon from "$lib/components/ui/icons/Icon.svelte";
  import ArrowLeft from "$lib/components/ui/icons/ArrowLeft.svelte";
  import AddCertificate from "$lib/components/AddCertificate.svelte";

  let menu: "certificates" | "alerts" | "cron" | undefined;
</script>

<div class="grid grid-cols-1 lg:grid-cols-[minmax(0,15rem)_1fr] gap-5 w-full">
  <div class="flex flex-col gap-3">
    {#if !menu}
      <AddDomain />
    {:else}
      <Button on:click={() => (menu = undefined)}>
        <Icon size="16">
          <ArrowLeft />
        </Icon>
        Back
      </Button>
    {/if}
    <div
      class="flex flex-col gap-2 rounded-md bg-white dark:bg-neutral-950 border border-neutral-800 dark:border-neutral-800 px-2 py-3"
    >
      <Button
        active={menu == "certificates"}
        on:click={() => (menu = "certificates")}
      >
        Certificates
      </Button>
      <Button active={menu == "alerts"} on:click={() => (menu = "alerts")}>
        Alerts
      </Button>
      <Button active={menu == "cron"} on:click={() => (menu = "cron")}>
        Cron
      </Button>
    </div>
    {#if menu == "certificates"}
      <AddCertificate />
    {/if}
  </div>
  <div class="min-h-80">
    {#if !menu}
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
    {/if}

    {#if menu}
      {#if menu == "certificates"}
        <Certificates />
      {:else if menu == "alerts"}
        <Alerts />
      {:else if menu == "cron"}
        <Cron />
      {/if}
    {/if}
  </div>
</div>
