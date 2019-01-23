<template>
  <div>
    <app-force-selector />
    <app-date-selector />
   
    <app-drop-down-filter
      v-for="(value, key, index) in categories"
      :key="index"
      :title="key"
      :filters="value"
    />
  </div>
</template>

<script>
import AppDateSelector from "@/components/AppDateSelector";
import AppDropDownFilter from "@/components/AppDropDownFilter";
import AppForceSelector from "@/components/AppForceSelector";

export default {
  components: {
    AppDateSelector,
    AppDropDownFilter,
    AppForceSelector,
  },
  methods: {
    generateSidebar() {
      this.$store.dispatch("fetchCategories");
    }
  },
  computed: {
    categories() {
      return this.$store.state.categories;
    }
  },
  
};
</script>

<style scoped>
</style>
