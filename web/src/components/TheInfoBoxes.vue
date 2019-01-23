<template>
  <div  v-if="loaded" class="tile is-ancestor">
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{totalCount}}</p>
        <p class="sub-title">Searches</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{arrestRate}}%</p>
        <p class="sub-title">Arrest Rate</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{positiveOutcome}}%</p>
        <p class="sub-title">Positive Outcome</p>
      </div>
    </div>
    <div class="tile is-parent">
      <div class="tile is-child box">
        <p class="title">{{noOutcome}}%</p>
        <p class="sub-title">No Outcome</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data(){
      return {}
  }, 
  computed:{
    loaded() {
      return this.$store.state.loaded
    },
    totalCount() {
      return this.$store.getters.formatCount
    },
    positiveOutcome() {
      return this.$store.getters.getPositiveOutcomePercent
    },
    noOutcome() {
      return this.$store.getters.getNoOutcomePercent
    },
    arrestRate() {
      return this.$store.getters.getArrestPercentage
    }
  },
  methods: {}
      
};
</script>

<style scoped>
p.title {
  margin-bottom: 5px
}
</style>
