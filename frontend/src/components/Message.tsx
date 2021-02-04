import { defineComponent, App, h, render } from 'vue'


export const MessageComponents = defineComponent({
  props: {
    msg: {
      type: String,
      required: true
    }
  },
  data() {
    return {

    }
  },
  render() {
    return (
      <div>
        <p>{this.msg}</p>
      </div>
    )
  }
})

interface MessagePulgin {
  (msg: string): void
}

declare module '@vue/runtime-core' {
  export interface ComponentCustomProperties {
    $message: MessagePulgin
  }
}

export default {
  install: (app: App) => {
    app.config.globalProperties.$message = (msg: string) => {
      const container = document.createElement("div")
      const vNode = h(MessageComponents, {
        msg: msg
      })
      render(vNode, container)
      document.body.appendChild(container.firstElementChild!)
    }
  }
}