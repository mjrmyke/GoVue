<template>
<div id="container">
  <div class="roomTitleContainer">
    <h2>{{ title }}</h2>
    <h3>Your status is: {{ status }} </h3>
     <span> Your name:<b-form-input class="nameInput" type="text" v-model="newName"/> <br/></span>
  </div>

    <div id="chatWrapper">
      <div id="chatContainer">
        <ul>
        <li v-for="item in responses">
          <span class="time">{{item.time}}</span> <span class="userName">{{item.from}}: {{item.system}}</span> {{item.message}}
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
      <b-form-input type="number" id="numDieInput" class="dieInput"  v-on:keyup.enter="customDie(numDieInput, typeDieInput, constantAdd)" v-model="numDieInput"/>
        d
      <b-form-input type="number" id="typeDieInput" class="dieInput" v-on:keyup.enter="customDie(numDieInput, typeDieInput, constantAdd)"  v-model="typeDieInput"/>
        +
      <b-form-input type="number" id="constantAdd" class="dieInput"  v-on:keyup.enter="customDie(numDieInput, typeDieInput, constantAdd)" v-model="constantAdd"/>
      <b-button class="customDieSubmit" @click="customDie(numDieInput, typeDieInput, constantAdd)">Custom</b-button>


      <b-button @click="emitDiceRoll(100)">d100</b-button>
      <b-button @click="emitDiceRoll(20)">d20</b-button>
      <b-button @click="emitDiceRoll(12)">d12</b-button>
      <b-button @click="emitDiceRoll(10)">d10</b-button>
      <b-button @click="emitDiceRoll(8)">d8</b-button>
      <b-button @click="emitDiceRoll(6)">d6</b-button>
      <b-button @click="emitDiceRoll(4)">d4</b-button>
    </div>
    
  <b-input-group  v-on:keyup.enter="emitEvent" class='input-group'>
    <b-form-input type="text" v-model="message"></b-form-input>
    <b-input-group-append>
      <b-btn  @click="emitEvent" variant="info">Send</b-btn>
    </b-input-group-append>
  </b-input-group>

  </div>
</div>
</template>

<script>
import LO from "lodash";
import cookies from "js-cookie";
// import jscolor from 'jscolor';
/* eslint-disable */
// window.jscolor = jscolor;

export default {
  name: "RollRoom",
  data: function() {
    return {
      message: "",
      title: "Roll Room",
      roomid: "",
      responses: [],
      id: this.id,
      name: "",
      status: "connecting",
      userList: [],
      numDieInput: 1,
      typeDieInput: 4,
      constantAdd: 0,
      newName: "",
      myColor: this.getRandomColor()
    };
  },
  created: function() {
    this.pageInit();
    this.connectionInit();
  },
  watch: {
    newName: LO.debounce(
      function() {
        if (this.newName.length > 0) {
          this.sendWSSystemMessage(
            "Changed name to: " + this.newName,
            this.newName
          );
          this.name = this.newName;
          cookies.set("name", this.name);
        }
      },
      1000,
      { maxWait: 2000 }
    )
  },
  methods: {
    emitEvent() {
      if (this.message === "") {
        return;
      }
      this.sendWSMessage(this.message);
      console.log("event emitted");
      this.message = "";
    },
    pageInit() {
      var savedName = cookies.get("name");
      this.roomid = window.location.hash.split("/")[
        window.location.hash.split("/").length - 1
      ];
      document.title = this.roomid;
      this.ws = new WebSocket(
        "ws://" +
          window.location.hostname +
          "/ws" +
          window.location.hash.split("#").join("")
      );
      this.windowLocation = window.location.hash.split("#");
      this.id = Math.floor(Math.random() * 100 + 1);

      if (savedName !== undefined) {
        this.name = savedName;
        this.newName = savedName;
      } else {
        this.name = "Guest" + this.id;
        this.newName = this.name;
      }
    },
    sendWSMessage(message) {
      this.ws.send(
        JSON.stringify({
          from: this.name,
          message: message,
          system: "",
          data: "",
          room: ""
        })
      );
    },
    sendWSSystemMessage(input, data) {
      this.ws.send(
        JSON.stringify({
          from: this.name,
          system: input,
          room: this.roomid,
          data: data,
          message: ""
        })
      );
    },
    emitDiceRoll(typeOfDice) {
      var tempMessage;
      tempMessage =
        "is rolling a d" +
        typeOfDice +
        " which rolled: " +
        Math.floor(Math.random() * typeOfDice + 1);
      this.sendWSMessage(tempMessage);
    },
    connectionInit() {
      this.ws.onopen = event => {
        console.log("Opened WS connection");
        this.sendWSSystemMessage("connected");
        this.status = "connected";
      };

      this.ws.onclose = event => {
        this.status = "disconnected";
      };

      this.ws.onmessage = event => {
        var element = document.getElementById("chatContainer");
        console.log("received message: " + event.data);
        this.x = JSON.parse(event.data);

        if (this.x.data > 0 || this.x.system.length > 0) {
          if (this.x.system === "connected") {
            this.userList.push(this.x.from);
            this.responses.push({
              time: this.getCurrentTime(),
              from: this.x.from,
              system: "has connected"
            });
          }

          if (LO.includes(this.x.system, "Changed name to: ")) {
            var index = this.userList.indexOf(this.x.from);
            Vue.set(this.userList, index, this.x.data);

            var index2 = this.userList.indexOf(this.x.from);
            Vue.delete(this.userList, index2);
            Vue.delete(this.userList, this.x.from);
          }

          if (this.x.system === "disconnected") {
            var index = this.userList.indexOf(this.x.from);
            Vue.delete(this.userList, index);
          }
        }

        if (
          !LO.includes(this.userList, this.x.from) &&
          this.x.system.length <= 0
        ) {
          this.userList.push(this.x.from);
        }

        this.x.time = this.getCurrentTime();

        this.responses.push(this.x);

        setTimeout(function() {
          element.scrollTop = element.scrollHeight;
        }, 25);
      };
    },
    getCurrentTime() {
      var tempdate = new Date();
      tempdate = tempdate.toTimeString();
      return (tempdate = tempdate.split(" ")[0]);
    },
    customDie(numberOfDie, typeOfDie, constant) {
      var dieRolls = [];
      var total = 0;
      var tempMessage;

      for (var i = 0; i < numberOfDie; i++) {
        var newRoll = Math.floor(Math.random() * typeOfDie + 1);
        dieRolls.push(newRoll);
      }

      total = LO.sum(dieRolls) + Number(constant);

      tempMessage =
        "rolled a " +
        total +
        " by rolling " +
        numberOfDie +
        "d" +
        typeOfDie +
        " + " +
        constant +
        "  rolls: (" +
        dieRolls +
        ")";

      this.sendWSMessage(tempMessage);
    },
    getRandomColor() {
      var letters = "0123456789ABCDEF";
      var color = "#";
      for (var i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }
      return color;
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style lang="scss" scoped>
#container {
  background-color: #42b983;
  width: 100%;
  height: calc(100% - 30px);
  display: flex;
  flex-direction: column;
}

#chatWrapper {
  box-sizing: border-box;
  background-color: green;
  min-width: 100vw;
  display: flex;
  flex-direction: row;
  flex: 1 0 calc(100% - 270px);
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
  justify-content: space-between;
  border-top: lightgrey 1px solid;
  border-bottom: lightgrey 1px solid;
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
  flex: 1 0 100px;
  background-color: cornflowerblue;
}

#userList {
  background-color: white;
  border-left: 10px solid;
  border-right: 25px solid;
  border-image: repeating-linear-gradient(to bottom, #42b983, cornflowerblue) 25;
  width: 10vw;
  min-width: 150px;
  z-index: 1;
}

#userList li {
  margin: 2px;
  padding: 1px;
  background-color: darkcyan;
  color: white;
}

.input-group {
  padding-top: 10px;
  width: 50%;
  margin-left: 25%;
}

.nameInput {
  width: 8vw;
  margin-bottom: 8px;
}
.customDieSubmit {
  margin-right: 100px;
}

.dieInput {
  width: 65px !important;
}

.time {
  padding-right: 4px;
  color: grey;
}

.diceContainer {
  padding-bottom: 5px;
}

.diceContainer input {
  width: 50px;
  padding-left: 10px;
  padding-right: 10px;
}

input {
  display: inline;
  height: 30px;
}

.btn {
  height: 30px;
  line-height: normal !important;
}

b-button:hover {
  border: 1px solid green;
  color: white;
}

.chatInput {
  width: 40vw;
  margin-bottom: 10px;
  margin-top: 5px;
}

.roomTitleContainer {
  display: flex;
  flex-direction: column;
  flex: 1 1 150px;
  height: 16vh;
}

.userName {
  color: blue;
}

h1,
h2 {
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
