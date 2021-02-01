import vueMarkdown from 'vue-markdown'
import { h } from "vue"

const oldRender = vueMarkdown.render
vueMarkdown.render = function () {
  let restul = null
  oldRender.call(this, (...args) => {
    restul = h(args[0], {
      innerHTML: args[1].domProps.innerHTML
    })
  })
  return restul
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

  this.watches.forEach((v) => {
    this.$watch(v, () => {
      this.$forceUpdate()
    })
  })
}

export default vueMarkdown