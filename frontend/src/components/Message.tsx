import { defineComponent, Transition, App, h, render, ref, createApp, nextTick, onMounted, onUnmounted } from 'vue'

export const MessageComponents = defineComponent({
  props: {
    msg: {
      type: String,
      required: true
    },
    time: {
      type: Number,
      required: true,
    }
  },
  setup(props, ctx) {

    const show = ref(false)

    nextTick().then(() => {
      show.value = true
    })

    const timer = setTimeout(() => {
      show.value = false
    }, props.time)

    onUnmounted(() => {
      clearTimeout(timer)
    })

    const onClose = () => {
      ctx.emit("close")
    }

    return () => {
      return (
        <Transition name="fade" onAfterLeave={onClose}>
          <div class="x-message" v-show={show.value}>
            <p>{props.msg}</p>
          </div>
        </Transition>
      )
    }
  },
})

type MessageOpts = {
  time?: number
}

interface MessagePulgin {
  (msg: string, opts?: MessageOpts): void
}

declare module '@vue/runtime-core' {
  export interface ComponentCustomProperties {
    $message: MessagePulgin
  }
}

function useMessageBox() {
  let boxEl: HTMLDivElement | null = null
  let num = 0
  const getBoxEl = () => {
    num += 1
    if (!boxEl) {
      boxEl = document.createElement("div")
      document.body.appendChild(boxEl)
    }
    return boxEl
  }

  const removeBoxEl = () => {
    num -= 1
    if (num <= 0 && boxEl) {
      num = 0
      document.body.removeChild(boxEl)
      boxEl = null
    }
  }
  return {
    getBoxEl,
    removeBoxEl
  }
}

const boxElState = useMessageBox()
let lastItemClose: null | (() => void) = null

function message(msg: string, _opts?: MessageOpts) {

  const opts = {
    time: 3000,
    ..._opts
  }
  const boxEl = boxElState.getBoxEl()
  const container = document.createElement('div')
  container.classList.add("x-message-box")

  lastItemClose?.()
  const onClose = () => {
    vm.unmount()
    boxEl.removeChild(container)
    boxElState.removeBoxEl()
    lastItemClose = null
  }
  lastItemClose = onClose
  const vm = createApp(MessageComponents, {
    msg: msg,
    time: opts.time,
    onClose: onClose
  })
  vm.mount(container)
  boxEl.appendChild(container)
}


export default {
  msg: message,
  install: (app: App) => {
    app.config.globalProperties.$message = (msg: string, opts?: MessageOpts) => {
      message(msg, opts)
    }
  },
  info: message
}
