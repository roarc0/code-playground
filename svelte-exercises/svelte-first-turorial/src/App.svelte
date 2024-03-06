<script>
  import Footer from "./lib/Footer.svelte";
  import Navbar from "./lib/Navbar.svelte";
  import Counter from "./lib/Counter/CounterDisplay.svelte";
  import TailwindCss from "./lib/TailwindCSS.svelte";
  import Adder from "./lib/Adder.svelte";
  import Answer from "./lib/Answer.svelte";
  import PackageInfo from "./lib/PackageInfo.svelte";
  import ThingRemover from "./lib/ThingRemover.svelte";
  import Awaiter from "./lib/Awaiter.svelte";
  import Greeter from "./lib/Greeter.svelte";
  import Input from "./lib/Input.svelte";
  import UglyCalculator from "./lib/UglyCalculator.svelte";
  import Spam from "./lib/Spam.svelte";
  import Clock from "./lib/Clock.svelte";

  const pkg = {
    name: "svelte",
    speed: "blazing",
    version: 4,
    website: "https://svelte.dev",
  };
</script>

<TailwindCss />

<main>
  <Navbar />
  <div class="hero min-h-screen bg-base-200">
    <div class="hero-content text-center">
      <div class="max-w-md">
        <!-- <SvelteLogo /> -->
        <Clock />

        <Answer />
        <Answer answer="42" />
        <PackageInfo {...pkg} />

        <div class="my-3">
          <Counter />
        </div>

        <div class="mb-3">
          <Adder />
        </div>

        <div class="mb-3">
          <ThingRemover />
        </div>

        <div class="mb-3">
          <Awaiter />
        </div>

        <Greeter
          on:message={(event) => {
            alert(event.detail.text);
          }}
        />

        <div class="mb-3">
          <Input />
        </div>

        <div class="mb-3">
          <UglyCalculator />
        </div>

        <div class="mb-3">
          <Spam />
        </div>
      </div>
    </div>
  </div>
  <Footer />
</main>
