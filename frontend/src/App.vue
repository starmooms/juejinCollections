<template>
  <img alt="Vue logo" src="./assets/logo.png" />
  <HelloWorld msg="Hello Vue 3.0 + Vite" />
</template>

<script lang="ts">
import HelloWorld from "./components/HelloWorld.vue";

export default {
  name: "App",
  components: {
    HelloWorld,
  },
  mounted() {

    (async () => {
      const data = await fetch("/api/abc?a=2", {
        method: 'POST',
        body: JSON.stringify({
          a: 1,
          b: 2
        })
      })
      console.log(data)
    })()
  }
};
</script>
