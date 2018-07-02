<template>
  <div class="hello">
    <h1>{{ title }}</h1>
      this page will have a list of rooms with rollers rolling, wip
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
  name: 'RoomList',
  data() {
    return {
      title: 'Roll Room Lister',
      ws: null
    };
  },
  created: function() {
      document.title =  window.location.hash.split('/')[window.location.hash.split('/').length-1];
  

  },
  methods: {
  //   emitEvent() {
  //     this.ws.send(JSON.stringify({from: this.name, message: this.message}));
  //     console.log('event emitted')
  //     this.message = "";
  //   },
  //   emitDiceRoll(typeOfDice) {
  //     var tempmessage;
  //     tempmessage = "is rolling a d" + typeOfDice + " which rolled: " + Math.floor((Math.random() * typeOfDice) + 1);
  //     this.sendWSMessage(tempmessage);
  //   },
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
