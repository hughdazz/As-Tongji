import { createRouter,createWebHashHistory} from "vue-router"

import initial_Page from '../page/initial_Page.vue'
import search from '../page/search.vue'
import chat from '../page/chat.vue'
import calender from '../page/calender.vue'

const router = createRouter({
  history: createWebHashHistory(),
  routes: [
		{
		  path:'/:pathMatch(.*)*',
		  redirect: '/initial_Page',
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

