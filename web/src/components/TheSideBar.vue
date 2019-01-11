<template>
  <div>
    <app-force-selector />
    <app-date-selector @date-range="updateDateRange"/>
    <app-check-box @update-checkbox="updateValues"/>
  </div>
</template>

<script>
import AppDateSelector from "@/components/AppDateSelector";
import AppForceSelector from "@/components/AppForceSelector";
import AppCheckBox from "@/components/AppCheckBox";

export default {
  components: {
    AppDateSelector,
    AppForceSelector,
    AppCheckBox
  },
  methods: {
    updateValues(e) {
      this.checkboxquery = e;
      this.$emit("update-query-string", this.requestParams);
    },
    updateDateRange(e) {
      this.daterange = e;
      this.$emit("update-query-string", this.requestParams);
    }
  },
  data() {
    return {
      checkboxquery:"",
      daterange: ""
    }
  },
  computed: {
   requestParams: function() {
     return this.checkboxquery + "&" + this.daterange 
     }
  }
};
</script>

<style>
</style>
