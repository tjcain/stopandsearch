<template>
  <ul>
    <li v-for="(filter, index) in filters" :key="index">
      <nav class="level is-mobile">
        <div class="level-left">
          <div class="level-item">
            <p class="menu-label">{{filter}}</p>
          </div>
        </div>
        <div class="level-right">
          <div class="level-item">
            <div class="field"></div>
            <b-checkbox
              type="is-success"
              :disabled="!active"
              v-model="checkboxGroup"
              :native-value="category+'='+filter.split(' ').join('+')"
            ></b-checkbox>
          </div>
        </div>
      </nav>
    </li>
  </ul>
</template>

<script>
export default {
  data() {
    return {
      checkboxGroup: []
    };
  },
  props: {
    filters: Array,
    category: String,
    active: Boolean
  },
  created() {
    this.checkboxGroup = this.$store.state.filters[this.category];
  },
  watch: {
    checkboxGroup() {
      if (this.active === true) {
        const filtersPayload = {
          key: this.category,
          parameters: this.checkboxGroup
        };
        const queryPayload = {
          key: this.category,
          parameters: this.checkboxGroup.join("&")
        };
        this.$store.dispatch("updateFilters", filtersPayload);
        this.$store.dispatch("updateURLQueryParameters", queryPayload);
      }
    }
  }
};
</script>

<style scoped>
@media only screen and (min-device-width: 1024px) {
  .menu-label {
    max-width: 12vw;
  }
}
li {
  padding: 5px 0;
  opacity: 1;
  transition: opacity 5s;
}
</style>
