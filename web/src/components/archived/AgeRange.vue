<template>
  <div>
    <line-example :chart-data="datacollection"></line-example>
  </div>
</template>

<script>
import axios from "axios";
import LineExample from "./Egchart.vue";
export default {
  props: ["list"],
  components: {
    LineExample
  },
  data() {
    return {
      returns: [],
      labels: [],
      values: [],
      loaded: false,
      datacollection: null
    };
  },
  mounted() {
    this.fillData();
  },
  methods: {
    fillData() {
      axios.get("api/stats/age_range?" + this.list).then(r => {
        this.datacollection = {
          labels: ["Under 10", "10-17", "18-24", "25-34", "Over 34"],
          datasets: [
            {
              label: "Age Range",
              backgroundColor: "#f87979",
              data: r.data.map(d => d.count)
            }
          ]
        };
      });
    }
  },
  watch: {
    list: function() {
      if (!this.list.includes("age_range=")) {
        return (this.datacollection = {
          labels: ["Under 10", "10-17", "18-24", "25-34", "Over 34"],
          datasets: [
            {
              label: "Age Range",
              backgroundColor: "#f87979",
              data: []
            }
          ]
        });
      }
      axios.get("api/stats/age_range?" + this.list).then(r => {
        let t = Object.values(r.data);
        let total = t.reduce((a, b) => a + b, 0);

        this.datacollection = {
          labels: ["Under 10", "10-17", "18-24", "25-34", "Over 34"],
          datasets: [
            {
              label: "Age Range",
              backgroundColor: "#f87979",
              data: [
                r.data["under 10"],
                r.data["10-17"],
                r.data["18-24"],
                r.data["25-34"],
                r.data["over 34"]
              ]
            }
          ]
        };
      });
    }
  }
};
</script>

<style>
</style>
