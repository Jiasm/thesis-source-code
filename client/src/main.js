import Vue from 'vue'
import VueRouter from 'vue-router'
import ElementUI from 'element-ui'
import App from './App.vue'
import Login from './components/Login.vue'
import List from './components/List.vue'
import Create from './components/Create.vue'
import Preview from './components/Preview.vue'
import Edit from './components/Edit.vue'
import Reget from './components/Reget.vue'
import UserLogin from './components/UserLogin.vue'
import UserSignin from './components/UserSignin.vue'
import 'element-ui/lib/theme-chalk/index.css'

Vue.use(ElementUI)
Vue.use(VueRouter)

Vue.config.productionTip = false

const routes = [
  { path: '/', component: Login },
  { path: '/list', component: List },
  { path: '/create', component: Create },
  { path: '/preview', component: Preview },
  { path: '/edit', component: Edit },
  { path: '/reget', component: Reget },
  { path: '/user-login', component: UserLogin },
  { path: '/user-signin', component: UserSignin },
]

const router = new VueRouter({
  routes,
})

new Vue({
  render: h => h(App),
  router,
}).$mount('#app')
