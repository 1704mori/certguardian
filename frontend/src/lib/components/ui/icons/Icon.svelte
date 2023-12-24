<!-- copied from lucide-svelte but modified -->

<script lang="ts">
  import type { SVGAttributes } from "svelte/elements";

  type Attrs = SVGAttributes<any>;

  export let color = "currentColor";
  export let size: number | string = 24;
  export let strokeWidth: number | string = 2;
  export let absoluteStrokeWidth: boolean = false;

  const defaultAttributes: Attrs = {
    xmlns: "http://www.w3.org/2000/svg",
    width: 24,
    height: 24,
    viewBox: "0 0 24 24",
    fill: "none",
    stroke: "currentColor",
    "stroke-width": 2,
    "stroke-linecap": "round",
    "stroke-linejoin": "round",
  };
</script>

<svg
  {...defaultAttributes}
  {...$$restProps}
  width={size}
  height={size}
  stroke={color}
  stroke-width={absoluteStrokeWidth
    ? (Number(strokeWidth) * 24) / Number(size)
    : strokeWidth}
  class={`${$$props.class ?? ""}`}
>
  <!-- <svelte:component this={iconNode} /> -->
  <slot />
</svg>
