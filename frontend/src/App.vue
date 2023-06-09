<template>
  <router-view></router-view>
</template>

<script lang="ts">
import { computed, defineComponent, nextTick, onMounted, ref } from "vue";
import { Article } from "./type";
import "./utils/api.ts";
interface Data {
  article: null | Article;
}

export default defineComponent({
  name: "App",
  setup() { },
});
</script>
