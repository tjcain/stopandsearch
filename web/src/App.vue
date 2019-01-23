<template>
  <div id="app">
    <nav-bar/>
    <router-view/>
    <the-footer/>
  </div>
</template>
<script>
import NavBar from "@/components/TheNavBar";
import TheFooter from "@/components/TheFooter";
export default {
  components: {
    NavBar,
    TheFooter
  },
  computed: {
    getQueryString() {
      return this.$store.getters.getQueryString;
    }
  },
  async created() {
   await this.$store.dispatch("fetchCategories");
   await this.$store.dispatch("fetchSearchData", this.getQueryString);
   await this.$store.dispatch("fetchCount", this.getQueryString);
  }
};
</script>

<style lang="scss">
@import "~bulma/sass/utilities/_all";

// $family-primary: 'Raleway';
// $family-secondary: 'Raleway';
$body-background-color: #f4f4f4;

@import "~bulma";
@import "~buefy/src/scss/buefy";
</style>