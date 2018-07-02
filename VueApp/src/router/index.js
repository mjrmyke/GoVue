import Vue from 'vue';
import Router from 'vue-router';
import HelloWorld from '@/components/HelloWorld';
import RollRoom from '@/components/RollRoom';
import RoomList from '@/components/RoomList';
import LO from 'lodash';



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
      meta: {title: 'Home'}
    },
    {
      path: '/roll',
      name: 'RoomList',
      component: RoomList,
      props: true,
      meta: {title: 'Home'}
    },
  ],
});