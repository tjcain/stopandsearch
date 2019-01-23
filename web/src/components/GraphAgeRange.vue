<template>
  <div>
    <line-example :options="options" :chart-data="datacollection"></line-example>
  </div>
</template>

<script>
import LineExample from "./Egchart.vue";
export default {
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
  computed: {
    chartData() {
      return this.$store.getters.getAgeRangeData;
    }
  },
  watch: {
    chartData() {
      this.fillData()
    }
  },
  methods: {
    fillData() {
      let chartInfo = {
        "under 10": 0,
        "10-17": 0,
        "18-24": 0,
        "25-34": 0,
        "over 34": 0,
        "Not Stated": 0
      };

      this.chartData.forEach(function(e) {
        chartInfo[e.name] = e.count;
      });

      this.datacollection = {
        labels: Object.keys(chartInfo),
        datasets: [
          {
            backgroundColor: "rgb(140, 103, 239)",
            data: Object.values(chartInfo)
          }
        ]
      };
    }
  },
  mounted() {
    this.fillData()
  }
};
</script>

<style scoped>
div {
  height: 30vh;
}
</style>
