<script lang="ts">
  export let data: any[] = [];
  export let columns: any[] = [];
  export let className: string = '';
  export let totalData: number;

  let generateGridColumns: string[];

  // find a better way to check if the value is an HTML element
  const isHTMLElement = (value: any) => {
    const htmlTags = ['input', 'button', 'select', 'textarea', 'div', 'p', 'span'];
    return htmlTags.includes(value?.type);
  };

  $: generateGridColumns = columns.map((column: any) => (["id", "actions"].includes(column.key) ? "auto" : "1fr"));
</script>

<div class="flex flex-col gap-1">
  <div class={`space-y-2 md:space-y-0 bg-neutral-50 dark:bg-neutral-950 border border-neutral-200 dark:border-neutral-900 rounded-lg divide-y divide-neutral-200 dark:divide-neutral-900 overflow-hidden max-h-80 overflow-y-scroll ${className}`}>
    {#if !totalData || !data.length}
      <div class="flex justify-between items-center px-4 py-2 bg-neutral-50 dark:bg-neutral-950">
        <p class="text-sm font-medium">no results found</p>
      </div>
    {/if}

    {#each data as item, index (item.id)}
      <div class={`bg-neutral-50 dark:bg-neutral-950 px-1 py-1.5 shadow-md flex flex-wrap ${generateGridColumns ? "grid" : ""} sm:gap-4`} style={`grid-template-columns: ${generateGridColumns.join(' ')}`}>
        {#each columns as column, columnIndex (column.key)}
          <div class={`w-full p-2 mt-auto ${isHTMLElement(item[column.key]) ? 'self-center truncate' : ''}`}>
            {#if typeof item[column.key] == 'string'}
              <p class="font-medium">{item[column.key]}</p>
            {:else if column.key === 'actions' && typeof item[column.key] != 'string'}
              <div class="w-full flex justify-center items-center">
                <svelte:component this={item[column.key]} />
              </div>
            {:else if typeof item[column.key] !== 'string' && typeof item[column.key] == 'object'}
              <div class={`flex flex-col justify-start items-stretch relative min-w-full ${item[column.key].rules}`}>
                <p class="text-sm font-semibold">{item[column.key].head}</p>
                <p class="text-sm truncate">
                  {column.modifyValue ? column.modifyValue(item[column.key].child) : item[column.key].child}
                </p>
              </div>
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
