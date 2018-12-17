<template>
    <section class="section">
        <div class="tile is-ancestor">
            <div class="tile is-parent is-4">
                <article class="tile is-child">
                <p class="title"> >> Welcome to the stop and search data tool. To get started, select your query parameters...</p>
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                    <parameter-selection />
                </article>
            </div>
        </div>

        <div class="tile is-ancestor">
            <div class="tile is-parent is-4">
                <article class="tile is-child">
                <p class="title"> >> Now select a column of the database table to examine...</p>
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                    <query-selection />
                </article>
            </div>
        </div>

        <div class="tile is-ancestor">
            <div class="tile is-parent is-4">
                <article class="tile is-child">
                <p class="title"> >> You can further refine your results by applying filters..</p>
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                  <filter-selection />  
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                  <filter-selection />  
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                  <filter-selection />  
                </article>
            </div>
            <div class="tile is-parent">
                <article class="tile is-child box">
                  <filter-selection />  
                </article>
            </div>
        </div>

       
    </section>
</template>

<script>
import VueTinytyper from "vue-tinytyper";
import ParameterSelection from "./ParameterSelection.vue";
import QuerySelection from "./QuerySelection.vue";
import FilterSelection from "./FilterSelection.vue";

export default {
  components: {
    VueTinytyper,
    ParameterSelection,
    QuerySelection,
    FilterSelection
  },
  data: function() {
    return {
      welcomeMessage: {
        text:
          ">> Welcome to the stop and search data tool. To get started, select your query parameters...",
        textSpeed: 10,
        cssClass: "title is-3 has-text-dark is-unselectable"
      },
      queryMessage: {
        text: ">> Now select a column of the database table to examine...",
        textSpeed: 10,
        cssClass: "title is-3 has-text-dark is-unselectable",
        display: false
      },
      filterMessage: {
        text: ">> You can further refine your results by applying filters..",
        textSpeed: 10,
        cssClass: "title is-3 has-text-dark is-unselectable",
        display: false
      }
    };
  },
  mounted() {
    // Set to true after 5 seconds
    setTimeout(() => {
      this.queryMessage.display = true;
    }, 3000);
    setTimeout(() => {
      this.filterMessage.display = true;
    }, 6000);
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@media screen and (min-width: 767px) {
  .padding-top {
    padding-top: 43px;
  }
}
</style>
