<script lang="ts">
  import AddDomain from "$lib/components/AddDomain.svelte";
  import Alerts from "$lib/components/Alerts.svelte";
  import Certificates from "$lib/components/Certificates.svelte";
  import Cron from "$lib/components/Cron.svelte";
  import Button from "$lib/components/ui/Button.svelte";
  import Table from "$lib/components/ui/Table.svelte";
  import { ArrowLeft } from "lucide-svelte";

  let menu: "certificates" | "alerts" | "cron" | undefined;
</script>

<div class="grid grid-cols-1 lg:grid-cols-[minmax(0,15rem)_1fr] gap-5 w-full">
  <div class="flex flex-col gap-3">
    <!-- <AddDomain /> -->
    {#if !menu}
      <AddDomain />
    {:else}
      <Button on:click={() => (menu = undefined)}>
        <ArrowLeft size="16" />
        Back
      </Button>
    {/if}
    <div
      class="flex flex-col gap-2 rounded-md bg-white dark:bg-neutral-950 border border-neutral-800 dark:border-neutral-800 px-2 py-3"
    >
      <Button on:click={() => (menu = "certificates")}>Certificates</Button>
      <Button on:click={() => (menu = "alerts")}>Alerts</Button>
      <Button on:click={() => (menu = "cron")}>Cron</Button>
    </div>
  </div>
  {#if !menu}
    <Table
      columns={[
        {
          name: "ID",
          key: "id",
        },
        {
          name: "Name",
          key: "name",
        },
        {
          name: "Status",
          key: "status",
        },
        {
          name: "created at",
          key: "created_at",
        },
      ]}
      totalData={1}
      data={[
        {
          id: 1,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 2,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 3,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 4,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 5,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 6,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 7,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 8,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
        {
          id: 9,
          name: "test",
          status: "active",
          created_at: "2021-08-01 00:00:00",
        },
      ]}
    />
  {/if}

  {#if menu}
    <div class="min-h-80">
      {#if menu == "certificates"}
        <Certificates />
      {:else if menu == "alerts"}
        <Alerts />
      {:else if menu == "cron"}
        <Cron />
      {/if}
    </div>
  {/if}
</div>
