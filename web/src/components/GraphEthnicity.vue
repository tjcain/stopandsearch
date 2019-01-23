<template>
  <div>
    <line-example :options="options" :chart-data="datacollection"></line-example>
  </div>
</template>

<script>
import LineExample from "./Egchart.vue";
export default {
  props: ["params"],
  components: {
    LineExample
  },
  computed: {
    chartData() {
      return this.$store.getters.getEthnicityData;
    }
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
        Asian: 0,
        Black: 0,
        "Chinese or other": 0,
        Mixed: 0,
        "Not stated": 0,
        White: 0
      };

      this.chartData.forEach(function(e) {
        chartInfo[e.name] = e.count;
      });

      this.datacollection = {
        labels: [
          ["Asian"],
          ["Black"],
          ["Chinese"],
          ["Mixed"],
          ["Not stated"],
          ["White"]
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
    chartData() {
      this.fillData();
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
 