<template>
  <div id="app">
    <div id="header">
      <div id="links">
        <b-link href="/">home</b-link>
        <b-link href="#/roll">roll rooms</b-link>
      </div>
      <div>
        Roll With Friends
      </div>
    </div>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
};


</script>

<style>
html, body {
  height: 100%;
}

#app {
  display: flex;
  flex-direction: column;
  margin: 0;
  height: 100%;
  width: 100%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#header {
  display: flex;
  height: 40px;
  justify-content: space-between;
  flex-direction: row;
  color: white;
  background-color: #336699;
  padding: 3px;
  border-bottom: 10px solid;
  border-image: repeating-linear-gradient(to top, #42b983, #336699) 20;
}

#header a {
  color: white;
  padding-left: 10px;
}


#header a:hover {
  color: repeating-linear-gradient(to bottom, #42b983, cornflowerblue) 25;
  border-image: repeating-linear-gradient(to right, opaque, #336699) 10;
}

</style>
