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
        <li v-for="item in responses">
          <span class="userName">{{item.from}}:</span> {{item.message}}
        </li>    

        </ul>
      </div>
      <div id="userList">
        <li v-for="item in userList">
              {{item}}
        </li>
      </div>
    </div>

  <div id="chatInputContainer">
    <div class="diceContainer">
      <input type="text" id="numDieInput" class="dieInput" v-model="numDieInput"/>
        d
      <input type="text" id="typeDieInput" class="dieInput" v-model="typeDieInput"/>
        +
      <input type="text" id="constantAdd" class="dieInput" v-model="constantAdd"/>
      <button class="customDieSubmit" @click="customDie(numDieInput, typeDieInput, constantAdd)">Custom</button>


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
      this.pageInit()
      this.connectionInit()
  },
  methods: {
    emitEvent() {
      this.ws.send(JSON.stringify({from: this.name, message: this.message}));
      console.log('event emitted')
      this.message = "";
    },
    pageInit() {
      document.title =  window.location.hash.split('/')[window.location.hash.split('/').length-1];
      this.ws = new WebSocket("ws://" + window.location.hostname +'/ws'+ window.location.hash.split('#').join(''))
      this.windowLocation = window.location.hash.split('#');
      this.id = Math.floor((Math.random() * 100) + 1);
      this.name = "Guest" + this.id;
    },
    sendWSMessage(message) {
      this.ws.send(JSON.stringify(
        { 
          from: this.name, 
          message: message
          }
        )
      );
    },
    emitDiceRoll(typeOfDice) {
      var tempMessage;
      tempMessage = "is rolling a d" + typeOfDice + " which rolled: " + Math.floor((Math.random() * typeOfDice) + 1);
      this.sendWSMessage(tempMessage);
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

        if (!this.userList.includes(this.x.from)) {
          this.userList.push(this.x.from);
        }

        this.responses.push(this.x);
        element.scrollTop = element.scrollHeight;
      }
    },
    customDie(numberOfDie, typeOfDie, constant) {
      var dieRolls = []
      var total = 0;
      var tempMessage;

      for (var i=0; i < numberOfDie; i++) {
        var newRoll = Math.floor((Math.random() * typeOfDie) + 1);
        dieRolls.push(newRoll);
      }
      
      total = LO.sum(dieRolls) + Number(constant);
      
      tempMessage = "roll a " + total + " by rolling " + numberOfDie + "d" + typeOfDie 
                    + " + " + constant + "  rolls: (" + dieRolls + ")";

      this.sendWSMessage(tempMessage);
    }

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
      userList: ['asd'],
      numDieInput: 1,
      typeDieInput: 4,
      constantAdd: 0,
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
  box-sizing: border-box;
  background-color: green;
  height: 75vh;
  min-width: 100vw;
  display: flex;
  flex-direction: row;
}

#chatContainer {  
  border-left: 25px solid;
  border-image: repeating-linear-gradient(to bottom, #42b983, cornflowerblue) 25;
  height: 100%;
  width: 90%;
  overflow-y: scroll;
  background-color: white;
}

#chatContainer li {
  border-top:lightgrey 1px solid;
  border-bottom:lightgrey 1px solid;
  background-color: lightcyan;
  padding-top: 2px;
  padding-bottom: 2px;
}

#chatContainer ul {
  text-align: left;
}

#chatInputContainer {
  padding: 1vh;
  width: 100%;
  background-color:cornflowerblue;
}

#userList{
  background-color: white;
  border-left: 10px solid;
  border-right: 25px solid;
  border-image: repeating-linear-gradient(to bottom, #42b983, cornflowerblue) 25;
  width: 10vw;
  z-index: 1;
}

.customDieSubmit {
  margin-right: 100px;
}

.dieInput {
  width:25px;
}

.diceContainer {
  padding-bottom: 5px;
}

.chatInput {
  width:  40vw;
  margin-bottom: 10px;  
}

.roomTitleContainer {
  height: 16vh;
}

.userName {
  color: blue;
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
