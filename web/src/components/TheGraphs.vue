<template>
       <div class="tile is-ancestor">
        <div class="tile is-parent">
          <div class="tile is-child box graph has-text-centered">
            <p class="title is-6">Age Range</p>
            <graph-age-range :params="queryParams"/>
          </div>
        </div>
        <div class="tile is-parent">
          <div class="tile is-child box graph has-text-centered">
            <p class="title is-6">Ethnicity</p>
            <graph-ethnicity :params="queryParams"/>
          </div>
        </div>
      </div>
</template>

<script>
import GraphAgeRange from "@/components/GraphAgeRange"
import GraphEthnicity from "@/components/GraphEthnicity"
export default {
    components: {
        GraphAgeRange,
        GraphEthnicity
    },
    props: [
        "queryParams"
    ]
}
</script>

<style scoped>


</style>
