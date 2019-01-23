<template>
  <div>
    <line-example :options="options" :chart-data="datacollection"></line-example>
  </div>
</template>

<script>
import LineExample from "./Egchart.vue";
export default {
  computed: {
    chartData() {
      return this.$store.getters.getOutcomeData;
    }
  },
  components: {
    LineExample
  },
  data() {
    return {
      loaded: false,
      datacollection: null,
      options: {
        maintainAspectRatio: false,
        legend: {
          display: false
        },
        scales: {
          xAxes: [
            {
              gridLines: {
                display: false
              },
              ticks: {
                fontSize: 9
              }
            }
          ],
          yAxes: [
            {
              gridLines: {
                display: false
              }
            }
          ]
        }
      }
    };
  },
  methods: {
    fillData() {
      let chartInfo = {
        "Arrest": 0,
        "Other": 0,
        "Community resolution": 0,
        "Nothing Found": 0,
        "Caution": 0,
        "Penalty Notice": 0,
        "Summons": 0
      };

      this.chartData.forEach(function(e) {
        chartInfo[e.name] = e.count;
      });
      this.datacollection = {
        labels: [
          ["Arrest"],
          ["Caution"],
          ["Comminty", "Resolution"],
          ["Nothing", "Found"],
          ["Other"],
          ["Penalty", "Notice"],
          ["Summons"]
        ],
        datasets: [
          {
            backgroundColor: "rgb(140, 103, 239)",
            data: Object.values(chartInfo)
          }
        ]
      };
    }
  },
  watch: {
    params: function() {
      this.fillData();
    }
  },
  mounted() {
    this.fillData();
  }
};
</script>

<style scoped>
div {
  height: 30vh;
}
</style>
