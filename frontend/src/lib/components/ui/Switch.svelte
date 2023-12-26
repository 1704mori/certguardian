<script>
  import { createEventDispatcher } from "svelte";

  export let value = false;
  const dispatch = createEventDispatcher();

  function toggle() {
    value = !value;
    dispatch("change", value);
  }
</script>

<label class="switch">
  <input type="checkbox" bind:checked={value} on:click={toggle} />
  <span class="slider round"></span>
</label>

<style>
  .switch {
    position: relative;
    display: inline-block;
    width: 44px;
    height: 24px;
  }

  .switch input {
    opacity: 0;
    width: 0;
    height: 0;
  }

  .slider {
    position: absolute;
    cursor: pointer;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    @apply bg-neutral-300 dark:bg-neutral-800;
    -webkit-transition: .2s;
    transition: .2s;
  }

  .slider:before {
    position: absolute;
    content: "";
    height: 20px;
    width: 20px;
    left: 2px;
    bottom: 2px;
    @apply bg-neutral-50 dark:bg-neutral-950;
    -webkit-transition: .2s;
    transition: .2s;
  }

  input:checked + .slider {
    /* background-color: #2196f3; */
    @apply bg-neutral-900 dark:bg-neutral-200;
  }

  input:focus + .slider {
    /* box-shadow: 0 0 1px #2196f3; */
    /* @apply ring-2 ring-offset-2 ring-sky-500; */
  }

  input:checked + .slider:before {
    -webkit-transform: translateX(20px);
    -ms-transform: translateX(20px);
    transform: translateX(20px);
  }

  /* Rounded sliders */
  .slider.round {
    border-radius: 34px;
  }

  .slider.round:before {
    border-radius: 50%;
  }
</style>
