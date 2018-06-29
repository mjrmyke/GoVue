<template>
  <div class="hello">
    <h1>{{ title }}</h1>
   Your name:<input type="text" v-model="name"/>

    <ul id="messages"></ul>
  <input type="text" v-model="message"/>
  <button @click="emitEvent">emit</button>
    <button @click="emitDiceRoll(100)">d100</button>
    <button @click="emitDiceRoll(20)">d20</button>
    <button @click="emitDiceRoll(12)">d12</button>
    <button @click="emitDiceRoll(10)">d10</button>
    <button @click="emitDiceRoll(8)">d8</button>
    <button @click="emitDiceRoll(6)">d6</button>
    <button @click="emitDiceRoll(4)">d4</button>

  <div>
    
    <ul>
        <div v-if="responses.length">
        <li v-for="item in responses">
          {{item.from}} said: {{ item.message }}
        </li>
        </div>

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
      document.title =  window.location.hash.split('/')[window.location.hash.split('/').length-1];
      this.ws = new WebSocket("ws://" + window.location.hostname +'/ws'+ window.location.hash.split('#').join(''))
      this.windowLocation = window.location.hash.split('#');
      this.id = Math.floor((Math.random() * 100) + 1);
      this.name = "Guest" + this.id;
    
      this.ws.onopen = function (event) {
        console.log("Opened WS connection");
      }
      
      this.ws.onmessage = event => {
        console.log("received message: " + event.data);
        this.x = JSON.parse(event.data);
          this.responses.push(this.x);
      }
  },
  methods: {
    emitEvent() {
      this.ws.send(JSON.stringify({from: this.name, message: this.message}));
      console.log('event emitted')
    },

    sendWSMessage(message) {
      this.ws.send(JSON.stringify({from: this.name, message: message}));
  },
    emitDiceRoll(typeOfDice) {
      var tempmessage;
      tempmessage = "is rolling a d" + typeOfDice + " which rolled: " + Math.floor((Math.random() * typeOfDice) + 1);
      this.sendWSMessage(tempmessage);
    },
  },
  data: function() {
    return {
      message : '',
      title: 'Roll Room',
      roomid: this.windowlocation,
      responses:[],
      id: this.id,
      name: '',
  

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
