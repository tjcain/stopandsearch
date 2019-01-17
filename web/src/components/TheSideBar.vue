<template>
  <div>
    <app-force-selector/>
    <app-date-selector @date-range="updateDateRange"/> 
    <app-drop-down-filter />


    <!-- <app-outcome-linked-radio @update-radio="updateRadio"/>
    <app-check-box @update-checkbox="updateValues"/> -->
  </div>
</template>

<script>
import AppDateSelector from "@/components/AppDateSelector";
import AppDropDownFilter from "@/components/AppDropDownFilter";
import AppForceSelector from "@/components/AppForceSelector";
import AppOutcomeLinkedRadio from "@/components/AppOutcomeLinkedRadio";
import AppCheckBox from "@/components/AppCheckBox";

export default {
  components: {
    AppDateSelector,
    AppDropDownFilter,
    AppForceSelector,
    AppCheckBox,
    AppOutcomeLinkedRadio
  },
  methods: {
    updateValues(e) {
      this.checkboxquery = e;
      this.$emit("update-query-string", this.requestParams);
    },
    updateDateRange(e) {
      this.daterange = e;
      this.$emit("update-query-string", this.requestParams);
    },
    updateRadio(e) {
      this.outcomelinked = e;
      this.$emit("update-query-string", this.requestParams);
    }
  },
  data() {
    return {
      checkboxquery: "",
      daterange: "",
      outcomelinked: ""
    };
  },
  computed: {
    requestParams: function() {
      if (this.outcomelinked === "") {
        return this.checkboxquery + "&" + this.daterange;
      }
      return (
        this.checkboxquery + "&" + this.daterange + "&" + this.outcomelinked
      );
    }
  }
};
</script>

<style scoped>

</style>
