<template>
  <div class="container padding-top">
    <!-- START SMALL BOXES -->
    <div class="tile is-ancestor">
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">34234</p>
          <p class="sub-title">Searches</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">12%</p>
          <p class="sub-title">Arrest Rate</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">31%</p>
          <p class="sub-title">Positive Outcome</p>
        </div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">53%</p>
          <p class="sub-title">No Outcome</p>
        </div>
      </div>
    </div>
    <!-- END SMALL BOXES -->
    <div class="tile is-ancestor">
      <div class="tile is-parent">
        <div class="tile is-child box" style="height:50px"></div>
      </div>
      <div class="tile is-parent">
        <div class="tile is-child box">
          <p class="title">GRAPH_ONE</p>
          <p class="content">
            <b>Selection:</b>
            {{ checkboxGroup }}
          </p>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import ParameterSelection from "./ParameterSelection.vue";
import QuerySelection from "./QuerySelection.vue";
import FilterSelection from "./FilterSelection.vue";
// import RandomChart from "./RandomChart.vue";
import LineExample from "./Egchart.vue";
import axios from "axios";

export default {
  components: {
    ParameterSelection,
    QuerySelection,
    FilterSelection,
    LineExample
  },
  data: function() {
    return {
      labels: [],
      values: [],
      loaded: false,
      checkboxGroup: ["under 10", "10-17", "18-24", "25-34", "over 34"]
    };
  },
  watch: {
    checkboxGroup: function() {
      this.loaded = false;
      let params = new URLSearchParams();
      this.checkboxGroup.forEach(r => {
        params.append("age_range", r);
      });
      let request = {
        params: params
      };
      console.log(request);
      axios.get("api/stats/ethnicity", request).then(r => {
        this.labels = r.data.map(d => d.name);
        this.values = r.data.map(d => d.count);
        this.loaded = true;
      });
    }
  },
  mounted() {
    var params = new URLSearchParams();
    params.append("age_range", "10-17");
    params.append("age_range", "18-24");
    let request = {
      params: params
    };

    this.loaded = false;
    axios.get("api/stats/ethnicity", request).then(r => {
      this.labels = r.data.map(d => d.name);
      this.values = r.data.map(d => d.count);
      this.loaded = true;
    });
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@media screen and (min-width: 767px) {
  .padding-top {
    padding-top: 20px;
  }
}
</style>
