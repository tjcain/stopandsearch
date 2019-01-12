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
      let chartInfo = {"Other":0,"Offensive weapons":0,"Drugs":0,"Going equipped":0,
      "Firearms":0,"Criminal damage":0,"Stolen property":0}

      axios.get("api/stats/columns/object_of_search?" + this.params).then(r => {
        r.data.forEach(function(e){
          chartInfo[e.name] = e.count
        })
        this.datacollection = {
          labels:[["Criminal", "damage"], ["Drugs"], ["Going", "Equipped"], ["Firearms"], 
      ["Offensive", "Weapons"], ["Other"], ["Stolen", "Property"]],
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
