<template>
  <div id="app">
    <div id="header">
      Roll With Friends
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
background-color: aqua;
}

</style>
