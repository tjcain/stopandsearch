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
      let chartInfo = {"Arrest":0,"Other":0,"Community resolution":0,"Nothing Found":0,
      "Caution":0,"Penalty Notice":0,"Summons":0}

      axios.get("api/stats/columns/outcome?" + this.params).then(r => {
        r.data.forEach(function(e){
          chartInfo[e.name] = e.count
        })
        this.datacollection = {
          labels:[["Arrest"], ["Caution"], ["Comminty", "Resolution"], ["Nothing", "Found"], 
          ["Other"],["Penalty", "Notice"], ["Summons"]],
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
