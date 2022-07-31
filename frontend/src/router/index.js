import Vue from 'vue'
import VueRouter from 'vue-router'
import DataSource from '../views/DataSource.vue'
import Page from '../views/Page.vue'
import DataTable from '../views/DataTable.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/datasource',
    name: 'DataSource',
    component: DataSource
  },
  {
    path: '/page',
    name: 'Page',
    component: Page,
  },
  {
    path: '/datatable',
    name: 'DataTable',
    component: DataTable
  },
  {
    path: '*',
    redirect: '/datasource'
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router