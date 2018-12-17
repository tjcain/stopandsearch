<template>
   <div>
       <nav class="navbar" role="navigation" aria-label="main navigation">
  <div class="navbar-brand">
    <router-link class="navbar-item" to="/"> 
      <img src="../../public/logo.png" width="112" height="28">
    </router-link>
  
    <a role="button" 
    class="navbar-burger burger" 
    aria-label="menu" aria-expanded="false" 
    @click="showNav = !showNav" 
    :class="{ 'is-active': showNav }">
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
      <span aria-hidden="true"></span>
    </a>
  </div>

  <div id="navbarBasicExample" class="navbar-menu" :class="{ 'is-active': showNav }">
    <div class="navbar-end">
        <router-link class="navbar-item" to="/about">About</router-link>
        <router-link class="navbar-item" to="/">Contact</router-link>
    </div>
  </div>
</nav>
   </div>
</template>

<script>
export default {
  data: function() {
    return {
      showNav: false
    };
  }
};
</script>

<style>
</style>
