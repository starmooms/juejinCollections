<template>
  <div class="w-100vw h-100vh flex justify-center items-center">
    <div class="w-800px max-w-full">
      <div class="bg-primary-300 b-rd-4">
        <ul class="flex justify-center p-2">
          <li v-for="item in tabList" :key="item.id" @click="selectTab(item)" :class="{ 'selected': item.isSelect }"
            class="p-2 cursor-pointer block text-center text-white flex-1 b-rd-4 h-selected:bg-primary hover:bg-primary/60">
            {{ item.tab }}
          </li>
        </ul>
      </div>
      <div class="h-600px my-4">
        <div class="tab1" v-if="curTab === 1">
          <div class="mx-auto max-w-600px">
            <input v-model.trim="keyword" @input="handleSearchList" type="text"
              class="b-rd-2 leading-40px b-none outline-none w-full px-4 text-xl" />
          </div>
        </div>
        <div class="tab2" v-if="curTab === 2">
          <div>
            <button
              class="text-white bg-primary py-2 px-3 cursor-pointer b-rd-2 border-none outline-none hover:bg-primary-500"
              @click="syncColletion">同步收藏集</button>
            <pre>
              <code class="block">{{ log }}</code>
            </pre>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { computed, nextTick, onMounted, onUnmounted, ref } from "vue";
import { searchArticleList } from "../utils/api";
import { useDebounce } from '../utils/util'

const log = ref('')

const curTab = ref(2);
const createTabItem = (id: number, tab: string) => {
  return { id, tab, isSelect: computed(() => curTab.value === id) }
}
const tabList = ref([
  createTabItem(1, '查找文章'),
  createTabItem(2, '同步收藏集'),
])

const selectTab = (tabItem: { id: number }) => {
  curTab.value = tabItem.id
}


const keyword = ref('')
const articleList = ref([])

const handleSearchList = () => useDebounce(async () => {
  const data = await searchArticleList({
    keyword: keyword.value
  })

}, 100)



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
    const data = JSON.parse(e.data)
    if (data.type === 'log') {
      log.value += `${data.data}\n`
    } else if (data.type === 'tip') {
      alert(data.data)
    }
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
