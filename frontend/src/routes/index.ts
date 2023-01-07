import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  // routes: [],
  routes: [
    {
      name: 'article',
      path: '/article/:articleId',
      component: () => import('../views/Article.vue')
    },
    {
      path: '/:path(.*)*',
      redirect: '/article/6844903974378668039'
    }
  ]
})

export {
  router
}