import Vue from 'vue'
import VueRouter from 'vue-router'
import ElementUI from 'element-ui'
import App from './App.vue'
import Login from './components/Login.vue'
import List from './components/List.vue'
import ProjectMemberList from './components/ProjectMemberList'
import GroupMemberList from './components/GroupMemberList'
import UserLogin from './components/UserLogin.vue'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)
Vue.use(VueRouter)

Vue.config.productionTip = false

const routes = [
  { path: '/', component: Login },
  { path: '/list', component: List },
  { path: '/project-manage', component: ProjectMemberList },
  { path: '/group-manage', component: GroupMemberList },
  { path: '/user-login', component: UserLogin },
]

const router = new VueRouter({
  routes,
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
