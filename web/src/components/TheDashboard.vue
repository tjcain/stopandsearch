<template>
<div>
  <div class="loading" v-if="!loaded">
    <b-loading :is-full-page="false" :active="!loaded"></b-loading>
  </div>
  <div v-else>
    <div class="columns">
      <div class="column is-3">
        <the-side-bar/>
      </div>

      <!--  -->
      <div class="column is-9">
        <the-info-boxes/>

        <the-graphs />
      </div>
    </div>
  </div>
</div>
</template>

<script>
import TheSideBar from "@/components/TheSideBar";
import TheGraphs from "@/components/TheGraphs";
import TheInfoBoxes from "@/components/TheInfoBoxes";

export default {
  components: {
    TheSideBar,
    TheGraphs,
    TheInfoBoxes
  },
  computed: {
    loaded() {
      return this.$store.state.loaded;
    },
    data() {
      return this.$store.state.searchData;
    },
    getQueryString() {
      return this.$store.getters.getQueryString;
    },
    age() {
      return this.$store.getters.getAgeRangeData;
    }
  },
  watch: {
    getQueryString() {
      this.$store.dispatch("fetchCount", this.getQueryString);
      this.$store.dispatch("fetchSearchData", this.getQueryString);
    }
  }
};
</script>

<style scoped>
.columns {
  margin: 0 5px;
}
.graph {
  height: 40vh;
}
.loading {
  min-height: 60vh;
  position: relative;
}
</style>
