<template>
  <div>
    <p class="menu-label">Police Force</p>
    <div class="dropdown">
      <div class="field">
        <b-select size="is-small" placeholder="West Midlands" disabled>
          <button class="button is-primary" slot="trigger">
            <span>West Midlands</span>
            <b-icon icon="menu-down"></b-icon>
          </button>
          <option>West Midlands</option>
        </b-select>
      </div>
    </div>
  </div>
</template>

<script>
import sourceData from "@/data.json";
export default {
  data() {
    return {
      options: sourceData.data,
      forces: sourceData.forces,
      forceSelection: "All Police Forces"
      //   columns: ["Outcomes", "Ethnicities", "Age groups", "Objects of search"]
    };
  }
};
</script>

<style scoped>
    .field:last-child {
    margin-bottom: 1px;
    }
    .field {
    margin: 1px 0px;
    }
    .checkboxes {
    max-height: 15vh;
    overflow-y: scroll;
    }
    .menu-label {
    margin: 1px 0
    }
</style>
