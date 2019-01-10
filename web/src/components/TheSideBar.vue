<template>
  <div>
    <app-force-selector :forces="queryParameters.forces"/>
    <app-date-selector @date-range="updateDateRange"/>
    <app-check-box
      v-for="(value, key) in queryParameters"
      :key="key"
      :boxes="queryParameters[key]"
      @change-values="updateValues($event, key)"
    />
  </div>
</template>

<script>
import AppDateSelector from "@/components/AppDateSelector";
import AppCheckBox from "@/components/AppCheckBox";
import AppForceSelector from "@/components/AppForceSelector";
export default {
  components: {
    AppDateSelector,
    AppCheckBox,
    AppForceSelector
  },
  data() {
    return {
      dateRange: "time=09-2015&time=12-2018",
      forceName: "west-midlands",
      queryParameters: {
        ageRange: {
          title: "Age Range",
          queryKey: "age_range",
          parameters: [
            { key: 1, label: "Under 10", value: "under+10" },
            { key: 2, label: "10 to 17", value: "10-17" },
            { key: 3, label: "18 to 24", value: "18-24" },
            { key: 4, label: "25 to 34", value: "24-34" },
            { key: 5, label: "Over 34", value: "over+34" },
            { key: 6, label: "Not Stated", value: "not+stated" }
          ],
          values: [
            "under+10",
            "10-17",
            "18-24",
            "24-34",
            "over+34",
            "not+stated"
          ]
        },
        ethnicity: {
          title: "Ethnicity",
          queryKey: "ethnicity",
          parameters: [
            { key: 1, label: "Black", value: "black" },
            { key: 2, label: "White", value: "white" },
            { key: 3, label: "Asian", value: "asian" }
          ],
          values: ["black", "white", "asian"]
        }
      }
    };
  },
  methods: {
    updateValues(e, k) {
      this.queryParameters[k].values = e;
      this.$emit("update-query-string", this.requestParams);
    },
    updateDateRange(e) {
      this.dateRange = e;
      this.$emit("update-query-string", this.requestParams);
    }
  },
  computed: {
    requestParams() {
      let arr = [];
      const keys = Object.keys(this.queryParameters);
      keys.forEach(e =>
        this.queryParameters[e].values.forEach(j =>
          arr.push(this.queryParameters[e].queryKey + "=" + j)
        )
      );
      // push date range onto the array
      arr.push(this.dateRange);
      // push force onto the array
      arr.push(this.forceName);
      return arr;
    }
  },
  created: function() {
    // `this` points to the vm instance
    this.$emit("update-query-string", this.requestParams);
  }
};
</script>

<style>
</style>
