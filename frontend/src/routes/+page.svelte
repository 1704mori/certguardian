<script lang="ts">
  import Table from "$lib/components/Table.svelte";
  import ThemeSwitch from "$lib/components/ThemeSwitch.svelte";
</script>

<ThemeSwitch />
<div class="w-full">
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
