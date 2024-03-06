<script>
  import { onMount } from "svelte";
  import { gradient } from "./gradient.js";

  onMount(() => {
    const canvas = document.querySelector("canvas");
    const context = canvas.getContext("2d");

    requestAnimationFrame(function loop(t) {
      requestAnimationFrame(loop);
      gradient(context, t);
    });
  });
</script>

<canvas width={32} height={32} />

<style>
  canvas {
    position: fixed;
    left: 30%;
    top: 0;
    width: 100%;
    height: 100%;
    background-color: #666;
    mask: url(./svelte-logo-mask.svg) 50% 50% no-repeat;
    mask-size: 20vmin;
    -webkit-mask: url(./svelte-logo-mask.svg) 50% 50% no-repeat;
    -webkit-mask-size: 20vmin;
  }
</style>
