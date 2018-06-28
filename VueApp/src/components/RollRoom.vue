<template>
  <div class="hello">
    <h1> ok</h1>
    <h1>{{ title }}</h1>

    <ul id="messages"></ul>
  <input type="text" v-model="message"/>
  <button @click="emitEvent">emit</button>
  <div>
    
    <ul>
        <li v-for="item in responses">
          {{ item.message }}
        </li>
    </ul>
  
  
  </div>
  </div>
</template>

<script>
/* eslint-disable */


export default {
  name: 'RollRoom',
  data() {
    return {
      title: 'Roll Room Test',
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
        this.responses.push({from: '', message: event.data});
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
      title: 'Roll Room',
      roomid: this.windowlocation,
      responses:[ {
        from: '',
        message: ''
      }]

        ,

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
  display: block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
