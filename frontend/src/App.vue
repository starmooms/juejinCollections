<template>
  <template v-if="article">
    <div style="width: 700px; margin: 0 auto">
      <h1>{{ article.title }}</h1>
      <p v-if="article.mark_content">dd</p>
      <VueMarkdown v-if="article.mark_content">{{
        article.mark_content
      }}</VueMarkdown>
      <article class="markdown-body" v-html="article.content"></article>
    </div>
  </template>
</template>

<script lang="ts">
import VueMarkdown from "./components/vueMarkdown";

export default {
  name: "App",
  components: {
    VueMarkdown,
  },
  data() {
    return {
      article: null
    }
  },
  methods: {
    async getArticle() {
      const response = await fetch("/api/getArticle?articleId=6844903974378668039", {
        method: 'GET'
      })
      const result = await response.json()
      if (result.status) {
        let article = result.data.article
        let articleId = article.article_id
        let replaceImgStr = (...args) => {
          if (args.length >= 4) {
            return `${args[1]}//localhost:8012/images/article/${articleId}?url=${encodeURIComponent(args[2])}${args[3]}`
          }
          return args[0]
        }

        if (article.mark_content) {
          article.mark_content = article.mark_content.replace(/(\!\[.*?\]\()(http\S+)(.*?\))/g, replaceImgStr)
        } else if (article.content) {
          article.content = article.content.replace(/(<img.*?src=")(http.*?)(".*?>)/g, replaceImgStr)
        }
        this.article = article
      }

      // console.log(response)
    }
  },
  mounted() {
    this.getArticle().then(() => {
      console.log(this.article)
    })
    // (async () => {
    //   // const data = await fetch("/api/getArticle?a=2", {
    //   //   method: 'POST',
    //   //   body: JSON.stringify({
    //   //     articleId: "6844904178075058189"
    //   //   })
    //   // })

    //   // if (result.status) {
    //   //   vm.article = result.data
    //   // }
    // })()
  }
};
</script>

<style>
.markdown-body {
  word-break: break-word;
  line-height: 1.75;
  font-weight: 400;
  font-size: 15px;
  overflow-x: hidden;
  color: #333;
}
.markdown-body h1,
.markdown-body h2,
.markdown-body h3,
.markdown-body h4,
.markdown-body h5,
.markdown-body h6 {
  line-height: 1.5;
  margin-top: 35px;
  margin-bottom: 10px;
  padding-bottom: 5px;
}
.markdown-body h1 {
  font-size: 30px;
  margin-bottom: 5px;
}
.markdown-body h2 {
  padding-bottom: 12px;
  font-size: 24px;
  border-bottom: 1px solid #ececec;
}
.markdown-body h3 {
  font-size: 18px;
  padding-bottom: 0;
}
.markdown-body h4 {
  font-size: 16px;
}
.markdown-body h5 {
  font-size: 15px;
}
.markdown-body h6 {
  margin-top: 5px;
}
.markdown-body p {
  line-height: inherit;
  margin-top: 22px;
  margin-bottom: 22px;
}
.markdown-body img {
  max-width: 100%;
}
.markdown-body hr {
  border: none;
  border-top: 1px solid #ddd;
  margin-top: 32px;
  margin-bottom: 32px;
}
.markdown-body code {
  word-break: break-word;
  border-radius: 2px;
  overflow-x: auto;
  background-color: #fff5f5;
  color: #ff502c;
  font-size: 0.87em;
  padding: 0.065em 0.4em;
}
.markdown-body code,
.markdown-body pre {
  font-family: Menlo, Monaco, Consolas, Courier New, monospace;
}
.markdown-body pre {
  overflow: auto;
  position: relative;
  line-height: 1.75;
}
.markdown-body pre > code {
  font-size: 12px;
  padding: 15px 12px;
  margin: 0;
  word-break: normal;
  display: block;
  overflow-x: auto;
  color: #333;
  background: #f8f8f8;
}
.markdown-body a {
  text-decoration: none;
  color: #0269c8;
  border-bottom: 1px solid #d1e9ff;
}
.markdown-body a:active,
.markdown-body a:hover {
  color: #275b8c;
}
.markdown-body table {
  display: inline-block !important;
  font-size: 12px;
  width: auto;
  max-width: 100%;
  overflow: auto;
  border: 1px solid #f6f6f6;
}
.markdown-body thead {
  background: #f6f6f6;
  color: #000;
  text-align: left;
}
.markdown-body tr:nth-child(2n) {
  background-color: #fcfcfc;
}
.markdown-body td,
.markdown-body th {
  padding: 12px 7px;
  line-height: 24px;
}
.markdown-body td {
  min-width: 120px;
}
.markdown-body blockquote {
  color: #666;
  padding: 1px 23px;
  margin: 22px 0;
  border-left: 4px solid #cbcbcb;
  background-color: #f8f8f8;
}
.markdown-body blockquote:after {
  display: block;
  content: "";
}
.markdown-body blockquote > p {
  margin: 10px 0;
}
.markdown-body ol,
.markdown-body ul {
  padding-left: 28px;
}
.markdown-body ol li,
.markdown-body ul li {
  margin-bottom: 0;
  list-style: inherit;
}
.markdown-body ol li .task-list-item,
.markdown-body ul li .task-list-item {
  list-style: none;
}
.markdown-body ol li .task-list-item ol,
.markdown-body ol li .task-list-item ul,
.markdown-body ul li .task-list-item ol,
.markdown-body ul li .task-list-item ul {
  margin-top: 0;
}
.markdown-body ol ol,
.markdown-body ol ul,
.markdown-body ul ol,
.markdown-body ul ul {
  margin-top: 3px;
}
.markdown-body ol li {
  padding-left: 6px;
}
@media (max-width: 720px) {
  .markdown-body h1 {
    font-size: 24px;
  }
  .markdown-body h2 {
    font-size: 20px;
  }
  .markdown-body h3 {
    font-size: 18px;
  }
}

.markdown-body pre,
.markdown-body pre > code.hljs {
  color: #333;
  background: #f8f8f8;
}
.hljs-comment,
.hljs-quote {
  color: #998;
  font-style: italic;
}
.hljs-keyword,
.hljs-selector-tag,
.hljs-subst {
  color: #333;
  font-weight: 700;
}
.hljs-literal,
.hljs-number,
.hljs-tag .hljs-attr,
.hljs-template-variable,
.hljs-variable {
  color: teal;
}
.hljs-doctag,
.hljs-string {
  color: #d14;
}
.hljs-section,
.hljs-selector-id,
.hljs-title {
  color: #900;
  font-weight: 700;
}
.hljs-subst {
  font-weight: 400;
}
.hljs-class .hljs-title,
.hljs-type {
  color: #458;
  font-weight: 700;
}
.hljs-attribute,
.hljs-name,
.hljs-tag {
  color: navy;
  font-weight: 400;
}
.hljs-link,
.hljs-regexp {
  color: #009926;
}
.hljs-bullet,
.hljs-symbol {
  color: #990073;
}
.hljs-built_in,
.hljs-builtin-name {
  color: #0086b3;
}
.hljs-meta {
  color: #999;
  font-weight: 700;
}
.hljs-deletion {
  background: #fdd;
}
.hljs-addition {
  background: #dfd;
}
.hljs-emphasis {
  font-style: italic;
}
.hljs-strong {
  font-weight: 700;
}
</style>
