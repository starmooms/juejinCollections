import { getCurrentInstance, onUnmounted } from "vue"

export function tryOnUnmounted(fn: () => any) {
  if (getCurrentInstance()) {
    onUnmounted(fn)
  }
}

export function useDebounce<T extends (...args: any[]) => any>(fn: T, wait: number) {
  let timer: null | number = null
  const clear = () => {
    if (timer !== null) {
      clearInterval(timer)
    }
    timer = null
  }

  tryOnUnmounted(clear)

  return (...args: Parameters<T>) => {
    clear()
    timer = setTimeout(() => {
      fn(...args)
    }, wait)
  }
}

export function getPromiseState<T = any>() {
  let resolveFn: (value: T | PromiseLike<T>) => void = null as any
  let rejectFn: (reason?: any) => void = null as any
  let isPending = true

  let wrapPending = <F extends Function>(fn: F): F => {
    return function (...params: any) {
      if (!isPending) {
        return
      }
      fn(...params)
      isPending = false
    } as any
  }
  const promiseState = new Promise<T>((resolve, reject) => {
    resolveFn = wrapPending(resolve)
    rejectFn = wrapPending(reject)
  })

  return {
    resolveFn,
    rejectFn,
    promiseState,
  }
}