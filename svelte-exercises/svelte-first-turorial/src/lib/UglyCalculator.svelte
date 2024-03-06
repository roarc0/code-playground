<script>
  let a = 1;
  let b = 2;
</script>

<label>
  <input
    class="range range-accent"
    type="number"
    bind:value={a}
    min="0"
    max="10"
  />
  <input
    class="range range-accent"
    type="range"
    bind:value={a}
    min="0"
    max="10"
  />
</label>

<label>
  <input
    class="range range-accent"
    type="number"
    bind:value={b}
    min="0"
    max="10"
  />
  <input
    class="range range-accent"
    type="range"
    bind:value={b}
    min="0"
    max="10"
  />
</label>

<p>{a} + {b} = {a + b}</p>
