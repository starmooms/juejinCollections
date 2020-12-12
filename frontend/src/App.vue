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
      const data = await fetch("/api/abc", {
        method: 'POST'
      })
      console.log(data)
    })()
  }
};
</script>
