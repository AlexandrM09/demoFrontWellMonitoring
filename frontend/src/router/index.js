import Vue from 'vue'
import Router from 'vue-router'
import First from '@/components/First'
import Main from '@/pages/index'
import page1 from '@/pages/page1/page1'
import page2 from '@/pages/page2/page2'
import page3 from '@/pages/page3/page3'
import page4 from '@/pages/page4/page4'
import page5 from '@/pages/page5/page5'
import page6 from '@/pages/page6/page6'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/test',
      name: 'First',
      component: First
    },
    {
      path: '/',
      name: 'Main',
      component: Main,
      children: [
        {
          path: 'page1',
          name: 'page1',
          component: page1
        },
        {
          path: 'page2',
          name: 'page2',
          component: page2
        },
        {
          path: 'page3',
          name: 'page3',
          component: page3
        },
        {
          path: 'page4',
          name: 'page4',
          component: page4
        },
        {
          path: 'page5',
          name: 'page5',
          component: page5
        },
        {
          path: 'page6',
          name: 'page6',
          component: page6
        }
      ]
    }
  ],
  mode: 'history'
})
