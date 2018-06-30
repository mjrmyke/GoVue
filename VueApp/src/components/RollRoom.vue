<template>
<div id="container">
  <div class="roomTitleContainer">
    <h2>{{ title }}</h2>
    <h3>Your status is: {{ status }} </h3>
      Your name:<input type="text" v-model="name"/>
  </div>

    <div id="chatWrapper">
      <div id="chatContainer">
        <ul>
            <li v-for="item in userList">
              asd
            </li>
        </ul>
      </div>
      <div id="userList">
        <li v-for="item in responses">
          {{item.name}}
        </li>      
      </div>
    </div>

  <div id="chatInputContainer">
    <div class="diceContainer">
      <button @click="emitDiceRoll(100)">d100</button>
      <button @click="emitDiceRoll(20)">d20</button>
      <button @click="emitDiceRoll(12)">d12</button>
      <button @click="emitDiceRoll(10)">d10</button>
      <button @click="emitDiceRoll(8)">d8</button>
      <button @click="emitDiceRoll(6)">d6</button>
      <button @click="emitDiceRoll(4)">d4</button>
    </div>

    <input type="text" class="chatInput" v-on:keyup.enter="emitEvent" v-model="message"/>
    <button @click="emitEvent">Send</button>
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


      this.connectionInit()


  },
  methods: {
    emitEvent() {
      this.ws.send(JSON.stringify({from: this.name, message: this.message}));
      console.log('event emitted')
      this.message = "";
    },
    sendWSMessage(message) {
      this.ws.send(JSON.stringify({from: this.name, message: message}));
    },
    emitDiceRoll(typeOfDice) {
      var tempmessage;
      tempmessage = "is rolling a d" + typeOfDice + " which rolled: " + Math.floor((Math.random() * typeOfDice) + 1);
      this.sendWSMessage(tempmessage);
    },
    connectionInit() {
      this.ws.onopen = event => {
        console.log("Opened WS connection");
        this.status = "connected";
      }

      this.ws.onclose = event => {
        this.status = "disconnected";
      }

      this.ws.onmessage = event => {
         var element = document.getElementById("chatContainer");
        console.log("received message: " + event.data);
        this.x = JSON.parse(event.data);
        this.responses.push(this.x);
        element.scrollTop = element.scrollHeight;
      }
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
      status: 'connecting',
      userList: [
        {
        name: "placeholder user"
        },
      ],
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

#container {
  background-color: #42b983;
  width: 100%;
  height: 80%;
  display: flex; 
  flex-direction: column;

  min-height: 80vh;

}

#chatWrapper {
  background-color: green;
  height: 75vh;
  min-width: 100vw;
  display: flex;
  flex-direction: row;
}

#chatContainer {  
  height: 100%;
  width: 90%;
  overflow-y: scroll;
  background-color:turquoise;
}

#chatContainer ul {
  text-align: left;
}

#chatInputContainer {
  padding-top: 1vh;
  width: 100%;
  background-color:cornflowerblue;
}
#userList{
  background-color: grey;
  width: 10vw;
  z-index: 1;
}

.chatInput {
  width:  40vw;
  margin-bottom: 10px;
}

.roomTitleContainer {
  height: 16vh;
}

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
