<template>
  <template v-if="article">
    <article
      class="article-container markdown-body"
      :class="{ 'is-markdown': article.isMarkdown }"
      ref="article"
      style="width: 700px; margin: 0 auto"
    >
      <h1>{{ article.title }}</h1>
      <VueMarkdown v-if="article.mark_content">
        {{ article.mark_content }}
      </VueMarkdown>
      <div v-else v-html="article.content"></div>
    </article>
  </template>
</template>

<script lang="ts">
import VueMarkdown from "./components/vueMarkdown";
// import hljs from 'highlight.js';
import prismjs from "prismjs"
import { defineComponent } from "vue";
import { article } from "./type"
import "./utils/api.ts"

interface Data {
  article: null | article
}

export default defineComponent<any, any, Data>({
  name: "App",
  components: {
    VueMarkdown,
  },
  data(): Data {
    return {
      article: null
    }
  },
  methods: {
    async getArticle() {
      // 6844904178075058189 6844903974378668039
      const response = await fetch("/api/getArticle?articleId=6844903974378668039", {
        method: 'GET'
      })
      const result = await response.json()
      if (result.status) {
        let article = result.data.article
        let articleId = article.article_id
        article.isMarkdown = !!article.mark_content
        let replaceImgStr = (...args: string[]) => {
          if (args.length >= 4) {
            return `${args[1]}//localhost:8012/images/article/${articleId}?url=${encodeURIComponent(args[2])}${args[3]}`
          }
          return args[0]
        }

        if (article.isMarkdown) {
          article.mark_content = article.mark_content.replace(/(\!\[.*?\]\()(http\S+)(.*?\))/g, replaceImgStr)
        } else {
          article.content = article.content.replace(/(<img.*?src=")(http.*?)(".*?>)/g, replaceImgStr)
        }
        this.article = article
      }
    }
  },
  mounted() {
    this.getArticle().then(() => {
      console.log(this.article)
    })
  }
});
</script>



<style lang="less">
</style>
