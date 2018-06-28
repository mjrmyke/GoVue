import Vue from 'vue';
import Router from 'vue-router';
import HelloWorld from '@/components/HelloWorld';
import RollRoom from '@/components/RollRoom';


Vue.use(Router);
window.Vue = Vue;

export default new Router({
  routes: [
    {
      path: '/',
      name: 'HelloWorld',
      component: HelloWorld,
    },
    {
      path: '/roll/:roomid',
      name: 'RollRoom',
      component: RollRoom,
      props: true,
    },
  ],
});
