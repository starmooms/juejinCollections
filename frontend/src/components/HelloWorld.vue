<template>
  <h1>{{ msg }}</h1>
  <button @click="count++">count is: {{ count }}</button>
  <p>
    Edit <code>components/HelloWorld.vue</code> to test hot module replacement.
  </p>
  <img
    src="//localhost:8012/images/article/6844904178075058189?url=https%3A%2F%2Fuser-gold-cdn.xitu.io%2F2020%2F6%2F1%2F1726f38e4afb7b1d%3FimageView2%2F0%2Fw%2F1280%2Fh%2F960%2Fformat%2Fwebp%2Fignore-error%2F1"
  />
</template>

<script lang="ts">
export default {
  name: 'HelloWorld',
  props: {
    msg: String
  },
  data() {
    return {
      count: 0
    }
  }
}
</script>
