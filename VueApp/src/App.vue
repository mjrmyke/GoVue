<template>
  <div id="app">
    <div id="header">
      <div id="links">
        <a href="/">home</a>
        <a href="#/roll">roll rooms</a>
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
#app {
  margin: 0;
  height: 100%;
  width: 100%;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

body {
  overflow: hidden;
  min-height: 100vh;
  min-width: 100vh;
  margin: 0;
}

#header {
  display: flex;
  justify-content: space-between;
  flex-direction: row;
  color: white;
  background-color: #336699;
  padding: 3px;
  border-bottom: 10px solid;
  border-image: repeating-linear-gradient(to top, #42b983, #336699) 20;
}

#header a {
  text-decoration: none;
  color: white;
  padding: 5px;
}

#header a:hover {
  border-bottom: 2px dashed;
  color: repeating-linear-gradient(to bottom, #42b983, cornflowerblue) 25;
  border-image: repeating-linear-gradient(to right, opaque, #336699) 10;
  text-decoration: none;
  padding: 5px;
}

</style>
