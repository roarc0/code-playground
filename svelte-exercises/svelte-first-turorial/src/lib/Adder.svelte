<script>
  let numbers = [1, 2, 3, 4];

  function addNumber() {
    numbers = [...numbers, numbers.length + 1];
  }

  $: sum = numbers.reduce((total, currentNumber) => total + currentNumber, 0);
</script>

<button class="btn btn-secondary my-2" on:click={addNumber}>
  Add a number
</button>
<p class="my-3">{numbers.join(" + ")} = {sum}</p>
