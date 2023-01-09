<template>
  <div>
    <button @click="syncColletion">同步收藏集</button>
  </div>
  <div>
    <input type="text">
  </div>
</template>

<script lang="ts" setup>
import { computed, nextTick, onMounted, ref } from "vue";
import { Article } from "../type";
import { useRoute } from "vue-router";
import { postSyncCollection } from "../utils/api";
import Message from "../components/Message";


const createWebStocket = () => {
  const baseUri = location.href.replace(/^http/, 'ws').replace(/\/$/, '')
  const wsuri = `${baseUri}/echo`
  var sock = new WebSocket(wsuri);

  sock.onopen = function () {
    console.log("connected to " + wsuri);
  }

  sock.onclose = function (e) {
    console.log("connection closed (" + e.code + ")");
  }

  sock.onmessage = function (e) {
    console.log("message received: " + e.data);
  }
}

const syncColletion = async () => {
  const result = await postSyncCollection()
  if (result.data.status) {
    Message.msg("正在同步收藏集...")
  }
}

onMounted(() => {
  createWebStocket()
  console.log('??')
})

</script>



<style lang="less">

</style>
