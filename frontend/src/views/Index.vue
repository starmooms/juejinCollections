<template>
  <div>
    <button @click="syncColletion">同步收藏集</button>
  </div>
  <div>
    <input type="text">
  </div>
</template>

<script lang="ts" setup>
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue";
import { Article } from "../type";
import { useRoute } from "vue-router";
import { postSyncCollection } from "../utils/api";
import Message from "../components/Message";


const createWebStocket = () => {
  // https://developer.mozilla.org/zh-CN/docs/Web/API/CloseEvent#status_codes
  const CODE = {
    CLOSE_NORMAL: 1000,
    CLOSE_ABNORMAL: 1006

  }
  const baseUri = location.href.replace(/^http/, 'ws').replace(/\/$/, '')
  const wsuri = `${baseUri}/echo`
  var sock = new WebSocket(wsuri);

  sock.addEventListener('open', () => {
    console.log("connected to " + wsuri);
  })

  sock.addEventListener('close', (e) => {
    console.log("connection closed (" + e.code + ")");
  })

  sock.addEventListener('message', (e) => {
    console.log(JSON.parse(e.data))
    console.log("message received: " + e.data, typeof e.data);
  })

  const send = (data: Record<any, any>) => {
    sock.send(JSON.stringify(data))
  }

  const close = () => {
    sock.close(CODE.CLOSE_NORMAL)
  }

  return {
    sock,
    send,
    close
  }

  // window.closeWW = function () {
  //   // https://developer.mozilla.org/zh-CN/docs/Web/API/WebSocket/close
  //   sock.close(CODE.CLOSE_NORMAL, 'client close')
  // }
  // window.sendWW = function (v) {
  //   // https://developer.mozilla.org/zh-CN/docs/Web/API/WebSocket/close
  //   sock.send(v)
  //   sock.close(CODE.CLOSE_NORMAL, 'client close')
  // }
}

const wsState = createWebStocket()

const syncColletion = async () => {
  wsState.send({
    type: 'action',
    data: 'run'
  })
  // const result = await postSyncCollection()
  // if (result.data.status) {
  //   Message.msg("正在同步收藏集...")
  // }
}

// onMounted(() => {
//   createWebStocket()
//   console.log('??')
// })

onUnmounted(() => {
  wsState.close()
})

</script>



<style lang="less">

</style>
