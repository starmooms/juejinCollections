import { createRouter, createWebHistory } from "vue-router"

const router = createRouter({
  history: createWebHistory(),
  // routes: [],
  routes: [
    {
      path: '/article/:articleId',
      component: () => import('../views/Article.vue')
    }
  ]
})

export {
  router
}