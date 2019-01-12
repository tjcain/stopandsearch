<template>
  <div>
    <line-example :options="options" :chart-data="datacollection"></line-example>
  </div>
</template>

<script>
import axios from "axios";
import LineExample from "./Egchart.vue";
export default {
  props: ["params"],
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
        xAxes: [{
            gridLines: {
                display:false
            },
                ticks: {
                fontSize: 10
            }
        }],
        yAxes: [{
            gridLines: {
                display:false
            }   
        }]
    }
       }
    };
  },
  methods: {
    fillData() {
      let chartInfo = {"under 10": 0, "10-17": 0, "18-24": 0, "25-34": 0, 
      "over 34": 0, "Not Stated": 0}
      axios.get("api/stats/columns/age_range?" + this.params).then(r => {
        r.data.forEach(function(e){
          chartInfo[e.name] = e.count
        })
        this.datacollection = {
          labels: Object.keys(chartInfo),
          datasets: [
            {
              label: "Age Range",
              backgroundColor: "#f87979",
              data: Object.values(chartInfo)
            }
          ]
        };
      });
    }
  },
  watch: {
    params: function() {
      this.fillData()
    }
  }
};
</script>

<style scoped>
div {
  height: 30vh
}
</style>
