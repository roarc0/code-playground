<script>
  import { createEventDispatcher } from "svelte";

  const dispatch = createEventDispatcher();

  function sayHello() {
    dispatch("message", {
      text: "Hello!",
    });
  }
</script>

<button class="btn btn-warning my-3" on:click={sayHello}>
  Click to say hello
</button>
