<script>
  let name = "";

  $: input_type = name !== "hello world" ? "input-primary" : "input-error";
</script>

<input
  type="text"
  placeholder="Don't write 'hello world'"
  class="my-3 input input-bordered {input_type} w-full max-w-xs"
  bind:value={name}
/>

<h1>Hello {name}!</h1>
