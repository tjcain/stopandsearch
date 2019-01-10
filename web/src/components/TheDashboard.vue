<template>
  <div>
    <div class="columns">
      <div class="column is-2 has-background-white">
        <the-side-bar @update-query-string="fetchData"/>
      </div>

      <!--  -->
      <div class="column has-background-light">{{queryParams}}</div>
    </div>
  </div>
</template>

<script>
import TheSideBar from "@/components/TheSideBar";
import axios from 'axios'

export default {
  components: {
    TheSideBar
  },
  data() {
    return {
      query: []
    };
  },
  computed: {
    queryParams() {
      return this.query.join("&");
    }
  },
  methods: {
    setQuery(e) {
      this.query = e;
    },
    fetchData(e) {
        this.query = e;
        console.log(e)
        axios.get("api/getcounts?" + this.queryParams)
        .then(r => console.log(r))
    }
  }
};
</script>

<style scoped>
.columns {
  margin: 0 5px;
}
.graph {
  height: 40vh;
}
</style>
