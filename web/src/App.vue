<template>
  <div id="app">
    <nav-bar/>
    <router-view/>
  </div>
</template>
<script>
import NavBar from "@/components/TheNavBar";
export default {
  components: {
    NavBar
  }
};
</script>

<style>
</style>
