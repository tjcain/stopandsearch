<template>
  <div class="box noselect">
    <nav class="level is-mobile">
      <div class="level-left" @click="showDateRange = !showDateRange">
        <div class="level-item">
            <span class="icon is-small has-text-primary"
            v-bind:class="[showDateRange ? 'rotated' : '']">
              <i class="fas fa-chevron-down"></i>
            </span>
        </div>
        <div class="level-item">
          <p class="menu-label">Date Range</p>
        </div>
      </div>

      <div class="level-right" @click="active = !active">
        <div class="level-item">
          <b-switch size="is-small" v-model="active" :value="true" type="is-success"></b-switch>
        </div>
      </div>
    </nav>

    <div v-if="showDateRange" class="slider">
      <div class="field">
        <vue-slider
          ref="slider4"
          v-bind="slideBarSettings"
          v-model="sliderModel"
          :disabled="!active"
        ></vue-slider>
      </div>
    </div>
  </div>
</template>

<script>
import slideBarSettings from "@/slideBarSettings.js";
import vueSlider from "vue-slider-component";
export default {
  components: {
    vueSlider
  },
  methods: {
    generateDateParams: function() {
      return this.dateValue;
    }
  },
  computed: {
    dateValue() {
      let from = this.sliderModel[0].split("-");
      from = from[0] + "-01-" + from[1];
      let to = this.sliderModel[1].split("-");
      to = to[0] + "-01-" + to[1];
      return "time=" + from + "&" + "time=" + to;
    }
  },
  data() {
    return {
      slideBarSettings: slideBarSettings,
      sliderModel: ["09-2015", "01-2019"],
      active: false,
      showDateRange: false
    };
  },
  watch: {
    sliderModel() {
      if (this.active) {
        const payload = {
          key: "time",
          parameters: this.dateValue
        };
        this.$store.dispatch("updateURLQueryParameters", payload);
      }
    },
    active() {
      if (this.active === false) {
        const payload = { key: "time", parameters: [] };
        this.$store.dispatch("updateURLQueryParameters", payload);
      } else {
        const payload = {
          key: "time",
          parameters: this.dateValue
        };
        this.$store.dispatch("updateURLQueryParameters", payload);
      }
    }
  }
};
</script>

<style scoped>
@media only screen and (min-device-width: 1024px) {
  .menu-label {
    max-width: 10vw;
  }
}

.level-left, .level-right {
  cursor: pointer;
}

.box {
  padding: 10px;
}
.slider {
  padding-top: 20px;
}

.rotated {
  transform: rotate(180deg); /* Equal to rotateZ(45deg) */
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

