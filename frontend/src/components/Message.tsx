import { defineComponent, App, h, render, ref, createApp } from 'vue'
console.log('33')

export const MessageComponents = defineComponent({
  props: {
    msg: {
      type: String,
      required: true
    }
  },
  setup(props) {
    return () => {
      return (
        <div>
          <p>{props.msg}</p>
        </div>
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

function message(msg: string, _opts?: MessageOpts) {
  const opts = {
    time: 2000,
    ..._opts
  }

  const boxEl = boxElState.getBoxEl()
  const container = document.createElement('div')

  const vm = createApp(MessageComponents, {
    msg: msg
  })
  vm.mount(container)
  boxEl.appendChild(container)

  setTimeout(() => {
    vm.unmount()
    boxEl.removeChild(container)
    boxElState.removeBoxEl()
  }, opts.time)
}


export default {
  msg: message,
  install: (app: App) => {
    app.config.globalProperties.$message = (msg: string, opts?: MessageOpts) => {
      message(msg, opts)
    }
  }
}