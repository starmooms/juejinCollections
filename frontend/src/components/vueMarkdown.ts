import vueMarkdown from 'vue-markdown'
import { h, nextTick, VNode } from "vue"
import * as Prism from "prismjs"
import "prismjs/plugins/line-numbers/prism-line-numbers.min"

const oldRender = vueMarkdown.render
vueMarkdown.render = function () {
  let result = null as unknown as VNode
  oldRender.call(this, (...args: any) => {
    result = h(args[0], {
      innerHTML: args[1].domProps.innerHTML
    })
  })
  nextTick(() => {
    let el = result.el as HTMLDivElement
    if (el) {
      el.classList.add('line-numbers')
      el.querySelectorAll('pre code').forEach(el => {
        if (el.classList.length === 0) {
          el.classList.add("language-text")
        }
      })
      Prism.highlightAllUnder(el)
    }
  })
  return result
}

vueMarkdown.beforeMount = function () {
  if (this.$slots.default) {
    this.sourceData = ''
    for (let slot of this.$slots.default()) {
      this.sourceData += slot.children
    }
  }

  this.$watch('source', () => {
    this.sourceData = this.prerender(this.source)
    this.$forceUpdate()
  })

  this.watches.forEach((v: any) => {
    this.$watch(v, () => {
      this.$forceUpdate()
    })
  })
}

export default vueMarkdown