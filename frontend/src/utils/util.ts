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
    clearInterval()
    timer = setTimeout(() => {
      fn(...args)
    }, wait)
  }
}