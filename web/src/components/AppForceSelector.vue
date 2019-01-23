<template>
  <div class="box noselect">
    <nav class="level is-mobile">
      <div class="level-left" @click="show = !show">
        <div class="level-item">
            <span class="icon is-small has-text-primary"
            :class="[this.show ? 'rotated' : '']">
              <i class="fas fa-chevron-down"></i>
            </span>
        </div>
        <div class="level-item">
          <p class="menu-label">Police Force</p>
        </div>
      </div>
      <div class="level-right">
        <!-- when enabled, add click event to level-right, same as other filters -->
        <div class="level-item">
          <b-switch size="is-small nudge-right" :value="true" type="is-success" disabled></b-switch>
        </div>
      </div>
    </nav>

    <div v-if="show" class="dropdown">
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
      forceSelection: "All Police Forces",
      show: false
      //   columns: ["Outcomes", "Ethnicities", "Age groups", "Objects of search"]
    };
  }
};
</script>

<style scoped>
@media only screen and (min-device-width: 1024px) {
  .menu-label {
    max-width: 10vw;
  }
}
.box {
  padding: 10px;
}
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
  margin: 1px 0;
}
.rotated {
  transform: rotate(180deg); /* Equal to rotateZ(45deg) */
}

.level-left, .level-right {
  cursor: pointer;
}

.noselect {
  -webkit-touch-callout: none; /* iOS Safari */
    -webkit-user-select: none; /* Safari */
     -khtml-user-select: none; /* Konqueror HTML */
       -moz-user-select: none; /* Firefox */
        -ms-user-select: none; /* Internet Explorer/Edge */
            user-select: none; /* Non-prefixed version, currently
                                  supported by Chrome and Opera */
}
</style>
