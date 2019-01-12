<template>
  <div class="tile is-ancestor">
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{count}}</p>
        <p class="sub-title">Searches</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{arrestsPercent}}%</p>
        <p class="sub-title">Arrest Rate</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{positiveOutcomePercent}}%</p>
        <p class="sub-title">Positive Outcome</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{negativeOutcomePercent}}%</p>
        <p class="sub-title">No Outcome</p>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios"
export default {
  props: ["queryParams"],
  data(){
      return {
          count: 0,
          arrestRate: 0,
          negativeOutcome: 0,
          positiveOutcome: 0
    }
  }, 
  computed:{
      arrestsPercent() {
          return Math.round((this.arrestRate / this.count) * 100)
      },
      negativeOutcomePercent() {
          return Math.round((this.negativeOutcome / this.count) * 100)
      },
       positiveOutcomePercent() {
          return Math.round((this.positiveOutcome / this.count) * 100)
      }
  },
  methods: {
      fetchCount() {
        axios.get("api/stats/count?" + this.queryParams).then(r => {
        this.count = r.data;
      });
      },
      fetchArrestRate(){
          axios.get("api/stats/columns/outcome?" + this.queryParams).then(r => {
              
              let i = 0
              r.data.forEach(function(n){
                if (n.name === "Arrest" || n.name === "Suspect arrested"){
                i = n.count
              }
          })
        this.arrestRate = i
          })
      },
      fetchOutcomes() {
          axios.get("api/stats/columns/search_happened?" + this.queryParams).then(r => {
              let t = 0
              let f = 0
              r.data.forEach(function(n){
              if (n.name === "true"){
                t = n.count
              }
              if (n.name === "false"){
                f = n.count
              }
          })
            this.positiveOutcome = t
            this.negativeOutcome = f
          })
      }
  },
  watch: {
      queryParams: function() {
          this.fetchCount()
          this.fetchArrestRate()
          this.fetchOutcomes()
      }
  }
};
</script>

<style>
</style>
