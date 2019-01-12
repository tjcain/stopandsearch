<template>
  <div>
    <div class="columns">
      <div class="column is-2 has-background-white">
        <the-side-bar @update-query-string="fetchData"/>
      </div>

      <!--  -->
      <div class="column is-10 has-background-light">
        <the-info-boxes :queryParams="query"/>
        <the-graphs :queryParams="query"/>
      </div>

    </div>
  </div>
</template>

<script>
import TheSideBar from "@/components/TheSideBar";
import TheGraphs from "@/components/TheGraphs";
import TheInfoBoxes from "@/components/TheInfoBoxes";
import axios from "axios"

export default {
  components: {
    TheSideBar,
    TheGraphs,
    TheInfoBoxes
  },
  data() {
    return {
      query: []
    };
  },
  computed: {

  },
  methods: {
    fetchData(e) {
        this.query = e
        // axios.get("api/getcounts?" + this.queryParams)
        // .then(r => console.log(r))
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
</style>
