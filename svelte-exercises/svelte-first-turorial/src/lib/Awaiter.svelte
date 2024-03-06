<script>
  // async function getRandomNumber() {
  //   const res = await fetch("/random-number");

  //   if (res.ok) {
  //     return await res.text();
  //   } else {
  //     throw new Error("Request failed");
  //   }
  // }
  async function mockFetchRandomNumber() {
    await new Promise((resolve) => setTimeout(resolve, 1000));
    const randomNumber = Math.floor(Math.random() * 100);
    return randomNumber.toString();
  }

  let promise = mockFetchRandomNumber();

  function handleClick() {
    promise = mockFetchRandomNumber();
  }
</script>

<button class="btn btn-info my-3" on:click={handleClick}>
  generate random number
</button>

{#await promise}
  <p class="text-2xl">...waiting</p>
{:then number}
  <p class="text-2xl">The number is {number}</p>
{:catch error}
  <p class="text-2xl">{error.message}</p>
{/await}
