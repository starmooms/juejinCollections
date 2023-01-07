<template>
  <template v-if="article">
    <article
      class="article-container markdown-body"
      :class="{ 'is-markdown': article.isMarkdown }"
      ref="articleRef"
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
import VueMarkdown from "../components/vueMarkdown";
// import hljs from 'highlight.js';
import prismjs from "prismjs";
import { computed, defineComponent, nextTick, onMounted, ref } from "vue";
import { Article } from "../type";
import { useRoute } from "vue-router";
import { getArticle } from "../utils/api";

interface Data {
  article: null | Article;
}

export default defineComponent({
  name: "App",
  components: {
    VueMarkdown,
  },
  setup() {
    const route = useRoute();
    console.log(route);
    const numc = ref(0);
    const article = ref<any>(null);
    const articleIdRef = computed(() => {
      const rId = route.params.articleId as string;
      console.log(rId);
      return rId || "6844903974378668039";
    });

    const getArticleData = async () => {
      // 6844904178075058189 6844903974378668039
      console.log(articleIdRef.value);
      const result = await getArticle({
        articleId: articleIdRef.value,
      });
      if (result.data.status) {
        let articleData = {
          ...result.data.data.article,
          isMarkdown: false,
        };
        let articleId = articleData.article_id;
        articleData.isMarkdown = !!articleData.mark_content;
        console.log(result.data.data, !!articleData.mark_content);

        let replaceImgStr = (...args: string[]) => {
          if (args.length >= 4) {
            // prettier-ignore
            return `${args[1]}//localhost:8012/images/article/${articleId}?url=${encodeURIComponent(args[2])}${args[3]}`;
          }
          return args[0];
        };

        if (articleData.isMarkdown) {
          articleData.mark_content = articleData.mark_content.replace(
            /(\!\[.*?\]\()(http\S+)(.*?\))/g,
            replaceImgStr
          );
        } else {
          articleData.content = articleData.content.replace(
            /(<img.*?src=")(http.*?)(".*?>)/g,
            replaceImgStr
          );
        }
        article.value = articleData;
      }
    };

    onMounted(async () => {
      await nextTick();
      getArticleData().then(() => {
        console.log(article.value);
      });
    });

    return { numc, article };
  },
});
</script>



<style lang="less">
</style>
