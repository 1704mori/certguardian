<script lang="ts">
  import { cn } from "$lib/utils";
  import type { HTMLButtonAttributes } from "svelte/elements";
  import { tv, type VariantProps } from "tailwind-variants";
  import Icon from "./icons/Icon.svelte";
  import Loader from "./icons/Loader.svelte";

  export { className as class };

  const buttonVariants = tv({
    base: "inline-flex items-center justify-center gap-2 rounded-md text-sm font-medium whitespace-nowrap ring-offset-lighter transition-colors focus-visible:outline-none focus-visible:ring-2 focus-visible:ring-500 focus-visible:ring-offset-2 disabled:pointer-events-none disabled:opacity-50",
    variants: {
      variant: {
        default:
          // "text-white bg-black border border-neutral-800 hover:border-neutral-600 focus-visible:ring-neutral-500",
          "bg-white dark:bg-black border border-neutral-800 dark:border-neutral-800 hover:border-neutral-500 dark:hover:border-neutral-600 focus-visible:ring-neutral-500",
        primary:
          "text-white bg-sky-600 hover:bg-sky-700 focus-visible:ring-sky-500",
        ghost:
          "bg-transparent border border-sky-600 hover:bg-sky-600/30 focus-visible:ring-sky-500",
      },
      size: {
        xs: "px-2.5 py-1.5 text-xs",
        sm: "px-3 py-2 text-sm",
        md: "px-4 py-2 text-sm",
        lg: "px-4 py-2 text-base",
      },
    },
    defaultVariants: {
      variant: "default",
      size: "md",
    },
  });

  type Variant = VariantProps<typeof buttonVariants>["variant"];
  type Size = VariantProps<typeof buttonVariants>["size"];

  type $$Props = {
    class?: string;
    variant?: Variant;
    size?: Size;
    loading?: boolean;
  } & HTMLButtonAttributes;

  export let variant: $$Props["variant"] = "default";
  export let size: $$Props["size"] = "md";
  export let loading: boolean = false;

  let className: $$Props["class"] = undefined;
</script>

<button
  class={cn(buttonVariants({ variant, size, className }))}
  {...$$restProps}
  on:click
  on:keydown
  disabled={$$props.disabled || loading}
>
  {#if loading}
    <!-- <Loader2 size="16" class="animate-spin" /> -->
    <Icon size="16" class="animate-spin">
      <Loader />
    </Icon>
  {/if}
  <slot />
</button>
