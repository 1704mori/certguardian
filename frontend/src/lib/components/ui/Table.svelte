<script lang="ts">
  import type { SvelteComponent } from "svelte";

  type Column = {
    key: string;
    label: string;
    modifyValue?: (value: any) => void;
  };

  type Actions = {
    component: any; //typeof SvelteComponent;
    props: {
      data: any;
    };
  };

  type CustomKey = {
    rules: string;
    head: string;
    child: string;
  };

  type Data = string | CustomKey | Actions | SvelteComponent;

  export let data: {
    [key: string]: Partial<Data>;
  }[];
  export let columns: Column[] = [];
  export let className: string = "";
  export let label = false;

  let generateGridColumns: string[];

  // find a better way to check if the value is an HTML element
  const isHTMLElement = (value: any) => {
    const htmlTags = [
      "input",
      "button",
      "select",
      "textarea",
      "div",
      "p",
      "span",
    ];
    return htmlTags.includes(value?.type);
  };

  function isActions(value: any): value is Actions {
    return value && typeof value === "object" && "component" in value;
  }

  function getActions(value: any): Actions {
    return value as Actions;
  }

  function getCustomKey(value: any): CustomKey {
    return value as CustomKey;
  }

  function asComponent(value: any): typeof SvelteComponent {
    return value as typeof SvelteComponent;
  }

  $: generateGridColumns = columns.map((column: any) =>
    ["id", "actions"].includes(column.key) ? "auto" : "1fr"
  );
</script>

<div class="flex flex-col gap-1">
  <div
    class={`space-y-2 md:space-y-0 bg-neutral-50 dark:bg-neutral-950 border border-neutral-200 dark:border-neutral-900 rounded-lg divide-y divide-neutral-200 dark:divide-neutral-900 overflow-hidden max-h-80 overflow-y-scroll ${className}`}
  >
    {#if !data.length}
      <div
        class="flex justify-between items-center px-4 py-2 bg-neutral-50 dark:bg-neutral-950"
      >
        <p class="text-sm font-medium">no results found</p>
      </div>
    {/if}

    {#each data as item, index (item.id)}
      <div
        class={`bg-neutral-50 dark:bg-neutral-950 px-1 py-1.5 shadow-md flex flex-wrap ${
          generateGridColumns ? "[@media(min-width:840px)]:grid" : ""
        } sm:gap-4`}
        style={`grid-template-columns: ${generateGridColumns.join(" ")}`}
      >
        {#each columns as column, columnIndex (column.key)}
          <div
            class={`w-full p-2 mt-auto max-md:[&:not(:last-child)]:border-b border-neutral-200 dark:border-neutral-900 self-center truncate ${
              isHTMLElement(item[column.key]) ? "self-center truncate" : ""
            }`}
          >
            {#if typeof item[column.key] == "string"}
              {#if label}
                <div
                  class="flex flex-col justify-start items-stretch relative min-w-full"
                >
                  <p class="text-sm font-semibold">
                    {column.label}
                  </p>
                  <p class="text-sm truncate">
                    {item[column.key]}
                  </p>
                </div>
              {:else}
                <p class="font-medium">{item[column.key]}</p>
              {/if}
            {:else if column.key === "actions" && isActions(item[column.key])}
              <div class="w-full flex justify-center items-center">
                <svelte:component
                  this={getActions(item[column.key]).component}
                  {...getActions(item[column.key]).props}
                />
              </div>
            {:else if typeof item[column.key] !== "string" && typeof item[column.key] == "object"}
              <div
                class={`flex flex-col justify-start items-stretch relative min-w-full ${
                  getCustomKey(item[column.key]).rules
                }`}
              >
                <p class="text-sm font-semibold">
                  {getCustomKey(item[column.key]).head}
                </p>
                <p class="text-sm truncate">
                  {column.modifyValue
                    ? getCustomKey(column.modifyValue(item[column.key])).child
                    : getCustomKey(item[column.key]).child}
                </p>
              </div>
            {:else if label}
              <div
                class="flex flex-col justify-start items-stretch relative min-w-full"
              >
                <p class="text-sm font-semibold">
                  {column.label}
                </p>
                <svelte:component this={asComponent(item[column.key])} />
              </div>
            {:else}
              <svelte:component this={asComponent(item[column.key])} />
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
