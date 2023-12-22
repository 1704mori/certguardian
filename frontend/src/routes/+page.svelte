<script lang="ts">
  import Button from "$lib/components/Button.svelte";
  import Table from "$lib/components/Table.svelte";
</script>

<div class="grid grid-cols-1 lg:grid-cols-[minmax(0,15rem)_1fr] gap-5 w-full">
  <div class="flex flex-col gap-3">
    <Button>Add site</Button>
    <div class="flex flex-col gap-2 rounded-md bg-white dark:bg-neutral-950 border border-neutral-800 dark:border-neutral-800 px-2 py-3">
      <Button>Alerts</Button>
      <Button>Cron</Button>
    </div>
  </div>
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
</div>
