<template>
  <div id="app">
    <nav-bar/>
    <router-view/>
    <the-footer />
  </div>
</template>
<script>
import NavBar from "@/components/TheNavBar";
import TheFooter from "@/components/TheFooter";
export default {
  components: {
    NavBar,
    TheFooter
  }
};
</script>

<style>
</style>
