<script>
  import Icon from "@iconify/svelte";
</script>

<div class="navbar bg-base-100">
  <div class="navbar-start">
    <div class="dropdown">
      <label tabIndex={0} class="btn btn-ghost btn-circle">
        <Icon icon="mdi-light:hamburger" class="text-3xl" />
      </label>
      <ul
        tabIndex={0}
        class="menu menu-sm dropdown-content mt-3 z-[1] p-2 shadow bg-base-100 rounded-box w-52"
      >
        <li><a>Homepage</a></li>
        <li><a>Portfolio</a></li>
        <li><a>About</a></li>
      </ul>
    </div>
  </div>
  <div class="navbar-center">
    <a class="btn btn-ghost normal-case text-xl">Svelte Exercises</a>
  </div>
  <div class="navbar-end">
    <button class="btn btn-ghost btn-circle">
      <Icon icon="mdi-light:magnify" class="text-2xl" />
    </button>
    <button class="btn btn-ghost btn-circle">
      <div class="indicator">
        <Icon icon="mdi-light:bell" class="text-2xl" />
        <span class="badge badge-xs badge-primary indicator-item" />
      </div>
    </button>
  </div>
</div>
