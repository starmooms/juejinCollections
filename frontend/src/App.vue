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
      // const data = await fetch("/api/getArticle?a=2", {
      //   method: 'POST',
      //   body: JSON.stringify({
      //     articleId: "6844904178075058189"
      //   })
      // })
      const data = await fetch("/api/getArticle?a=2", {
        method: 'GET'
      })
      console.log(data)
    })()
  }
};
</script>
