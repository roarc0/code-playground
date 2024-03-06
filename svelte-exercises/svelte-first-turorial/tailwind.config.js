export default {
  plugins: [require("@tailwindcss/typography"),require("daisyui")],
    theme: {
      extend: {},
    },
  content: ["./index.html",'./src/**/*.{svelte,js,ts}'],
  daisyui: {
    themes: ["light", "dark", "cupcake"],
  },
}
