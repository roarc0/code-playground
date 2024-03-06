<script>
  import { count } from "./store";
  import { onDestroy } from "svelte";

  let count_value;
  const unsubscribe = count.subscribe((value) => {
    count_value = value;
  });
  onDestroy(unsubscribe);

  $: {
    console.log(`the count is ${count_value}`);
    console.log(`this will also be logged whenever count changes`);
    if (count_value >= 15) {
      alert("count is dangerously high!");
      count.set(0);
    }
  }
  $: doubled = count_value * 2;
</script>

<button class="btn btn-primary mr-1" on:click={count.increment}>+</button>
<button class="btn btn-primary mr-1" on:click={count.decrement}>-</button>
<button class="btn btn-primary mr-1" on:click={count.reset}>reset</button>
<p class="my-3">{count_value} doubled is {doubled}</p>
{#if count_value > 10}
  <p class="mb-3">{count_value} is greater than 10</p>
{:else if count_value < 5}
  <p class="mb-3">{count_value} is less than 5</p>
{:else}
  <p>{count_value} is between 0 and 10</p>
{/if}
