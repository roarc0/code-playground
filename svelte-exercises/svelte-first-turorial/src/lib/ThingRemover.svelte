<script>
  import Thing from "./Thing.svelte";

  let things = [
    { id: 1, name: "apple" },
    { id: 2, name: "banana" },
    { id: 3, name: "carrot" },
    { id: 4, name: "doughnut" },
    { id: 5, name: "egg" },
  ];

  function removeThing() {
    things = things.slice(1);
  }

  function addApple() {
    things.push({ id: things.length + 1, name: "apple" });
    things = things;
  }
</script>

{#if things.length > 0}
  <button class="btn btn-neutral my-2" on:click={removeThing}>
    Remove first thing
  </button>
{:else}
  <button class="btn btn-neutral my-2" on:click={addApple}> Add apple </button>
{/if}

{#each things as thing (thing.id)}
  <Thing name={thing.name} />
{/each}
