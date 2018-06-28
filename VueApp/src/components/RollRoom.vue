<template>
  <div class="hello">
    <h1> ok</h1>
    <h1>{{ roomid }}</h1>
    <h2>Hiya</h2>
    <ul>
      <li>
        <a
          href="https://vuejs.org"
          target="_blank"
        >
          Core Docs
        </a>
      </li>
      <li>
        <a
          href="https://forum.vuejs.org"
          target="_blank"
        >
          Forum
        </a>
      </li>
      <li>
        <a
          href="https://chat.vuejs.org"
          target="_blank"
        >
          Community Chat
        </a>
      </li>
      <li>
        <a
          href="https://twitter.com/vuejs"
          target="_blank"
        >
          Twitter
        </a>
      </li>
      <br>
      <li>
        <a
          href="http://vuejs-templates.github.io/webpack/"
          target="_blank"
        >
          Docs for This Template
        </a>
      </li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li>
        <a
          href="http://router.vuejs.org/"
          target="_blank"
        >
          vue-router
        </a>
      </li>
      <li>
        <a
          href="http://vuex.vuejs.org/"
          target="_blank"
        >
          vuex
        </a>
      </li>
      <li>
        <a
          href="http://vue-loader.vuejs.org/"
          target="_blank"
        >
          vue-loader
        </a>
      </li>
      <li>
        <a
          href="https://github.com/vuejs/awesome-vue"
          target="_blank"
        >
          awesome-vue
        </a>
      </li>
    </ul>
    <ul id="messages"></ul>
  <input type="text" v-model="message"/>
  <button @click="emitEvent">emit</button>
  <div><span>Response is:</span>{{response}}</div>
  </div>
</template>

<script>
/* eslint-disable */


export default {
  name: 'RollRoom',
  data() {
    return {
      msg: 'Roll Room Test',
      ws: null
    };
  },
  created: function() {
      this.ws = new WebSocket("ws://" + window.location.hostname +'/ws'+ window.location.hash.split('#').join(''))
      this.windowLocation = window.location.hash.split('#');
    
      this.ws.onopen = function (event) {
        console.log("Opened WS connection");
      }
      
      this.ws.onmessage = event => {
        console.log("received message: " + event.data);
        this.response = event.data;
      }
  },
  methods: {
    emitEvent() {
      this.ws.send(this.message);
      console.log('event emitted')
    }

  },
  data: function() {
    return {
      message : '',
      roomid: this.windowlocation,
      response: '',

    }
  }
};


</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
