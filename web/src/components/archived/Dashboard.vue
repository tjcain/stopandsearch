<template>
  <div class="columns">
    <div class="column is-2 has-background-white">
      <dashboard-side-options v-model="urlparams" v-on:params="urlparams = $event"></dashboard-side-options>
    </div>
    <div class="column has-background-light">
      <div class="tile is-ancestor">
        <div class="tile is-parent">
          <div class="tile is-child box">
            <p class="title">{{ count }}</p>
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
          <div class="tile is-child box graph">
            <age-range :list="urlparams"/>
          </div>
        </div>
        <div class="tile is-parent">
          <div class="tile is-child box">
            <p class="title">GRAPH_ONE</p>
            <p class="content"></p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
import DashboardSideOptions from "@/components/DashboardSideOptions.vue";
import AgeRange from "@/components/AgeRange.vue";
export default {
  data() {
    return {
      urlparams: "",
      count: 0
    };
  },
  components: {
    DashboardSideOptions,
    AgeRange
  },
  mounted() {
    this.getCount();
  },
  watch: {
    urlparams: function() {
      this.getCount();
    }
  },
  methods: {
    getCount: function() {
      console.log("count");
      axios.get("api/count?" + this.urlparams).then(r => {
        this.count = r.data;
        console.log(r.data);
      });
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
