import { createRouter,createWebHashHistory} from "vue-router"

import initial_Page from '../page/initial_Page.vue'
import search from '../page/search.vue'
import chat from '../page/chat.vue'
import calender from '../page/calender.vue'
import login from '../page/login.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [//初始登录跳转有问题
		{
		  path:'/:pathMatch(.*)*',
		  redirect: '/login',
		},
		{
		  path: '/login',
		  name:"login",
		  component:login
		},
		{ 
		  path: '/initial_Page',
		  name:"initial_Page",
		  component:initial_Page
		},
		{ 
		  path: '/search',
		  name:"search",
		  component:search
		},
		{ 
		  path: '/chat',
		  name:"chat",
		  component:chat
		},
		{
		  path: '/calender',
		  name:"calender",
		  component:calender
		}
	]
});
 
export default router;

