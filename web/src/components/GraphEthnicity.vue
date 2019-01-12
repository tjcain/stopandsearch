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
            },
            
        }]
    }
       }
    };
  },
  methods: {
    fillData() {
      let chartInfo = {"Asian": 0, "Black": 0, "Chinese or other": 0, "Mixed": 0, 
      "Not stated": 0, "White": 0}
      axios.get("api/stats/columns/ethnicity?" + this.params).then(r => {
        r.data.forEach(function(e){
          chartInfo[e.name] = e.count
        })
        this.datacollection = {
          labels:[["Asian"], ["Black"], ["Chinese"], ["Mixed"], 
      ["Not stated"], ["White"]],
          datasets: [
            {
              label: "Ethnicity",
              backgroundColor: "rgb(140, 103, 239)",
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
