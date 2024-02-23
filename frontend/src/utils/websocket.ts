import { getPromiseState, tryOnUnmounted } from "./util";
import Message from "../components/Message";

type StockRespone<T = any> = {
  type: string
  data: T
}

type CreateWsParams = {
  url: string
  onMessage?: <T extends any>(v: StockRespone<T>) => void
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/CloseEvent#status_codes
const CODE = {
  CLOSE_NORMAL: 1000,
  CLOSE_ABNORMAL: 1006
}

function initWebStocket(params: {
  wsUrl: string,
}) {

  const { wsUrl } = params
  const stock = new WebSocket(wsUrl);
  const { rejectFn, resolveFn, promiseState } = getPromiseState<WebSocket>()

  stock.addEventListener('open', () => {
    console.log("connected to " + wsUrl);
    resolveFn(stock)
  })

  stock.addEventListener('error', (err) => {
    console.log("ws error", err);
  })

  stock.addEventListener('close', (e) => {
    console.log("connection closed (" + e.code + ")");
    rejectFn(new Error("stock is closed"))
  })

  return {
    stock,
    openWait: promiseState
  }
}

export const createWebStocket = (params: CreateWsParams) => {
  const { url, onMessage } = params
  const state = {
    wsUrl: `${location.protocol === "https:" ? "wws" : "ws"}://${location.host}${params.url}`,
    connect: true,
    retriesTimer: 0,
  }

  let stockState = null as unknown as ReturnType<typeof initWebStocket>
  let stock = null as unknown as ReturnType<typeof initWebStocket>['stock']

  const isCloseStock = () => {
    const closeState: number[] = [WebSocket.CLOSING, WebSocket.CLOSED]
    return !stock || closeState.includes(stock.readyState)
  }

  const clearRetries = () => {
    clearTimeout(state.retriesTimer)
  }

  const init = () => {
    if (!state.connect || !isCloseStock()) {
      return
    }
    clearRetries()
    stockState = initWebStocket({
      wsUrl: state.wsUrl
    })
    stock = stockState.stock

    stock.addEventListener("close", () => {
      state.retriesTimer = setTimeout(() => {
        init()
      }, 3000)
    })

    stock.addEventListener("message", (e) => {
      const data = JSON.parse(e.data) as StockRespone
      if (data.type === 'tip') {
        alert(data.data)
      }
      onMessage?.(data)
    })
  }
  init()


  const getStock = async () => {
    // 不等待 马上重连
    if (isCloseStock()) {
      init()
    }
    return stockState.openWait
  }

  const send = async (data: Record<any, any>) => {
    try {
      const stock = await getStock()
      stock.send(JSON.stringify(data))
    } catch (err) {
      console.log("handle err ???")
      Message.msg(`err ${(err as any).message}`)
      throw err
    }
  }

  const close = () => {
    state.connect = false
    clearRetries()
    if (isCloseStock()) {
      return
    }
    stock.close(CODE.CLOSE_NORMAL)
  }

  tryOnUnmounted(() => {
    close()
  })

  return {
    stock,
    send,
    close
  }
}