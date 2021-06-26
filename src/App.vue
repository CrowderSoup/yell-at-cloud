<template>
  <div id="app">
    <Header />
    <Yell />
  </div>
</template>

<script>
import Header from './components/Header.vue'
import Yell from './components/yell.vue'

export default {
  name: 'app',
  components: {
    Header,
    Yell,
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
